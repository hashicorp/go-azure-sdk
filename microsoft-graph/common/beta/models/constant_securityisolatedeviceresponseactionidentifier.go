package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityIsolateDeviceResponseActionIdentifier string

const (
	SecurityIsolateDeviceResponseActionIdentifierdeviceId SecurityIsolateDeviceResponseActionIdentifier = "DeviceId"
)

func PossibleValuesForSecurityIsolateDeviceResponseActionIdentifier() []string {
	return []string{
		string(SecurityIsolateDeviceResponseActionIdentifierdeviceId),
	}
}

func parseSecurityIsolateDeviceResponseActionIdentifier(input string) (*SecurityIsolateDeviceResponseActionIdentifier, error) {
	vals := map[string]SecurityIsolateDeviceResponseActionIdentifier{
		"deviceid": SecurityIsolateDeviceResponseActionIdentifierdeviceId,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityIsolateDeviceResponseActionIdentifier(input)
	return &out, nil
}
