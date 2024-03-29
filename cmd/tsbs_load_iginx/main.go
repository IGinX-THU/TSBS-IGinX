// bulk_load_iginx loads an Iginx daemon with data from stdin.
//
// The caller is responsible for assuring that the database is empty before
// bulk load.
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/blagojts/viper"
	"github.com/spf13/pflag"
	"github.com/timescale/tsbs/internal/utils"
	"github.com/timescale/tsbs/load"
	"github.com/timescale/tsbs/pkg/targets"
	"github.com/timescale/tsbs/pkg/targets/constants"
	"github.com/timescale/tsbs/pkg/targets/initializers"
)

// Program option vars:
var (
	doAbortOnExist       bool
	connectionSocketList []string
	defaultTruck         = "unknown"
	defaultTagK          = []string{"fleet", "driver", "model", "device_version"}
	defaultTagV          = []string{"unknown", "unknown", "unknown", "unknown"}
)

// Global vars
var (
	loader  load.BenchmarkRunner
	config  load.BenchmarkRunnerConfig
	bufPool sync.Pool
	target  targets.ImplementedTarget
)

// allows for testing
var fatal = log.Fatalf

// Parse args:
func init() {
	target = initializers.GetTarget(constants.FormatIginx)
	config = load.BenchmarkRunnerConfig{}

	config.AddToFlagSet(pflag.CommandLine)
	target.TargetSpecificFlags("", pflag.CommandLine)
	pflag.Parse()

	err := utils.SetupConfigFile()

	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	if err := viper.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("unable to decode config: %s", err))
	}

	var connectionStrings = viper.GetString("connStr")

	connectionSocketList = strings.Split(connectionStrings, ",")
	if len(connectionSocketList) == 0 {
		log.Fatal("missing 'connStr' flag")
	}

	config.HashWorkers = false
	loader = load.GetBenchmarkRunner(config)
}

type benchmark struct{}

func (b *benchmark) GetDataSource() targets.DataSource {
	return &fileDataSource{scanner: bufio.NewScanner(os.Stdin)}
}

func (b *benchmark) GetBatchFactory() targets.BatchFactory {
	return &factory{}
}

func (b *benchmark) GetPointIndexer(_ uint) targets.PointIndexer {
	return &targets.ConstantIndexer{}
}

func (b *benchmark) GetProcessor() targets.Processor {
	return &processor{}
}

func (b *benchmark) GetDBCreator() targets.DBCreator {
	return &dbCreator{}
}

func main() {
	bufPool = sync.Pool{
		New: func() interface{} {
			return bytes.NewBuffer(make([]byte, 0, 4*1024*1024))
		},
	}

	loader.RunBenchmark(&benchmark{})
}
