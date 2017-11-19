package model

import (
	"time"

	"github.com/Jsharkc/RedPacket/general"
	"github.com/Jsharkc/RedPacket/orm"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
)

type timerWorkServiceProvider struct{}

var (
	// TimerWorkService handles operations
	TimerWorkService = &timerWorkServiceProvider{}

	TimerWorkDispatchChan  chan *RedPack
	TimerWorkChan          []chan *RedPack
	timerWorkCount         = 0
	timerWorkDispatchIndex = 0
)

// Init - initial timer work
func (ts *timerWorkServiceProvider) Init(dispatchCache, workCache, number int) {
	var (
		mRedPackList []RedPack
		err          error
	)
	TimerWorkDispatchChan = make(chan *RedPack, dispatchCache)
	TimerWorkChan = make([]chan *RedPack, number)
	timerWorkCount = number

	for index := 0; index < number; index++ {
		TimerWorkChan[index] = make(chan *RedPack, workCache)
		startWorker(index)
	}

	startDispatch()

	mRedPackList, err = getActiveRedPack()
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			panic(err)
		}
	}

	for index, _ := range mRedPackList {
		TimerWorkDispatchChan <- &mRedPackList[index]
	}
}

// startDispatch - start dispatch
func startDispatch() {
	go func() {
		for redPack := range TimerWorkDispatchChan {
			TimerWorkChan[timerWorkDispatchIndex] <- redPack
			timerWorkDispatchIndex = (timerWorkDispatchIndex + 1) % timerWorkCount
		}
	}()
}

// startWorker - start timer work
func startWorker(chanIndex int) {
	go func() {
		var (
			mRedPack RedPack
			mTimer   = time.NewTimer(time.Hour)
			mDur     time.Duration
			err      error
		)
	outer:
		for redPack := range TimerWorkChan[chanIndex] {
			tmpIndex := 0

			mDur = redPack.Created.Add(time.Hour * general.RPActiveTimeInHour).Sub(time.Now())

			if mDur > 0 {
				mTimer.Reset(mDur)
				<-mTimer.C
			}

			for ; tmpIndex < general.TimerWorkErrorTryTime; tmpIndex++ {
				tx := orm.DBConn.Begin()
				err = tx.Raw("SELECT * FROM redpack WHERE id = ? AND status = ? FOR UPDATE", redPack.ID, general.RPGrab).Scan(&mRedPack).Error

				if err != nil {
					tx.Rollback()
					log.Error(err)

					mTimer.Reset(time.Duration(tmpIndex*general.TimerWorkErrorTryInterval) * time.Millisecond)
					<-mTimer.C

					continue
				}

				switch {
				case mRedPack.Number > 0:
					err = tx.Model(mRedPack).Update(general.TBStatus, general.RPRefund).Error
				case mRedPack.Number == 0:
					err = tx.Model(mRedPack).Update(general.TBStatus, general.RPFinish).Error
				}

				if err != nil {
					tx.Rollback()
					log.Error(err)

					mTimer.Reset(time.Duration(tmpIndex*general.TimerWorkErrorTryInterval) * time.Millisecond)
					<-mTimer.C

					continue
				}

				tx.Commit()
				continue outer
			}
		}
	}()
}
