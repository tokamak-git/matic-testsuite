#!/usr/bin/env sh

set -x #echo on

NODE_DIR=/root/.bor
DATA_DIR=/root/.bor/data

ADDRESSES=(
  "0x4551942266A683a3Bd875Fe716798D3E2eB6B122"
  "0x0D98C85A9DA99651f48964Ac71ABB7c74856b8c1"
  "0xf522FE4B332D896B9dCa6A83CA4A40067bbEBb81"
  "0x437d2eA2De07cDdA3f6D72EfB0CCcB4b7538951C"
  "0x3B9DaBCa84Db19B3F77baEB5c1696Ac0d695bE8d"
)

INDEX=$1;
ADDRESS=${ADDRESSES[$INDEX]};

NODE_DIR=/root/.bor
DATA_DIR=/root/.bor/data

docker-compose run -d --name bor$INDEX bor$INDEX sh -c "
touch /root/logs/bor.log
bor --datadir $DATA_DIR \
  --port 30303 \
  --heimdall http://heimdall$INDEX:1317 \
  --rpc --rpcaddr '0.0.0.0' \
  --ws --wsport 8546 \
  --rpcvhosts '*' \
  --rpccorsdomain '*' \
  --rpcport 8545 \
  --ipcpath $DATA_DIR/bor.ipc \
  --rpcapi 'personal,db,eth,net,web3,txpool,miner,admin,bor' \
  --syncmode 'full' \
  --networkid '15001' \
  --miner.gaslimit '2000000000' \
  --txpool.nolocals \
  --txpool.accountslots '128' \
  --txpool.globalslots '20000' \
  --txpool.lifetime '0h16m0s' \
  --unlock $ADDRESS \
  --keystore $NODE_DIR/keystore \
  --password $NODE_DIR/password.txt \
  --allow-insecure-unlock \
  --mine > /root/logs/bor.log 2>&1 &
tail -f /root/logs/bor.log
"

