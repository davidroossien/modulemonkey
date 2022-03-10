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

$HOME/gowork<br>
    /bin<br>
    /pkg<br>
    /src<br>

In my local env, I modified:

.zshrc<br>
    source ~/.bash_profile<br>
    
.bash_profile<br>
    export GOPATH="$HOME/gowork:$HOME/gowork/src"<br>
    export GOBIN="$HOME/gowork/bin"<br>
    export PATH="$PATH:$GOPATH:$GOPATH/bin"<br>
    source ~/.bashrc<br>
    
.bashrc<br>
    export PATH="${PATH}:~/usr/local/opt"<br>
    export GO111MODULE=on<br>
    alias gow='cd $HOME/gowork'<br>
    alias gos='cd $HOME/gowork/src/github.com/davidroossien'<br>

go env

GO111MODULE="on"<br>
GOARCH="amd64"<br>
GOBIN="/Users/davidroossien/gowork/bin"<br>
GOCACHE="/Users/davidroossien/Library/Caches/go-build"<br>
GOENV="/Users/davidroossien/Library/Application Support/go/env"<br>
GOEXE=""<br>
GOEXPERIMENT=""<br>
GOFLAGS=""<br>
GOHOSTARCH="amd64"<br>
GOHOSTOS="darwin"<br>
GOINSECURE=""<br>
GOMODCACHE="/Users/davidroossien/gowork/pkg/mod"<br>
GONOPROXY=""<br>
GONOSUMDB=""<br>
GOOS="darwin"<br>
GOPATH="/Users/davidroossien/gowork:/Users/davidroossien/gowork/src"<br>
GOPRIVATE=""<br>
GOPROXY="https://proxy.golang.org,direct"<br>
GOROOT="/usr/local/Cellar/go/1.17.6/libexec"<br>
GOSUMDB="sum.golang.org"<br>
GOTMPDIR=""<br>
GOTOOLDIR="/usr/local/Cellar/go/1.17.6/libexec/pkg/tool/darwin_amd64"<br>
GOVCS=""<br>
GOVERSION="go1.17.6"<br>
GCCGO="gccgo"<br>
AR="ar"<br>
CC="clang"<br>
CXX="clang++"<br>
CGO_ENABLED="1"<br>
GOMOD="/dev/null"<br>
CGO_CFLAGS="-g -O2"<br>
CGO_CPPFLAGS=""<br>
CGO_CXXFLAGS="-g -O2"<br>
CGO_FFLAGS="-g -O2"<br>
CGO_LDFLAGS="-g -O2"<br>
PKG_CONFIG="pkg-config"<br>
GOGCCFLAGS="-fPIC -arch x86_64 -m64 -pthread -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fdebug-prefix-map=/var/folders/7q/n28537fs6d14c7bwpb51r8n40000gn/T/go-build1179182478=/tmp/go-build -gno-record-gcc-switches -fno-common"<br>
