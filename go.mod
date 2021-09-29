module http-theft-bank

go 1.14

replace http-theft-bank => ./

require (
	github.com/dgraph-io/ristretto v0.1.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.7.4
	github.com/muxih4ck/Go-Web-Application-Template v0.0.0-20201003120115-ddad4e8b14dd
	github.com/satori/go.uuid v1.2.0
	github.com/shirou/gopsutil v3.21.8+incompatible
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.9.0
	github.com/teris-io/shortid v0.0.0-20171029131806-771a37caa5cf
	github.com/tklauser/go-sysconf v0.3.9 // indirect
	github.com/willf/pad v0.0.0-20200313202418-172aa767f2a4
	go.uber.org/zap v1.19.1
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)
