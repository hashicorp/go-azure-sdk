package sqlpoolssensitivitylabels

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecommendedSensitivityLabelUpdateKind string

const (
	RecommendedSensitivityLabelUpdateKindDisable RecommendedSensitivityLabelUpdateKind = "disable"
	RecommendedSensitivityLabelUpdateKindEnable  RecommendedSensitivityLabelUpdateKind = "enable"
)

func PossibleValuesForRecommendedSensitivityLabelUpdateKind() []string {
	return []string{
		string(RecommendedSensitivityLabelUpdateKindDisable),
		string(RecommendedSensitivityLabelUpdateKindEnable),
	}
}

func parseRecommendedSensitivityLabelUpdateKind(input string) (*RecommendedSensitivityLabelUpdateKind, error) {
	vals := map[string]RecommendedSensitivityLabelUpdateKind{
		"disable": RecommendedSensitivityLabelUpdateKindDisable,
		"enable":  RecommendedSensitivityLabelUpdateKindEnable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecommendedSensitivityLabelUpdateKind(input)
	return &out, nil
}

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

type SensitivityLabelSource string

const (
	SensitivityLabelSourceCurrent     SensitivityLabelSource = "current"
	SensitivityLabelSourceRecommended SensitivityLabelSource = "recommended"
)

func PossibleValuesForSensitivityLabelSource() []string {
	return []string{
		string(SensitivityLabelSourceCurrent),
		string(SensitivityLabelSourceRecommended),
	}
}

func parseSensitivityLabelSource(input string) (*SensitivityLabelSource, error) {
	vals := map[string]SensitivityLabelSource{
		"current":     SensitivityLabelSourceCurrent,
		"recommended": SensitivityLabelSourceRecommended,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SensitivityLabelSource(input)
	return &out, nil
}

type SensitivityLabelUpdateKind string

const (
	SensitivityLabelUpdateKindRemove SensitivityLabelUpdateKind = "remove"
	SensitivityLabelUpdateKindSet    SensitivityLabelUpdateKind = "set"
)

func PossibleValuesForSensitivityLabelUpdateKind() []string {
	return []string{
		string(SensitivityLabelUpdateKindRemove),
		string(SensitivityLabelUpdateKindSet),
	}
}

func parseSensitivityLabelUpdateKind(input string) (*SensitivityLabelUpdateKind, error) {
	vals := map[string]SensitivityLabelUpdateKind{
		"remove": SensitivityLabelUpdateKindRemove,
		"set":    SensitivityLabelUpdateKindSet,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SensitivityLabelUpdateKind(input)
	return &out, nil
}
