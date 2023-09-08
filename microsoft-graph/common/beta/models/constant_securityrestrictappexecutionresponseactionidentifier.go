package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityRestrictAppExecutionResponseActionIdentifier string

const (
	SecurityRestrictAppExecutionResponseActionIdentifierdeviceId SecurityRestrictAppExecutionResponseActionIdentifier = "DeviceId"
)

func PossibleValuesForSecurityRestrictAppExecutionResponseActionIdentifier() []string {
	return []string{
		string(SecurityRestrictAppExecutionResponseActionIdentifierdeviceId),
	}
}

func parseSecurityRestrictAppExecutionResponseActionIdentifier(input string) (*SecurityRestrictAppExecutionResponseActionIdentifier, error) {
	vals := map[string]SecurityRestrictAppExecutionResponseActionIdentifier{
		"deviceid": SecurityRestrictAppExecutionResponseActionIdentifierdeviceId,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityRestrictAppExecutionResponseActionIdentifier(input)
	return &out, nil
}
