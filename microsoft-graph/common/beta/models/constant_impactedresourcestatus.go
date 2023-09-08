package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImpactedResourceStatus string

const (
	ImpactedResourceStatusactive            ImpactedResourceStatus = "Active"
	ImpactedResourceStatuscompletedBySystem ImpactedResourceStatus = "CompletedBySystem"
	ImpactedResourceStatuscompletedByUser   ImpactedResourceStatus = "CompletedByUser"
	ImpactedResourceStatusdismissed         ImpactedResourceStatus = "Dismissed"
	ImpactedResourceStatuspostponed         ImpactedResourceStatus = "Postponed"
)

func PossibleValuesForImpactedResourceStatus() []string {
	return []string{
		string(ImpactedResourceStatusactive),
		string(ImpactedResourceStatuscompletedBySystem),
		string(ImpactedResourceStatuscompletedByUser),
		string(ImpactedResourceStatusdismissed),
		string(ImpactedResourceStatuspostponed),
	}
}

func parseImpactedResourceStatus(input string) (*ImpactedResourceStatus, error) {
	vals := map[string]ImpactedResourceStatus{
		"active":            ImpactedResourceStatusactive,
		"completedbysystem": ImpactedResourceStatuscompletedBySystem,
		"completedbyuser":   ImpactedResourceStatuscompletedByUser,
		"dismissed":         ImpactedResourceStatusdismissed,
		"postponed":         ImpactedResourceStatuspostponed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ImpactedResourceStatus(input)
	return &out, nil
}
