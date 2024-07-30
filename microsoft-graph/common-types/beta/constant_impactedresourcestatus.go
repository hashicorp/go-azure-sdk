package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImpactedResourceStatus string

const (
	ImpactedResourceStatus_Active            ImpactedResourceStatus = "active"
	ImpactedResourceStatus_CompletedBySystem ImpactedResourceStatus = "completedBySystem"
	ImpactedResourceStatus_CompletedByUser   ImpactedResourceStatus = "completedByUser"
	ImpactedResourceStatus_Dismissed         ImpactedResourceStatus = "dismissed"
	ImpactedResourceStatus_Postponed         ImpactedResourceStatus = "postponed"
)

func PossibleValuesForImpactedResourceStatus() []string {
	return []string{
		string(ImpactedResourceStatus_Active),
		string(ImpactedResourceStatus_CompletedBySystem),
		string(ImpactedResourceStatus_CompletedByUser),
		string(ImpactedResourceStatus_Dismissed),
		string(ImpactedResourceStatus_Postponed),
	}
}

func (s *ImpactedResourceStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseImpactedResourceStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseImpactedResourceStatus(input string) (*ImpactedResourceStatus, error) {
	vals := map[string]ImpactedResourceStatus{
		"active":            ImpactedResourceStatus_Active,
		"completedbysystem": ImpactedResourceStatus_CompletedBySystem,
		"completedbyuser":   ImpactedResourceStatus_CompletedByUser,
		"dismissed":         ImpactedResourceStatus_Dismissed,
		"postponed":         ImpactedResourceStatus_Postponed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ImpactedResourceStatus(input)
	return &out, nil
}
