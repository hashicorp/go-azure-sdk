package alerts

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlertSeverity string

const (
	AlertSeverityCritical      AlertSeverity = "Critical"
	AlertSeverityInformational AlertSeverity = "Informational"
	AlertSeverityWarning       AlertSeverity = "Warning"
)

func PossibleValuesForAlertSeverity() []string {
	return []string{
		string(AlertSeverityCritical),
		string(AlertSeverityInformational),
		string(AlertSeverityWarning),
	}
}

func (s *AlertSeverity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAlertSeverity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAlertSeverity(input string) (*AlertSeverity, error) {
	vals := map[string]AlertSeverity{
		"critical":      AlertSeverityCritical,
		"informational": AlertSeverityInformational,
		"warning":       AlertSeverityWarning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AlertSeverity(input)
	return &out, nil
}
