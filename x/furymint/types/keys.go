package types

const (
	ModuleName   = "furymint"
	StoreKey     = ModuleName
	RouterKey    = ModuleName
	QuerierRoute = ModuleName
	MemStoreKey  = "mem_" + ModuleName

	EtpStorePrefix = StoreKey + ":etp:"
	CreditsDenom   = "ufusd"
	BondDenom      = "ufury"

	QueryGetEtpRest         = "etp"
	QueryGetallEtpsRest     = "etps"
	QueryGetEtpsByOwnerRest = "etpsOwner"
	QueryConversionRateRest = "conversion_rate"
	QueryFreezePeriodRest   = "freeze_period"
	QueryGetParamsRest      = "params"

	MsgTypeMintFUSD   = "mintFUSD"
	MsgTypeBurnFUSD   = "burnFUSD"
	MsgTypeSetParams = "setParams"
)
