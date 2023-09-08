package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BitLockerSystemDrivePolicyStartupAuthenticationTpmKeyUsage string

const (
	BitLockerSystemDrivePolicyStartupAuthenticationTpmKeyUsageallowed       BitLockerSystemDrivePolicyStartupAuthenticationTpmKeyUsage = "Allowed"
	BitLockerSystemDrivePolicyStartupAuthenticationTpmKeyUsageblocked       BitLockerSystemDrivePolicyStartupAuthenticationTpmKeyUsage = "Blocked"
	BitLockerSystemDrivePolicyStartupAuthenticationTpmKeyUsagenotConfigured BitLockerSystemDrivePolicyStartupAuthenticationTpmKeyUsage = "NotConfigured"
	BitLockerSystemDrivePolicyStartupAuthenticationTpmKeyUsagerequired      BitLockerSystemDrivePolicyStartupAuthenticationTpmKeyUsage = "Required"
)

func PossibleValuesForBitLockerSystemDrivePolicyStartupAuthenticationTpmKeyUsage() []string {
	return []string{
		string(BitLockerSystemDrivePolicyStartupAuthenticationTpmKeyUsageallowed),
		string(BitLockerSystemDrivePolicyStartupAuthenticationTpmKeyUsageblocked),
		string(BitLockerSystemDrivePolicyStartupAuthenticationTpmKeyUsagenotConfigured),
		string(BitLockerSystemDrivePolicyStartupAuthenticationTpmKeyUsagerequired),
	}
}

func parseBitLockerSystemDrivePolicyStartupAuthenticationTpmKeyUsage(input string) (*BitLockerSystemDrivePolicyStartupAuthenticationTpmKeyUsage, error) {
	vals := map[string]BitLockerSystemDrivePolicyStartupAuthenticationTpmKeyUsage{
		"allowed":       BitLockerSystemDrivePolicyStartupAuthenticationTpmKeyUsageallowed,
		"blocked":       BitLockerSystemDrivePolicyStartupAuthenticationTpmKeyUsageblocked,
		"notconfigured": BitLockerSystemDrivePolicyStartupAuthenticationTpmKeyUsagenotConfigured,
		"required":      BitLockerSystemDrivePolicyStartupAuthenticationTpmKeyUsagerequired,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BitLockerSystemDrivePolicyStartupAuthenticationTpmKeyUsage(input)
	return &out, nil
}
