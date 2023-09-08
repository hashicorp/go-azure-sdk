package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SynchronizationStatusCode string

const (
	SynchronizationStatusCodeActive        SynchronizationStatusCode = "Active"
	SynchronizationStatusCodeNotConfigured SynchronizationStatusCode = "NotConfigured"
	SynchronizationStatusCodeNotRun        SynchronizationStatusCode = "NotRun"
	SynchronizationStatusCodePaused        SynchronizationStatusCode = "Paused"
	SynchronizationStatusCodeQuarantine    SynchronizationStatusCode = "Quarantine"
)

func PossibleValuesForSynchronizationStatusCode() []string {
	return []string{
		string(SynchronizationStatusCodeActive),
		string(SynchronizationStatusCodeNotConfigured),
		string(SynchronizationStatusCodeNotRun),
		string(SynchronizationStatusCodePaused),
		string(SynchronizationStatusCodeQuarantine),
	}
}

func parseSynchronizationStatusCode(input string) (*SynchronizationStatusCode, error) {
	vals := map[string]SynchronizationStatusCode{
		"active":        SynchronizationStatusCodeActive,
		"notconfigured": SynchronizationStatusCodeNotConfigured,
		"notrun":        SynchronizationStatusCodeNotRun,
		"paused":        SynchronizationStatusCodePaused,
		"quarantine":    SynchronizationStatusCodeQuarantine,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SynchronizationStatusCode(input)
	return &out, nil
}
