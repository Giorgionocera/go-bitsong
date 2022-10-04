<!--
order: 1
-->

# Concepts

## Non-Fungible Tokens

BitSong *NFTs*, conceptually based on the [ERC-721 Standard](https://ethereum.org/it/developers/docs/standards/tokens/erc-721), are chain-native **non-fungible tokens** for the BitSong platforms, which are very strongly integrated with other featuress like *FanTokens*.

Thanks to this feature, the ensamble of actors in the content creation industry are provided with a new set of tools which empower their economy. With this module, any content creator, like star performers, actors, designers, musicians, photographers, writers, and really many more, are able to provide the world their new *NFT* projects.

Those projects can be used to create loyalty programs providing users with very exclusive contents, or experiences. In such projects, *NFTs* can be used as ticket for events, as pieces of art, or even as memorabilia. And these are only some examples. Here, together with the [*FanTokens*](../../fantoken/spec/README.md), you can build everything. The limit is your imagination.

In the design of the `nft` module functionalities, big part of the reasonings were based on the [Metaplex NFT platform](https://docs.metaplex.com/architecture/deep_dive/overview).

### Collection
Each BitSong *NFT* belongs to a `Collection`. So a *collection* is a group of *NFTs*.
This is defined as:
| Attribute | Type | Description |
| --------------------- | ---------------------------- | ---------------------------------------------- |
| id | `uint64` | It is an autoincrement integer number globally identifying a *Collection*.|
| symbol | `string` | It is chosen by the user and can be any string. It should follow the ISO standard for the alphabetic code (e.g., USD, EUR, BTSG, etc.).|
| uri | `string` | It is chosen by the user. It should be a link to a resource which contains a set of information linked to the *Collection*. We suggest to follow the [OpenSea `Metadata`](https://docs.opensea.io/docs/contract-level-metadata) schema as standard. It can also be empty. |
| is_mutable | `bool` | It indicates whether or not the *Collection* can be modified. Once true, it cannot be flipped anymore. |
| update_authority | `string` |  It is the address of the authority for the *Collection* managment (i.e. also adding new *NFTs* to the *Collection*). It can be changed to trasfer the management ability during the time. At the moment, this value can also be empty (which corresponds to a behavior like the one where `is_mutable = false`).|

### Metadata
Each BitSong *NFT* is defined through a set of `Metadata`.
These make use of `Creator` and `MasterEdition` objects

#### Creator
`Metadata` can be "created" by many users, each identified by its address. This users can only be added by the *NFT* creator on the platform (which corresponds to a single address), and can be verified. The verified creators, will be used to distribute the earn for secondary sells.
Each `Creator` is made up of:

| Attribute | Type | Description |
| --------------------- | ---------------------------- | ---------------------------------------------- |
| address | `string` | It is the address of the *Creator* who, if verified, will receive `share` of the earns for secondary sells.|
| verified | `bool` | It is false by default. Since it becomes true, allows the *Creator* to "claim" the earns share from secondary sells. It can be verified only by the *Creator* through a `SignMetadata` message.|
| share | `uint32` | It corresponds to the quantity (in percentage) of share for seconday sells, the *Creator* will get with respect to the other *Creators*. It is a number in the range [`0`;`10000`] (which corresponds to [`0.00`;`100.00`]).|

#### MasterEdition
In some cases, like event ticketing one, you could have the need to have a set of *NFTs* which shares the metadata. In such a scenario, it is important to define the maximum number of "copies" (supply) of an *NFT* that can be printed, and the current supply. For this reason, the `MasterEdition` object is made up of:

| Attribute | Type | Description |
| --------------------- | ---------------------------- | ---------------------------------------------- |
| max_supply | `uint64` | It is the maximum number of *NFTs* that can be generated for this specific `Metadata`.|
| supply | `uint64` | It is the number of currently minted *NFTs* for this specific `Metadata`.|

#### Metadata
Once defined the `Creator` and the `MasterEdition` objects, it is possible to define the **Metadata** one. It is characterized by:
| Attribute | Type | Description |
| --------------------- | ---------------------------- | ---------------------------------------------- |
| id | `uint64` | It is an autoincrement integer number globally identifying a collection. |
| coll_id | `uint64` | It is an external ID to the `Collection`. It is used to make the *NFTs* belong to a `Collection`.|
| name | `string` | It is chosen by the user. It should correspond to the long name the user want to associate to the *NFT*. At the moment, this value can be any string. |
| uri | `string` | It is chosen by the user. It should be a link to a resource which contains a set of information linked to the *NFT*. We suggest to follow the [OpenSea `Metadata`](https://docs.opensea.io/docs/metadata-standards) schema as standard. It can also be empty. |
| seller_fee_basis_points | `uint32` | It corresponds to the quantity (in percentage) of share for seconday sells, the creators will get with respect to the earn value. It is a number in the range [`0`;`10000`] (which corresponds to [`0.00`;`100.00`]). |
| primary_sale_happened | `bool` | It indicates whether or not the first sale of this *NFT Metadata* happened. |
| is_mutable | `bool` | It indicates whether or not the *Metadata* can be modified. Once true, it cannot be flipped anymore. |
| creators | `repeated Creator` | It is an array of `Creators` that, if verified, will receive the earn shares for secondary sells. |
| metadata_authority | `string` | It is the address of the authority for the *NFT Metadata* managment. It can be changed to trasfer the management ability during the time. At the moment, this value can also be empty (which corresponds to a behavior like the one where `is_mutable = false`). |
| mint_authority | `string` | It is the address of the minter for the *NFTs*. It can be changed to trasfer the minting ability of the *NFT* during the time. At the moment, this value can also be empty (which corresponds to a behavior in which no one is able to mint anymore). |
| master_edition | `MasterEdition` | It is a `MasterEdition` object containing info about the ability to **print** multiple editions of this *NFT*. |

### NFT
At this point it is possible to define a BitSong *NFT*. Each *NFT* is globally identified by the triad `CollID:MetadataID:Seq`. Generally, a BitSong *NFT* is defined through:

| Attribute | Type | Description |
| --------------------- | ---------------------------- | ---------------------------------------------- |
| coll_id | `uint64` | It is an external ID to the `Collection`. It is used to make the *NFTs* belong to a `Collection`.|
| metadata_id | `uint64` | It is an external ID to the `Metadata`. It is used to make the *NFTs* belong to a `Metadata`.|
| seq | `uint64` | It is the `sequence number` of the *NFT*, which is greaten than `0` only for `print` (or editions/copies) of the *NFT* (in the case it allows multiple editions).|
| owner | `string` | It is the `owner` of the *NFT*. It can be changed through an *NFT* transfer operation. At the moment, this value can also be empty (which corresponds to a behavior in which no one is the owner of the NFT anymore).|

## Main steps in NFTs project delivery

At this point, after defining all the principal elements of an *NFT project*, it is possible to analyze which are the main steps in an *NFTs* project delivery on the BitSong ecosystem.

### 1. Create a **Collection**
Once the `Collection Metadata` are produced (we suggest to follow the [OpenSea standard](https://docs.opensea.io/docs/contract-level-metadata)) and stored in decentralized file networks so that can't be modified by a central party (e.g., on IPFS or Arweave), you can create a *collection* which will store the group of *NFTs*. The command from the CLI is `bitsongd tx nft create-collection`, give it a look.

### 2. Prepare the **Metadata** for "Normal" NFT
Now it is the turn of the `NFT Metadata`. Once you produced them (we suggest to follow the [OpenSea standard](https://docs.opensea.io/docs/metadata-standard)), you need to store them in decentralized file networks so that can't be modified by a central party (e.g., on IPFS or Arweave).

### 3. Create "Normal" NFT
At this point, it is possible to create a new *NFT*. Here you need to define all the information about the *NFT*, like the `Collection` the *NFT* belongs to, the information about the `Creators`, the `Metadata` uri and so on. That's all.
If you want to try, the command from the CLI is `bitsongd tx nft create-nft`.

### 4. Prepare the **Metadata** for "multiple-edition" NFT
When you want to add another *NFT* to the `Collection` (this time, a multiple-edition one), you should start from the `NFT Metadata`. As usual, you need to store them.

### 5. Create the "Master Edition" NFT
The next step is to create a new *NFT*, but this time, you need to specify a `MaxSupply` value. The command from the CLI is again the `bitsongd tx nft create-nft`, but pay attention to the flag `--master-edition-max-supply`.

### 6. Sign **Metadata** 
As one of the creators of the *NFT*, you should verify your address in order to recevie shares of the earn after the first sale. To do this, you can simply run the `bitsongd tx nft sign-metadata` command from the CLI.

### 7. Print "Edition" of the NFT
Once all the previous steps have been successfully completed, you are ready to print one edition of the *NFT* you already created, to any valid address, by simply using the following command `bitsongd tx nft print-edition`.