<!--
order: 2
-->

# Messages

## Mint Commercio Cash Credit (FUSD)


### Protobuf message

```protobuf
message MsgMintFUSD {
  string depositor = 1;
  repeated cosmos.base.v1beta1.Coin deposit_amount = 2;
  string ID = 3;
}
```

### Transaction message
To mint FUSD you need to create and sign the following message:
  
```json
{
  "type": "commercio/MsgMintFUSD",
  "value": {
    "depositor": "<user address>",
    "deposited_amount": [
      {
        "denom": "<token denom to be deposited>",
        "amount": "<amount to be deposited>"
      }
    ],
    "id": "<Mint UUID>"
  }
}
```


#### Fields requirements
| Field | Required | Limit/Format |
| :---: | :------: | :------: |
| `depositor` | Yes | bech32 | 
| `deposited_amount` | Yes |  | 
| `id` | Yes | [uuid-v4](https://en.wikipedia.org/wiki/Universally_unique_identifier) | 

### Action type
If you want to [list past transactions](../../../docs/developers/listing-transactions.md) including this kind of message,
you need to use the following `message.action` value: 

```
mintFUSD
```  


## Burn Commercio Cash Credit (FUSD)


### Protobuf message

```protobuf
message MsgBurnFUSD {
  string signer = 1;
  cosmos.base.v1beta1.Coin amount = 2;
  string ID = 3;
}
```

### Transaction message

To burn previously minteted FUSD you need to create and sign the following message:

```json
{
  "type": "commercio/MsgBurnFUSD",
  "value": {
    "signer": "<user address>",
    "amount": {
      "denom": "<token denom to be burned>",
      "amount": "<amount to be burned>"
    },
    "id": "<Mint UUID>"
  }
}
```
#### Fields requirements
| Field | Required | Limit/Format |
| :---: | :------: | :------: |
| `signer` | Yes | bech32 | 
| `amount` | Yes | |
| `id` | Yes | [uuid-v4](https://en.wikipedia.org/wiki/Universally_unique_identifier) |


### Action type
If you want to [list past transactions](../../../docs/developers/listing-transactions.md) including this kind of message,
you need to use the following `message.action` value: 

```
burnFUSD
```



## Set Parameters (Conversion Rate & Freeze Period)

:::warning  
This transaction type is accessible only to the [government](../../government/spec/README.md).  
Trying to perform this transaction without being the government will result in an error.  
:::


### Protobuf message

```protobuf
message MsgSetParams {
  string signer = 1;
  Params params = 2;
}
```

Params type definition

```protobuf
message Params {
  option (gogoproto.equal) = true;

  string conversion_rate = 1 [
    (gogoproto.nullable) = false,
    (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Dec"
  ];
  google.protobuf.Duration freeze_period = 2
      [ (gogoproto.nullable) = false, (gogoproto.stdduration) = true ];
}
```

### Transaction message

To set module params you need to create and sign the following message:

```json
{
  "type": "commercio/MsgSetParams",
  "value": {
    "signer": "<government address>",
    "params": {
      "conversion_rate": "<floating-point collateral rate>",
      "freeze_period": "<nono seconds freeze period>"
    },
  }
}
```

##### Fields requirements
| Field | Required | Limit/Format |
| :---: | :------: | :------: |
| `signer` | Yes | bech32 | 
| `params` | Yes | |



#### Action type
If you want to [list past transactions](../../../docs/developers/listing-transactions.md) including this kind of message,
you need to use the following `message.action` value: 

```
setParams
```