# 使用文档
## 安装aft
git clone https://github.com/bananafund/AFToken <br>
cd AFToken <br>
### MAC/Linux下执行命令：
make all
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
nohup aft --datadir ${AFT_HOME}/data --syncmode "full" --networkid 1108785740937 --identity "aft1" --rpc --rpcaddr "0.0.0.0" --rpccorsdomain "*" --rpcapi "web3,admin,eth,txpool,miner,personal" -bootnodes "enode://4ef2069de70a83956432d8845a9f8e344d0a4504ec57f2b021690e1eb4a854719758a11d00aa0b3f83a7d8f77a0d62a7cf94cec4eb0a9d2a8bfb625a5dcce612@101.132.74.194:30303" &
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
