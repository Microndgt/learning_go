MAC中安装Delve调试器
---

`brew install go-delve/delve/delve`运行这个命令进行安装，但是可能会出现以下错误

```
Last 15 lines from /Users/kevin/Library/Logs/Homebrew/delve/02.sudo:
2017-10-17 12:04:16 +0800

sudo
security
add-trusted-cert
-d
-r
trustRoot
-k
/Library/Keychains/System.keychain
dlv-cert.cer


If reporting this issue please do so at (not Homebrew/brew or Homebrew/core):
https://github.com/go-delve/homebrew-delve/issues
```

这时候进入`cd /Users/kevin/Library/Caches/Homebrew/`，将缓存的`delve-1.0.0-rc.1.tar.gz`解压`tar -xzvf delve-1.0.0-rc.1.tar.gz`，之后进入`cd delve-1.0.0-rc.1/scripts`，执行`./gencert.sh`，输入密码即可将证书生成。

然后运行`brew install go-delve/delve/delve`即可完成安装。执行`dlv version`可看到版本号即为成功。

然后在VSCode中在`main.go`中运行F5开始调试，这时候会报`go: GOPATH entry is relative; must be absolute path: "undefined".`，这是因为没有设置`GOPATH`

在`.bash_profile`文件中加入:

```
export GOPATH=/Users/kevin/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN
```

然后`source .bash_profile`即可，可以使用`go env`来查看环境变量

再点击F5进行调试，会出现下面的错误：

```
2017/10/17 12:24:19 server.go:73: Using API v1
2017/10/17 12:24:19 debugger.go:97: launching process with args: [/Users/kevin/go/src/learning_go/debug]
could not launch process: exec: "lldb-server": executable file not found in $PATH
Process exiting with code: 1
```

这个问题在使用0.12.2或者以上的delve的Mac用户上会出现，不知道为什么，但是执行`xcode-select --install`就可以解决问题。

到此就可以正常调试啦！

运行
---

`go run main.go`