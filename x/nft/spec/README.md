# `nft`

## Abstract

This document specifies the `nft` module of the BitSong chain.

The `nft` module enables the BitSong chain to support non-fungible tokens. Thanks to this feature, the ensamble of actors in the content creation industry can support their economy by launching their *NFT* projects. It is quite likely that one of the main reasons, which have enabled a large adoption of *NFT* technology in the universe of content creation, is precisely that creators see the profits from their activities being cut to the bone by the platforms that provide the services.

In this sense, the ability to represent anything unique gives those actors a power never seen before. But it doesn't end there. The most interesting thing is that beyond the content creation sector, this technology can find applications in a lot of other environments! 

Since *NFTs* allows to represent the ownership of unique items, in fact, they permit to tokenize art, collectibles, ticket for events and very much more. 

Thanks to this module, anyone inside the BitSong ecosystem can start creating their *nft collections* and mint them within for very low fees.

### The power of NFT

*NFT* (**Non-Fungible Token**) are particular type of token used to represent somthing that is *unique*. **Non-fungible**, in fact, is an economic term generally used to describe things which have been considered in their identity and are not replaceable with any other asset.

To make it clearer the difference between *fungible* and *non-fungible* tokens we can use an example.
Let's assume that Alice has got a 5 euro banknote. Also Bob has got a 5 euro bankonote. If Alice exchange this bankonote with bob, in the end, nothing changes. This is possible because money is a **fungible** commodity, and each banknote is exchangeable for any other if it is part of the same denomination.

On the other hand, Charlie has a very rare limited edition pokemon collection card, produced in 2005 and kept in perfect condition. All these conditions make the card a very rare, if not unique and, as a consequence, **non-fungible**.

For this reason, *NFTs* are characterized by *unicity*, since each *NFT* is unique. There also exists type of *NFT* which allows a predetermined number of "copies", where each copy is unique, but they share a set of feature. Each *NFT* is linked to a *owner* and the ownership can be transferred to another user only by the actual owner.

In this way, the user who own an *NFT* can easily demonstrate the ownership. He can also sell it (by making a profit), transfer it, or even hold it forever.

Similarly, who create an *NFT*, can easily prove the fact he's the creator, he can determine the scarcity level and, he can also earn royalties for secondary market sells, every time the *NFT* is sold.

### NFT in BitSong

The `nft` module for the BitSong ecosystem is based on the concept of *NFT* platform proposed by [**Metaplex**](https://docs.metaplex.com/architecture/deep_dive/overview). Thanks to this module, BitSong enable a very large set of applications for its users.

Thanks to their versatility, BitSong *NFT* can be used as well in the art industry (as a piece of digital art), as in the entertainment universe (for the implementation of ticketing systems). Or more in the experiences world (by incarnating the concept of memorabilia).

Each *NFT*, which is characterized by a `sequence` number, belongs to a `collection` (with a `symbol`, a `name` and a `uri` that, until the collection `is mutable`, can be changed by an `authority`), is defined through a set of `metadata` (with a `name`, a `uri`, a list of `creators` and an attribute to manage the royalties. All this, until the metadata `is mutable` and was not sold yet, can be changed by an `authority`, and in the case of an NFT allowing "copies", can be minted by another authority) and is held by an `owner`.
For this reason, in fact, each *NFT* is identified by an ID which is made up of the triad `CollectionID:MetadataID:SequenceNumber`.

Finally, thanks to the `nft` module, users on BitSong can:

- manage *NFTs*, issuing, minting, updating, and transferring them;
- build applications that use the *NFTs* API to create completely new customizable systems for a very huge list of applications.

Features that may be added in the future are described in Future Improvements.

## Table of Contents

1. **[Concepts](01_concepts.md)**
   
2. **[State](02_state.md)**
   
     <!--
     State Transitions
     -->
     <!--
     Keeper
     -->
3. **[Messages](03_messages.md)**
   
     <!--
     Begin-Block
     -->
     <!--
     End-Block
     -->
4. **[Events](04_events.md)**
   
5. **[Parameters](05_parameters.md)**
   <!--
   Test Cases
   -->
   <!--
   Benchmarks
   -->
6. **[Client](06_client.md)**   
7. **[Future Improvements](07_future_improvements.md)**