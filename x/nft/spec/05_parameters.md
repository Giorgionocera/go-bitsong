<!-- 
order: 5
-->

# Parameters

This section corresponds to a module-wide configuration structure that stores system parameters. In particular, it defines the overall nft module functioning and contains the **issueFee** for the *NFT*. Such an implementation allows governance to decide the issue fee the users have to pay to perform these operations with the tokens, in an arbitrary way - since proposals can modify it.

| Key        | Type     | Value                                   |
| ---------- | -------- | --------------------------------------- |
| IssueFee | sdk.Coin | {"denom": "ubtsg", "amount": "1000000"} |