## 根据设备生成唯一ID

```go
go run main.go
```

## Window平台交叉编译

```bash
./build.bat

编译参数为 -ldflags="-s -w"
```

![image.png](assets/image.png)

## 编译后执行文件压缩体积

```bash
upx -9 deviceId_windows_amd64.exe
```

## 各平台可执行文件

```plaintext
output
 ├── deviceId_darwin_amd64
 ├── deviceId_linux_amd64
 └── deviceId_windows_amd64.exe
