# go-single-bypass

一个简单的免杀，主要参考go-mimikatz，使其技术能够运用到CS的shellcode上。另外添加了upx、sgn和garble进行混淆和压缩

## 免责声明

该工具仅用于安全研究，禁止使用工具发起非法攻击等违法行为，造成的后果使用者负责

## 如何使用

将CS stagerless生成的beacon.bin放到packer文件夹下，在主目录输入：

```go
go generate
```

主目录生成的cold.exe即为免杀马，需添加启动参数，启动参数为小时分钟日期，若运行木马的时间为12:21，4月15日，则启动密码为122115

注意:

* 需下载好go环境，目前测试是在windows+64位环境下可正常生成
* 生成的中间文件可能不免杀，在生成的过程中需要将杀软关闭！
* upx.exe和sgn.exe是github上下载的，若不放心可自行替换即可
* 如果下载依赖过慢配置镜像`go env -w GOPROXY=https://goproxy.cn,direct`，如果是第一次安装go建议先执行该命令

## 更新

2022/04/21 添加启动参数对抗沙箱，添加测试截图

2022/04/20 发布

## 参考

https://github.com/vyrus001/go-mimikatz

https://github.com/burrowers/garble

https://github.com/upx/upx

https://github.com/EgeBalci/sgn

## 测试效果：

测试时间：2022/04/21

![image-20220421152204905](README.assets\image-20220421152204905.png)

360效果时好时坏，只能说不要一个马用很久吧

![image-20220421153149751](README.assets\image-20220421153149751.png)

![image-20220421153247347](README.assets\image-20220421153247347.png)

火绒

![image-20220421154009202](README.assets\image-20220421154009202.png)

windows defender

![image-20220421154326483](README.assets\image-20220421154326483.png)

卡巴免费版

![image-20220421154802358](README.assets\image-20220421154802358.png)

