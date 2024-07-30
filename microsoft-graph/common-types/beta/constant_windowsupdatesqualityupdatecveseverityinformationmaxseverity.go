package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesQualityUpdateCveSeverityInformationMaxSeverity string

const (
	WindowsUpdatesQualityUpdateCveSeverityInformationMaxSeverity_Critical  WindowsUpdatesQualityUpdateCveSeverityInformationMaxSeverity = "critical"
	WindowsUpdatesQualityUpdateCveSeverityInformationMaxSeverity_Important WindowsUpdatesQualityUpdateCveSeverityInformationMaxSeverity = "important"
	WindowsUpdatesQualityUpdateCveSeverityInformationMaxSeverity_Moderate  WindowsUpdatesQualityUpdateCveSeverityInformationMaxSeverity = "moderate"
)

func PossibleValuesForWindowsUpdatesQualityUpdateCveSeverityInformationMaxSeverity() []string {
	return []string{
		string(WindowsUpdatesQualityUpdateCveSeverityInformationMaxSeverity_Critical),
		string(WindowsUpdatesQualityUpdateCveSeverityInformationMaxSeverity_Important),
		string(WindowsUpdatesQualityUpdateCveSeverityInformationMaxSeverity_Moderate),
	}
}

func (s *WindowsUpdatesQualityUpdateCveSeverityInformationMaxSeverity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsUpdatesQualityUpdateCveSeverityInformationMaxSeverity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsUpdatesQualityUpdateCveSeverityInformationMaxSeverity(input string) (*WindowsUpdatesQualityUpdateCveSeverityInformationMaxSeverity, error) {
	vals := map[string]WindowsUpdatesQualityUpdateCveSeverityInformationMaxSeverity{
		"critical":  WindowsUpdatesQualityUpdateCveSeverityInformationMaxSeverity_Critical,
		"important": WindowsUpdatesQualityUpdateCveSeverityInformationMaxSeverity_Important,
		"moderate":  WindowsUpdatesQualityUpdateCveSeverityInformationMaxSeverity_Moderate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUpdatesQualityUpdateCveSeverityInformationMaxSeverity(input)
	return &out, nil
}
