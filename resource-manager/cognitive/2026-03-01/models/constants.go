package models

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeprecationStatus string

const (
	DeprecationStatusPlanned   DeprecationStatus = "Planned"
	DeprecationStatusTentative DeprecationStatus = "Tentative"
)

func PossibleValuesForDeprecationStatus() []string {
	return []string{
		string(DeprecationStatusPlanned),
		string(DeprecationStatusTentative),
	}
}

func (s *DeprecationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeprecationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeprecationStatus(input string) (*DeprecationStatus, error) {
	vals := map[string]DeprecationStatus{
		"planned":   DeprecationStatusPlanned,
		"tentative": DeprecationStatusTentative,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeprecationStatus(input)
	return &out, nil
}

type ModelLifecycleStatus string

const (
	ModelLifecycleStatusDeprecated         ModelLifecycleStatus = "Deprecated"
	ModelLifecycleStatusDeprecating        ModelLifecycleStatus = "Deprecating"
	ModelLifecycleStatusGenerallyAvailable ModelLifecycleStatus = "GenerallyAvailable"
	ModelLifecycleStatusLegacy             ModelLifecycleStatus = "Legacy"
	ModelLifecycleStatusPreview            ModelLifecycleStatus = "Preview"
	ModelLifecycleStatusStable             ModelLifecycleStatus = "Stable"
)

func PossibleValuesForModelLifecycleStatus() []string {
	return []string{
		string(ModelLifecycleStatusDeprecated),
		string(ModelLifecycleStatusDeprecating),
		string(ModelLifecycleStatusGenerallyAvailable),
		string(ModelLifecycleStatusLegacy),
		string(ModelLifecycleStatusPreview),
		string(ModelLifecycleStatusStable),
	}
}

func (s *ModelLifecycleStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseModelLifecycleStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseModelLifecycleStatus(input string) (*ModelLifecycleStatus, error) {
	vals := map[string]ModelLifecycleStatus{
		"deprecated":         ModelLifecycleStatusDeprecated,
		"deprecating":        ModelLifecycleStatusDeprecating,
		"generallyavailable": ModelLifecycleStatusGenerallyAvailable,
		"legacy":             ModelLifecycleStatusLegacy,
		"preview":            ModelLifecycleStatusPreview,
		"stable":             ModelLifecycleStatusStable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ModelLifecycleStatus(input)
	return &out, nil
}
