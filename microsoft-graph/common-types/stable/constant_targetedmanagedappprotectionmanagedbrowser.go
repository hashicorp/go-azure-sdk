package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TargetedManagedAppProtectionManagedBrowser string

const (
	TargetedManagedAppProtectionManagedBrowser_MicrosoftEdge TargetedManagedAppProtectionManagedBrowser = "microsoftEdge"
	TargetedManagedAppProtectionManagedBrowser_NotConfigured TargetedManagedAppProtectionManagedBrowser = "notConfigured"
)

func PossibleValuesForTargetedManagedAppProtectionManagedBrowser() []string {
	return []string{
		string(TargetedManagedAppProtectionManagedBrowser_MicrosoftEdge),
		string(TargetedManagedAppProtectionManagedBrowser_NotConfigured),
	}
}

func (s *TargetedManagedAppProtectionManagedBrowser) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTargetedManagedAppProtectionManagedBrowser(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTargetedManagedAppProtectionManagedBrowser(input string) (*TargetedManagedAppProtectionManagedBrowser, error) {
	vals := map[string]TargetedManagedAppProtectionManagedBrowser{
		"microsoftedge": TargetedManagedAppProtectionManagedBrowser_MicrosoftEdge,
		"notconfigured": TargetedManagedAppProtectionManagedBrowser_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TargetedManagedAppProtectionManagedBrowser(input)
	return &out, nil
}
