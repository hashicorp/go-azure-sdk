package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppProtectionManagedBrowser string

const (
	ManagedAppProtectionManagedBrowser_MicrosoftEdge ManagedAppProtectionManagedBrowser = "microsoftEdge"
	ManagedAppProtectionManagedBrowser_NotConfigured ManagedAppProtectionManagedBrowser = "notConfigured"
)

func PossibleValuesForManagedAppProtectionManagedBrowser() []string {
	return []string{
		string(ManagedAppProtectionManagedBrowser_MicrosoftEdge),
		string(ManagedAppProtectionManagedBrowser_NotConfigured),
	}
}

func (s *ManagedAppProtectionManagedBrowser) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedAppProtectionManagedBrowser(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedAppProtectionManagedBrowser(input string) (*ManagedAppProtectionManagedBrowser, error) {
	vals := map[string]ManagedAppProtectionManagedBrowser{
		"microsoftedge": ManagedAppProtectionManagedBrowser_MicrosoftEdge,
		"notconfigured": ManagedAppProtectionManagedBrowser_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedAppProtectionManagedBrowser(input)
	return &out, nil
}
