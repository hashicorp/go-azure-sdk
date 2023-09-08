package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityGoogleCloudResourceEvidenceLocationType string

const (
	SecurityGoogleCloudResourceEvidenceLocationTypeglobal   SecurityGoogleCloudResourceEvidenceLocationType = "Global"
	SecurityGoogleCloudResourceEvidenceLocationTyperegional SecurityGoogleCloudResourceEvidenceLocationType = "Regional"
	SecurityGoogleCloudResourceEvidenceLocationTypeunknown  SecurityGoogleCloudResourceEvidenceLocationType = "Unknown"
	SecurityGoogleCloudResourceEvidenceLocationTypezonal    SecurityGoogleCloudResourceEvidenceLocationType = "Zonal"
)

func PossibleValuesForSecurityGoogleCloudResourceEvidenceLocationType() []string {
	return []string{
		string(SecurityGoogleCloudResourceEvidenceLocationTypeglobal),
		string(SecurityGoogleCloudResourceEvidenceLocationTyperegional),
		string(SecurityGoogleCloudResourceEvidenceLocationTypeunknown),
		string(SecurityGoogleCloudResourceEvidenceLocationTypezonal),
	}
}

func parseSecurityGoogleCloudResourceEvidenceLocationType(input string) (*SecurityGoogleCloudResourceEvidenceLocationType, error) {
	vals := map[string]SecurityGoogleCloudResourceEvidenceLocationType{
		"global":   SecurityGoogleCloudResourceEvidenceLocationTypeglobal,
		"regional": SecurityGoogleCloudResourceEvidenceLocationTyperegional,
		"unknown":  SecurityGoogleCloudResourceEvidenceLocationTypeunknown,
		"zonal":    SecurityGoogleCloudResourceEvidenceLocationTypezonal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityGoogleCloudResourceEvidenceLocationType(input)
	return &out, nil
}
