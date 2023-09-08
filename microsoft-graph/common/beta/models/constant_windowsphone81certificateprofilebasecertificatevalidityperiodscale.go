package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPhone81CertificateProfileBaseCertificateValidityPeriodScale string

const (
	WindowsPhone81CertificateProfileBaseCertificateValidityPeriodScaledays   WindowsPhone81CertificateProfileBaseCertificateValidityPeriodScale = "Days"
	WindowsPhone81CertificateProfileBaseCertificateValidityPeriodScalemonths WindowsPhone81CertificateProfileBaseCertificateValidityPeriodScale = "Months"
	WindowsPhone81CertificateProfileBaseCertificateValidityPeriodScaleyears  WindowsPhone81CertificateProfileBaseCertificateValidityPeriodScale = "Years"
)

func PossibleValuesForWindowsPhone81CertificateProfileBaseCertificateValidityPeriodScale() []string {
	return []string{
		string(WindowsPhone81CertificateProfileBaseCertificateValidityPeriodScaledays),
		string(WindowsPhone81CertificateProfileBaseCertificateValidityPeriodScalemonths),
		string(WindowsPhone81CertificateProfileBaseCertificateValidityPeriodScaleyears),
	}
}

func parseWindowsPhone81CertificateProfileBaseCertificateValidityPeriodScale(input string) (*WindowsPhone81CertificateProfileBaseCertificateValidityPeriodScale, error) {
	vals := map[string]WindowsPhone81CertificateProfileBaseCertificateValidityPeriodScale{
		"days":   WindowsPhone81CertificateProfileBaseCertificateValidityPeriodScaledays,
		"months": WindowsPhone81CertificateProfileBaseCertificateValidityPeriodScalemonths,
		"years":  WindowsPhone81CertificateProfileBaseCertificateValidityPeriodScaleyears,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPhone81CertificateProfileBaseCertificateValidityPeriodScale(input)
	return &out, nil
}
