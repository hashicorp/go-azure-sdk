package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataIndustryDataRunStatisticsStatus string

const (
	IndustryDataIndustryDataRunStatisticsStatus_Completed             IndustryDataIndustryDataRunStatisticsStatus = "completed"
	IndustryDataIndustryDataRunStatisticsStatus_CompletedWithErrors   IndustryDataIndustryDataRunStatisticsStatus = "completedWithErrors"
	IndustryDataIndustryDataRunStatisticsStatus_CompletedWithWarnings IndustryDataIndustryDataRunStatisticsStatus = "completedWithWarnings"
	IndustryDataIndustryDataRunStatisticsStatus_Failed                IndustryDataIndustryDataRunStatisticsStatus = "failed"
	IndustryDataIndustryDataRunStatisticsStatus_Running               IndustryDataIndustryDataRunStatisticsStatus = "running"
)

func PossibleValuesForIndustryDataIndustryDataRunStatisticsStatus() []string {
	return []string{
		string(IndustryDataIndustryDataRunStatisticsStatus_Completed),
		string(IndustryDataIndustryDataRunStatisticsStatus_CompletedWithErrors),
		string(IndustryDataIndustryDataRunStatisticsStatus_CompletedWithWarnings),
		string(IndustryDataIndustryDataRunStatisticsStatus_Failed),
		string(IndustryDataIndustryDataRunStatisticsStatus_Running),
	}
}

func (s *IndustryDataIndustryDataRunStatisticsStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIndustryDataIndustryDataRunStatisticsStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIndustryDataIndustryDataRunStatisticsStatus(input string) (*IndustryDataIndustryDataRunStatisticsStatus, error) {
	vals := map[string]IndustryDataIndustryDataRunStatisticsStatus{
		"completed":             IndustryDataIndustryDataRunStatisticsStatus_Completed,
		"completedwitherrors":   IndustryDataIndustryDataRunStatisticsStatus_CompletedWithErrors,
		"completedwithwarnings": IndustryDataIndustryDataRunStatisticsStatus_CompletedWithWarnings,
		"failed":                IndustryDataIndustryDataRunStatisticsStatus_Failed,
		"running":               IndustryDataIndustryDataRunStatisticsStatus_Running,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IndustryDataIndustryDataRunStatisticsStatus(input)
	return &out, nil
}
