package consts

const (
	// ExpirationRestoreWorkerCount specifies the number of workers to use while
	// restoring leases into the expiration manager
	ExpirationRestoreWorkerCount = 64

	// NamespaceHeaderName is the header set to specify which namespace the
	// request is indented for.
	NamespaceHeaderName = "X-Vault-Namespace"

	// AuthHeaderName is the name of the header containing the token.
	AuthHeaderName = "X-Vault-Token"

	// RequestHeaderName is the name of the header used by the Agent for
	// SSRF protection.
	RequestHeaderName = "X-Vault-Request"

	// PerformanceReplicationALPN is the negotiated protocol used for
	// performance replication.
	PerformanceReplicationALPN = "replication_v1"

	// DRReplicationALPN is the negotiated protocol used for dr replication.
	DRReplicationALPN = "replication_dr_v1"

	PerfStandbyALPN = "perf_standby_v1"

	RequestForwardingALPN = "req_fw_sb-act_v1"

	RaftStorageALPN = "raft_storage_v1"

	// ReplicationResolverALPN is the negotiated protocol used for
	// resolving replicaiton addresses
	ReplicationResolverALPN = "replication_resolver_v1"

	VaultDisableFilePermissionsCheckEnv = "VAULT_DISABLE_FILE_PERMISSIONS_CHECK"
)
