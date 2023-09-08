package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10AssociatedAppsAppType string

const (
	Windows10AssociatedAppsAppTypedesktop   Windows10AssociatedAppsAppType = "Desktop"
	Windows10AssociatedAppsAppTypeuniversal Windows10AssociatedAppsAppType = "Universal"
)

func PossibleValuesForWindows10AssociatedAppsAppType() []string {
	return []string{
		string(Windows10AssociatedAppsAppTypedesktop),
		string(Windows10AssociatedAppsAppTypeuniversal),
	}
}

func parseWindows10AssociatedAppsAppType(input string) (*Windows10AssociatedAppsAppType, error) {
	vals := map[string]Windows10AssociatedAppsAppType{
		"desktop":   Windows10AssociatedAppsAppTypedesktop,
		"universal": Windows10AssociatedAppsAppTypeuniversal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10AssociatedAppsAppType(input)
	return &out, nil
}
