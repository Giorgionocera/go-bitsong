<!--
order: 2
-->
# State
## State Objects

The `x/nft` module keeps the following objects in state:

| State Object | Description | Key | Value | Store |
| ------------ | ----------- | --- | ----- | ----- |
| `NFT`        | NFT bytecode | `[]byte{1} + []byte{coll_id} + []byte{metadata_id} + []byte{seq}` | `[]byte{nft}` | KV |
| `NFTByOwner` | NFT id bytecode by Owner | `[]byte{2} + []byte{owner} + []byte{coll_id} + []byte{metadata_id} + []byte{seq}` | `[]byte{ntf_identifier}` | KV |
| `Metadata` | Metadata bytecode | `[]byte{3} + []byte{coll_id} + []byte{metadata_id}` | `[]byte{metadata}` | KV |
| `Collection` | Collection bytecode | `[]byte{4} + []byte{coll_id}` | `[]byte{collection}`| KV |
| `LastMetadataId` | Last Metadata id (together with respective collection id) bytecode | `[]byte{5} + []byte{coll_id}` | `[]byte{{coll_id, metadata_id}}` | KV |
| `LastCollectionId` | Last Collection id bytecode | `[]byte{6}` | `[]byte{coll_id}` | KV |


### NFT

A `NFT` is a single unit of a non-fungible token. It is defined by `coll_id`, `metadata_id`, `seq` and `owner`.

The string identifier of NFT (nftId) is expressed as `{coll_id}:{metadata_id}:{seq}`

```protobuf
message NFT {
  // ID of the collection the NFT belongs to
  uint64 coll_id = 1;

  // ID of the metadata relative to the NFT
  uint64 metadata_id = 2;

  // Sequence number for the NFT. It is greater than 0 only for a print of a multiprint edition NFT
  uint64 seq = 3; 

  // Address of the NFT owner
  string owner = 4; 
}
```

### NFT by Owner

`NFTByOwner` is an additional state object for querying NFTs by their owner.

### Metadata

`Metadata` are the metadata of each NFT, both they are `Prints` of `MasterEdition` and `Normal` NFTs.
Metadata has `Creators` and `MasterEdition` object integrated for print ability.
The latter involves `supply` and `max_supply` fields and, everytime a new print is created, the supply is increased by one, and the new `NFT` object has a new `seq` number that makes it unique. Print numbers cannot exceed `max_supply`.

```protobuf
message MasterEdition {
  // Current supply for a multiple-edition NFT
  uint64 supply = 1; 

  // Maximum supply for an NFT. If it is a "normal" NFT, this value must be
  // equal to 1
  uint64 max_supply = 2; 
}

message Metadata {
  // ID of the metadata
  uint64 id = 1; 

  // ID of the collection the NFT (Metadata) belongs to
  uint64 coll_id = 2; 

  // Name for the NFT
  string name = 3; 

  // URI pointing to a JSON representing the asset (preferable on a 
  // decentralized file network)
  string uri = 4; 

  // Royalty percentage for the verified creators group in all secondary 
  // sales [0;10000] (corresponds to [0.00;100.00])
  uint32 seller_fee_basis_points = 5; 

  // True if primary sale happened. Once true, cannot be flipped anymore. 
  // Default is false
  bool primary_sale_happened = 6; 

  // True if the Metadata are mutable. Once false, cannot be flipped anymore. 
  // Default is false
  bool is_mutable = 7; 

  // Array of creators. They, once verified, will receive shares of earn for 
  // secondary sales.
  repeated Creator creators = 8 [ (gogoproto.nullable) = false ]; 

  // Address of the wallet who can update these Metadata (while is_mutable is 
  // true)
  string metadata_authority = 9; 

  // Address of the wallet who can mint new NFT from these Metadata (whle 
  // max_supply of master_edition is not reached)
  string mint_authority = 10; 

  // Master edition configuration
  MasterEdition master_edition = 11; 
}

message Creator {
  // Address of the creator
  string address = 1; 

  // If the creator has completed the verification process. IMPORTANT: WHILE 
  // FALSE, THE CREATOR WILL NOT RECEIVE ANY SHARE FOR SECONDARY SALES
  bool verified = 2; 

  // Share value for the creator. In percentage, NOT basis points
  uint32 share = 3; 
}
```

### Collection

A `Collection` is a collection of NFTs. It is defined by `id`, `symbol`, `name`, `uri`, `is_mutable`, `update_authority` fields.

```protobuf
message Collection {
  // ID of the metadata
  uint64 id = 1; 
  
  // Symbol for the Collection
  string symbol = 2; 
  
  // Name for the Collection
  string name = 3; 

  // URI pointing to a JSON representing the Collection (preferable on a 
  // decentralized file network)
  string uri = 4; 
  
  // True if the Collection is mutable. Once false, cannot be flipped anymore.
  // Default is false
  bool is_mutable = 5; 
  
  // Address of the wallet who can update this Collection (while is_mutable is
  // true)
  string update_authority = 6; 
}
```

### LastMetadataIdInfo

`LastMetadataIdInfo` tracks last metadata id per each collection to avoid duplication in metadata ids.

```protobuf
message LastMetadataIdInfo {
  // Collection ID
  uint64 coll_id = 1; 
  
  // Last Metadata ID
  uint64 last_metadata_id = 2; 
}
```

## Genesis State

The `x/nft` module's `GenesisState` defines the state necessary for initializing the chain from a previous exported height. It is defined as:

```protobuf
message GenesisState {
  Params params = 1 [ (gogoproto.nullable) = false ];

  repeated Metadata metadata = 2 [ (gogoproto.nullable) = false ];

  repeated LastMetadataIdInfo last_metadata_ids = 3 [ (gogoproto.nullable) = false ];

  repeated NFT nfts = 4 [ (gogoproto.nullable) = false ];

  repeated Collection collections = 5 [ (gogoproto.nullable) = false ];

  uint64 last_collection_id = 6;
}
```