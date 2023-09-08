package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10ImportedPFXCertificateProfileCertificateValidityPeriodScale string

const (
	Windows10ImportedPFXCertificateProfileCertificateValidityPeriodScaledays   Windows10ImportedPFXCertificateProfileCertificateValidityPeriodScale = "Days"
	Windows10ImportedPFXCertificateProfileCertificateValidityPeriodScalemonths Windows10ImportedPFXCertificateProfileCertificateValidityPeriodScale = "Months"
	Windows10ImportedPFXCertificateProfileCertificateValidityPeriodScaleyears  Windows10ImportedPFXCertificateProfileCertificateValidityPeriodScale = "Years"
)

func PossibleValuesForWindows10ImportedPFXCertificateProfileCertificateValidityPeriodScale() []string {
	return []string{
		string(Windows10ImportedPFXCertificateProfileCertificateValidityPeriodScaledays),
		string(Windows10ImportedPFXCertificateProfileCertificateValidityPeriodScalemonths),
		string(Windows10ImportedPFXCertificateProfileCertificateValidityPeriodScaleyears),
	}
}

func parseWindows10ImportedPFXCertificateProfileCertificateValidityPeriodScale(input string) (*Windows10ImportedPFXCertificateProfileCertificateValidityPeriodScale, error) {
	vals := map[string]Windows10ImportedPFXCertificateProfileCertificateValidityPeriodScale{
		"days":   Windows10ImportedPFXCertificateProfileCertificateValidityPeriodScaledays,
		"months": Windows10ImportedPFXCertificateProfileCertificateValidityPeriodScalemonths,
		"years":  Windows10ImportedPFXCertificateProfileCertificateValidityPeriodScaleyears,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10ImportedPFXCertificateProfileCertificateValidityPeriodScale(input)
	return &out, nil
}
