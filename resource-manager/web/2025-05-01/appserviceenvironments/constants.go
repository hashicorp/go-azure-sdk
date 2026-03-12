package appserviceenvironments

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComputeModeOptions string

const (
	ComputeModeOptionsDedicated ComputeModeOptions = "Dedicated"
	ComputeModeOptionsDynamic   ComputeModeOptions = "Dynamic"
	ComputeModeOptionsShared    ComputeModeOptions = "Shared"
)

func PossibleValuesForComputeModeOptions() []string {
	return []string{
		string(ComputeModeOptionsDedicated),
		string(ComputeModeOptionsDynamic),
		string(ComputeModeOptionsShared),
	}
}

func (s *ComputeModeOptions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseComputeModeOptions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseComputeModeOptions(input string) (*ComputeModeOptions, error) {
	vals := map[string]ComputeModeOptions{
		"dedicated": ComputeModeOptionsDedicated,
		"dynamic":   ComputeModeOptionsDynamic,
		"shared":    ComputeModeOptionsShared,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ComputeModeOptions(input)
	return &out, nil
}
