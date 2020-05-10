package openwtester

import (
	"github.com/assetsadapterstore/pizzachain-adapter/pizzachain"
	"github.com/blocktree/openwallet/v2/log"
	"github.com/blocktree/openwallet/v2/openw"
)

func init() {
	//注册钱包管理工具
	log.Notice("Wallet Manager Load Successfully.")
	cache := pizzachain.NewCacheManager()

	openw.RegAssets(pizzachain.Symbol, pizzachain.NewWalletManager(&cache))
}
