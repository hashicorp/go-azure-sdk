package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataFileValidateOperationStatus string

const (
	IndustryDataFileValidateOperationStatus_Failed     IndustryDataFileValidateOperationStatus = "failed"
	IndustryDataFileValidateOperationStatus_NotStarted IndustryDataFileValidateOperationStatus = "notStarted"
	IndustryDataFileValidateOperationStatus_Running    IndustryDataFileValidateOperationStatus = "running"
	IndustryDataFileValidateOperationStatus_Succeeded  IndustryDataFileValidateOperationStatus = "succeeded"
)

func PossibleValuesForIndustryDataFileValidateOperationStatus() []string {
	return []string{
		string(IndustryDataFileValidateOperationStatus_Failed),
		string(IndustryDataFileValidateOperationStatus_NotStarted),
		string(IndustryDataFileValidateOperationStatus_Running),
		string(IndustryDataFileValidateOperationStatus_Succeeded),
	}
}

func (s *IndustryDataFileValidateOperationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIndustryDataFileValidateOperationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIndustryDataFileValidateOperationStatus(input string) (*IndustryDataFileValidateOperationStatus, error) {
	vals := map[string]IndustryDataFileValidateOperationStatus{
		"failed":     IndustryDataFileValidateOperationStatus_Failed,
		"notstarted": IndustryDataFileValidateOperationStatus_NotStarted,
		"running":    IndustryDataFileValidateOperationStatus_Running,
		"succeeded":  IndustryDataFileValidateOperationStatus_Succeeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IndustryDataFileValidateOperationStatus(input)
	return &out, nil
}
