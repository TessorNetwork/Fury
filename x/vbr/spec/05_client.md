<!--
order: 5
-->

# Client

## Transactions

### Deposit
Increments the block rewards pool's liquidity by the given amount denom
#### CLI


```bash
furyd tx vbr deposit \
  [amount denom]
```

**Parameters:**

| Parameter | Description |
| :------- | :---------- | 
| `amount denom`               | Amount of tokens to increment reward pool  |


### Set Params

#### CLI

```bash
furyd tx vbr set-params \
  [epoch_identifier] \
  [earn_rate]
```

**Parameters:**

| Parameter | Description |
| :------- | :---------- | 
| `epoch_identifier`         | Epoch reward identifier: minute/hour/day/week/month  |
| `earn_rate`         | Earn reward rate  |



## Queries

### Getting the total Pool Funds

#### CLI

```bash
furyd query vbr pool-funds
```

#### gRPC
Endpoint:

```
tessornetwork.fury.vbr.Query/GetBlockRewardsPoolFunds
```

##### Example

```bash
grpcurl -plaintext \
    localhost:9090 \
    tessornetwork.fury.vbr.Query/GetBlockRewardsPoolFunds
```

##### Response
```json
{
  "funds": [
    {
      "denom": "ufury",
      "amount": "10161937347246000000000000000000"
    }
  ]
}
```

#### REST

Endpoint:
   
```
/commercionetwork/vbr/funds
```

##### Example

Getting all the block rewards pool Funds:

```
http://localhost:1317/commercionetwork/vbr/funds
```

##### Response
```json
{
  "funds": [
    {
      "denom": "ufury",
      "amount": "10000000.000000000000000000"
    }
  ]
}
```

### Reading the Params

#### CLI

```bash
furyd query vbr get-params
```

#### gRPC
Endpoint:

```
tessornetwork.fury.vbr.Query/GetParams
```

##### Example

```bash
grpcurl -plaintext \
    localhost:9090 \
    tessornetwork.fury.vbr.Query/GetParams
```

##### Response
```json
{
  "params": {
    "distrEpochIdentifier": "day",
    "earnRate": "500000000000000000"
  }
}
```

#### REST

Endpoint:
   
```
/commercionetwork/vbr/params
```

##### Example

Getting the parameters:

```
http://localhost:1317/commercionetwork/vbr/params
```

##### Response
```json
{
  "params": {
    "distr_epoch_identifier": "day",
    "earn_rate": "0.500000000000000000"
  }
}
```