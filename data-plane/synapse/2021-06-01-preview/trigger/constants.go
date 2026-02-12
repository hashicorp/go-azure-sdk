package trigger

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PipelineReferenceType string

const (
	PipelineReferenceTypePipelineReference PipelineReferenceType = "PipelineReference"
)

func PossibleValuesForPipelineReferenceType() []string {
	return []string{
		string(PipelineReferenceTypePipelineReference),
	}
}

func (s *PipelineReferenceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePipelineReferenceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePipelineReferenceType(input string) (*PipelineReferenceType, error) {
	vals := map[string]PipelineReferenceType{
		"pipelinereference": PipelineReferenceTypePipelineReference,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PipelineReferenceType(input)
	return &out, nil
}

type TriggerReferenceType string

const (
	TriggerReferenceTypeTriggerReference TriggerReferenceType = "TriggerReference"
)

func PossibleValuesForTriggerReferenceType() []string {
	return []string{
		string(TriggerReferenceTypeTriggerReference),
	}
}

func (s *TriggerReferenceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTriggerReferenceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTriggerReferenceType(input string) (*TriggerReferenceType, error) {
	vals := map[string]TriggerReferenceType{
		"triggerreference": TriggerReferenceTypeTriggerReference,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TriggerReferenceType(input)
	return &out, nil
}

type TriggerRuntimeState string

const (
	TriggerRuntimeStateDisabled TriggerRuntimeState = "Disabled"
	TriggerRuntimeStateStarted  TriggerRuntimeState = "Started"
	TriggerRuntimeStateStopped  TriggerRuntimeState = "Stopped"
)

func PossibleValuesForTriggerRuntimeState() []string {
	return []string{
		string(TriggerRuntimeStateDisabled),
		string(TriggerRuntimeStateStarted),
		string(TriggerRuntimeStateStopped),
	}
}

func (s *TriggerRuntimeState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTriggerRuntimeState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTriggerRuntimeState(input string) (*TriggerRuntimeState, error) {
	vals := map[string]TriggerRuntimeState{
		"disabled": TriggerRuntimeStateDisabled,
		"started":  TriggerRuntimeStateStarted,
		"stopped":  TriggerRuntimeStateStopped,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TriggerRuntimeState(input)
	return &out, nil
}

type TumblingWindowFrequency string

const (
	TumblingWindowFrequencyHour   TumblingWindowFrequency = "Hour"
	TumblingWindowFrequencyMinute TumblingWindowFrequency = "Minute"
	TumblingWindowFrequencyMonth  TumblingWindowFrequency = "Month"
)

func PossibleValuesForTumblingWindowFrequency() []string {
	return []string{
		string(TumblingWindowFrequencyHour),
		string(TumblingWindowFrequencyMinute),
		string(TumblingWindowFrequencyMonth),
	}
}

func (s *TumblingWindowFrequency) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTumblingWindowFrequency(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTumblingWindowFrequency(input string) (*TumblingWindowFrequency, error) {
	vals := map[string]TumblingWindowFrequency{
		"hour":   TumblingWindowFrequencyHour,
		"minute": TumblingWindowFrequencyMinute,
		"month":  TumblingWindowFrequencyMonth,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TumblingWindowFrequency(input)
	return &out, nil
}
