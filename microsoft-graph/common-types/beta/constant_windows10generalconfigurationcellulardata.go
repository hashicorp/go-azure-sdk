package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationCellularData string

const (
	Windows10GeneralConfigurationCellularData_Allowed       Windows10GeneralConfigurationCellularData = "allowed"
	Windows10GeneralConfigurationCellularData_Blocked       Windows10GeneralConfigurationCellularData = "blocked"
	Windows10GeneralConfigurationCellularData_NotConfigured Windows10GeneralConfigurationCellularData = "notConfigured"
	Windows10GeneralConfigurationCellularData_Required      Windows10GeneralConfigurationCellularData = "required"
)

func PossibleValuesForWindows10GeneralConfigurationCellularData() []string {
	return []string{
		string(Windows10GeneralConfigurationCellularData_Allowed),
		string(Windows10GeneralConfigurationCellularData_Blocked),
		string(Windows10GeneralConfigurationCellularData_NotConfigured),
		string(Windows10GeneralConfigurationCellularData_Required),
	}
}

func (s *Windows10GeneralConfigurationCellularData) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10GeneralConfigurationCellularData(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10GeneralConfigurationCellularData(input string) (*Windows10GeneralConfigurationCellularData, error) {
	vals := map[string]Windows10GeneralConfigurationCellularData{
		"allowed":       Windows10GeneralConfigurationCellularData_Allowed,
		"blocked":       Windows10GeneralConfigurationCellularData_Blocked,
		"notconfigured": Windows10GeneralConfigurationCellularData_NotConfigured,
		"required":      Windows10GeneralConfigurationCellularData_Required,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationCellularData(input)
	return &out, nil
}
