/*
 * Copyright 2018 The openwallet Authors
 * This file is part of the openwallet library.
 *
 * The openwallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The openwallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */

package openwtester

import (
	"github.com/astaxie/beego/config"
	"github.com/blocktree/openwallet/v2/openw"
	"path/filepath"
	"testing"

	"github.com/blocktree/openwallet/v2/log"
	"github.com/blocktree/openwallet/v2/openwallet"
)

func TestWalletManager_GetTransactions(t *testing.T) {
	tm := testInitWalletManager()
	list, err := tm.GetTransactions(testApp, 0, -1, "", false)
	if err != nil {
		log.Error("GetTransactions failed, unexpected error:", err)
		return
	}
	for i, tx := range list {
		log.Info("trx[", i, "] :", tx)
	}
	log.Info("trx count:", len(list))
}

func TestWalletManager_GetTransactionByWxID(t *testing.T) {
	tm := testInitWalletManager()
	wxID := openwallet.GenTransactionWxID(&openwallet.Transaction{
		TxID: "4aabaedba12594e869b99916dca8619132a96b7ea00a90f497f57d52c2c2fa68",
		Coin: openwallet.Coin{
			Symbol:     "PIZ",
			IsContract: false,
			ContractID: "",
		},
	})
	log.Info("wxID:", wxID)
	tx, err := tm.GetTransactionByWxID(testApp, wxID)
	if err != nil {
		log.Error("GetTransactionByTxID failed, unexpected error:", err)
		return
	}
	log.Info("tx:", tx)
}

func TestWalletManager_GetAssetsAccountBalance(t *testing.T) {
	tm := testInitWalletManager()
	walletID := "W2kgcfhPLc6LQeLHVK8CtztC9ab6bsuF2p"
	accountID := "CokfTqE8bEYPjPqENsjhu5hc1cTdyUkjWRt2TvRTxy2t"

	balance, err := tm.GetAssetsAccountBalance(testApp, walletID, accountID)
	if err != nil {
		log.Error("GetAssetsAccountBalance failed, unexpected error:", err)
		return
	}
	log.Info("balance:", balance)
}

func TestWalletManager_GetAssetsAccountTokenBalance(t *testing.T) {
	tm := testInitWalletManager()
	walletID := "W2kgcfhPLc6LQeLHVK8CtztC9ab6bsuF2p"
	accountID := "CokfTqE8bEYPjPqENsjhu5hc1cTdyUkjWRt2TvRTxy2t"
	//accountID := "EURz2Yj9nBk1ZJ9NAZii4sTWXSmMEEUVeodBMxQaxmdg"
	contract := openwallet.SmartContract{
		Address:  "token.pex:PEX",
		Protocol: "multiple-token",
		Symbol:   "PEX",
		Name:     "PEX",
		Decimals: 4,
	}

	balance, err := tm.GetAssetsAccountTokenBalance(testApp, walletID, accountID, contract)
	if err != nil {
		log.Error("GetAssetsAccountTokenBalance failed, unexpected error:", err)
		return
	}
	log.Info("balance:", balance.Balance)
}

func TestWalletManager_GetEstimateFeeRate(t *testing.T) {
	tm := testInitWalletManager()
	coin := openwallet.Coin{
		Symbol: "PIZ",
	}
	feeRate, unit, err := tm.GetEstimateFeeRate(coin)
	if err != nil {
		log.Error("GetEstimateFeeRate failed, unexpected error:", err)
		return
	}
	log.Std.Info("feeRate: %s %s/%s", feeRate, coin.Symbol, unit)
}

func TestGetAccountTokenBalance(t *testing.T) {
	symbol := "PIZ"
	assetsMgr, err := openw.GetAssetsAdapter(symbol)
	if err != nil {
		log.Error(symbol, "is not support")
		return
	}
	//读取配置
	absFile := filepath.Join(configFilePath, symbol+".ini")

	c, err := config.NewConfig("ini", absFile)
	if err != nil {
		return
	}
	assetsMgr.LoadAssetsConfig(c)
	sm := assetsMgr.GetSmartContractDecoder()

	contract := openwallet.SmartContract{
		Address:  "tonydchan123:ZING",
		Protocol: "",
		Symbol:   "PIZ",
		Name:     "ZING",
		Decimals: 8,
	}

	contractID := openwallet.GenContractID(contract.Symbol, contract.Address)
	log.Infof("contractID = %s", contractID)
	log.Infof("BalanceModelType = %v", assetsMgr.BalanceModelType())

	addrs := []string{
		"pizjiahua222",
	}

	balances, err := sm.GetTokenBalanceByAddress(contract, addrs...)
	if err != nil {
		log.Errorf(err.Error())
		return
	}
	for _, b := range balances {
		log.Infof("balance[%s] = %s", b.Balance.Address, b.Balance.Balance)
		log.Infof("UnconfirmBalance[%s] = %s", b.Balance.Address, b.Balance.UnconfirmBalance)
		log.Infof("ConfirmBalance[%s] = %s", b.Balance.Address, b.Balance.ConfirmBalance)
	}
}

func TestGetAddressVerify(t *testing.T) {
	symbol := "PIZ"
	assetsMgr, err := openw.GetAssetsAdapter(symbol)
	if err != nil {
		log.Error(symbol, "is not support")
		return
	}
	//读取配置
	absFile := filepath.Join(configFilePath, symbol+".ini")

	c, err := config.NewConfig("ini", absFile)
	if err != nil {
		return
	}
	assetsMgr.LoadAssetsConfig(c)
	addrDec := assetsMgr.GetAddressDecoderV2()

	flag := addrDec.AddressVerify("pizopenw2222")
	log.Infof("flag: %v", flag)

}
