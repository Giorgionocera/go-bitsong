<!-- 
order: 6
-->

# Parameters

This section corresponds to a module-wide configuration structure that stores system parameters. In particular, it defines the overall merkledrop module functioning and contains the **creationFee** for the _merkledrop_. Such an implementation allows governance to decide the creation fee, in an arbitrary way - since proposals can modify it.

| Key         | Type             | Value                                     |
| ----------- | ---------------- | ----------------------------------------- |
| CreationFee | sdk.NewInt64Coin | {"denom": "ubtsg", "amount": "100000000"} |