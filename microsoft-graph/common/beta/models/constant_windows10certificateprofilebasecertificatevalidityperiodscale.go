package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10CertificateProfileBaseCertificateValidityPeriodScale string

const (
	Windows10CertificateProfileBaseCertificateValidityPeriodScaledays   Windows10CertificateProfileBaseCertificateValidityPeriodScale = "Days"
	Windows10CertificateProfileBaseCertificateValidityPeriodScalemonths Windows10CertificateProfileBaseCertificateValidityPeriodScale = "Months"
	Windows10CertificateProfileBaseCertificateValidityPeriodScaleyears  Windows10CertificateProfileBaseCertificateValidityPeriodScale = "Years"
)

func PossibleValuesForWindows10CertificateProfileBaseCertificateValidityPeriodScale() []string {
	return []string{
		string(Windows10CertificateProfileBaseCertificateValidityPeriodScaledays),
		string(Windows10CertificateProfileBaseCertificateValidityPeriodScalemonths),
		string(Windows10CertificateProfileBaseCertificateValidityPeriodScaleyears),
	}
}

func parseWindows10CertificateProfileBaseCertificateValidityPeriodScale(input string) (*Windows10CertificateProfileBaseCertificateValidityPeriodScale, error) {
	vals := map[string]Windows10CertificateProfileBaseCertificateValidityPeriodScale{
		"days":   Windows10CertificateProfileBaseCertificateValidityPeriodScaledays,
		"months": Windows10CertificateProfileBaseCertificateValidityPeriodScalemonths,
		"years":  Windows10CertificateProfileBaseCertificateValidityPeriodScaleyears,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10CertificateProfileBaseCertificateValidityPeriodScale(input)
	return &out, nil
}
