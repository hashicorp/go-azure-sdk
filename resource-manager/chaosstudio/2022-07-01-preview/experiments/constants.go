package experiments

import (
	"encoding/json"
	"fmt"
	"strings"
)

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

func (s *SelectorType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForSelectorType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = SelectorType(decoded)
	return nil
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

func (s *TargetReferenceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForTargetReferenceType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = TargetReferenceType(decoded)
	return nil
}
