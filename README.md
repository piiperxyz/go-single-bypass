# go-single-bypass

一个简单的免杀，主要参考go-mimikatz，使其技术能够运用到CS的shellcode上。另外添加了upx、sgn和garble进行混淆和压缩

## 免责声明

该工具仅用于安全研究，禁止使用工具发起非法攻击等违法行为，造成的后果使用者负责

## 如何使用

将CS stagerless生成的beacon.bin放到packer文件夹下，在主目录输入：

```go
go generate
```

主目录生成的cold.exe即为免杀马

注意:

* 需下载好go环境，目前测试是在windows+64位环境下可正常生成
* 生成的中间文件可能不免杀，在生成的过程中需要将杀软关闭！
* upx.exe和sgn.exe是github上下载的，若不放心可自行替换即可
* 如果下载依赖过慢配置镜像`go env -w GOPROXY=https://goproxy.cn,direct`，如果是第一次安装go建议先执行该命令

## 参考

https://github.com/vyrus001/go-mimikatz

https://github.com/burrowers/garble

https://github.com/upx/upx

https://github.com/EgeBalci/sgn
