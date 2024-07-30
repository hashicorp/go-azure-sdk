package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationDiagnosticsDataSubmissionMode string

const (
	Windows10GeneralConfigurationDiagnosticsDataSubmissionMode_Basic       Windows10GeneralConfigurationDiagnosticsDataSubmissionMode = "basic"
	Windows10GeneralConfigurationDiagnosticsDataSubmissionMode_Enhanced    Windows10GeneralConfigurationDiagnosticsDataSubmissionMode = "enhanced"
	Windows10GeneralConfigurationDiagnosticsDataSubmissionMode_Full        Windows10GeneralConfigurationDiagnosticsDataSubmissionMode = "full"
	Windows10GeneralConfigurationDiagnosticsDataSubmissionMode_None        Windows10GeneralConfigurationDiagnosticsDataSubmissionMode = "none"
	Windows10GeneralConfigurationDiagnosticsDataSubmissionMode_UserDefined Windows10GeneralConfigurationDiagnosticsDataSubmissionMode = "userDefined"
)

func PossibleValuesForWindows10GeneralConfigurationDiagnosticsDataSubmissionMode() []string {
	return []string{
		string(Windows10GeneralConfigurationDiagnosticsDataSubmissionMode_Basic),
		string(Windows10GeneralConfigurationDiagnosticsDataSubmissionMode_Enhanced),
		string(Windows10GeneralConfigurationDiagnosticsDataSubmissionMode_Full),
		string(Windows10GeneralConfigurationDiagnosticsDataSubmissionMode_None),
		string(Windows10GeneralConfigurationDiagnosticsDataSubmissionMode_UserDefined),
	}
}

func (s *Windows10GeneralConfigurationDiagnosticsDataSubmissionMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10GeneralConfigurationDiagnosticsDataSubmissionMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10GeneralConfigurationDiagnosticsDataSubmissionMode(input string) (*Windows10GeneralConfigurationDiagnosticsDataSubmissionMode, error) {
	vals := map[string]Windows10GeneralConfigurationDiagnosticsDataSubmissionMode{
		"basic":       Windows10GeneralConfigurationDiagnosticsDataSubmissionMode_Basic,
		"enhanced":    Windows10GeneralConfigurationDiagnosticsDataSubmissionMode_Enhanced,
		"full":        Windows10GeneralConfigurationDiagnosticsDataSubmissionMode_Full,
		"none":        Windows10GeneralConfigurationDiagnosticsDataSubmissionMode_None,
		"userdefined": Windows10GeneralConfigurationDiagnosticsDataSubmissionMode_UserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationDiagnosticsDataSubmissionMode(input)
	return &out, nil
}
