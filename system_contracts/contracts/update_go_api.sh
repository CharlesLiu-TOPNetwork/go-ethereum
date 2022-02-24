#! /bin/bash

if [ ! -f "../../build/bin/abigen" ];then
    echo "not find abigen. try RUN \`make all\` first" 
    exit
else
    echo "find abigen continue."
fi

mkdir -p go-api
mkdir -p go-api/tokenexchange
mkdir -p go-api/crosschain

../../build/bin/abigen --sol TokenExchange.sol --pkg tokenexchange --out go-api/tokenexchange/TokenExchange.go
../../build/bin/abigen --sol CrossChain.sol --pkg crosschain --out go-api/crosschain/CrossChain.go