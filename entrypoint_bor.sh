#!/bin/bash
set -e 

echo $(which bor)
bor --heimdall $HEIMDALL_URL --datadir $DATA_DIR init $NODE_DIR/genesis.json 

bor --datadir $DATA_DIR \\
  --port 30303 \\
  --heimdall $HEIMDALL_URL \\
  --rpc --rpcaddr '0.0.0.0' \\
  --ws --wsport 8546 \\
  --rpcvhosts '*' \\
  --rpccorsdomain * \\
  --rpcport 8545 \\
  --ipcpath $DATA_DIR/bor.ipc \\
  --rpcapi 'personal,db,eth,net,web3,txpool,miner,admin,bor' \\
  --syncmode full \\
  --networkid 15001 \\
  --miner.gaslimit 2000000000 \\
  --txpool.nolocals \\
  --txpool.accountslots 128 \\
  --txpool.globalslots 20000 \\
  --txpool.lifetime 0h16m0s \\
  --unlock $ADDRESS \\
  --keystore $NODE_DIR/keystore \\
  --password $NODE_DIR/password.txt \\
  --allow-insecure-unlock  --mine
