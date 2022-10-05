<!--
order: 2
-->

# State
## State Objects

The `x/fantoken` module keeps the following objects in state:

| State Object | Description | Key | Value | Store |
| ------------ | ----------- | --- | ----- | ----- |
| `FanTokenForDenom` | FanToken bytecode by Denom | `[]byte{1} + []byte{denom}` | `[]byte{fantoken}` | KV |
| `FanTokens` | FanTokens bytecode | `[]byte{2} + []byte{owner} + []byte{denom}` | `[]byte{denom}` | KV |


### FanToken

A `FanToken` is a new fungible FanToken on the BitSong ecosystem. It makes use of `metadata` defined as:

#### Metadata
```protobuf
message Metadata {
  // The name of the fantoken (eg: Kitty Punk)
  string name = 1;

  // The token symbol usually shown on exchanges (eg: KITTY)
  string symbol = 2;

  // The URI to a document (on or off-chain) that contains additional 
  // information. Optional.
  string uri = 3 [ (gogoproto.customname) = "URI" ];

  // The address of the wallet allowed to set a new uri
  string authority = 4;
}
```

In particular, a `FanToken` is made up of:

```protobuf
message FanToken {
  // The string name of the given denom unit (e.g ft<hash>).
  string denom = 1;

  // The maximum supply value of mintable tokens from its definition.
  string max_supply = 2 [
    (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Int",
    (gogoproto.moretags) = "yaml:\"max_supply\"",
    (gogoproto.nullable) = false
  ];

  // // The address of the wallet allowed to mint new FanToken
  string minter = 3;

  // Metadata object for the FanToken
  Metadata meta_data = 4 [
    (gogoproto.moretags) = "yaml:\"meta_data\"",
    (gogoproto.nullable) = false
  ];
}
```

### FanToken by Denom

`FanTokenByDenom` is an additional state object for querying a FanToken by its `denom`.


## Genesis State

The `x/fantoken` module's `GenesisState` defines the state necessary for initializing the chain from a previous exported height. It is defined as:

```protobuf
message GenesisState {
  Params params = 1 [ (gogoproto.nullable) = false ];

  repeated FanToken fan_tokens = 2 [ (gogoproto.nullable) = false ];
}
```