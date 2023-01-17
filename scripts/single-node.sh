#!/bin/sh

#set -o errexit -o nounset
set -o errexit

CHAINID=$1
GENACCT=$2
HOMEFURY=$3
BINFURY=$4

if [ -z "$1" ]; then
  echo "Need to input chain id..."
  exit 1
fi

if [ -z "$2" ]; then
  echo "Need to input genesis account address..."
  exit 1
fi


if [ -z "$3" ]; then
  echo "Need to input home chain..."
  exit 1
fi

if [ -z "$4" ]; then
  BINFURY="furyd"
fi




# Build genesis file incl account for passed address
coins="10000000000ufusd,100000000000stake"
echo "[OK] inizializzo il genesis"
$BINFURY init --chain-id $CHAINID $CHAINID --home $HOMEFURY
echo "[OK] aggiungo account di test al keyring"
$BINFURY keys add validator --keyring-backend="test"
echo "[OK] aggiungo account di test alla chain"
$BINFURY add-genesis-account $($BINFURY keys show validator -a --keyring-backend="test") $coins --home $HOMEFURY
echo "[OK] aggiungo account di in input"
$BINFURY add-genesis-account $GENACCT $coins --home $HOMEFURY
echo "[OK] aggiungo government"
$BINFURY set-genesis-government-address $GENACCT --home $HOMEFURY
echo "[OK] aggiungo vbr pool"
$BINFURY set-genesis-vbr-pool-amount 1000000000stake --home $HOMEFURY
echo "[OK] aggiungo vbr pool"
$BINFURY set-genesis-vbr-reward-rate 0.01 --home $HOMEFURY
echo "[OK] aggiungo validatore"
$BINFURY gentx validator 5000000000stake --keyring-backend="test" --keyring-dir="~/.fury" --chain-id $CHAINID --home $HOMEFURY
$BINFURY collect-gentxs --home $HOMEFURY

# Set proper defaults and change ports
sed 's#"tcp://127.0.0.1:26657"#"tcp://0.0.0.0:26657"#g' \
   $HOMEFURY/config/config.toml > \
   $HOMEFURY/config/config.toml.tmp
sed 's/timeout_commit = "5s"/timeout_commit = "1s"/g' \
   $HOMEFURY/config/config.toml.tmp > \
   $HOMEFURY/config/config.toml.tmp2
sed 's/timeout_propose = "3s"/timeout_propose = "1s"/g' \
   $HOMEFURY/config/config.toml.tmp2 > \
   $HOMEFURY/config/config.toml.tmp
sed 's/index_all_keys = false/index_all_keys = true/g' \
   $HOMEFURY/config/config.toml.tmp > \
   $HOMEFURY/config/config.toml
# Start the fury
$BINFURY start --pruning=nothing --home $HOMEFURY