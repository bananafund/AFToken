package main

import (
	"gopkg.in/urfave/cli.v1"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/core"
	"encoding/json"
	"github.com/ethereum/go-ethereum/log"
)

var genesisBytes = []byte(`{
  "config": {
        "chainId": 1108785740937,
        "homesteadBlock": 0,
        "eip155Block": 0,
        "eip158Block": 0
  },
  "alloc": {
      "0xc3895de2fecf31bcb7f5d5a7c2b23360d9984589": {"balance": "500000000000000000000000000"}
  },
  "coinbase"   : "0xc3895de2fecf31bcb7f5d5a7c2b23360d9984589",
  "difficulty" : "0x20000",
  "extraData"  : "",
  "gasLimit"   : "0x7a1200",
  "nonce"      : "0x0000000000000042",
  "mixhash"    : "0x0000000000000000000000000000000000000000000000000000000000000000",
  "parentHash" : "0x0000000000000000000000000000000000000000000000000000000000000000",
  "timestamp"  : "0x00"
}`)

/**
 * 自定义创世块函数。
 * 参考自initGenesis函数，区别是initCustomGenesis是从字符串json中初始化创世块。
 */
func initCustomGenesis(ctx *cli.Context) {
	genesis := new(core.Genesis)
	json.Unmarshal(genesisBytes, genesis)

	stack := makeFullNode(ctx)
	for _, name := range []string{"chaindata", "lightchaindata"} {
		dbdir := stack.ResolvePath(name)
		if !common.FileExist(dbdir) {
			chaindb, err := stack.OpenDatabase(name, 0, 0)
			if err != nil {
				utils.Fatalf("Failed to open database: %v", err)
			}
			_, hash, err := core.SetupGenesisBlock(chaindb, genesis)
			if err != nil {
				utils.Fatalf("Failed to write genesis block: %v", err)
			}
			log.Info("Successfully wrote genesis state", "database", name, "hash", hash)
			chaindb.Close()
		}
	}
}
