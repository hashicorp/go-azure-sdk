package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureADJoinPolicyAppliesTo string

const (
	AzureADJoinPolicyAppliesTo_All      AzureADJoinPolicyAppliesTo = "all"
	AzureADJoinPolicyAppliesTo_None     AzureADJoinPolicyAppliesTo = "none"
	AzureADJoinPolicyAppliesTo_Selected AzureADJoinPolicyAppliesTo = "selected"
)

func PossibleValuesForAzureADJoinPolicyAppliesTo() []string {
	return []string{
		string(AzureADJoinPolicyAppliesTo_All),
		string(AzureADJoinPolicyAppliesTo_None),
		string(AzureADJoinPolicyAppliesTo_Selected),
	}
}

func (s *AzureADJoinPolicyAppliesTo) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAzureADJoinPolicyAppliesTo(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAzureADJoinPolicyAppliesTo(input string) (*AzureADJoinPolicyAppliesTo, error) {
	vals := map[string]AzureADJoinPolicyAppliesTo{
		"all":      AzureADJoinPolicyAppliesTo_All,
		"none":     AzureADJoinPolicyAppliesTo_None,
		"selected": AzureADJoinPolicyAppliesTo_Selected,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AzureADJoinPolicyAppliesTo(input)
	return &out, nil
}
