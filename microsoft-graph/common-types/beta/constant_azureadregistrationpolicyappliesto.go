package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureADRegistrationPolicyAppliesTo string

const (
	AzureADRegistrationPolicyAppliesTo_All      AzureADRegistrationPolicyAppliesTo = "all"
	AzureADRegistrationPolicyAppliesTo_None     AzureADRegistrationPolicyAppliesTo = "none"
	AzureADRegistrationPolicyAppliesTo_Selected AzureADRegistrationPolicyAppliesTo = "selected"
)

func PossibleValuesForAzureADRegistrationPolicyAppliesTo() []string {
	return []string{
		string(AzureADRegistrationPolicyAppliesTo_All),
		string(AzureADRegistrationPolicyAppliesTo_None),
		string(AzureADRegistrationPolicyAppliesTo_Selected),
	}
}

func (s *AzureADRegistrationPolicyAppliesTo) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAzureADRegistrationPolicyAppliesTo(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAzureADRegistrationPolicyAppliesTo(input string) (*AzureADRegistrationPolicyAppliesTo, error) {
	vals := map[string]AzureADRegistrationPolicyAppliesTo{
		"all":      AzureADRegistrationPolicyAppliesTo_All,
		"none":     AzureADRegistrationPolicyAppliesTo_None,
		"selected": AzureADRegistrationPolicyAppliesTo_Selected,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AzureADRegistrationPolicyAppliesTo(input)
	return &out, nil
}
