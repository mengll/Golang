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

go mod edit -replace=golang.org/x/mod=github.com/golang/mod@latest
go mod edit -replace=gopkg.in/telegram-bot-api.v4=github.com/go-telegram-bot-api/telegram-bot-api@v4.6.4
go mod edit -replace=gopkg.in/src-d/go-git-fixtures.v3@v3.5.0=github.com/src-d/go-git-fixtures@latest
