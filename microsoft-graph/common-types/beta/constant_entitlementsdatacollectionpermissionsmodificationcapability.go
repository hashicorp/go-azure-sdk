package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementsDataCollectionPermissionsModificationCapability string

const (
	EntitlementsDataCollectionPermissionsModificationCapability_Enabled               EntitlementsDataCollectionPermissionsModificationCapability = "enabled"
	EntitlementsDataCollectionPermissionsModificationCapability_NoRecentDataCollected EntitlementsDataCollectionPermissionsModificationCapability = "noRecentDataCollected"
	EntitlementsDataCollectionPermissionsModificationCapability_NotConfigured         EntitlementsDataCollectionPermissionsModificationCapability = "notConfigured"
)

func PossibleValuesForEntitlementsDataCollectionPermissionsModificationCapability() []string {
	return []string{
		string(EntitlementsDataCollectionPermissionsModificationCapability_Enabled),
		string(EntitlementsDataCollectionPermissionsModificationCapability_NoRecentDataCollected),
		string(EntitlementsDataCollectionPermissionsModificationCapability_NotConfigured),
	}
}

func (s *EntitlementsDataCollectionPermissionsModificationCapability) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEntitlementsDataCollectionPermissionsModificationCapability(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEntitlementsDataCollectionPermissionsModificationCapability(input string) (*EntitlementsDataCollectionPermissionsModificationCapability, error) {
	vals := map[string]EntitlementsDataCollectionPermissionsModificationCapability{
		"enabled":               EntitlementsDataCollectionPermissionsModificationCapability_Enabled,
		"norecentdatacollected": EntitlementsDataCollectionPermissionsModificationCapability_NoRecentDataCollected,
		"notconfigured":         EntitlementsDataCollectionPermissionsModificationCapability_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EntitlementsDataCollectionPermissionsModificationCapability(input)
	return &out, nil
}
