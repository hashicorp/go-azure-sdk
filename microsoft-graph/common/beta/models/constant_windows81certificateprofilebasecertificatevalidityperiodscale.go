package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows81CertificateProfileBaseCertificateValidityPeriodScale string

const (
	Windows81CertificateProfileBaseCertificateValidityPeriodScaledays   Windows81CertificateProfileBaseCertificateValidityPeriodScale = "Days"
	Windows81CertificateProfileBaseCertificateValidityPeriodScalemonths Windows81CertificateProfileBaseCertificateValidityPeriodScale = "Months"
	Windows81CertificateProfileBaseCertificateValidityPeriodScaleyears  Windows81CertificateProfileBaseCertificateValidityPeriodScale = "Years"
)

func PossibleValuesForWindows81CertificateProfileBaseCertificateValidityPeriodScale() []string {
	return []string{
		string(Windows81CertificateProfileBaseCertificateValidityPeriodScaledays),
		string(Windows81CertificateProfileBaseCertificateValidityPeriodScalemonths),
		string(Windows81CertificateProfileBaseCertificateValidityPeriodScaleyears),
	}
}

func parseWindows81CertificateProfileBaseCertificateValidityPeriodScale(input string) (*Windows81CertificateProfileBaseCertificateValidityPeriodScale, error) {
	vals := map[string]Windows81CertificateProfileBaseCertificateValidityPeriodScale{
		"days":   Windows81CertificateProfileBaseCertificateValidityPeriodScaledays,
		"months": Windows81CertificateProfileBaseCertificateValidityPeriodScalemonths,
		"years":  Windows81CertificateProfileBaseCertificateValidityPeriodScaleyears,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows81CertificateProfileBaseCertificateValidityPeriodScale(input)
	return &out, nil
}
