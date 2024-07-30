package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataIndustryDataRunActivityStatus string

const (
	IndustryDataIndustryDataRunActivityStatus_Completed             IndustryDataIndustryDataRunActivityStatus = "completed"
	IndustryDataIndustryDataRunActivityStatus_CompletedWithErrors   IndustryDataIndustryDataRunActivityStatus = "completedWithErrors"
	IndustryDataIndustryDataRunActivityStatus_CompletedWithWarnings IndustryDataIndustryDataRunActivityStatus = "completedWithWarnings"
	IndustryDataIndustryDataRunActivityStatus_Failed                IndustryDataIndustryDataRunActivityStatus = "failed"
	IndustryDataIndustryDataRunActivityStatus_InProgress            IndustryDataIndustryDataRunActivityStatus = "inProgress"
	IndustryDataIndustryDataRunActivityStatus_Skipped               IndustryDataIndustryDataRunActivityStatus = "skipped"
)

func PossibleValuesForIndustryDataIndustryDataRunActivityStatus() []string {
	return []string{
		string(IndustryDataIndustryDataRunActivityStatus_Completed),
		string(IndustryDataIndustryDataRunActivityStatus_CompletedWithErrors),
		string(IndustryDataIndustryDataRunActivityStatus_CompletedWithWarnings),
		string(IndustryDataIndustryDataRunActivityStatus_Failed),
		string(IndustryDataIndustryDataRunActivityStatus_InProgress),
		string(IndustryDataIndustryDataRunActivityStatus_Skipped),
	}
}

func (s *IndustryDataIndustryDataRunActivityStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIndustryDataIndustryDataRunActivityStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIndustryDataIndustryDataRunActivityStatus(input string) (*IndustryDataIndustryDataRunActivityStatus, error) {
	vals := map[string]IndustryDataIndustryDataRunActivityStatus{
		"completed":             IndustryDataIndustryDataRunActivityStatus_Completed,
		"completedwitherrors":   IndustryDataIndustryDataRunActivityStatus_CompletedWithErrors,
		"completedwithwarnings": IndustryDataIndustryDataRunActivityStatus_CompletedWithWarnings,
		"failed":                IndustryDataIndustryDataRunActivityStatus_Failed,
		"inprogress":            IndustryDataIndustryDataRunActivityStatus_InProgress,
		"skipped":               IndustryDataIndustryDataRunActivityStatus_Skipped,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IndustryDataIndustryDataRunActivityStatus(input)
	return &out, nil
}
