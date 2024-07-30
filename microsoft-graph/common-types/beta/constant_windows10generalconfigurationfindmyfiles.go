package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationFindMyFiles string

const (
	Windows10GeneralConfigurationFindMyFiles_Disabled      Windows10GeneralConfigurationFindMyFiles = "disabled"
	Windows10GeneralConfigurationFindMyFiles_Enabled       Windows10GeneralConfigurationFindMyFiles = "enabled"
	Windows10GeneralConfigurationFindMyFiles_NotConfigured Windows10GeneralConfigurationFindMyFiles = "notConfigured"
)

func PossibleValuesForWindows10GeneralConfigurationFindMyFiles() []string {
	return []string{
		string(Windows10GeneralConfigurationFindMyFiles_Disabled),
		string(Windows10GeneralConfigurationFindMyFiles_Enabled),
		string(Windows10GeneralConfigurationFindMyFiles_NotConfigured),
	}
}

func (s *Windows10GeneralConfigurationFindMyFiles) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10GeneralConfigurationFindMyFiles(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10GeneralConfigurationFindMyFiles(input string) (*Windows10GeneralConfigurationFindMyFiles, error) {
	vals := map[string]Windows10GeneralConfigurationFindMyFiles{
		"disabled":      Windows10GeneralConfigurationFindMyFiles_Disabled,
		"enabled":       Windows10GeneralConfigurationFindMyFiles_Enabled,
		"notconfigured": Windows10GeneralConfigurationFindMyFiles_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationFindMyFiles(input)
	return &out, nil
}
