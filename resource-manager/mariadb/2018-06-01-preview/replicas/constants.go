package replicas

import "strings"

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

func parseGeoRedundantBackup(input string) (*GeoRedundantBackup, error) {
	vals := map[string]GeoRedundantBackup{
		"disabled": GeoRedundantBackupDisabled,
		"enabled":  GeoRedundantBackupEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GeoRedundantBackup(input)
	return &out, nil
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

func parseMinimalTlsVersionEnum(input string) (*MinimalTlsVersionEnum, error) {
	vals := map[string]MinimalTlsVersionEnum{
		"tlsenforcementdisabled": MinimalTlsVersionEnumTLSEnforcementDisabled,
		"tls1_1":                 MinimalTlsVersionEnumTLSOneOne,
		"tls1_2":                 MinimalTlsVersionEnumTLSOneTwo,
		"tls1_0":                 MinimalTlsVersionEnumTLSOneZero,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MinimalTlsVersionEnum(input)
	return &out, nil
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

func parseServerState(input string) (*ServerState, error) {
	vals := map[string]ServerState{
		"disabled": ServerStateDisabled,
		"dropping": ServerStateDropping,
		"ready":    ServerStateReady,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServerState(input)
	return &out, nil
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

func parseServerVersion(input string) (*ServerVersion, error) {
	vals := map[string]ServerVersion{
		"10.3": ServerVersionOneZeroPointThree,
		"10.2": ServerVersionOneZeroPointTwo,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServerVersion(input)
	return &out, nil
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

func parseSkuTier(input string) (*SkuTier, error) {
	vals := map[string]SkuTier{
		"basic":           SkuTierBasic,
		"generalpurpose":  SkuTierGeneralPurpose,
		"memoryoptimized": SkuTierMemoryOptimized,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SkuTier(input)
	return &out, nil
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

func parseSslEnforcementEnum(input string) (*SslEnforcementEnum, error) {
	vals := map[string]SslEnforcementEnum{
		"disabled": SslEnforcementEnumDisabled,
		"enabled":  SslEnforcementEnumEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SslEnforcementEnum(input)
	return &out, nil
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

func parseStorageAutogrow(input string) (*StorageAutogrow, error) {
	vals := map[string]StorageAutogrow{
		"disabled": StorageAutogrowDisabled,
		"enabled":  StorageAutogrowEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StorageAutogrow(input)
	return &out, nil
}
