# CommercioMint

The `commerciomint` module is the one that allows you to create Exchange Trade Position (*ETPs*) using your 
Commercio.network tokens (*ufury*) in order to get Commercio Cash Credits (*ufusd*) in return.

A *Exchange Trade Position* (*ETP*) is a core component of the Commercio Network blockchain whose purpose is to
create Commercio Cash Credits (`ufusd`) in exchange for Commercio Tokens (`ufury`) which it then holds in
escrow until the borrowed Commercio Cash Credits are returned.

In simple words, opening a ETP allows you to exchange any amount of `ufury` to get relative the amount of `ufusd` with relative Conversion Rate value. 
For example, if you open a ETP lending `100 ufury` with 1.1 Conversion Rate value will result in you receiving `90 ufusd` (approximation by default).  
Initial Conversion Rate value is 1.0.       

## Transactions

### Mint Commercio Cash Credit (FUSD)

#### Transaction message
To mint FUSD you need to create and sign the following message:
  
```json
{
  "type": "commercio/MsgMintFUSD",
  "value": {
    "deposited_amount": [
      {
        "amount": "<amount to be deposited>",
        "denom": "<token denom to be deposited>"
      }
    ],
    "depositor": "<user address>",
    "id": "<Mint UUID>"
  }
}
```


##### Fields requirements
| Field | Required | Limit/Format |
| :---: | :------: | :------: |
| `deposited_amount` | Yes |  | 
| `depositor` | Yes | bech32 | 
| `id` | Yes | [uuid-v4](https://en.wikipedia.org/wiki/Universally_unique_identifier) | 

#### Action type
If you want to [list past transactions](../../../developers/listing-transactions.md) including this kind of message,
you need to use the following `message.action` value: 

```
mintFUSD
```  






### Burn Commercio Cash Credit (FUSD)

#### Transaction message

To burn previously minteted FUSD you need to create and sign the following message:

```json
{
  "type": "commercio/MsgBurnFUSD",
  "value": {
    "signer": "<user address>",
    "amount": {
      "amount": "<amount to be burned>",
      "denom": "<token denom to be burned>"
    },
    "id": "<Mint UUID>"
  }
}
```

##### Fields requirements
| Field | Required | Limit/Format |
| :---: | :------: | :------: |
| `signer` | Yes | bech32 | 
| `amount` | Yes | |
| `id` | Yes | [uuid-v4](https://en.wikipedia.org/wiki/Universally_unique_identifier) |


#### Action type
If you want to [list past transactions](../../../developers/listing-transactions.md) including this kind of message,
you need to use the following `message.action` value: 

```
burnFUSD
```



### Set FUSD conversion rate

:::warning  
This transaction type is accessible only to the [government](../../government/README.md).  
Trying to perform this transaction without being the government will result in an error.  
:::

#### Transaction message

To set the FUSD conversion rate you need to create and sign the following message:

```json
{
  "type": "commercio/MsgSetFUSDConversionRate",
  "value": {
    "signer": "<government address>",
    "rate": "<floating-point collateral rate>"
  }
}
```

#### Action type
If you want to [list past transactions](../../../developers/listing-transactions.md) including this kind of message,
you need to use the following `message.action` value: 

```
setEtpsConversionRate
```

### Set FUSD freeze period

:::warning  
This transaction type is accessible only to the [government](../../government/README.md).  
Trying to perform this transaction without being the government will result in an error.  
:::

#### Transaction message

To set the FUSD freeze period you need to create and sign the following message:

```json
{
  "type": "commercio/MsgSetFUSDFreezePeriod",
  "value": {
    "signer": "<government address>",
    "freeze_period": "<nono seconds freeze period>"
  }
}
```

#### Action type
If you want to [list past transactions](../../../developers/listing-transactions.md) including this kind of message,
you need to use the following `message.action` value: 

```
setEtpsFreezePeriod
```






## Queries

### Reading all Exchange Trade Position (ETP) opened by a user

#### CLI

```sh
fycli query commerciomint get-etps [user-addr]
```

#### REST

Endpoint:
   
```
/commerciomint/etps/${address}
```

Parameters:

| Parameter | Description |
| :-------: | :---------- | 
| `address` | Address of the user for which to read all the ETPs |

##### Example

Getting ETPs opened by `did:com:15erw8aqttln5semks0vnqjy9yzrygzmjwh7vke`:

```
http://localhost:1317/commerciomint/etps/did:com:15erw8aqttln5semks0vnqjy9yzrygzmjwh7vke
```

#### Response
```json
{
  "height": "0",
  "result": [
    {
      "credits": {
        "amount": "500000",
        "denom": "ufusd"
      },
      "collateral": "450000",
      "exchange_rate": "0.900000000000000000",
      "owner": "did:com:15erw8aqttln5semks0vnqjy9yzrygzmjwh7vke/1570177686",
      "id": "83672b49-c2a1-4ce3-a52a-859039b1231e",
      "created_at": "2021-03-21T22:42:39.805871642Z"
    }
  ]
}
```

### Reading the current FUSD conversion rate

#### CLI

```bash
fycli query commerciomint conversion-rate
```

#### REST

Endpoint:
   
```
/commerciomint/conversion_rate
```

##### Example

Getting the current FUSD conversion rate:

```
http://localhost:1317/commerciomint/conversion_rate
```

#### Response
```json
{
  "height": "0",
  "result": "1.000000000000000000"
}
```



### Reading the current FUSD freeze period
#### CLI

```bash
fycli query commerciomint freeze-period
```

#### REST

Endpoint:
   
```
/commerciomint/freeze_period
```

##### Example

Getting the current FUSD freeze period:

```
http://localhost:1317/commerciomint/freeze_period
```

#### Response
```json
{
  "height": "0",
  "result": "1814400000000000"
}
```