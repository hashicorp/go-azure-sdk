package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsManagedDeviceManagementFeatures string

const (
	WindowsManagedDeviceManagementFeatures_MicrosoftManagedDesktop WindowsManagedDeviceManagementFeatures = "microsoftManagedDesktop"
	WindowsManagedDeviceManagementFeatures_None                    WindowsManagedDeviceManagementFeatures = "none"
)

func PossibleValuesForWindowsManagedDeviceManagementFeatures() []string {
	return []string{
		string(WindowsManagedDeviceManagementFeatures_MicrosoftManagedDesktop),
		string(WindowsManagedDeviceManagementFeatures_None),
	}
}

func (s *WindowsManagedDeviceManagementFeatures) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsManagedDeviceManagementFeatures(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsManagedDeviceManagementFeatures(input string) (*WindowsManagedDeviceManagementFeatures, error) {
	vals := map[string]WindowsManagedDeviceManagementFeatures{
		"microsoftmanageddesktop": WindowsManagedDeviceManagementFeatures_MicrosoftManagedDesktop,
		"none":                    WindowsManagedDeviceManagementFeatures_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsManagedDeviceManagementFeatures(input)
	return &out, nil
}
