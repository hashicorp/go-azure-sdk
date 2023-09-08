package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityRunAntivirusScanResponseActionIdentifier string

const (
	SecurityRunAntivirusScanResponseActionIdentifierdeviceId SecurityRunAntivirusScanResponseActionIdentifier = "DeviceId"
)

func PossibleValuesForSecurityRunAntivirusScanResponseActionIdentifier() []string {
	return []string{
		string(SecurityRunAntivirusScanResponseActionIdentifierdeviceId),
	}
}

func parseSecurityRunAntivirusScanResponseActionIdentifier(input string) (*SecurityRunAntivirusScanResponseActionIdentifier, error) {
	vals := map[string]SecurityRunAntivirusScanResponseActionIdentifier{
		"deviceid": SecurityRunAntivirusScanResponseActionIdentifierdeviceId,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityRunAntivirusScanResponseActionIdentifier(input)
	return &out, nil
}
