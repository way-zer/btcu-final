## BTCU 毕业设计
实现版权上链(sole,leveldb)
### 进度
- [x] 搭建网络(wz)
- [x] 编写链码(wz)
- [x] 完成阶段性测试(wz)
- [x] 编写服务端clientSDK并完成测试(wz)
- [x] 编写服务端(wxm)
- [x] 编写前端(wxm)
- [ ] 整体测试
### 使用方法(链码部分)
环境版本1.4.3  
需手动安装相关docker镜像  
./bin下需要3个可执行文件 configtxgen configtxlator cryptogen  
启动方式与byfn相似  
```shell
sh manage.sh generate #(可选)
sh manage.sh up 
sh manage.sh down 
```
### 项目结构介绍
- 根目录为fabric网络相关问题
- ./chaincode 为链码源码
- ./client 为服务端部分(内含前端)
- ./clientSDK/ccSDK.go 为链码clientSDK
- ./clientSDK/sdkConfig.yaml 为链码clientSDK配置文件
- ./clientSDK/test 为链码clientSDK测试程序
- ./其他文件及文件夹均与网络有关
### 链码函数介绍
- register 版权登记 参数 Hash Name Author Press Data || 防覆写 写入Hash->Data (Name,Author,Press)->Hash两条记录
- query 查询 参数 Hash 返回 Data || 查询Hash->Data
- queryHash 查询 参数 Name Author Press 返回 Hash || 查询(Name,Author,Press)->Hash
### 启动方法(完整)
```shell
 #工作目录为项目根目录,即当前README所在目录
 #./bin下需要3个可执行文件 configtxgen configtxlator cryptogen (设计fabric版本1.4.3)
 sh manage.sh up #启动网络
 go mod verify #安装golang项目依赖
 #配置 ./server/conf/app.conf 中数据库信息(默认为root:123456@127.0.0.1:3306/copyright)
 cd server
 go run . #启动服务端(前端为8080端口)
```
