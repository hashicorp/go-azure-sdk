package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkScepCertificateProfileKeyUsage string

const (
	AndroidForWorkScepCertificateProfileKeyUsagedigitalSignature AndroidForWorkScepCertificateProfileKeyUsage = "DigitalSignature"
	AndroidForWorkScepCertificateProfileKeyUsagekeyEncipherment  AndroidForWorkScepCertificateProfileKeyUsage = "KeyEncipherment"
)

func PossibleValuesForAndroidForWorkScepCertificateProfileKeyUsage() []string {
	return []string{
		string(AndroidForWorkScepCertificateProfileKeyUsagedigitalSignature),
		string(AndroidForWorkScepCertificateProfileKeyUsagekeyEncipherment),
	}
}

func parseAndroidForWorkScepCertificateProfileKeyUsage(input string) (*AndroidForWorkScepCertificateProfileKeyUsage, error) {
	vals := map[string]AndroidForWorkScepCertificateProfileKeyUsage{
		"digitalsignature": AndroidForWorkScepCertificateProfileKeyUsagedigitalSignature,
		"keyencipherment":  AndroidForWorkScepCertificateProfileKeyUsagekeyEncipherment,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkScepCertificateProfileKeyUsage(input)
	return &out, nil
}
