package replicas

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GeoRedundantBackup string

const (
	GeoRedundantBackupDisabled GeoRedundantBackup = "Disabled"
	GeoRedundantBackupEnabled  GeoRedundantBackup = "Enabled"
)

func PossibleValuesForGeoRedundantBackup() []string {
	return []string{
		string(GeoRedundantBackupDisabled),
		string(GeoRedundantBackupEnabled),
	}
}

type MinimalTlsVersionEnum string

const (
	MinimalTlsVersionEnumTLSEnforcementDisabled MinimalTlsVersionEnum = "TLSEnforcementDisabled"
	MinimalTlsVersionEnumTLSOneOne              MinimalTlsVersionEnum = "TLS1_1"
	MinimalTlsVersionEnumTLSOneTwo              MinimalTlsVersionEnum = "TLS1_2"
	MinimalTlsVersionEnumTLSOneZero             MinimalTlsVersionEnum = "TLS1_0"
)

func PossibleValuesForMinimalTlsVersionEnum() []string {
	return []string{
		string(MinimalTlsVersionEnumTLSEnforcementDisabled),
		string(MinimalTlsVersionEnumTLSOneOne),
		string(MinimalTlsVersionEnumTLSOneTwo),
		string(MinimalTlsVersionEnumTLSOneZero),
	}
}

type ServerState string

const (
	ServerStateDisabled ServerState = "Disabled"
	ServerStateDropping ServerState = "Dropping"
	ServerStateReady    ServerState = "Ready"
)

func PossibleValuesForServerState() []string {
	return []string{
		string(ServerStateDisabled),
		string(ServerStateDropping),
		string(ServerStateReady),
	}
}

type ServerVersion string

const (
	ServerVersionOneZeroPointThree ServerVersion = "10.3"
	ServerVersionOneZeroPointTwo   ServerVersion = "10.2"
)

func PossibleValuesForServerVersion() []string {
	return []string{
		string(ServerVersionOneZeroPointThree),
		string(ServerVersionOneZeroPointTwo),
	}
}

type SkuTier string

const (
	SkuTierBasic           SkuTier = "Basic"
	SkuTierGeneralPurpose  SkuTier = "GeneralPurpose"
	SkuTierMemoryOptimized SkuTier = "MemoryOptimized"
)

func PossibleValuesForSkuTier() []string {
	return []string{
		string(SkuTierBasic),
		string(SkuTierGeneralPurpose),
		string(SkuTierMemoryOptimized),
	}
}

type SslEnforcementEnum string

const (
	SslEnforcementEnumDisabled SslEnforcementEnum = "Disabled"
	SslEnforcementEnumEnabled  SslEnforcementEnum = "Enabled"
)

func PossibleValuesForSslEnforcementEnum() []string {
	return []string{
		string(SslEnforcementEnumDisabled),
		string(SslEnforcementEnumEnabled),
	}
}

type StorageAutogrow string

const (
	StorageAutogrowDisabled StorageAutogrow = "Disabled"
	StorageAutogrowEnabled  StorageAutogrow = "Enabled"
)

func PossibleValuesForStorageAutogrow() []string {
	return []string{
		string(StorageAutogrowDisabled),
		string(StorageAutogrowEnabled),
	}
}
