#!/bin/bash

GETH_HOME=$1
GENESIS_FILE=${GETH_HOME}/genesis.json
GETH_DATA_PATH=${GETH_HOME}/data

if [ ! -d "${GETH_DATA_PATH}/geth/chaindata" ]
then
    echo "init genesis"
    geth --datadir ${GETH_DATA_PATH} init $GENESIS_FILE
fi

nohup geth --datadir ${GETH_DATA_PATH} --syncmode "full" --networkid 1108785740937 --identity "aft1" --rpc --rpcaddr "0.0.0.0" --rpccorsdomain "*" --rpcapi "web3,admin,eth,txpool,miner,personal" 1>> geth.log 2>&1 &
