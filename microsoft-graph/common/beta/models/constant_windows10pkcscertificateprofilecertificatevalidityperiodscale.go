package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10PkcsCertificateProfileCertificateValidityPeriodScale string

const (
	Windows10PkcsCertificateProfileCertificateValidityPeriodScaledays   Windows10PkcsCertificateProfileCertificateValidityPeriodScale = "Days"
	Windows10PkcsCertificateProfileCertificateValidityPeriodScalemonths Windows10PkcsCertificateProfileCertificateValidityPeriodScale = "Months"
	Windows10PkcsCertificateProfileCertificateValidityPeriodScaleyears  Windows10PkcsCertificateProfileCertificateValidityPeriodScale = "Years"
)

func PossibleValuesForWindows10PkcsCertificateProfileCertificateValidityPeriodScale() []string {
	return []string{
		string(Windows10PkcsCertificateProfileCertificateValidityPeriodScaledays),
		string(Windows10PkcsCertificateProfileCertificateValidityPeriodScalemonths),
		string(Windows10PkcsCertificateProfileCertificateValidityPeriodScaleyears),
	}
}

func parseWindows10PkcsCertificateProfileCertificateValidityPeriodScale(input string) (*Windows10PkcsCertificateProfileCertificateValidityPeriodScale, error) {
	vals := map[string]Windows10PkcsCertificateProfileCertificateValidityPeriodScale{
		"days":   Windows10PkcsCertificateProfileCertificateValidityPeriodScaledays,
		"months": Windows10PkcsCertificateProfileCertificateValidityPeriodScalemonths,
		"years":  Windows10PkcsCertificateProfileCertificateValidityPeriodScaleyears,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10PkcsCertificateProfileCertificateValidityPeriodScale(input)
	return &out, nil
}
