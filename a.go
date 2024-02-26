package a

import (
	"github.com/ethereum/go-ethereum/common"
)

func Addr() string {
	return common.HexToAddress("0x0").String()
}
