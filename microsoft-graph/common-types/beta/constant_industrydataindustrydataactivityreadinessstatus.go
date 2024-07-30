package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataIndustryDataActivityReadinessStatus string

const (
	IndustryDataIndustryDataActivityReadinessStatus_Disabled IndustryDataIndustryDataActivityReadinessStatus = "disabled"
	IndustryDataIndustryDataActivityReadinessStatus_Expired  IndustryDataIndustryDataActivityReadinessStatus = "expired"
	IndustryDataIndustryDataActivityReadinessStatus_Failed   IndustryDataIndustryDataActivityReadinessStatus = "failed"
	IndustryDataIndustryDataActivityReadinessStatus_NotReady IndustryDataIndustryDataActivityReadinessStatus = "notReady"
	IndustryDataIndustryDataActivityReadinessStatus_Ready    IndustryDataIndustryDataActivityReadinessStatus = "ready"
)

func PossibleValuesForIndustryDataIndustryDataActivityReadinessStatus() []string {
	return []string{
		string(IndustryDataIndustryDataActivityReadinessStatus_Disabled),
		string(IndustryDataIndustryDataActivityReadinessStatus_Expired),
		string(IndustryDataIndustryDataActivityReadinessStatus_Failed),
		string(IndustryDataIndustryDataActivityReadinessStatus_NotReady),
		string(IndustryDataIndustryDataActivityReadinessStatus_Ready),
	}
}

func (s *IndustryDataIndustryDataActivityReadinessStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIndustryDataIndustryDataActivityReadinessStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIndustryDataIndustryDataActivityReadinessStatus(input string) (*IndustryDataIndustryDataActivityReadinessStatus, error) {
	vals := map[string]IndustryDataIndustryDataActivityReadinessStatus{
		"disabled": IndustryDataIndustryDataActivityReadinessStatus_Disabled,
		"expired":  IndustryDataIndustryDataActivityReadinessStatus_Expired,
		"failed":   IndustryDataIndustryDataActivityReadinessStatus_Failed,
		"notready": IndustryDataIndustryDataActivityReadinessStatus_NotReady,
		"ready":    IndustryDataIndustryDataActivityReadinessStatus_Ready,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IndustryDataIndustryDataActivityReadinessStatus(input)
	return &out, nil
}
