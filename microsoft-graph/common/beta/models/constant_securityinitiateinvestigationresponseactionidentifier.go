package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityInitiateInvestigationResponseActionIdentifier string

const (
	SecurityInitiateInvestigationResponseActionIdentifierdeviceId SecurityInitiateInvestigationResponseActionIdentifier = "DeviceId"
)

func PossibleValuesForSecurityInitiateInvestigationResponseActionIdentifier() []string {
	return []string{
		string(SecurityInitiateInvestigationResponseActionIdentifierdeviceId),
	}
}

func parseSecurityInitiateInvestigationResponseActionIdentifier(input string) (*SecurityInitiateInvestigationResponseActionIdentifier, error) {
	vals := map[string]SecurityInitiateInvestigationResponseActionIdentifier{
		"deviceid": SecurityInitiateInvestigationResponseActionIdentifierdeviceId,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityInitiateInvestigationResponseActionIdentifier(input)
	return &out, nil
}
