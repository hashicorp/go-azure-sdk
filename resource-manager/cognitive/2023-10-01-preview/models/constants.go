package models

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ModelLifecycleStatus string

const (
	ModelLifecycleStatusGenerallyAvailable ModelLifecycleStatus = "GenerallyAvailable"
	ModelLifecycleStatusPreview            ModelLifecycleStatus = "Preview"
)

func PossibleValuesForModelLifecycleStatus() []string {
	return []string{
		string(ModelLifecycleStatusGenerallyAvailable),
		string(ModelLifecycleStatusPreview),
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
		"generallyavailable": ModelLifecycleStatusGenerallyAvailable,
		"preview":            ModelLifecycleStatusPreview,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ModelLifecycleStatus(input)
	return &out, nil
}
