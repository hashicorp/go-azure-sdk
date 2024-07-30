package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SynchronizationQuarantineReason string

const (
	SynchronizationQuarantineReason_EncounteredBaseEscrowThreshold       SynchronizationQuarantineReason = "EncounteredBaseEscrowThreshold"
	SynchronizationQuarantineReason_EncounteredEscrowProportionThreshold SynchronizationQuarantineReason = "EncounteredEscrowProportionThreshold"
	SynchronizationQuarantineReason_EncounteredQuarantineException       SynchronizationQuarantineReason = "EncounteredQuarantineException"
	SynchronizationQuarantineReason_EncounteredTotalEscrowThreshold      SynchronizationQuarantineReason = "EncounteredTotalEscrowThreshold"
	SynchronizationQuarantineReason_IngestionInterrupted                 SynchronizationQuarantineReason = "IngestionInterrupted"
	SynchronizationQuarantineReason_QuarantinedOnDemand                  SynchronizationQuarantineReason = "QuarantinedOnDemand"
	SynchronizationQuarantineReason_TooManyDeletes                       SynchronizationQuarantineReason = "TooManyDeletes"
	SynchronizationQuarantineReason_Unknown                              SynchronizationQuarantineReason = "Unknown"
)

func PossibleValuesForSynchronizationQuarantineReason() []string {
	return []string{
		string(SynchronizationQuarantineReason_EncounteredBaseEscrowThreshold),
		string(SynchronizationQuarantineReason_EncounteredEscrowProportionThreshold),
		string(SynchronizationQuarantineReason_EncounteredQuarantineException),
		string(SynchronizationQuarantineReason_EncounteredTotalEscrowThreshold),
		string(SynchronizationQuarantineReason_IngestionInterrupted),
		string(SynchronizationQuarantineReason_QuarantinedOnDemand),
		string(SynchronizationQuarantineReason_TooManyDeletes),
		string(SynchronizationQuarantineReason_Unknown),
	}
}

func (s *SynchronizationQuarantineReason) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSynchronizationQuarantineReason(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSynchronizationQuarantineReason(input string) (*SynchronizationQuarantineReason, error) {
	vals := map[string]SynchronizationQuarantineReason{
		"encounteredbaseescrowthreshold":       SynchronizationQuarantineReason_EncounteredBaseEscrowThreshold,
		"encounteredescrowproportionthreshold": SynchronizationQuarantineReason_EncounteredEscrowProportionThreshold,
		"encounteredquarantineexception":       SynchronizationQuarantineReason_EncounteredQuarantineException,
		"encounteredtotalescrowthreshold":      SynchronizationQuarantineReason_EncounteredTotalEscrowThreshold,
		"ingestioninterrupted":                 SynchronizationQuarantineReason_IngestionInterrupted,
		"quarantinedondemand":                  SynchronizationQuarantineReason_QuarantinedOnDemand,
		"toomanydeletes":                       SynchronizationQuarantineReason_TooManyDeletes,
		"unknown":                              SynchronizationQuarantineReason_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SynchronizationQuarantineReason(input)
	return &out, nil
}
