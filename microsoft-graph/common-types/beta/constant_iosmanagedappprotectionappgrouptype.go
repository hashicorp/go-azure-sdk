package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosManagedAppProtectionAppGroupType string

const (
	IosManagedAppProtectionAppGroupType_AllApps              IosManagedAppProtectionAppGroupType = "allApps"
	IosManagedAppProtectionAppGroupType_AllCoreMicrosoftApps IosManagedAppProtectionAppGroupType = "allCoreMicrosoftApps"
	IosManagedAppProtectionAppGroupType_AllMicrosoftApps     IosManagedAppProtectionAppGroupType = "allMicrosoftApps"
	IosManagedAppProtectionAppGroupType_SelectedPublicApps   IosManagedAppProtectionAppGroupType = "selectedPublicApps"
)

func PossibleValuesForIosManagedAppProtectionAppGroupType() []string {
	return []string{
		string(IosManagedAppProtectionAppGroupType_AllApps),
		string(IosManagedAppProtectionAppGroupType_AllCoreMicrosoftApps),
		string(IosManagedAppProtectionAppGroupType_AllMicrosoftApps),
		string(IosManagedAppProtectionAppGroupType_SelectedPublicApps),
	}
}

func (s *IosManagedAppProtectionAppGroupType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosManagedAppProtectionAppGroupType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosManagedAppProtectionAppGroupType(input string) (*IosManagedAppProtectionAppGroupType, error) {
	vals := map[string]IosManagedAppProtectionAppGroupType{
		"allapps":              IosManagedAppProtectionAppGroupType_AllApps,
		"allcoremicrosoftapps": IosManagedAppProtectionAppGroupType_AllCoreMicrosoftApps,
		"allmicrosoftapps":     IosManagedAppProtectionAppGroupType_AllMicrosoftApps,
		"selectedpublicapps":   IosManagedAppProtectionAppGroupType_SelectedPublicApps,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosManagedAppProtectionAppGroupType(input)
	return &out, nil
}
