package v3_0_0

const (
	ModuleName   = "furykyc"
	StoreKey     = ModuleName
	RouterKey    = ModuleName
	QuerierRoute = ModuleName
	MemStoreKey  = "mem_" + ModuleName

	// --- Keeper
	MembershipsStorageKey  = StoreKey + ":storage:"
	TrustedSignersStoreKey = StoreKey + ":signers"
	InviteStorePrefix      = "invite:"

	// --- Messages
	MsgTypeInviteUser                = "inviteUser"
	MsgTypesDepositIntoLiquidityPool = "depositIntoLiquidityPool"
	MsgTypeAddTsp                    = "addTsp"
	MsgTypeRemoveTsp                 = "removeTsp"
	MsgTypeBuyMembership             = "buyMembership"
	MsgTypeSetMembership             = "setMembership"
	MsgTypeRemoveMembership          = "removeMembership"

	QueryGetInvites                 = "invites"
	QueryGetInvite                  = "invite"
	QueryGetTrustedServiceProviders = "tsps"
	QueryGetPoolFunds               = "poolFunds"
	QueryGetMembership              = "membership"
	QueryGetMemberships             = "memberships"
	QueryGetTspMemberships          = "sold"
)
