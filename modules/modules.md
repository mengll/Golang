replace golang.org/x/text => github.com/golang/text latest

replace golang.org/x/net => github.com/golang/net latest

replace golang.org/x/crypto => github.com/golang/crypto latest

replace golang.org/x/tools => github.com/golang/tools latest

replace golang.org/x/sync => github.com/golang/sync latest

replace golang.org/x/sys => github.com/golang/sys latest

replace cloud.google.com/go => github.com/googleapis/google-cloud-go latest

replace google.golang.org/genproto => github.com/google/go-genproto latest

replace golang.org/x/exp => github.com/golang/exp latest

replace golang.org/x/time => github.com/golang/time latest

replace golang.org/x/oauth2 => github.com/golang/oauth2 latest

replace golang.org/x/lint => github.com/golang/lint latest

replace google.golang.org/grpc => github.com/grpc/grpc-go latest

replace google.golang.org/api => github.com/googleapis/google-api-go-client latest

replace google.golang.org/appengine => github.com/golang/appengine latest

replace golang.org/x/mobile => github.com/golang/mobile latest

replace golang.org/x/image => github.com/golang/image latest





###go代理配置地址
  export GOPROXY=https://goproxy.cn
###gomod的开启的方式
  GO111MODULE=off禁用模块支持，编译时会从GOPATH和vendor文件夹中查找包。
  GO111MODULE=on启用模块支持，编译时会忽略GOPATH和vendor文件夹，只根据 go.mod下载依赖。
  GO111MODULE=auto，当项目在$GOPATH/src外且项目根目录有go.mod文件时，开启模块支持。

### gomod 常用命令

  go mod download    下载依赖的module到本地cache（默认为$GOPATH/pkg/mod目录）
  go mod edit        编辑go.mod文件
  go mod graph       打印模块依赖图
  go mod init        初始化当前文件夹, 创建go.mod文件
  go mod tidy        增加缺少的module，删除无用的module
  go mod vendor      将依赖复制到vendor下
  go mod verify      校验依赖
  go mod why         解释为什么需要依赖
  go mod edit -fmt   格式化文件
  go list -m all     显示依赖关系
  go list -m -json all   显示详细的依赖关系

### 依赖的添加
  go mod edit -require=golang.org/x/text
### 移除文件依赖
  go mod edit -droprequire=golang.org/x/text

###替换依赖
go mod edit -replace golang.org/x/text=github.com/golang/text@latest
go mod edit -replace=golang.org/x/mod=github.com/golang/mod@latest
go mod edit -replace=gopkg.in/telegram-bot-api.v4=github.com/go-telegram-bot-api/telegram-bot-api@v4.6.4
go mod edit -replace=gopkg.in/src-d/go-git-fixtures.v3@v3.5.0=github.com/src-d/go-git-fixtures@latest





