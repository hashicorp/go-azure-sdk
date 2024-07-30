package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultManagedAppProtectionManagedBrowser string

const (
	DefaultManagedAppProtectionManagedBrowser_MicrosoftEdge DefaultManagedAppProtectionManagedBrowser = "microsoftEdge"
	DefaultManagedAppProtectionManagedBrowser_NotConfigured DefaultManagedAppProtectionManagedBrowser = "notConfigured"
)

func PossibleValuesForDefaultManagedAppProtectionManagedBrowser() []string {
	return []string{
		string(DefaultManagedAppProtectionManagedBrowser_MicrosoftEdge),
		string(DefaultManagedAppProtectionManagedBrowser_NotConfigured),
	}
}

func (s *DefaultManagedAppProtectionManagedBrowser) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDefaultManagedAppProtectionManagedBrowser(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDefaultManagedAppProtectionManagedBrowser(input string) (*DefaultManagedAppProtectionManagedBrowser, error) {
	vals := map[string]DefaultManagedAppProtectionManagedBrowser{
		"microsoftedge": DefaultManagedAppProtectionManagedBrowser_MicrosoftEdge,
		"notconfigured": DefaultManagedAppProtectionManagedBrowser_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultManagedAppProtectionManagedBrowser(input)
	return &out, nil
}
