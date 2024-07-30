package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataValidateOperationStatus string

const (
	IndustryDataValidateOperationStatus_Failed     IndustryDataValidateOperationStatus = "failed"
	IndustryDataValidateOperationStatus_NotStarted IndustryDataValidateOperationStatus = "notStarted"
	IndustryDataValidateOperationStatus_Running    IndustryDataValidateOperationStatus = "running"
	IndustryDataValidateOperationStatus_Succeeded  IndustryDataValidateOperationStatus = "succeeded"
)

func PossibleValuesForIndustryDataValidateOperationStatus() []string {
	return []string{
		string(IndustryDataValidateOperationStatus_Failed),
		string(IndustryDataValidateOperationStatus_NotStarted),
		string(IndustryDataValidateOperationStatus_Running),
		string(IndustryDataValidateOperationStatus_Succeeded),
	}
}

func (s *IndustryDataValidateOperationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIndustryDataValidateOperationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIndustryDataValidateOperationStatus(input string) (*IndustryDataValidateOperationStatus, error) {
	vals := map[string]IndustryDataValidateOperationStatus{
		"failed":     IndustryDataValidateOperationStatus_Failed,
		"notstarted": IndustryDataValidateOperationStatus_NotStarted,
		"running":    IndustryDataValidateOperationStatus_Running,
		"succeeded":  IndustryDataValidateOperationStatus_Succeeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IndustryDataValidateOperationStatus(input)
	return &out, nil
}
