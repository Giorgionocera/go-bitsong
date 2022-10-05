<!-- 
order: 5
-->

# Parameters

This section corresponds to a module-wide configuration structure that stores system parameters. In particular, it defines the overall fantoken module functioning and contains the **issueFee**, **mintFee** and **burnFee** for the _FanToken_. Such an implementation allows governance to decide the issue fee, but also the mint and burn fees the users have to pay to perform these operations with the tokens, in an arbitrary way - since proposals can modify it.

| Key        | Type     | Value                                   |
| ---------- | -------- | --------------------------------------- |
| IssueFee | sdk.Coin | {"denom": "ubtsg", "amount": "1000000"} |
| MintFee | sdk.Coin | {"denom": "ubtsg", "amount": "0"} |
| BurnFee | sdk.Coin | {"denom": "ubtsg", "amount": "0"} |