modulemonkey in golang

I created this to document the setup of my golang development environment 
on my mac.

It provides makefiles with some of the commands I found handy.

Refer to the makefile in /prototype (the only bin).

Basic commands:

1. make lint (uses golangci-lint)
2. make build (binary is created in src folder)
3. make install (binary is created in $GOBIN)
4. make deploy (binary is created in $GOBIN/modulemonkey/prototype)
5. make docker (docker image is put into local docker server repo)

It assumes the standard golang local project structure

$HOME/gowork
    /bin
    /pkg
    /src

In my local env, I modified:

.zshrc
    source ~/.bash_profile
.bash_profile
    export GOPATH="$HOME/gowork:$HOME/gowork/src"
    export GOBIN="$HOME/gowork/bin"
    export PATH="$PATH:$GOPATH:$GOPATH/bin"
    source ~/.bashrc
.bashrc
    export PATH="${PATH}:~/usr/local/opt"
    export GO111MODULE=on
    alias gow='cd $HOME/gowork'
    alias gos='cd $HOME/gowork/src/github.com/davidroossien'

go env

GO111MODULE="on"
GOARCH="amd64"
GOBIN="/Users/davidroossien/gowork/bin"
GOCACHE="/Users/davidroossien/Library/Caches/go-build"
GOENV="/Users/davidroossien/Library/Application Support/go/env"
GOEXE=""
GOEXPERIMENT=""
GOFLAGS=""
GOHOSTARCH="amd64"
GOHOSTOS="darwin"
GOINSECURE=""
GOMODCACHE="/Users/davidroossien/gowork/pkg/mod"
GONOPROXY=""
GONOSUMDB=""
GOOS="darwin"
GOPATH="/Users/davidroossien/gowork:/Users/davidroossien/gowork/src"
GOPRIVATE=""
GOPROXY="https://proxy.golang.org,direct"
GOROOT="/usr/local/Cellar/go/1.17.6/libexec"
GOSUMDB="sum.golang.org"
GOTMPDIR=""
GOTOOLDIR="/usr/local/Cellar/go/1.17.6/libexec/pkg/tool/darwin_amd64"
GOVCS=""
GOVERSION="go1.17.6"
GCCGO="gccgo"
AR="ar"
CC="clang"
CXX="clang++"
CGO_ENABLED="1"
GOMOD="/dev/null"
CGO_CFLAGS="-g -O2"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-g -O2"
CGO_FFLAGS="-g -O2"
CGO_LDFLAGS="-g -O2"
PKG_CONFIG="pkg-config"
GOGCCFLAGS="-fPIC -arch x86_64 -m64 -pthread -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fdebug-prefix-map=/var/folders/7q/n28537fs6d14c7bwpb51r8n40000gn/T/go-build1179182478=/tmp/go-build -gno-record-gcc-switches -fno-common"