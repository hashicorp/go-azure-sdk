package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BitLockerSystemDrivePolicyStartupAuthenticationTpmPinUsage string

const (
	BitLockerSystemDrivePolicyStartupAuthenticationTpmPinUsageallowed       BitLockerSystemDrivePolicyStartupAuthenticationTpmPinUsage = "Allowed"
	BitLockerSystemDrivePolicyStartupAuthenticationTpmPinUsageblocked       BitLockerSystemDrivePolicyStartupAuthenticationTpmPinUsage = "Blocked"
	BitLockerSystemDrivePolicyStartupAuthenticationTpmPinUsagenotConfigured BitLockerSystemDrivePolicyStartupAuthenticationTpmPinUsage = "NotConfigured"
	BitLockerSystemDrivePolicyStartupAuthenticationTpmPinUsagerequired      BitLockerSystemDrivePolicyStartupAuthenticationTpmPinUsage = "Required"
)

func PossibleValuesForBitLockerSystemDrivePolicyStartupAuthenticationTpmPinUsage() []string {
	return []string{
		string(BitLockerSystemDrivePolicyStartupAuthenticationTpmPinUsageallowed),
		string(BitLockerSystemDrivePolicyStartupAuthenticationTpmPinUsageblocked),
		string(BitLockerSystemDrivePolicyStartupAuthenticationTpmPinUsagenotConfigured),
		string(BitLockerSystemDrivePolicyStartupAuthenticationTpmPinUsagerequired),
	}
}

func parseBitLockerSystemDrivePolicyStartupAuthenticationTpmPinUsage(input string) (*BitLockerSystemDrivePolicyStartupAuthenticationTpmPinUsage, error) {
	vals := map[string]BitLockerSystemDrivePolicyStartupAuthenticationTpmPinUsage{
		"allowed":       BitLockerSystemDrivePolicyStartupAuthenticationTpmPinUsageallowed,
		"blocked":       BitLockerSystemDrivePolicyStartupAuthenticationTpmPinUsageblocked,
		"notconfigured": BitLockerSystemDrivePolicyStartupAuthenticationTpmPinUsagenotConfigured,
		"required":      BitLockerSystemDrivePolicyStartupAuthenticationTpmPinUsagerequired,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BitLockerSystemDrivePolicyStartupAuthenticationTpmPinUsage(input)
	return &out, nil
}
