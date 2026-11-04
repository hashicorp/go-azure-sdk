package datamaskingpolicies

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataMaskingState string

const (
	DataMaskingStateDisabled DataMaskingState = "Disabled"
	DataMaskingStateEnabled  DataMaskingState = "Enabled"
)

func PossibleValuesForDataMaskingState() []string {
	return []string{
		string(DataMaskingStateDisabled),
		string(DataMaskingStateEnabled),
	}
}

func (s *DataMaskingState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDataMaskingState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDataMaskingState(input string) (*DataMaskingState, error) {
	vals := map[string]DataMaskingState{
		"disabled": DataMaskingStateDisabled,
		"enabled":  DataMaskingStateEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DataMaskingState(input)
	return &out, nil
}
