package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerEnrollmentProfileEnrollmentTokenType string

const (
	AndroidDeviceOwnerEnrollmentProfileEnrollmentTokenTypecorporateOwnedDedicatedDeviceWithAzureADSharedMode AndroidDeviceOwnerEnrollmentProfileEnrollmentTokenType = "CorporateOwnedDedicatedDeviceWithAzureADSharedMode"
	AndroidDeviceOwnerEnrollmentProfileEnrollmentTokenTypedefault                                            AndroidDeviceOwnerEnrollmentProfileEnrollmentTokenType = "Default"
)

func PossibleValuesForAndroidDeviceOwnerEnrollmentProfileEnrollmentTokenType() []string {
	return []string{
		string(AndroidDeviceOwnerEnrollmentProfileEnrollmentTokenTypecorporateOwnedDedicatedDeviceWithAzureADSharedMode),
		string(AndroidDeviceOwnerEnrollmentProfileEnrollmentTokenTypedefault),
	}
}

func parseAndroidDeviceOwnerEnrollmentProfileEnrollmentTokenType(input string) (*AndroidDeviceOwnerEnrollmentProfileEnrollmentTokenType, error) {
	vals := map[string]AndroidDeviceOwnerEnrollmentProfileEnrollmentTokenType{
		"corporateowneddedicateddevicewithazureadsharedmode": AndroidDeviceOwnerEnrollmentProfileEnrollmentTokenTypecorporateOwnedDedicatedDeviceWithAzureADSharedMode,
		"default": AndroidDeviceOwnerEnrollmentProfileEnrollmentTokenTypedefault,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerEnrollmentProfileEnrollmentTokenType(input)
	return &out, nil
}
