# Installing a full node
After you've setup your hardware following the [hardware requirements](hardware-requirements.md) you are now ready to
become a Commercio.network full node. 

:::warning  
Make sure you have read the [hardware requirements](hardware-requirements.md) before starting  
:::

## 1. Installing the software requirements
In order to update the OS so that you can work properly, execute the following commands:

```bash
apt update && apt upgrade -y
apt install -y git gcc make unzip
snap install --classic go

export NODENAME="<your-moniker>"

echo 'export GOPATH="$HOME/go"' >> ~/.profile
echo 'export PATH="$GOPATH/bin:$PATH"' >> ~/.profile
echo 'export PATH="$PATH:/snap/bin"' >> ~/.profile
echo "export NODENAME=\"$NODENAME\"" >> ~/.profile

source ~/.profile
```

## 2. Chain selection
Before installing the node, please select which chain you would like to connect to (for example **testent7000**)

```bash
rm -rf commercio-chains
mkdir commercio-chains && cd commercio-chains
git clone https://github.com/commercionetwork/chains.git .
cd commercio-<chain-version>
```

:::tip
Always remember to pick the latest chain version listed inside [chains repo](https://github.com/commercionetwork/chains) 
::: 

## 3. Install binaries, genesis file and setup configuration

Compile binaries 

```bash
pkill fyd
pkill fycli
git init . 
git remote add origin https://github.com/tessornetwork/fury.git
git pull
git checkout tags/$(cat .data | grep -oP 'Release\s+\K\S+')
go mod verify
make install
```

Test if you have the correct binaries version:

```bash
fyd version
# Should output the same version written inside the .data file.
# cat .data | grep -oP 'Release\s+\K\S+'
```

Setup the validator node name. We will use the same name for node as well as the wallet key:

```bash
export CHAINID=commercio-$(cat .data | grep -oP 'Name\s+\K\S+')
cat <<EOF >> ~/.profile
export CHAINID="$CHAINID"
EOF
```

Init the `.fyd` folder with the basic configuration

:::warning  
At this point there may be some differences if you are using `KMS` with `HSM`. Specifications will be published shortly.
:::

```bash
fyd unsafe-reset-all
# If you get a error because .fyd folder is not present don't worry 

fyd init $NODENAME
# If you get a error because .fyd folder is present don't worry 
```

Install `genesis.json` file

```bash
pkill fyd
rm -rf ~/.fyd/config/genesis.json
cp genesis.json ~/.fyd/config
```

Change the persistent peers inside `config.toml` file

```bash
sed -e "s|persistent_peers = \".*\"|persistent_peers = \"$(cat .data | grep -oP 'Persistent peers\s+\K\S+')\"|g" ~/.fyd/config/config.toml > ~/.fyd/config/config.toml.tmp
mv ~/.fyd/config/config.toml.tmp  ~/.fyd/config/config.toml
```

Change the seeds inside the `config.toml` file
```bash
sed -e "s|seeds = \".*\"|seeds = \"$(cat .data | grep -oP 'Seeds\s+\K\S+')\"|g" ~/.fyd/config/config.toml > ~/.fyd/config/config.toml.tmp
mv ~/.fyd/config/config.toml.tmp  ~/.fyd/config/config.toml
```

## 4. Configure the service

```bash
tee /etc/systemd/system/fyd.service > /dev/null <<EOF  
[Unit]
Description=Commercio Node
After=network-online.target

[Service]
User=root
ExecStart=/root/go/bin/fyd start
Restart=always
RestartSec=3
LimitNOFILE=4096

[Install]
WantedBy=multi-user.target
EOF
```

**Optional**. You can quick sync with the follow procedure:
```bash
wget "https://quicksync.commercio.network/$CHAINID.latest.tgz" -P ~/.fyd/
# Check if the checksum matches the one present inside https://quicksync.commercio.network
cd ~/.fyd/
tar -zxf $(echo $CHAINID).latest.tgz
```


Now you can start you full node. Enable the newly created server and try starting it using:

```bash
# Start the node  
systemctl enable fyd  
systemctl start fyd
```

Control if the sync was started. Use `Ctrl + C` to interrupt the `journalctl` command

```bash
journalctl -u fyd -f
# OUTPUT SHOULD BE LIKE BELOW
#
# Aug 13 16:30:20 commerciotestnet-node4 fyd[351]: I[2019-08-13|16:30:20.722] Executed block                               module=state height=1 validTxs=0 invalidTxs=0
# Aug 13 16:30:20 commerciotestnet-node4 fyd[351]: I[2019-08-13|16:30:20.728] Committed state                              module=state height=1 txs=0 appHash=9815044185EB222CE9084AA467A156DFE6B4A0B1BAAC6751DE86BB31C83C4B08
# Aug 13 16:30:20 commerciotestnet-node4 fyd[351]: I[2019-08-13|16:30:20.745] Executed block                               module=state height=2 validTxs=0 invalidTxs=0
# Aug 13 16:30:20 commerciotestnet-node4 fyd[351]: I[2019-08-13|16:30:20.751] Committed state                              module=state height=2 txs=0 appHash=96BFD9C8714A79193A7913E5F091470691B195E1E6F028BC46D6B1423F7508A5
# Aug 13 16:30:20 commerciotestnet-node4 fyd[351]: I[2019-08-13|16:30:20.771] Executed block                               module=state height=3 validTxs=0 invalidTxs=0
```

## 6. Start the REST API
Each full node can start up its own REST API service. 
This will allow it to expose some endpoints that can be used in order to query the chain state at any moment. 

If you want to start such a service, you need to run the following command

```
fycli config chain-id $CHAINID
fycli rest-server
``` 

This will start up the REST server and make it reachable using the port `1317`.     
**From here, if you want you can use services such as [Nginx](https://www.nginx.com/) in order to make it available to other devices.**

## Next step
Now that you are a Commercio.network full node, if you want you can become a validator.
If you wish to do so, please read the [*Becoming a validator* guide](validator-node-installation.md).
