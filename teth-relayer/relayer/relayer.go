package relayer

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/telosprotocol/teth-relayer/common"
	config "github.com/telosprotocol/teth-relayer/config"
	"github.com/telosprotocol/teth-relayer/executor"
)

type Relayer struct {
	db           *gorm.DB
	cfg          *config.Config
	topExecutor  *executor.TopExecutor
	tethExecutor *executor.TethExecutor
}

func NewRelayer(db *gorm.DB, cfg *config.Config, topExecutor *executor.TopExecutor, tethExecutor *executor.TethExecutor) *Relayer {
	return &Relayer{
		db:           db,
		cfg:          cfg,
		topExecutor:  topExecutor,
		tethExecutor: tethExecutor,
	}
}

func (r *Relayer) Start(startHeight uint64) {
	go r.relayerDaemon(startHeight)
}

func (r *Relayer) relayerDaemon(startHeight uint64) {
	// var tashSet *common.TaskSet
	height := startHeight
	common.Logger.Info("Start relayer daemon")
	for {
		// bytes, _ := r.topExecutor.GetBlockByHeight("Ta0000@0", height)
		latest_height, err := r.topExecutor.GetLatestBlockHeight("Ta0000@0")
		if err != nil {
			common.Logger.Error(err.Error())
			return
		}

		if err != nil {
			common.Logger.Error(err.Error())
			time.Sleep(time.Duration(5000 * int64(time.Millisecond)))
			continue
		}
		if latest_height < height {
			common.Logger.Info(fmt.Sprintf("Height %d not update\n", height))
			time.Sleep(time.Duration(5000 * int64(time.Millisecond)))
			continue
		}

		common.Logger.Info(fmt.Sprintf("Relaying %d block...\n", height))
		_, err = r.tethExecutor.RelayCrossChainPackage(0x1, 0, height)
		if err != nil {
			common.Logger.Error(err.Error())
			time.Sleep(time.Duration(5000 * int64(time.Millisecond)))
			continue
		}
		common.Logger.Info(fmt.Sprintf("Relay height %d block success\n", height))

		height++
		time.Sleep(time.Duration(5000 * int64(time.Millisecond)))
		continue
	}
}
