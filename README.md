## BTCU 毕业设计
实现版权上链(sole,leveldb)
### 进度
- [x] 搭建网络(wz)
- [x] 编写链码(wz)
- [x] 完成阶段性测试(wz)
- [ ] 编写服务端
- [ ] 服务端测试
- [ ] 编写前端
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
