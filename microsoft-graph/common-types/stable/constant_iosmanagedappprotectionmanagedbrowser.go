package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosManagedAppProtectionManagedBrowser string

const (
	IosManagedAppProtectionManagedBrowser_MicrosoftEdge IosManagedAppProtectionManagedBrowser = "microsoftEdge"
	IosManagedAppProtectionManagedBrowser_NotConfigured IosManagedAppProtectionManagedBrowser = "notConfigured"
)

func PossibleValuesForIosManagedAppProtectionManagedBrowser() []string {
	return []string{
		string(IosManagedAppProtectionManagedBrowser_MicrosoftEdge),
		string(IosManagedAppProtectionManagedBrowser_NotConfigured),
	}
}

func (s *IosManagedAppProtectionManagedBrowser) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosManagedAppProtectionManagedBrowser(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosManagedAppProtectionManagedBrowser(input string) (*IosManagedAppProtectionManagedBrowser, error) {
	vals := map[string]IosManagedAppProtectionManagedBrowser{
		"microsoftedge": IosManagedAppProtectionManagedBrowser_MicrosoftEdge,
		"notconfigured": IosManagedAppProtectionManagedBrowser_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosManagedAppProtectionManagedBrowser(input)
	return &out, nil
}
