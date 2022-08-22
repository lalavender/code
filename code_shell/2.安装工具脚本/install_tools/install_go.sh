#!/bin/bash
GO_FILE=go1.17.5.linux-amd64.tar.gz
[ -f %${GO_FILE} ] $$ wget https://studygolang.com/dl/golang/${GO_FILE} || echo "${GO_FILE} is not  exist"
rm -rf /usr/local/go
pwd
tar -C /usr/local -xzf ${GO_FILE}
echo "export PATH=$PATH:/usr/local/go/bin" >>$HOME/.bashrc
source $HOME/.bashrc
go version