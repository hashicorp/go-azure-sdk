package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlertTemplateSeverity string

const (
	AlertTemplateSeverity_High          AlertTemplateSeverity = "high"
	AlertTemplateSeverity_Informational AlertTemplateSeverity = "informational"
	AlertTemplateSeverity_Low           AlertTemplateSeverity = "low"
	AlertTemplateSeverity_Medium        AlertTemplateSeverity = "medium"
	AlertTemplateSeverity_Unknown       AlertTemplateSeverity = "unknown"
)

func PossibleValuesForAlertTemplateSeverity() []string {
	return []string{
		string(AlertTemplateSeverity_High),
		string(AlertTemplateSeverity_Informational),
		string(AlertTemplateSeverity_Low),
		string(AlertTemplateSeverity_Medium),
		string(AlertTemplateSeverity_Unknown),
	}
}

func (s *AlertTemplateSeverity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAlertTemplateSeverity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAlertTemplateSeverity(input string) (*AlertTemplateSeverity, error) {
	vals := map[string]AlertTemplateSeverity{
		"high":          AlertTemplateSeverity_High,
		"informational": AlertTemplateSeverity_Informational,
		"low":           AlertTemplateSeverity_Low,
		"medium":        AlertTemplateSeverity_Medium,
		"unknown":       AlertTemplateSeverity_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AlertTemplateSeverity(input)
	return &out, nil
}
