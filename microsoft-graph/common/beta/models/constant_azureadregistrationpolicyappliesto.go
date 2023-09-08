package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureADRegistrationPolicyAppliesTo string

const (
	AzureADRegistrationPolicyAppliesToall      AzureADRegistrationPolicyAppliesTo = "All"
	AzureADRegistrationPolicyAppliesTonone     AzureADRegistrationPolicyAppliesTo = "None"
	AzureADRegistrationPolicyAppliesToselected AzureADRegistrationPolicyAppliesTo = "Selected"
)

func PossibleValuesForAzureADRegistrationPolicyAppliesTo() []string {
	return []string{
		string(AzureADRegistrationPolicyAppliesToall),
		string(AzureADRegistrationPolicyAppliesTonone),
		string(AzureADRegistrationPolicyAppliesToselected),
	}
}

func parseAzureADRegistrationPolicyAppliesTo(input string) (*AzureADRegistrationPolicyAppliesTo, error) {
	vals := map[string]AzureADRegistrationPolicyAppliesTo{
		"all":      AzureADRegistrationPolicyAppliesToall,
		"none":     AzureADRegistrationPolicyAppliesTonone,
		"selected": AzureADRegistrationPolicyAppliesToselected,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AzureADRegistrationPolicyAppliesTo(input)
	return &out, nil
}
