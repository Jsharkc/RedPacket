package utils

import (
	"math/rand"
	"time"

	"github.com/Jsharkc/RedPacket/general"
)

// GenePwd - generate password
func GenePwd() string {
	var target = make([]byte, general.TargetPWDNum)
	rand.Seed(time.Now().UnixNano())

	for index := 0; index < general.TargetPWDNum; index++ {
		target[index] = general.LetterSource[rand.Int31n(62)]
	}

	return string(target)
}
