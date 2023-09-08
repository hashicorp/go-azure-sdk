package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows81SCEPCertificateProfileKeyUsage string

const (
	Windows81SCEPCertificateProfileKeyUsagedigitalSignature Windows81SCEPCertificateProfileKeyUsage = "DigitalSignature"
	Windows81SCEPCertificateProfileKeyUsagekeyEncipherment  Windows81SCEPCertificateProfileKeyUsage = "KeyEncipherment"
)

func PossibleValuesForWindows81SCEPCertificateProfileKeyUsage() []string {
	return []string{
		string(Windows81SCEPCertificateProfileKeyUsagedigitalSignature),
		string(Windows81SCEPCertificateProfileKeyUsagekeyEncipherment),
	}
}

func parseWindows81SCEPCertificateProfileKeyUsage(input string) (*Windows81SCEPCertificateProfileKeyUsage, error) {
	vals := map[string]Windows81SCEPCertificateProfileKeyUsage{
		"digitalsignature": Windows81SCEPCertificateProfileKeyUsagedigitalSignature,
		"keyencipherment":  Windows81SCEPCertificateProfileKeyUsagekeyEncipherment,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows81SCEPCertificateProfileKeyUsage(input)
	return &out, nil
}
