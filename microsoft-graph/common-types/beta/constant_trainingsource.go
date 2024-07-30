package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TrainingSource string

const (
	TrainingSource_Global  TrainingSource = "global"
	TrainingSource_Tenant  TrainingSource = "tenant"
	TrainingSource_Unknown TrainingSource = "unknown"
)

func PossibleValuesForTrainingSource() []string {
	return []string{
		string(TrainingSource_Global),
		string(TrainingSource_Tenant),
		string(TrainingSource_Unknown),
	}
}

func (s *TrainingSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTrainingSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTrainingSource(input string) (*TrainingSource, error) {
	vals := map[string]TrainingSource{
		"global":  TrainingSource_Global,
		"tenant":  TrainingSource_Tenant,
		"unknown": TrainingSource_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TrainingSource(input)
	return &out, nil
}
