package main

import (
	"flag"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/telosprotocol/teth-relayer/admin"
	"github.com/telosprotocol/teth-relayer/common"
	config "github.com/telosprotocol/teth-relayer/config"
	"github.com/telosprotocol/teth-relayer/executor"
	"github.com/telosprotocol/teth-relayer/relayer"
)

const (
	flagConfigPath = "config-path"
)

func initFlags() {
	flag.String(flagConfigPath, "", "config file path")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		panic(err)
	}
}

func printUsage() {
	fmt.Print("usage: ./teth-relayer --config-path configFile\n")
}

func main() {
	initFlags()

	configFilePath := viper.GetString(flagConfigPath)
	if configFilePath == "" {
		printUsage()
		return
	}
	cfg := config.ParseConfigFromFile(configFilePath)

	if cfg == nil {
		fmt.Println("failed to get configuration")
		return
	}

	common.InitLogger(&cfg.LogConfig)

	var db *gorm.DB
	// if cfg.DBConfig.DBPath != "" {
	// 	var err error
	// 	db, err = gorm.Open(cfg.DBConfig.Dialect, cfg.DBConfig.DBPath)
	// 	if err != nil {
	// 		panic(fmt.Sprintf("open db error, err=%s", err.Error()))
	// 	}
	// 	defer db.Close()
	// 	model.InitTables(db)
	// }

	adm := admin.NewAdmin(db, cfg)
	go adm.Serve()

	topExecutor, err := executor.NewTopExecutor(cfg, config.Network)
	if err != nil {
		common.Logger.Error(err.Error())
		return
	}
	tethExecutor, err := executor.NewTethExecutor(db, cfg)
	if err != nil {
		common.Logger.Error(err.Error())
		return
	}

	startHeight, err := topExecutor.GetLatestBlockHeight("Ta0000@0")
	if err != nil {
		common.Logger.Error(err.Error())
		return
	}

	relayerInstance := relayer.NewRelayer(db, cfg, topExecutor, tethExecutor)
	common.Logger.Info("Starting relayer")
	relayerInstance.Start(uint64(startHeight))

	select {}
}
