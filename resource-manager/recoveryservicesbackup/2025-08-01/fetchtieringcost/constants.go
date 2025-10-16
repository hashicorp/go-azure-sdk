package fetchtieringcost

import (
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecoveryPointTierType string

const (
	RecoveryPointTierTypeArchivedRP RecoveryPointTierType = "ArchivedRP"
	RecoveryPointTierTypeHardenedRP RecoveryPointTierType = "HardenedRP"
	RecoveryPointTierTypeInstantRP  RecoveryPointTierType = "InstantRP"
	RecoveryPointTierTypeInvalid    RecoveryPointTierType = "Invalid"
)

func PossibleValuesForRecoveryPointTierType() []string {
	return []string{
		string(RecoveryPointTierTypeArchivedRP),
		string(RecoveryPointTierTypeHardenedRP),
		string(RecoveryPointTierTypeInstantRP),
		string(RecoveryPointTierTypeInvalid),
	}
}

func parseRecoveryPointTierType(input string) (*RecoveryPointTierType, error) {
	vals := map[string]RecoveryPointTierType{
		"archivedrp": RecoveryPointTierTypeArchivedRP,
		"hardenedrp": RecoveryPointTierTypeHardenedRP,
		"instantrp":  RecoveryPointTierTypeInstantRP,
		"invalid":    RecoveryPointTierTypeInvalid,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecoveryPointTierType(input)
	return &out, nil
}

type RehydrationPriority string

const (
	RehydrationPriorityHigh     RehydrationPriority = "High"
	RehydrationPriorityStandard RehydrationPriority = "Standard"
)

func PossibleValuesForRehydrationPriority() []string {
	return []string{
		string(RehydrationPriorityHigh),
		string(RehydrationPriorityStandard),
	}
}

func parseRehydrationPriority(input string) (*RehydrationPriority, error) {
	vals := map[string]RehydrationPriority{
		"high":     RehydrationPriorityHigh,
		"standard": RehydrationPriorityStandard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RehydrationPriority(input)
	return &out, nil
}
