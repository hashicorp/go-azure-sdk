package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BitLockerRecoveryOptionsRecoveryPasswordUsage string

const (
	BitLockerRecoveryOptionsRecoveryPasswordUsageallowed       BitLockerRecoveryOptionsRecoveryPasswordUsage = "Allowed"
	BitLockerRecoveryOptionsRecoveryPasswordUsageblocked       BitLockerRecoveryOptionsRecoveryPasswordUsage = "Blocked"
	BitLockerRecoveryOptionsRecoveryPasswordUsagenotConfigured BitLockerRecoveryOptionsRecoveryPasswordUsage = "NotConfigured"
	BitLockerRecoveryOptionsRecoveryPasswordUsagerequired      BitLockerRecoveryOptionsRecoveryPasswordUsage = "Required"
)

func PossibleValuesForBitLockerRecoveryOptionsRecoveryPasswordUsage() []string {
	return []string{
		string(BitLockerRecoveryOptionsRecoveryPasswordUsageallowed),
		string(BitLockerRecoveryOptionsRecoveryPasswordUsageblocked),
		string(BitLockerRecoveryOptionsRecoveryPasswordUsagenotConfigured),
		string(BitLockerRecoveryOptionsRecoveryPasswordUsagerequired),
	}
}

func parseBitLockerRecoveryOptionsRecoveryPasswordUsage(input string) (*BitLockerRecoveryOptionsRecoveryPasswordUsage, error) {
	vals := map[string]BitLockerRecoveryOptionsRecoveryPasswordUsage{
		"allowed":       BitLockerRecoveryOptionsRecoveryPasswordUsageallowed,
		"blocked":       BitLockerRecoveryOptionsRecoveryPasswordUsageblocked,
		"notconfigured": BitLockerRecoveryOptionsRecoveryPasswordUsagenotConfigured,
		"required":      BitLockerRecoveryOptionsRecoveryPasswordUsagerequired,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BitLockerRecoveryOptionsRecoveryPasswordUsage(input)
	return &out, nil
}
