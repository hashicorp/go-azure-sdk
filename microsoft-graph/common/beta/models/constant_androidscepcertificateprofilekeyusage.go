package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidScepCertificateProfileKeyUsage string

const (
	AndroidScepCertificateProfileKeyUsagedigitalSignature AndroidScepCertificateProfileKeyUsage = "DigitalSignature"
	AndroidScepCertificateProfileKeyUsagekeyEncipherment  AndroidScepCertificateProfileKeyUsage = "KeyEncipherment"
)

func PossibleValuesForAndroidScepCertificateProfileKeyUsage() []string {
	return []string{
		string(AndroidScepCertificateProfileKeyUsagedigitalSignature),
		string(AndroidScepCertificateProfileKeyUsagekeyEncipherment),
	}
}

func parseAndroidScepCertificateProfileKeyUsage(input string) (*AndroidScepCertificateProfileKeyUsage, error) {
	vals := map[string]AndroidScepCertificateProfileKeyUsage{
		"digitalsignature": AndroidScepCertificateProfileKeyUsagedigitalSignature,
		"keyencipherment":  AndroidScepCertificateProfileKeyUsagekeyEncipherment,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidScepCertificateProfileKeyUsage(input)
	return &out, nil
}
