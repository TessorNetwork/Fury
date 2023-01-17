<!--
order: 3
-->

# Events

The commerciomint module emits the following events:

## Handlers

### MsgMintFUSD

| Type           | Attribute Key    | Attribute Value      |
| --------       | -------------    | ------------------   |
| new_position   | depositor        | {depositorAddress}   |
| new_position   | amount_deposited | {ucomAmount}         |
| new_position   | minted_coins     | {creditsCoins}       |
| new_position   | position_id      | {position_id}        |
| new_position   | timestamp        | {position_createdAt} |
| transfer (ufury) | recipient     | {moduleAddress}   |
| transfer (ufury) | sender        | {depositorAddress} |
| transfer (ufury) | amount        | {ucomAmount}      |
| transfer (ufusd) | recipient     | {depositorAddress}   |
| transfer (ufusd) | sender        | {moduleAddress} |
| transfer (ufusd) | amount        | {creditsCoins}      |
| message        | action           | mintFUSD              |
| message        | sender           | {senderAddress}      |
| message        | sender           | {depositorAddress}      |
| message        | sender           | {moduleAddress}      |

### MsgBurnFUSD (WIP)

| Type       | Attribute Key    | Attribute Value   |
| --------   | -------------    |----------------   |
| burned_ccc | position_id      | {position_id}     |
| burned_ccc | sender           | {userAddress}     |
| burned_ccc | amount           | {burnAmount}      |
| burned_ccc | position_deleted | {bool}            |
| message    | action           | burnFUSD           |
| message    | sender           | {senderAddress}   |

### MsgSetParams

| Type                | Attribute Key       | Attribute Value |
| ------------------- | ------------------- | --------------- |
| new_params          | params              | {params}        |
| message             | action              | setParams      |
| message             | sender              | {senderAddress} |
