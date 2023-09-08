package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityIsolateDeviceResponseActionIsolationType string

const (
	SecurityIsolateDeviceResponseActionIsolationTypefull      SecurityIsolateDeviceResponseActionIsolationType = "Full"
	SecurityIsolateDeviceResponseActionIsolationTypeselective SecurityIsolateDeviceResponseActionIsolationType = "Selective"
)

func PossibleValuesForSecurityIsolateDeviceResponseActionIsolationType() []string {
	return []string{
		string(SecurityIsolateDeviceResponseActionIsolationTypefull),
		string(SecurityIsolateDeviceResponseActionIsolationTypeselective),
	}
}

func parseSecurityIsolateDeviceResponseActionIsolationType(input string) (*SecurityIsolateDeviceResponseActionIsolationType, error) {
	vals := map[string]SecurityIsolateDeviceResponseActionIsolationType{
		"full":      SecurityIsolateDeviceResponseActionIsolationTypefull,
		"selective": SecurityIsolateDeviceResponseActionIsolationTypeselective,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityIsolateDeviceResponseActionIsolationType(input)
	return &out, nil
}
