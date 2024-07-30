package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppDependencyDependencyType string

const (
	MobileAppDependencyDependencyType_AutoInstall MobileAppDependencyDependencyType = "autoInstall"
	MobileAppDependencyDependencyType_Detect      MobileAppDependencyDependencyType = "detect"
)

func PossibleValuesForMobileAppDependencyDependencyType() []string {
	return []string{
		string(MobileAppDependencyDependencyType_AutoInstall),
		string(MobileAppDependencyDependencyType_Detect),
	}
}

func (s *MobileAppDependencyDependencyType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMobileAppDependencyDependencyType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMobileAppDependencyDependencyType(input string) (*MobileAppDependencyDependencyType, error) {
	vals := map[string]MobileAppDependencyDependencyType{
		"autoinstall": MobileAppDependencyDependencyType_AutoInstall,
		"detect":      MobileAppDependencyDependencyType_Detect,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MobileAppDependencyDependencyType(input)
	return &out, nil
}
