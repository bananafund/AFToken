# 使用文档
## 安装geth
git clone https://github.com/bananafund/AFToken <br>
cd AFToken/go-ethereum && make geth <br>
并将geth所在路径加到PATH环境变量中。
## 准备geth运行环境
export GETH_HOME=/Users/work/Documents/geth <br>
mkdir -p ${GETH_HOME}/data <br>
将源码中的genesis.json以及编译生成的geth文件拷贝到GETH_HOME目录。
## 初始化创世快
cd ${GETH_HOME} && geth --datadir ${GETH_HOME}/data init genesis.json
## 启动geth客户端
nohup geth --datadir /home/work/geth/data --syncmode "full" --networkid 1108785740937 --identity "aft1" --rpc --rpcaddr "0.0.0.0" --rpccorsdomain "*" --rpcapi "web3,admin,eth,txpool,miner,personal" -bootnodes "enode://d02e8ed777fbd943bdf314c1e0d7e30d39a1e67af39db00ea4ea0908940647f5606b44071d257248da22c15c4b421fe42ca6700f4f0000163aed5366bd33d1ef@101.132.74.194:30303" &
## 连接本地节点，进入web3 console终端
geth attach http://127.0.0.1:8545
## 新建账号，密码是123456
personal.newAccount("123456")
## 设置挖矿账号为上一步新建的账号
miner.setEtherbase("0x123456")
## 启动挖矿
miner.start()
## 停止挖矿
miner.stop()
