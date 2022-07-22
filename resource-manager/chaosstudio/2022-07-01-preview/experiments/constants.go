package experiments

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SelectorType string

const (
	SelectorTypeList    SelectorType = "List"
	SelectorTypePercent SelectorType = "Percent"
	SelectorTypeRandom  SelectorType = "Random"
	SelectorTypeTag     SelectorType = "Tag"
)

func PossibleValuesForSelectorType() []string {
	return []string{
		string(SelectorTypeList),
		string(SelectorTypePercent),
		string(SelectorTypeRandom),
		string(SelectorTypeTag),
	}
}

func parseSelectorType(input string) (*SelectorType, error) {
	vals := map[string]SelectorType{
		"list":    SelectorTypeList,
		"percent": SelectorTypePercent,
		"random":  SelectorTypeRandom,
		"tag":     SelectorTypeTag,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SelectorType(input)
	return &out, nil
}

type TargetReferenceType string

const (
	TargetReferenceTypeChaosTarget TargetReferenceType = "ChaosTarget"
)

func PossibleValuesForTargetReferenceType() []string {
	return []string{
		string(TargetReferenceTypeChaosTarget),
	}
}

func parseTargetReferenceType(input string) (*TargetReferenceType, error) {
	vals := map[string]TargetReferenceType{
		"chaostarget": TargetReferenceTypeChaosTarget,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TargetReferenceType(input)
	return &out, nil
}
