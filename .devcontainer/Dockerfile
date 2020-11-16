FROM golang

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
ENV GOSUMDB=sum.golang.google.cn

RUN sed -i 's/deb.debian.org/mirrors.ustc.edu.cn/g' /etc/apt/sources.list \
    && sed -i 's|security.debian.org/debian-security|mirrors.ustc.edu.cn/debian-security|g' /etc/apt/sources.list \
    && sed -i 's/http/https/g' /etc/apt/sources.list \
    # Configure apt, install packages and tools
    && apt-get update \
    && apt-get -y install --no-install-recommends apt-utils dialog 2>&1 \
    #
    # Verify git, process tools, lsb-release (common in install instructions for CLIs) installed
    && apt-get -y install git openssh-client iproute2 procps lsb-release zip locales zsh \
    # Set git EOL always = lf
    && git config --global core.eol lf \
    && git config --global core.autocrlf input \
    && git config --global credential.helper store \
    && git config --global user.name "mengino" \
    && git config --global user.email "mengino@outlook.com" \
    # Install Go tools w/module support
    && mkdir -p /tmp/gotools \
    && cd /tmp/gotools \
    && export GOPATH=/tmp/gotools \
    && export GOCACHE=/tmp/gotools/cache \
    && go get -v \
        golang.org/x/tools/gopls \
        honnef.co/go/tools/... \
        golang.org/x/tools/cmd/gorename \
        golang.org/x/tools/cmd/goimports \
        golang.org/x/tools/cmd/guru \
        golang.org/x/lint/golint \
        github.com/mdempsky/gocode \
        github.com/cweill/gotests/... \
        github.com/haya14busa/goplay/cmd/goplay \
        github.com/sqs/goreturns \
        github.com/josharian/impl \
        github.com/davidrjenni/reftools/cmd/fillstruct \
        github.com/uudashr/gopkgs/v2/cmd/gopkgs \
        github.com/ramya-rao-a/go-outline \
        github.com/acroca/go-symbols \
        github.com/godoctor/godoctor \
        github.com/rogpeppe/godef \
        github.com/zmb3/gogetdoc \
        github.com/fatih/gomodifytags \
        github.com/mgechev/revive \
        github.com/go-delve/delve/cmd/dlv\
    #
    # Install gocode-gomod
    && GO111MODULE=auto go get -x -d github.com/samblerre/gocode 2>&1 \
    && GO111MODULE=auto go build -o gocode-gomod github.com/stamblerre/gocode \
    #
    # Install golangci-lint
    && curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin 2>&1 \
    #
    # Clean up
    && mv /tmp/gotools/bin/* $GOPATH/bin/ \
    && mv gocode-gomod $GOPATH/bin/ \
    && apt-get autoremove -y \
    && apt-get clean -y \
    && rm -rf /var/lib/apt/lists/* /go/src /tmp/gotools
