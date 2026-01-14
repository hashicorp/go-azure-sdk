package tuningoptionsoperationgroup

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecommendationTypeEnum string

const (
	RecommendationTypeEnumAnalyzeTable RecommendationTypeEnum = "AnalyzeTable"
	RecommendationTypeEnumCreateIndex  RecommendationTypeEnum = "CreateIndex"
	RecommendationTypeEnumDropIndex    RecommendationTypeEnum = "DropIndex"
	RecommendationTypeEnumReIndex      RecommendationTypeEnum = "ReIndex"
)

func PossibleValuesForRecommendationTypeEnum() []string {
	return []string{
		string(RecommendationTypeEnumAnalyzeTable),
		string(RecommendationTypeEnumCreateIndex),
		string(RecommendationTypeEnumDropIndex),
		string(RecommendationTypeEnumReIndex),
	}
}

func (s *RecommendationTypeEnum) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRecommendationTypeEnum(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRecommendationTypeEnum(input string) (*RecommendationTypeEnum, error) {
	vals := map[string]RecommendationTypeEnum{
		"analyzetable": RecommendationTypeEnumAnalyzeTable,
		"createindex":  RecommendationTypeEnumCreateIndex,
		"dropindex":    RecommendationTypeEnumDropIndex,
		"reindex":      RecommendationTypeEnumReIndex,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecommendationTypeEnum(input)
	return &out, nil
}

type RecommendationTypeParameterEnum string

const (
	RecommendationTypeParameterEnumAnalyzeTable RecommendationTypeParameterEnum = "AnalyzeTable"
	RecommendationTypeParameterEnumCreateIndex  RecommendationTypeParameterEnum = "CreateIndex"
	RecommendationTypeParameterEnumDropIndex    RecommendationTypeParameterEnum = "DropIndex"
	RecommendationTypeParameterEnumReIndex      RecommendationTypeParameterEnum = "ReIndex"
)

func PossibleValuesForRecommendationTypeParameterEnum() []string {
	return []string{
		string(RecommendationTypeParameterEnumAnalyzeTable),
		string(RecommendationTypeParameterEnumCreateIndex),
		string(RecommendationTypeParameterEnumDropIndex),
		string(RecommendationTypeParameterEnumReIndex),
	}
}

func (s *RecommendationTypeParameterEnum) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRecommendationTypeParameterEnum(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRecommendationTypeParameterEnum(input string) (*RecommendationTypeParameterEnum, error) {
	vals := map[string]RecommendationTypeParameterEnum{
		"analyzetable": RecommendationTypeParameterEnumAnalyzeTable,
		"createindex":  RecommendationTypeParameterEnumCreateIndex,
		"dropindex":    RecommendationTypeParameterEnumDropIndex,
		"reindex":      RecommendationTypeParameterEnumReIndex,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecommendationTypeParameterEnum(input)
	return &out, nil
}

type TuningOptionParameterEnum string

const (
	TuningOptionParameterEnumIndex TuningOptionParameterEnum = "index"
	TuningOptionParameterEnumTable TuningOptionParameterEnum = "table"
)

func PossibleValuesForTuningOptionParameterEnum() []string {
	return []string{
		string(TuningOptionParameterEnumIndex),
		string(TuningOptionParameterEnumTable),
	}
}

func (s *TuningOptionParameterEnum) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTuningOptionParameterEnum(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTuningOptionParameterEnum(input string) (*TuningOptionParameterEnum, error) {
	vals := map[string]TuningOptionParameterEnum{
		"index": TuningOptionParameterEnumIndex,
		"table": TuningOptionParameterEnumTable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TuningOptionParameterEnum(input)
	return &out, nil
}
