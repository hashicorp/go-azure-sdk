package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10XSCEPCertificateProfileCertificateValidityPeriodScale string

const (
	Windows10XSCEPCertificateProfileCertificateValidityPeriodScaledays   Windows10XSCEPCertificateProfileCertificateValidityPeriodScale = "Days"
	Windows10XSCEPCertificateProfileCertificateValidityPeriodScalemonths Windows10XSCEPCertificateProfileCertificateValidityPeriodScale = "Months"
	Windows10XSCEPCertificateProfileCertificateValidityPeriodScaleyears  Windows10XSCEPCertificateProfileCertificateValidityPeriodScale = "Years"
)

func PossibleValuesForWindows10XSCEPCertificateProfileCertificateValidityPeriodScale() []string {
	return []string{
		string(Windows10XSCEPCertificateProfileCertificateValidityPeriodScaledays),
		string(Windows10XSCEPCertificateProfileCertificateValidityPeriodScalemonths),
		string(Windows10XSCEPCertificateProfileCertificateValidityPeriodScaleyears),
	}
}

func parseWindows10XSCEPCertificateProfileCertificateValidityPeriodScale(input string) (*Windows10XSCEPCertificateProfileCertificateValidityPeriodScale, error) {
	vals := map[string]Windows10XSCEPCertificateProfileCertificateValidityPeriodScale{
		"days":   Windows10XSCEPCertificateProfileCertificateValidityPeriodScaledays,
		"months": Windows10XSCEPCertificateProfileCertificateValidityPeriodScalemonths,
		"years":  Windows10XSCEPCertificateProfileCertificateValidityPeriodScaleyears,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10XSCEPCertificateProfileCertificateValidityPeriodScale(input)
	return &out, nil
}
