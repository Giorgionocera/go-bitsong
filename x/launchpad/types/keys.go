package types

const (
	// ModuleName is the name of the module
	ModuleName = "launchpad"

	// StoreKey is the string store representation
	StoreKey string = ModuleName

	// QuerierRoute is the querier route for the module
	QuerierRoute string = ModuleName

	// RouterKey is the msg router key for the module
	RouterKey string = ModuleName
)

var (
	PrefixLaunchPad           = []byte{0x01}
	PrefixLaunchPadByEndTime  = []byte{0x02}
	PrefixMintableMetadataIds = []byte{0x03}
)