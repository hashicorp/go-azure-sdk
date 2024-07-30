package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AdminConsentShareAPNSData string

const (
	AdminConsentShareAPNSData_Granted       AdminConsentShareAPNSData = "granted"
	AdminConsentShareAPNSData_NotConfigured AdminConsentShareAPNSData = "notConfigured"
	AdminConsentShareAPNSData_NotGranted    AdminConsentShareAPNSData = "notGranted"
)

func PossibleValuesForAdminConsentShareAPNSData() []string {
	return []string{
		string(AdminConsentShareAPNSData_Granted),
		string(AdminConsentShareAPNSData_NotConfigured),
		string(AdminConsentShareAPNSData_NotGranted),
	}
}

func (s *AdminConsentShareAPNSData) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAdminConsentShareAPNSData(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAdminConsentShareAPNSData(input string) (*AdminConsentShareAPNSData, error) {
	vals := map[string]AdminConsentShareAPNSData{
		"granted":       AdminConsentShareAPNSData_Granted,
		"notconfigured": AdminConsentShareAPNSData_NotConfigured,
		"notgranted":    AdminConsentShareAPNSData_NotGranted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AdminConsentShareAPNSData(input)
	return &out, nil
}
