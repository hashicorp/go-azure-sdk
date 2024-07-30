package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesUpdateManagementEnrollmentUpdateCategory string

const (
	WindowsUpdatesUpdateManagementEnrollmentUpdateCategory_Driver  WindowsUpdatesUpdateManagementEnrollmentUpdateCategory = "driver"
	WindowsUpdatesUpdateManagementEnrollmentUpdateCategory_Feature WindowsUpdatesUpdateManagementEnrollmentUpdateCategory = "feature"
	WindowsUpdatesUpdateManagementEnrollmentUpdateCategory_Quality WindowsUpdatesUpdateManagementEnrollmentUpdateCategory = "quality"
)

func PossibleValuesForWindowsUpdatesUpdateManagementEnrollmentUpdateCategory() []string {
	return []string{
		string(WindowsUpdatesUpdateManagementEnrollmentUpdateCategory_Driver),
		string(WindowsUpdatesUpdateManagementEnrollmentUpdateCategory_Feature),
		string(WindowsUpdatesUpdateManagementEnrollmentUpdateCategory_Quality),
	}
}

func (s *WindowsUpdatesUpdateManagementEnrollmentUpdateCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsUpdatesUpdateManagementEnrollmentUpdateCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsUpdatesUpdateManagementEnrollmentUpdateCategory(input string) (*WindowsUpdatesUpdateManagementEnrollmentUpdateCategory, error) {
	vals := map[string]WindowsUpdatesUpdateManagementEnrollmentUpdateCategory{
		"driver":  WindowsUpdatesUpdateManagementEnrollmentUpdateCategory_Driver,
		"feature": WindowsUpdatesUpdateManagementEnrollmentUpdateCategory_Feature,
		"quality": WindowsUpdatesUpdateManagementEnrollmentUpdateCategory_Quality,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUpdatesUpdateManagementEnrollmentUpdateCategory(input)
	return &out, nil
}
