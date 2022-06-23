# State

## NFT

A `NFT` is a single unit of a non-fungible token. It stores `id`, `owner` and `metadata_id`.

```protobuf
message NFT {
  uint64 id = 1;
  string owner = 2;
  uint64 metadata_id = 3;
  uint64 edition = 4;
}
```

- NFT: `0x01 | format(id) -> NFT`
- NFT by Owner: `0x02 | owner | format(id) -> nft_id`
- LastNftId `0x06 -> id`

## Metadata

A `Metadata` is a metadata that is attached to an nft.

```protobuf
message MasterEdition {
  uint64 supply = 1;
  uint64 max_supply = 2;
}

message Data {
  /// The name of the asset
  string name = 1;
  /// The symbol for the asset
  string symbol = 2;
  /// URI pointing to JSON representing the asset
  string uri = 3;
  /// Royalty basis points that goes to creators in secondary sales (0-10000)
  uint32 seller_fee_basis_points = 4;
  /// Array of creators, optional
  repeated Creator creators = 5;
}

message Metadata {
  uint64 id = 1;
  string update_authority = 2;
  string mint = 3;
  Data data = 4;
  // Immutable, once flipped, all sales of this metadata are considered
  // secondary.
  bool primary_sale_happened = 5;
  // Whether or not the data struct is mutable, default is not
  bool is_mutable = 6;
  MasterEdition master_edition = 7;
}

message Creator {
  string address = 1;
  bool verified = 2;
  // In percentages, NOT basis points ;) Watch out!
  uint32 share = 3;
}
```

- Metadata: `0x03 | format(id) -> Metadata`
- LastMetadataId `0x07 -> id`

### Edition

Metadata has `MasterEdition` object integrated for print ability.
It involves `supply` and `max_supply` fields.
When new print is created, supply is increased and new `NFT` object with unique `edition` is created.
Print cannot exceed `max_supply`.

## Collection

A `Collection` is a collection of nfts on certain criteria. It stores `id`, `name`, `uri`, `update_authority` fields.
The nfts on certain collection can be found by querying `CollectionRecord` objects with certain `collection_id`.

```protobuf
message Collection {
  uint64 id = 1;
  string name = 2;
  string uri = 3;
  string update_authority = 4;
}
```

- Collection: `0x04 | format(id) -> Collection`
- LastCollectionId `0x08 -> id`

```protobuf
message CollectionRecord {
  uint64 nft_id = 1;
  uint64 collection_id = 2;
}
```

- CollectionRecord: `0x05 | format(collection_id) | format(nft_id) -> CollectionRecord`

## Params

Params is a module-wide configuration structure that stores nft module's system parameters.

- Params: `Paramsspace("nft") -> Params`

```protobuf
// Params defines nft module's parameters
message Params {
  option (gogoproto.equal) = true;
  option (gogoproto.goproto_stringer) = false;

  cosmos.base.v1beta1.Coin issue_price = 1 [
    (gogoproto.moretags) = "yaml:\"issue_price\"",
    (gogoproto.nullable) = false
  ];
}
```