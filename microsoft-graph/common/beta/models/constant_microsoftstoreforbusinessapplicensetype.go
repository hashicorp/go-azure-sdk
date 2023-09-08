package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftStoreForBusinessAppLicenseType string

const (
	MicrosoftStoreForBusinessAppLicenseTypeoffline MicrosoftStoreForBusinessAppLicenseType = "Offline"
	MicrosoftStoreForBusinessAppLicenseTypeonline  MicrosoftStoreForBusinessAppLicenseType = "Online"
)

func PossibleValuesForMicrosoftStoreForBusinessAppLicenseType() []string {
	return []string{
		string(MicrosoftStoreForBusinessAppLicenseTypeoffline),
		string(MicrosoftStoreForBusinessAppLicenseTypeonline),
	}
}

func parseMicrosoftStoreForBusinessAppLicenseType(input string) (*MicrosoftStoreForBusinessAppLicenseType, error) {
	vals := map[string]MicrosoftStoreForBusinessAppLicenseType{
		"offline": MicrosoftStoreForBusinessAppLicenseTypeoffline,
		"online":  MicrosoftStoreForBusinessAppLicenseTypeonline,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MicrosoftStoreForBusinessAppLicenseType(input)
	return &out, nil
}
