package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BitLockerSystemDrivePolicyStartupAuthenticationTpmPinAndKeyUsage string

const (
	BitLockerSystemDrivePolicyStartupAuthenticationTpmPinAndKeyUsageallowed       BitLockerSystemDrivePolicyStartupAuthenticationTpmPinAndKeyUsage = "Allowed"
	BitLockerSystemDrivePolicyStartupAuthenticationTpmPinAndKeyUsageblocked       BitLockerSystemDrivePolicyStartupAuthenticationTpmPinAndKeyUsage = "Blocked"
	BitLockerSystemDrivePolicyStartupAuthenticationTpmPinAndKeyUsagenotConfigured BitLockerSystemDrivePolicyStartupAuthenticationTpmPinAndKeyUsage = "NotConfigured"
	BitLockerSystemDrivePolicyStartupAuthenticationTpmPinAndKeyUsagerequired      BitLockerSystemDrivePolicyStartupAuthenticationTpmPinAndKeyUsage = "Required"
)

func PossibleValuesForBitLockerSystemDrivePolicyStartupAuthenticationTpmPinAndKeyUsage() []string {
	return []string{
		string(BitLockerSystemDrivePolicyStartupAuthenticationTpmPinAndKeyUsageallowed),
		string(BitLockerSystemDrivePolicyStartupAuthenticationTpmPinAndKeyUsageblocked),
		string(BitLockerSystemDrivePolicyStartupAuthenticationTpmPinAndKeyUsagenotConfigured),
		string(BitLockerSystemDrivePolicyStartupAuthenticationTpmPinAndKeyUsagerequired),
	}
}

func parseBitLockerSystemDrivePolicyStartupAuthenticationTpmPinAndKeyUsage(input string) (*BitLockerSystemDrivePolicyStartupAuthenticationTpmPinAndKeyUsage, error) {
	vals := map[string]BitLockerSystemDrivePolicyStartupAuthenticationTpmPinAndKeyUsage{
		"allowed":       BitLockerSystemDrivePolicyStartupAuthenticationTpmPinAndKeyUsageallowed,
		"blocked":       BitLockerSystemDrivePolicyStartupAuthenticationTpmPinAndKeyUsageblocked,
		"notconfigured": BitLockerSystemDrivePolicyStartupAuthenticationTpmPinAndKeyUsagenotConfigured,
		"required":      BitLockerSystemDrivePolicyStartupAuthenticationTpmPinAndKeyUsagerequired,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BitLockerSystemDrivePolicyStartupAuthenticationTpmPinAndKeyUsage(input)
	return &out, nil
}
