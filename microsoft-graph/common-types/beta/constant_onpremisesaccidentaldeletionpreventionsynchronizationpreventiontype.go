package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnPremisesAccidentalDeletionPreventionSynchronizationPreventionType string

const (
	OnPremisesAccidentalDeletionPreventionSynchronizationPreventionType_Disabled             OnPremisesAccidentalDeletionPreventionSynchronizationPreventionType = "disabled"
	OnPremisesAccidentalDeletionPreventionSynchronizationPreventionType_EnabledForCount      OnPremisesAccidentalDeletionPreventionSynchronizationPreventionType = "enabledForCount"
	OnPremisesAccidentalDeletionPreventionSynchronizationPreventionType_EnabledForPercentage OnPremisesAccidentalDeletionPreventionSynchronizationPreventionType = "enabledForPercentage"
)

func PossibleValuesForOnPremisesAccidentalDeletionPreventionSynchronizationPreventionType() []string {
	return []string{
		string(OnPremisesAccidentalDeletionPreventionSynchronizationPreventionType_Disabled),
		string(OnPremisesAccidentalDeletionPreventionSynchronizationPreventionType_EnabledForCount),
		string(OnPremisesAccidentalDeletionPreventionSynchronizationPreventionType_EnabledForPercentage),
	}
}

func (s *OnPremisesAccidentalDeletionPreventionSynchronizationPreventionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOnPremisesAccidentalDeletionPreventionSynchronizationPreventionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOnPremisesAccidentalDeletionPreventionSynchronizationPreventionType(input string) (*OnPremisesAccidentalDeletionPreventionSynchronizationPreventionType, error) {
	vals := map[string]OnPremisesAccidentalDeletionPreventionSynchronizationPreventionType{
		"disabled":             OnPremisesAccidentalDeletionPreventionSynchronizationPreventionType_Disabled,
		"enabledforcount":      OnPremisesAccidentalDeletionPreventionSynchronizationPreventionType_EnabledForCount,
		"enabledforpercentage": OnPremisesAccidentalDeletionPreventionSynchronizationPreventionType_EnabledForPercentage,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OnPremisesAccidentalDeletionPreventionSynchronizationPreventionType(input)
	return &out, nil
}
