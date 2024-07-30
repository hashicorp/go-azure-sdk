package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceUpdateMessageCategory string

const (
	ServiceUpdateMessageCategory_PlanForChange     ServiceUpdateMessageCategory = "planForChange"
	ServiceUpdateMessageCategory_PreventOrFixIssue ServiceUpdateMessageCategory = "preventOrFixIssue"
	ServiceUpdateMessageCategory_StayInformed      ServiceUpdateMessageCategory = "stayInformed"
)

func PossibleValuesForServiceUpdateMessageCategory() []string {
	return []string{
		string(ServiceUpdateMessageCategory_PlanForChange),
		string(ServiceUpdateMessageCategory_PreventOrFixIssue),
		string(ServiceUpdateMessageCategory_StayInformed),
	}
}

func (s *ServiceUpdateMessageCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseServiceUpdateMessageCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseServiceUpdateMessageCategory(input string) (*ServiceUpdateMessageCategory, error) {
	vals := map[string]ServiceUpdateMessageCategory{
		"planforchange":     ServiceUpdateMessageCategory_PlanForChange,
		"preventorfixissue": ServiceUpdateMessageCategory_PreventOrFixIssue,
		"stayinformed":      ServiceUpdateMessageCategory_StayInformed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServiceUpdateMessageCategory(input)
	return &out, nil
}
