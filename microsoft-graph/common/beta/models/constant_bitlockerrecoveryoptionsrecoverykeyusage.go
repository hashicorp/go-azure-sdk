package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BitLockerRecoveryOptionsRecoveryKeyUsage string

const (
	BitLockerRecoveryOptionsRecoveryKeyUsageallowed       BitLockerRecoveryOptionsRecoveryKeyUsage = "Allowed"
	BitLockerRecoveryOptionsRecoveryKeyUsageblocked       BitLockerRecoveryOptionsRecoveryKeyUsage = "Blocked"
	BitLockerRecoveryOptionsRecoveryKeyUsagenotConfigured BitLockerRecoveryOptionsRecoveryKeyUsage = "NotConfigured"
	BitLockerRecoveryOptionsRecoveryKeyUsagerequired      BitLockerRecoveryOptionsRecoveryKeyUsage = "Required"
)

func PossibleValuesForBitLockerRecoveryOptionsRecoveryKeyUsage() []string {
	return []string{
		string(BitLockerRecoveryOptionsRecoveryKeyUsageallowed),
		string(BitLockerRecoveryOptionsRecoveryKeyUsageblocked),
		string(BitLockerRecoveryOptionsRecoveryKeyUsagenotConfigured),
		string(BitLockerRecoveryOptionsRecoveryKeyUsagerequired),
	}
}

func parseBitLockerRecoveryOptionsRecoveryKeyUsage(input string) (*BitLockerRecoveryOptionsRecoveryKeyUsage, error) {
	vals := map[string]BitLockerRecoveryOptionsRecoveryKeyUsage{
		"allowed":       BitLockerRecoveryOptionsRecoveryKeyUsageallowed,
		"blocked":       BitLockerRecoveryOptionsRecoveryKeyUsageblocked,
		"notconfigured": BitLockerRecoveryOptionsRecoveryKeyUsagenotConfigured,
		"required":      BitLockerRecoveryOptionsRecoveryKeyUsagerequired,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BitLockerRecoveryOptionsRecoveryKeyUsage(input)
	return &out, nil
}
