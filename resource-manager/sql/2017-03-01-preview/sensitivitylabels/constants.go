package sensitivitylabels

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SensitivityLabelRank string

const (
	SensitivityLabelRankCritical SensitivityLabelRank = "Critical"
	SensitivityLabelRankHigh     SensitivityLabelRank = "High"
	SensitivityLabelRankLow      SensitivityLabelRank = "Low"
	SensitivityLabelRankMedium   SensitivityLabelRank = "Medium"
	SensitivityLabelRankNone     SensitivityLabelRank = "None"
)

func PossibleValuesForSensitivityLabelRank() []string {
	return []string{
		string(SensitivityLabelRankCritical),
		string(SensitivityLabelRankHigh),
		string(SensitivityLabelRankLow),
		string(SensitivityLabelRankMedium),
		string(SensitivityLabelRankNone),
	}
}

func (s *SensitivityLabelRank) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSensitivityLabelRank(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSensitivityLabelRank(input string) (*SensitivityLabelRank, error) {
	vals := map[string]SensitivityLabelRank{
		"critical": SensitivityLabelRankCritical,
		"high":     SensitivityLabelRankHigh,
		"low":      SensitivityLabelRankLow,
		"medium":   SensitivityLabelRankMedium,
		"none":     SensitivityLabelRankNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SensitivityLabelRank(input)
	return &out, nil
}
