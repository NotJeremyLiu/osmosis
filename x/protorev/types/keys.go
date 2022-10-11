package types

import (
	"fmt"
)

const (
	// ModuleName defines the module name
	ModuleName = "protorev"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_protorev"
)

const (
	prefixNeedToArb = iota + 1
	prefixArbDetails
	prefixPairsToIDs
	prefixTokenToTokens
	prefixPoolToRoutes
)

var (
	// ProtoRev Code
	KeyNeedToArb     = []byte{prefixNeedToArb}
	KeyArbDeatils    = []byte{prefixArbDetails}
	KeyPairsToIDs    = []byte{prefixPairsToIDs}
	KeyTokenToTokens = []byte{prefixTokenToTokens}
	KeyPoolToRoutes  = []byte{prefixPoolToRoutes}
)

func GetTokenStoreKey(token string) []byte {
	return []byte(fmt.Sprintf("token/%s", token))
}

func GetConnectedTokensPoolIDsStoreKey(tokenA, tokenB string) []byte {
	// Compare tokenA and tokenB to see which one is alphabetically first
	if tokenA < tokenB {
		return []byte(fmt.Sprintf("connected_tokens/%s/%s", tokenA, tokenB))
	} else {
		return []byte(fmt.Sprintf("connected_tokens/%s/%s", tokenB, tokenA))
	}
}

func GetPoolRoutesStoreKey(poolId uint64) []byte {
	return []byte(fmt.Sprintf("pool_routes/%d", poolId))
}

func KeyPrefix(p string) []byte {
	return []byte(p)
}
