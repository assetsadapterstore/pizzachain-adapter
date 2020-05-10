module github.com/assetsadapterstore/pizzachain-adapter

go 1.12

require (
	github.com/astaxie/beego v1.12.0
	github.com/blocktree/go-owcdrivers v1.2.0
	github.com/blocktree/go-owcrypt v1.1.1
	github.com/blocktree/openwallet/v2 v2.0.4
	github.com/eoscanada/eos-go v0.8.16
	github.com/imroc/req v0.2.4
	github.com/shopspring/decimal v0.0.0-20200105231215-408a2507e114
)

replace github.com/eoscanada/eos-go => github.com/blocktree/eos-go v0.8.13-blocktree
