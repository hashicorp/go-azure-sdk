package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileScepCertificateProfileKeyUsage string

const (
	AndroidWorkProfileScepCertificateProfileKeyUsagedigitalSignature AndroidWorkProfileScepCertificateProfileKeyUsage = "DigitalSignature"
	AndroidWorkProfileScepCertificateProfileKeyUsagekeyEncipherment  AndroidWorkProfileScepCertificateProfileKeyUsage = "KeyEncipherment"
)

func PossibleValuesForAndroidWorkProfileScepCertificateProfileKeyUsage() []string {
	return []string{
		string(AndroidWorkProfileScepCertificateProfileKeyUsagedigitalSignature),
		string(AndroidWorkProfileScepCertificateProfileKeyUsagekeyEncipherment),
	}
}

func parseAndroidWorkProfileScepCertificateProfileKeyUsage(input string) (*AndroidWorkProfileScepCertificateProfileKeyUsage, error) {
	vals := map[string]AndroidWorkProfileScepCertificateProfileKeyUsage{
		"digitalsignature": AndroidWorkProfileScepCertificateProfileKeyUsagedigitalSignature,
		"keyencipherment":  AndroidWorkProfileScepCertificateProfileKeyUsagekeyEncipherment,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileScepCertificateProfileKeyUsage(input)
	return &out, nil
}
