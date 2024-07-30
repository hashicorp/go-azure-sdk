package informationprotectionpolicylabel

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContentInfoFormat string

const (
	ContentInfoFormatDefault ContentInfoFormat = "default"
	ContentInfoFormatEmail   ContentInfoFormat = "email"
)

func PossibleValuesForContentInfoFormat() []string {
	return []string{
		string(ContentInfoFormatDefault),
		string(ContentInfoFormatEmail),
	}
}

func (s *ContentInfoFormat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseContentInfoFormat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseContentInfoFormat(input string) (*ContentInfoFormat, error) {
	vals := map[string]ContentInfoFormat{
		"default": ContentInfoFormatDefault,
		"email":   ContentInfoFormatEmail,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ContentInfoFormat(input)
	return &out, nil
}

type ContentInfoState string

const (
	ContentInfoStateMotion ContentInfoState = "motion"
	ContentInfoStateRest   ContentInfoState = "rest"
	ContentInfoStateUse    ContentInfoState = "use"
)

func PossibleValuesForContentInfoState() []string {
	return []string{
		string(ContentInfoStateMotion),
		string(ContentInfoStateRest),
		string(ContentInfoStateUse),
	}
}

func (s *ContentInfoState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseContentInfoState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseContentInfoState(input string) (*ContentInfoState, error) {
	vals := map[string]ContentInfoState{
		"motion": ContentInfoStateMotion,
		"rest":   ContentInfoStateRest,
		"use":    ContentInfoStateUse,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ContentInfoState(input)
	return &out, nil
}

type LabelingOptionsAssignmentMethod string

const (
	LabelingOptionsAssignmentMethodAuto       LabelingOptionsAssignmentMethod = "auto"
	LabelingOptionsAssignmentMethodPrivileged LabelingOptionsAssignmentMethod = "privileged"
	LabelingOptionsAssignmentMethodStandard   LabelingOptionsAssignmentMethod = "standard"
)

func PossibleValuesForLabelingOptionsAssignmentMethod() []string {
	return []string{
		string(LabelingOptionsAssignmentMethodAuto),
		string(LabelingOptionsAssignmentMethodPrivileged),
		string(LabelingOptionsAssignmentMethodStandard),
	}
}

func (s *LabelingOptionsAssignmentMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLabelingOptionsAssignmentMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLabelingOptionsAssignmentMethod(input string) (*LabelingOptionsAssignmentMethod, error) {
	vals := map[string]LabelingOptionsAssignmentMethod{
		"auto":       LabelingOptionsAssignmentMethodAuto,
		"privileged": LabelingOptionsAssignmentMethodPrivileged,
		"standard":   LabelingOptionsAssignmentMethodStandard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LabelingOptionsAssignmentMethod(input)
	return &out, nil
}
