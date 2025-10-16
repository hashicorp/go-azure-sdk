package tuningoptions

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecommendationType string

const (
	RecommendationTypeAnalyzeTable RecommendationType = "AnalyzeTable"
	RecommendationTypeCreateIndex  RecommendationType = "CreateIndex"
	RecommendationTypeDropIndex    RecommendationType = "DropIndex"
	RecommendationTypeReIndex      RecommendationType = "ReIndex"
)

func PossibleValuesForRecommendationType() []string {
	return []string{
		string(RecommendationTypeAnalyzeTable),
		string(RecommendationTypeCreateIndex),
		string(RecommendationTypeDropIndex),
		string(RecommendationTypeReIndex),
	}
}

func (s *RecommendationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRecommendationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRecommendationType(input string) (*RecommendationType, error) {
	vals := map[string]RecommendationType{
		"analyzetable": RecommendationTypeAnalyzeTable,
		"createindex":  RecommendationTypeCreateIndex,
		"dropindex":    RecommendationTypeDropIndex,
		"reindex":      RecommendationTypeReIndex,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecommendationType(input)
	return &out, nil
}

type TuningOption string

const (
	TuningOptionIndex TuningOption = "index"
	TuningOptionTable TuningOption = "table"
)

func PossibleValuesForTuningOption() []string {
	return []string{
		string(TuningOptionIndex),
		string(TuningOptionTable),
	}
}

func (s *TuningOption) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTuningOption(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTuningOption(input string) (*TuningOption, error) {
	vals := map[string]TuningOption{
		"index": TuningOptionIndex,
		"table": TuningOptionTable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TuningOption(input)
	return &out, nil
}
