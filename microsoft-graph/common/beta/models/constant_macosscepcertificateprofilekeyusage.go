package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSScepCertificateProfileKeyUsage string

const (
	MacOSScepCertificateProfileKeyUsagedigitalSignature MacOSScepCertificateProfileKeyUsage = "DigitalSignature"
	MacOSScepCertificateProfileKeyUsagekeyEncipherment  MacOSScepCertificateProfileKeyUsage = "KeyEncipherment"
)

func PossibleValuesForMacOSScepCertificateProfileKeyUsage() []string {
	return []string{
		string(MacOSScepCertificateProfileKeyUsagedigitalSignature),
		string(MacOSScepCertificateProfileKeyUsagekeyEncipherment),
	}
}

func parseMacOSScepCertificateProfileKeyUsage(input string) (*MacOSScepCertificateProfileKeyUsage, error) {
	vals := map[string]MacOSScepCertificateProfileKeyUsage{
		"digitalsignature": MacOSScepCertificateProfileKeyUsagedigitalSignature,
		"keyencipherment":  MacOSScepCertificateProfileKeyUsagekeyEncipherment,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSScepCertificateProfileKeyUsage(input)
	return &out, nil
}
