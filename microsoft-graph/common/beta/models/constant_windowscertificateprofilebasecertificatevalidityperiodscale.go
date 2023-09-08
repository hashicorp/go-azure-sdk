package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsCertificateProfileBaseCertificateValidityPeriodScale string

const (
	WindowsCertificateProfileBaseCertificateValidityPeriodScaledays   WindowsCertificateProfileBaseCertificateValidityPeriodScale = "Days"
	WindowsCertificateProfileBaseCertificateValidityPeriodScalemonths WindowsCertificateProfileBaseCertificateValidityPeriodScale = "Months"
	WindowsCertificateProfileBaseCertificateValidityPeriodScaleyears  WindowsCertificateProfileBaseCertificateValidityPeriodScale = "Years"
)

func PossibleValuesForWindowsCertificateProfileBaseCertificateValidityPeriodScale() []string {
	return []string{
		string(WindowsCertificateProfileBaseCertificateValidityPeriodScaledays),
		string(WindowsCertificateProfileBaseCertificateValidityPeriodScalemonths),
		string(WindowsCertificateProfileBaseCertificateValidityPeriodScaleyears),
	}
}

func parseWindowsCertificateProfileBaseCertificateValidityPeriodScale(input string) (*WindowsCertificateProfileBaseCertificateValidityPeriodScale, error) {
	vals := map[string]WindowsCertificateProfileBaseCertificateValidityPeriodScale{
		"days":   WindowsCertificateProfileBaseCertificateValidityPeriodScaledays,
		"months": WindowsCertificateProfileBaseCertificateValidityPeriodScalemonths,
		"years":  WindowsCertificateProfileBaseCertificateValidityPeriodScaleyears,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsCertificateProfileBaseCertificateValidityPeriodScale(input)
	return &out, nil
}
