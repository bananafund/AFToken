# 使用文档
## ubuntu下安装go 1.10环境
adduser work <br>
cd /home/work <br>
wget https://dl.google.com/go/go1.10.3.linux-amd64.tar.gz <br>
tar -zxvf go1.10.3.linux-amd64.tar.gz <br>
mv go goroot <br>
mkdir -p /home/work/go <br>
打开~/.bashrc，在结尾添加以下内容：<br>
export GOROOT=/home/work/goroot <br>
export PATH=$GOROOT/bin:$PATH <br>
export GOPATH=/home/work/go <br>
export PATH=/home/work/go/src/github.com/bananafund/AFToken/build/bin:$PATH <br>
source ~/.bashrc <br>

## ubuntu下安装aft
cd /home/work/go <br>
go get -u github.com/bananafund/AFToken <br>
cd /home/work/go/src/github.com/bananafund/AFToken <br>
make all <br>
打开~/.bashrc，在结尾添加以下内容：<br>
export PATH=/home/work/go/src/github.com/bananafund/AFToken/build/bin:$PATH <br>
source ~/.bashrc <br>

### Windows下执行命令：
go install -v ./cmd/...
### 将aft所在路径加到PATH环境变量中。
## 准备aft运行环境
### Mac/Linux下执行如下命令：
export AFT_HOME=/Users/work/Documents/aft <br>
mkdir -p ${AFT_HOME}/data <br>
### Windows下新建aft目录作为AFT_HOME，并在aft目录下新建data目录
将源码中的genesis.json以及编译生成的aft文件拷贝到AFT_HOME目录。
## 初始化创世快
cd ${AFT_HOME} && aft --datadir ${aft_HOME}/data init genesis.json
## 启动aft客户端
nohup aft --datadir ${AFT_HOME}/data --syncmode "full" --networkid 1108785740937 --identity "aft1" --rpc --rpcaddr "0.0.0.0" --rpccorsdomain "*" --rpcapi "web3,admin,eth,txpool,miner,personal" -bootnodes "enode://184276db387c64eb2bfec864ac3762cafad5d4187a0c1c7d1cfc1639fdd9926c20c7c999a6131002902f4ee6b6b48018ee9095d99de401350de21bf7ad00822d@47.101.36.51:30303" &
## 连接本地节点，进入web3 console终端
aft attach http://127.0.0.1:8545
## 新建账号，密码是123456
personal.newAccount("123456")
## 设置挖矿账号为上一步新建的账号
miner.setEtherbase("0x123456")
## 启动挖矿
miner.start()
## 停止挖矿
miner.stop()
