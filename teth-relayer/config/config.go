package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/ethereum/go-ethereum/common"
)

type ChannelConfig struct {
	ChannelID      int8           `json:"channel_id"`
	Method         string         `json:"method"`
	ABIName        string         `json:"abi_name"`
	ContractAddr   common.Address `json:"contract_addr"`
	SequenceMethod string         `json:"sequence_method"`
}

type Config struct {
	CrossChainConfig CrossChainConfig `json:"cross_chain_config"`
	TopConfig        TopConfig        `json:"top_config"`
	TethConfig       TethConfig       `json:"teth_config"`
	LogConfig        LogConfig        `json:"log_config"`
	AdminConfig      AdminConfig      `json:"admin_config"`
	DBConfig         DBConfig         `json:"db_config"`
}

type CrossChainConfig struct {
	SourceChainID uint16 `json:"source_chain_id"`
	DestChainID   uint16 `json:"dest_chain_id"`
}

func (cfg *CrossChainConfig) Validate() {
}

type AdminConfig struct {
	ListenAddr string `json:"listen_addr"`
}

func (cfg *AdminConfig) Validate() {
	if cfg.ListenAddr == "" {
		panic("listen address should not be empty")
	}
}

type TopConfig struct {
	RpcAddrs []string `json:"rpc_addrs"`
}

func (cfg *TopConfig) Validate() {
	if len(cfg.RpcAddrs) == 0 {
		panic("rpc endpoint of Top chain should not be empty")
	}
}

type TethConfig struct {
	PrivateKey string   `json:"private_key"`
	Providers  []string `json:"providers"`
	GasLimit   uint64   `json:"gas_limit"`
	GasPrice   uint64   `json:"gas_price"`
}

func (cfg *TethConfig) Validate() {
	if len(cfg.Providers) == 0 {
		panic("provider address of teth should not be empty")
	}
	if cfg.GasLimit == 0 {
		panic("gas_limit of teth should be larger than 0")
	}
}

type LogConfig struct {
	Level                        string `json:"level"`
	Filename                     string `json:"filename"`
	MaxFileSizeInMB              int    `json:"max_file_size_in_mb"`
	MaxBackupsOfLogFiles         int    `json:"max_backups_of_log_files"`
	MaxAgeToRetainLogFilesInDays int    `json:"max_age_to_retain_log_files_in_days"`
	UseConsoleLogger             bool   `json:"use_console_logger"`
	UseFileLogger                bool   `json:"use_file_logger"`
	Compress                     bool   `json:"compress"`
}

func (cfg *LogConfig) Validate() {
	if cfg.UseFileLogger {
		if cfg.Filename == "" {
			panic("filename should not be empty if use file logger")
		}
		if cfg.MaxFileSizeInMB <= 0 {
			panic("max_file_size_in_mb should be larger than 0 if use file logger")
		}
		if cfg.MaxBackupsOfLogFiles <= 0 {
			panic("max_backups_off_log_files should be larger than 0 if use file logger")
		}
	}
}

type DBConfig struct {
	Dialect string `json:"dialect"`
	DBPath  string `json:"db_path"`
}

func (cfg *DBConfig) Validate() {
	if cfg.Dialect != DBDialectMysql && cfg.Dialect != DBDialectSqlite3 {
		panic(fmt.Sprintf("only %s and %s supported", DBDialectMysql, DBDialectSqlite3))
	}
}

func (cfg *Config) Validate() {
	cfg.CrossChainConfig.Validate()
	cfg.AdminConfig.Validate()
	cfg.LogConfig.Validate()
	cfg.TopConfig.Validate()
	cfg.TethConfig.Validate()
	cfg.DBConfig.Validate()
}

func ParseConfigFromJson(content string) *Config {
	var config Config
	if err := json.Unmarshal([]byte(content), &config); err != nil {
		panic(err)
	}
	return &config
}

func ParseConfigFromFile(filePath string) *Config {
	bz, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	var config Config
	if err := json.Unmarshal(bz, &config); err != nil {
		panic(err)
	}

	config.Validate()

	return &config
}
