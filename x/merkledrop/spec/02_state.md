<!--
order: 2
-->

# State
## State Objects

The `x/merkledrop` module keeps the following objects in state:

| State Object | Description | Key | Value | Store |
| ------------ | ----------- | --- | ----- | ----- |
| `MerkleDrop` | MerkleDrop bytecode | `[]byte{1:merkledrop_id}` | `[]byte{merkledrop}` | KV |
| `MerkleDropByOwner` | MerkleDrop id bytecode by Owner | `[]byte{2:owner:merkledrop_id}` | `[]byte{merkledrop_id}` | KV |
| `LastMerkleDropId` | Last MerkleDrop id bytecode | `[]byte{3}` | `[]byte{merkledrop_id}` | KV |
| `ClaimedMerkleDrop` | MerkleDrop claimed by Index | `[]byte{4:merkledrop_id:index}` | `[]byte{1}` | KV |
| `MerkleDropByEndHeight` | MerkleDrop by End Height | `[]byte{10:block_height:merkledrop_id}` | `[]byte{1}` | KV |


### MerkleDrop

A `MerkleDrop` is a new MerkleTree structure that allows the users to claim airdrop. It makes use of:

```protobuf
message Merkledrop {
	// ID of the merkledrop
	uint64 id = 1;

	// Root of the merkledrop
	string merkle_root = 2 [ (gogoproto.moretags) = "yaml:\"merkle_root\"" ];

	// Starting block height for claim the airdrop
	int64 start_height = 3;

	// Ending block height for claim the airdrop
	int64 end_height = 4;

	// Denom to distribute through the merkledrop
	string denom = 5;

	// Total amount to distribuite through the merkledrop
	string amount = 6 [
		(gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Int",
		(gogoproto.nullable) = false
	];

	// Total amount already claimed through the merkledrop
	string claimed = 7 [
		(gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Int",
		(gogoproto.nullable) = false
	];

	// Owner of the merkledrop
	string owner = 8;
}
```

### MerkleDrop by Owner and EndHeight, Claimed MerkleDrop

`MerkleDropByOwner` and `MerkleDropByEndHeight` are additional state objects for querying MerkleDrop by their `Owner` or by their `EndHeight`. Similarly the `ClaimedMerkleDrop` is used for querying which index claimed an `MerkleDrop`.


## Genesis State

The `x/merkledrop` module's `GenesisState` defines the state necessary for initializing the chain from a previous exported height. 

It is makes use of the `Indexes` structure:
```protobuf
message Indexes {
  uint64 merkledrop_id = 1 [ (gogoproto.moretags) = "yaml:\"mdi\"" ];
  repeated uint64 index = 2 [ (gogoproto.moretags) = "yaml:\"i\"" ];
}
```

And is defined  as:
```protobuf
message GenesisState {
  uint64 last_merkledrop_id = 1;

  repeated Merkledrop merkledrops = 2 [ (gogoproto.nullable) = false ];

  repeated Indexes indexes = 3;

  Params params = 4 [ (gogoproto.nullable) = false ];
}
```