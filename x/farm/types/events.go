package types

// farm module event types
const (
	EventTypeCreatePool = "create-pool"
	EventTypeProvide    = "provide"
	EventTypeLock       = "lock"
	EventTypeUnlock     = "unlock"
	EventTypeClaim      = "claim"

	AttributeKeyAddress            = "address"
	AttributeKeyPool               = "pool"
	AttributeKeyStartHeightToYield = "start_height_to_yield"
	AttributeKeyYiledPerBlock      = "yield_per_block"
	AttributeKeyLockToken          = "lock_token"
	AttributeKeyYieldToken         = "yield_token"

	AttributeValueCategory = ModuleName
)
