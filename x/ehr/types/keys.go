package types

const (
	// ModuleName defines the module name
	ModuleName = "ehr"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_ehr"

	// Version defines the current version the IBC module supports
	Version = "ehr-1"

	// PortID is the default port id that module binds to
	PortID = "ehr"

	EhrIdKey = "ehr-id"
	EhrKey   = "ehr-record"

	PendingTransferUrl         = "http://localhost:1319/"
	OrganizationPrivateKeyPath = "privatekey.pem"
	OrganizationAddress        = "master14rt2mngt5lrvtsw3s77uk90hs5j30xlk5mgms0"
)

var (
	// PortKey defines the key to store the port ID in store
	PortKey = KeyPrefix("ehr-port-")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
