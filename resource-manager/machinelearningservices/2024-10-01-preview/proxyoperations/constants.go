package proxyoperations

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiagnoseResultLevel string

const (
	DiagnoseResultLevelError       DiagnoseResultLevel = "Error"
	DiagnoseResultLevelInformation DiagnoseResultLevel = "Information"
	DiagnoseResultLevelWarning     DiagnoseResultLevel = "Warning"
)

func PossibleValuesForDiagnoseResultLevel() []string {
	return []string{
		string(DiagnoseResultLevelError),
		string(DiagnoseResultLevelInformation),
		string(DiagnoseResultLevelWarning),
	}
}

func (s *DiagnoseResultLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDiagnoseResultLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDiagnoseResultLevel(input string) (*DiagnoseResultLevel, error) {
	vals := map[string]DiagnoseResultLevel{
		"error":       DiagnoseResultLevelError,
		"information": DiagnoseResultLevelInformation,
		"warning":     DiagnoseResultLevelWarning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DiagnoseResultLevel(input)
	return &out, nil
}
