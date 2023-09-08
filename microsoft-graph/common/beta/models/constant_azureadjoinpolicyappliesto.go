package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureAdJoinPolicyAppliesTo string

const (
	AzureAdJoinPolicyAppliesToall      AzureAdJoinPolicyAppliesTo = "All"
	AzureAdJoinPolicyAppliesTonone     AzureAdJoinPolicyAppliesTo = "None"
	AzureAdJoinPolicyAppliesToselected AzureAdJoinPolicyAppliesTo = "Selected"
)

func PossibleValuesForAzureAdJoinPolicyAppliesTo() []string {
	return []string{
		string(AzureAdJoinPolicyAppliesToall),
		string(AzureAdJoinPolicyAppliesTonone),
		string(AzureAdJoinPolicyAppliesToselected),
	}
}

func parseAzureAdJoinPolicyAppliesTo(input string) (*AzureAdJoinPolicyAppliesTo, error) {
	vals := map[string]AzureAdJoinPolicyAppliesTo{
		"all":      AzureAdJoinPolicyAppliesToall,
		"none":     AzureAdJoinPolicyAppliesTonone,
		"selected": AzureAdJoinPolicyAppliesToselected,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AzureAdJoinPolicyAppliesTo(input)
	return &out, nil
}
