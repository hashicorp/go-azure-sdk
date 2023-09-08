package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPhone81SCEPCertificateProfileKeyUsage string

const (
	WindowsPhone81SCEPCertificateProfileKeyUsagedigitalSignature WindowsPhone81SCEPCertificateProfileKeyUsage = "DigitalSignature"
	WindowsPhone81SCEPCertificateProfileKeyUsagekeyEncipherment  WindowsPhone81SCEPCertificateProfileKeyUsage = "KeyEncipherment"
)

func PossibleValuesForWindowsPhone81SCEPCertificateProfileKeyUsage() []string {
	return []string{
		string(WindowsPhone81SCEPCertificateProfileKeyUsagedigitalSignature),
		string(WindowsPhone81SCEPCertificateProfileKeyUsagekeyEncipherment),
	}
}

func parseWindowsPhone81SCEPCertificateProfileKeyUsage(input string) (*WindowsPhone81SCEPCertificateProfileKeyUsage, error) {
	vals := map[string]WindowsPhone81SCEPCertificateProfileKeyUsage{
		"digitalsignature": WindowsPhone81SCEPCertificateProfileKeyUsagedigitalSignature,
		"keyencipherment":  WindowsPhone81SCEPCertificateProfileKeyUsagekeyEncipherment,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPhone81SCEPCertificateProfileKeyUsage(input)
	return &out, nil
}
