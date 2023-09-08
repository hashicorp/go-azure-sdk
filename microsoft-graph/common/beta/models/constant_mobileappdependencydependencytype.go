package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppDependencyDependencyType string

const (
	MobileAppDependencyDependencyTypeautoInstall MobileAppDependencyDependencyType = "AutoInstall"
	MobileAppDependencyDependencyTypedetect      MobileAppDependencyDependencyType = "Detect"
)

func PossibleValuesForMobileAppDependencyDependencyType() []string {
	return []string{
		string(MobileAppDependencyDependencyTypeautoInstall),
		string(MobileAppDependencyDependencyTypedetect),
	}
}

func parseMobileAppDependencyDependencyType(input string) (*MobileAppDependencyDependencyType, error) {
	vals := map[string]MobileAppDependencyDependencyType{
		"autoinstall": MobileAppDependencyDependencyTypeautoInstall,
		"detect":      MobileAppDependencyDependencyTypedetect,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MobileAppDependencyDependencyType(input)
	return &out, nil
}
