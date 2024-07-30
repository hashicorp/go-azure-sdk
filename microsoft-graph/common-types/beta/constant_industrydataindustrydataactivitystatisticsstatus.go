package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataIndustryDataActivityStatisticsStatus string

const (
	IndustryDataIndustryDataActivityStatisticsStatus_Completed             IndustryDataIndustryDataActivityStatisticsStatus = "completed"
	IndustryDataIndustryDataActivityStatisticsStatus_CompletedWithErrors   IndustryDataIndustryDataActivityStatisticsStatus = "completedWithErrors"
	IndustryDataIndustryDataActivityStatisticsStatus_CompletedWithWarnings IndustryDataIndustryDataActivityStatisticsStatus = "completedWithWarnings"
	IndustryDataIndustryDataActivityStatisticsStatus_Failed                IndustryDataIndustryDataActivityStatisticsStatus = "failed"
	IndustryDataIndustryDataActivityStatisticsStatus_InProgress            IndustryDataIndustryDataActivityStatisticsStatus = "inProgress"
	IndustryDataIndustryDataActivityStatisticsStatus_Skipped               IndustryDataIndustryDataActivityStatisticsStatus = "skipped"
)

func PossibleValuesForIndustryDataIndustryDataActivityStatisticsStatus() []string {
	return []string{
		string(IndustryDataIndustryDataActivityStatisticsStatus_Completed),
		string(IndustryDataIndustryDataActivityStatisticsStatus_CompletedWithErrors),
		string(IndustryDataIndustryDataActivityStatisticsStatus_CompletedWithWarnings),
		string(IndustryDataIndustryDataActivityStatisticsStatus_Failed),
		string(IndustryDataIndustryDataActivityStatisticsStatus_InProgress),
		string(IndustryDataIndustryDataActivityStatisticsStatus_Skipped),
	}
}

func (s *IndustryDataIndustryDataActivityStatisticsStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIndustryDataIndustryDataActivityStatisticsStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIndustryDataIndustryDataActivityStatisticsStatus(input string) (*IndustryDataIndustryDataActivityStatisticsStatus, error) {
	vals := map[string]IndustryDataIndustryDataActivityStatisticsStatus{
		"completed":             IndustryDataIndustryDataActivityStatisticsStatus_Completed,
		"completedwitherrors":   IndustryDataIndustryDataActivityStatisticsStatus_CompletedWithErrors,
		"completedwithwarnings": IndustryDataIndustryDataActivityStatisticsStatus_CompletedWithWarnings,
		"failed":                IndustryDataIndustryDataActivityStatisticsStatus_Failed,
		"inprogress":            IndustryDataIndustryDataActivityStatisticsStatus_InProgress,
		"skipped":               IndustryDataIndustryDataActivityStatisticsStatus_Skipped,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IndustryDataIndustryDataActivityStatisticsStatus(input)
	return &out, nil
}
