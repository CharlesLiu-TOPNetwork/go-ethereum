#! /bin/bash

# kill geth process:
ps -ef | grep 'geth' | grep -v 'grep' | awk -F ' ' '{print $2}' | xargs kill -9

sleep 0.5

# clean old datas:
echo "clean datas and re-run"
rm -rf ./data/*

sleep 0.5

# check if geth is in this dir else cp from `../build/bin/geth`:
if [ ! -f "./geth" ];then
    echo "not find geth. try RUN \`cp ../../build/bin/geth .\` first" 
    exit
else
    echo "find geth continue."
fi

# import genesis account by private keys
echo "import default genesis account: "
./geth --datadir data/node0 account import keystores/prikey_0.hex --password passwd
./geth --datadir data/node1 account import keystores/prikey_1.hex --password passwd
./geth --datadir data/node2 account import keystores/prikey_2.hex --password passwd

# import genesis json:
echo "import genesis json: "
./geth --datadir data/node0 init poa_genesis.json
./geth --datadir data/node1 init poa_genesis.json
./geth --datadir data/node2 init poa_genesis.json

# start geth at background:
echo "start geth at background: "
./geth --datadir data/node0 --networkid 616 --port 10010 --nodiscover --allow-insecure-unlock --unlock '0' --password ./passwd --http --http.port 8540 --http.api eth,web3,admin,net,miner,personal --rpc.allow-unprotected-txs --mine --miner.threads 1 --miner.gasprice 100 --syncmode full 2>>data/node0/teth.log &
./geth --datadir data/node1 --networkid 616 --port 10011 --nodiscover --allow-insecure-unlock --unlock '0' --password ./passwd --http --http.port 8541 --http.api eth,web3,admin,net,miner,personal --rpc.allow-unprotected-txs --mine --miner.threads 1 --miner.gasprice 100 --syncmode full 2>>data/node1/teth.log &
./geth --datadir data/node2 --networkid 616 --port 10012 --nodiscover --allow-insecure-unlock --unlock '0' --password ./passwd --http --http.port 8542 --http.api eth,web3,admin,net,miner,personal --rpc.allow-unprotected-txs --mine --miner.threads 1 --miner.gasprice 100 --syncmode full 2>>data/node2/teth.log &

# get enode info:
enode_info0=`./geth attach http://127.0.0.1:8540 --exec "admin.nodeInfo.enode"`
echo "enode 0: $enode_info0"
enode_info1=`./geth attach http://127.0.0.1:8541 --exec "admin.nodeInfo.enode"`
echo "enode 1: $enode_info1"
enode_info2=`./geth attach http://127.0.0.1:8542 --exec "admin.nodeInfo.enode"`
echo "enode 2: $enode_info2"

# node connect each other:
echo "genesis node connect to each other"
connect_result=`./geth attach http://127.0.0.1:8540 --exec "admin.addPeer($enode_info1)"`
connect_result=`./geth attach http://127.0.0.1:8540 --exec "admin.addPeer($enode_info2)"`
connect_result=`./geth attach http://127.0.0.1:8541 --exec "admin.addPeer($enode_info2)"`
# todo check eq = true
# todo make it double loop statement.

# check peers number:
echo "check peers number: "
peers_count0=`./geth attach http://127.0.0.1:8540 --exec "net.peerCount"`
peers_count1=`./geth attach http://127.0.0.1:8541 --exec "net.peerCount"`
peers_count2=`./geth attach http://127.0.0.1:8542 --exec "net.peerCount"`
# todo check eq = 2 

# start poa mining:
# `--mine --miner.thread 1` do the same thing 
# echo "start poa mining:"
# ./geth attach http://127.0.0.1:8540 --exec "miner.start(1)"
# ./geth attach http://127.0.0.1:8541 --exec "miner.start(1)"
# ./geth attach http://127.0.0.1:8542 --exec "miner.start(1)"

sleep 1
# init contracts:
# eth.sendTransaction({from:eth.accounts[0],to:"0000000000000000000000000000000000000100",data:"0xe1c7392a"})
res=`./geth attach http://127.0.0.1:8540 --exec "eth.sendTransaction({from:eth.accounts[0],to:\"0000000000000000000000000000000000000100\",data:\"0xe1c7392a\"})"`