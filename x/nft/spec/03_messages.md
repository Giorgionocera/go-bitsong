# Messages

Messages (`msg`s) are objects that trigger state transitions. Messages are wrapped in transactions (`tx`s) that clients submit to the network. The BitSong SDK wraps and unwraps `nft` module messages from transactions.

## MsgCreateCollection

`MsgCreateCollection` is a message to create a new Collection which will groups a set of NFTs.

```protobuf
message MsgCreateCollection {
  // Sender of the message
  string sender = 1;

  // Symbol of the new Collection
  string symbol = 2;

  // Name of the new Collection
  string name = 3;

  // URI pointing to a JSON representing the Collection (preferable on a 
  // decentralized file network)
  string uri = 4;

  // True if the Collection is mutable. Once false, cannot be flipped anymore.
  bool is_mutable = 5;

  // Address of the wallet who can update this Collection (while is_mutable is
  // true)
  string update_authority = 6;
}

message MsgCreateCollectionResponse {
  // Newly created Collection ID
  uint64 id = 1;
}
```

Steps:

1. Get the next unique `collection_id` by using `last_collection_id` (it is an autoincrement integer);
2. Update the state with the new `last_collection_id`;
3. Create the new Collection object with got `collection_id`, `msg.Symbol` as symbol, `msg.Name` as name, `msg.Uri` as URI, `msg.IsMutable` as value for the is_mutable value, and `msg.UpdateAuthority` as authority for the Collection;
4. Update the state with the new `Collection`;
5. Emit event for new Collection creation;
6. Return newly create `collection_id`.

## MsgCreateNFT

`MsgCreateNFT` is a message for creating an NFT with specific Metadata.
The execution of this message will produce the creation of a new `Metadata` object and the creation of a new NFT associated to this newly created Metadata.
It returns the just created NFT and Metadata ids.
Here, the `sender` is set as NFT creator (`Authority` and `Minter`) and NFT `Owner` after the successful execution.

```protobuf
message MsgCreateNFT {
  // Sender of the message
  string sender = 1;

  // Metadata for the new NFT
  Metadata Metadata = 2  [ (gogoproto.nullable) = false ];
}
message MsgCreateNFTResponse {
  // ID of the newly created NFT
  string id = 1;

  // Collection of the newly created NFT
  uint64 coll_id = 2;

  // Metadata id linked to the newly created NFT
  uint64 metadata_id = 3;
}
```

Steps:

1. Ensure Collection with id value `msg.Metadata.CollId` exists;
2. Ensure `msg.Sender` is the actual `Authority` of the Collection;
3. Get the next unique `metadata_id` by using `last_metadata_id` (it is an autoincrement integer);
4. Update the state with the new `last_metadata_id`;
5. Create `Metadata` for the new NFT object;
6. For each creator, set the `verified` field to false (since no creator validated the new Metadata yet);
7. Manage the `MasterEdition` for "Normal NFT", by using a `MaxSupply = Supply = 1`;
8. Update the state with the new `Metadata`;
9. Emit event for Metadata creation;
10. Pay the NFT *issue fee* if the parameter is a positive value;
11. Create an NFT object that has as owner the `msg.Sender`, belongs to the Collection `msg.CollId`, is linked to the newly created `metadata_id` and has 0 as sequence number `Seq`;
12. Update the state with the new `NFT`;
13. Emit event for NFT creation;
14. Return newly created NFT `id` and `metadata_id`.

## MsgPrintEdition

`MsgPrintEdition` is a message for printing a new edition for a *multiple-copies* NFT. It allows to create copies for a `MasterEdition` which not reached the `MaxSupply` yet.
Editions can only be printed by `Minter`.

```protobuf
message MsgPrintEdition {
  // Sender of the message
  string sender = 1;

  // Collection of the NFT you want to print
  uint64 coll_id = 2;

  // Metadata of the NFT you want to print
  uint64 metadata_id = 3;

  // Address of the wallet you want to print the NFT 
  string owner = 4;
}
message MsgPrintEditionResponse {
  // ID of the newly created NFT
  string id = 1;

  // Collection of the newly created NFT
  uint64 coll_id = 2;

  // Metadata id linked to the newly created NFT
  uint64 metadata_id = 3;
}
```

Steps:

1. Ensure that Metadata with id `msg.MetadataId` exists;
2. Ensure that the master edition NFT for such a Metadata exists (the one with `Seq = 0`);
3. Ensure that message is executed by the Metadata `Minter`;
4. Ensure that `MasterEdition` attribute for Metadata is valid;
5. Ensure that the current supply of the printed NFT is lower than `MaxSupply`;
6. Obtain Seq number for the new Metadata;
7. Update the state with the new `Metadata` including the incremented supply;
8. Pay the NFT *issue fee* if the parameter is a positive value;
9. Create a new NFT with `msg.CollId`, `msg.MetadataId`, new edition number and `msg.Owner`
10. Update the state with the new `NFT`;
11. Emit event for print edition;
12. Return NFT identifier.

## MsgTransferNFT

`MsgTransferNFT` is a message for updating the owner of an NFT.

```protobuf
message MsgTransferNFT {
  // Sender of the message
  string sender = 1;

  // ID of the NFT to transfer
  string id = 2;

  // New Owner of the NFT
  string new_owner = 3;
}
```

Steps:

1. Ensure that NFT with id `msg.Id` exists;
2. Ensure that the `msg.Sender` is the actual `Owner` of the NFT;
3. Set the owner of NFT to `msg.NewOwner`;
4. Update the state with the updated `NFT`;
5. Emit event for an NFT transfer.

## MsgSignMetadata

`MsgSignMetadata` is a message for allowing the creator verifying the Metadata.
IMPORTANT: This is a mandatory step to get shares for secondary sells.
Once it is executed, the `Verified` field of the specific creator for the specific Metadata become true.

```protobuf
message MsgSignMetadata {
  // Sender of the message
  string sender = 1;

  // ID of the Collection to which the Metadata belongs
  uint64 coll_id = 2;

  // ID of the Metadata to sign
  uint64 metadata_id = 3;
}
```

Steps:

1. Ensure that Metadata with id `msg.MetadataId` exists;
2. Ensure that `msg.Sender` is one of the creators of Metadata;
3. Set `Verified` attribute to true for such a Creator;
4. Update the state with the updated Metadata;
5. Emit event for Metadata sign.

## MsgUpdateMetadata

`MsgUpdateMetadata` is a message for updating Metadata. It can be run only by the Metadata UpdateAuthority.
`Name`, `URI`, `SellerFeeBasisPoints` and `Creators` fields can be changed when the Metadata has `IsMutable` flag as true.
IMPORTANT: AFTER THIS UPDATE, ALL THE CREATORS WILL BE AUTOMATICALLY "UNSIGNED" AND THEY WILL NOT RECEIVE SHARES UP TO A NEW SIGN.

```protobuf
message MsgUpdateMetadata {
  // Sender of the message
  string sender = 1;

  // ID of the Collection to which the Metadata belongs
  uint64 coll_id = 2;

  // ID of the Metadata to update
  uint64 metadata_id = 3;

  // New name of the asset
  string name = 4;
  
  // New URI of the asset pointing to a JSON representing the asset (
  // preferable on a decentralized file network)
  string uri = 5;

  // New value for royalty percentage for the verified creators group in all 
  // secondary sales [0;10000] (corresponds to [0.00;100.00])
  uint32 seller_fee_basis_points = 6;
  
  // New array of creators. They, once verified, will receive shares of earn 
  // for secondary sales. IMPORTANT: They will be automatically "unsigned"
  repeated Creator creators = 7
      [ (gogoproto.nullable) = false ];
}
```

Steps:

1. Ensure that Metadata with id `msg.MetadataId` exists;
2. Ensure that Metadata is mutable;
3. Ensure `msg.Sender` is authority for the Metadata;
4. Set passed `Name`, `Uri`, `SellerFeeBasisPoints` and `Creators` for the Metadata;
5. Set `verified` attribute to `false` for all creators;
4. Update the state with the updated Metadata;
7. Emit event for Metadata update.

## MsgUpdateMetadataAuthority

`MsgUpdateMetadataAuthority` is a message for updating Metadata authority address.

```protobuf
message MsgUpdateMetadataAuthority {
  // Sender of the message
  string sender = 1;

  // ID of the Collection to which the Metadata belongs
  uint64 coll_id = 2;

  // ID of the Metadata to update
  uint64 metadata_id = 3;

  // Address of the new authority
  string new_authority = 4;
}
```

Steps:

1. Ensure Metadata with id `msg.MetadataId` exists;
2. Ensure `msg.Sender` is the actual authority for the Metadata;
3. Set Metadata `UpdateAuthority` with `NewAuthority`;
4. Update the state with the updated Metadata;
5. Emit event for authority update for the Metadata.

## MsgUpdateMintAuthority

`MsgUpdateMetadataAuthority` is a message for updating the Minter address.

```protobuf
message MsgUpdateMintAuthority {
  // Sender of the message
  string sender = 1;

  // ID of the Collection to which the Metadata belongs
  uint64 coll_id = 2;

  // ID of the Metadata to update
  uint64 metadata_id = 3;

  // Address of the new authority
  string new_authority = 4;
}
```

Steps:

1. Ensure Metadata with id `msg.MetadataId` exists;
2. Ensure `msg.Sender` is the actual authority for the Metadata;
3. Set Metadata `MintAuthority` with `NewAuthority`;
4. Update the state with the updated Metadata;
5. Emit event for mint authority update for the Metadata.

## MsgUpdateCollectionAuthority

`MsgUpdateCollectionAuthority` is a message to update Collection authority to a new one.
It should be executed by Collection authority.

```protobuf
message MsgUpdateCollectionAuthority {
  // Sender of the message
  string sender = 1;

  // ID of the Collection to update
  uint64 collection_id = 2;

  // Address of the new authority
  string new_authority = 4;
}
```

Steps:

1. Ensure Collection exists with id `msg.CollectionId`;
2. Ensure `msg.Sender` is the actual authority for the Collection;
3. Set Collection authority with `msg.NewAuthority`;
4. Store updated Collection object into storage
4. Update the state with the updated Collection;
5. Emit event for Collection authority update.
