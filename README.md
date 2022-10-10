### Gener

基于xo、proteus等开源项目实现的数据库字段Go模型结构体和Proto原型文件生成器，用于快速开发业务。

### 安装
```
$ go get -u -v github.com/Smilefish2/gener
```

### 使用
```
// 运行环境检查命令检测所有必要条件是否可用
$ gener doctor 
// 运行生成命令
$ gener gen 
```

### 感谢
* [xo](https://github.com/xo/xo) - Command line tool to generate idiomatic Go code for SQL databases
* [proteus](https://github.com/src-d/proteus) - Generate .proto files from Go source code.

### License

MIT, see [LICENSE](/LICENSE)