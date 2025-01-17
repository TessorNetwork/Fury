<!--
order: 0
title: Commerciomint Overview
parent:
  title: "commerciomint"
-->

# CommercioMint 

## Abstract

This document specifies the commerciomint module of the Commercio Network.

The `commerciomint` module is the one that allows you to create Exchange Trade Position (*ETPs*) using your 
Commercio.network tokens (*ufury*) in order to get Commercio Cash Credits (*ufusd*) in return.

A *Exchange Trade Position* (*ETP*) is a core component of the Commercio Network blockchain whose purpose is to
create Commercio Cash Credits (`ufusd`) in exchange for Commercio Tokens (`ufury`) which it then holds in
escrow until the borrowed Commercio Cash Credits are returned.

In simple words, opening an ETP allows you to exchange any amount of `ufury` to get relative the amount of `ufusd` with relative Conversion Rate value. 
For example, if you open an ETP lending `100 ufury` with 1.1 Conversion Rate value will result in you receiving `90 ufusd` (approximation by default).  
Initial Conversion Rate value in Params is 1. 

## Contents

1. **[State](01_state.md)**
2. **[Messages](02_messages.md)**
   - [Mint Commercio Cash Credit (FUSD)](02_messages.md#mint-commercio-cash-credit-(FUSD))
   - [Burn Commercio Cash Credit (FUSD)](02_messages.md#burn-commercio-cash-credit-(FUSD))
   - [Set Parameters (Conversion Rate & Freeze Period)](02_messages.md#set-parameters-(conversion-rate-&-freeze-period))
3. **[Events](03_events.md)**
   - [Handlers](03_events.md#handlers)
4. **[Parameters](04_params.md)**
5. **[Client](05_client.md)**