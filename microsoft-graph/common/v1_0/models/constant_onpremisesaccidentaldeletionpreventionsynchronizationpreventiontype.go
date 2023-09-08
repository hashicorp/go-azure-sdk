package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnPremisesAccidentalDeletionPreventionSynchronizationPreventionType string

const (
	OnPremisesAccidentalDeletionPreventionSynchronizationPreventionTypedisabled             OnPremisesAccidentalDeletionPreventionSynchronizationPreventionType = "Disabled"
	OnPremisesAccidentalDeletionPreventionSynchronizationPreventionTypeenabledForCount      OnPremisesAccidentalDeletionPreventionSynchronizationPreventionType = "EnabledForCount"
	OnPremisesAccidentalDeletionPreventionSynchronizationPreventionTypeenabledForPercentage OnPremisesAccidentalDeletionPreventionSynchronizationPreventionType = "EnabledForPercentage"
)

func PossibleValuesForOnPremisesAccidentalDeletionPreventionSynchronizationPreventionType() []string {
	return []string{
		string(OnPremisesAccidentalDeletionPreventionSynchronizationPreventionTypedisabled),
		string(OnPremisesAccidentalDeletionPreventionSynchronizationPreventionTypeenabledForCount),
		string(OnPremisesAccidentalDeletionPreventionSynchronizationPreventionTypeenabledForPercentage),
	}
}

func parseOnPremisesAccidentalDeletionPreventionSynchronizationPreventionType(input string) (*OnPremisesAccidentalDeletionPreventionSynchronizationPreventionType, error) {
	vals := map[string]OnPremisesAccidentalDeletionPreventionSynchronizationPreventionType{
		"disabled":             OnPremisesAccidentalDeletionPreventionSynchronizationPreventionTypedisabled,
		"enabledforcount":      OnPremisesAccidentalDeletionPreventionSynchronizationPreventionTypeenabledForCount,
		"enabledforpercentage": OnPremisesAccidentalDeletionPreventionSynchronizationPreventionTypeenabledForPercentage,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OnPremisesAccidentalDeletionPreventionSynchronizationPreventionType(input)
	return &out, nil
}
