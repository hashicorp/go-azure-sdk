package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidCertificateProfileBaseCertificateValidityPeriodScale string

const (
	AndroidCertificateProfileBaseCertificateValidityPeriodScaledays   AndroidCertificateProfileBaseCertificateValidityPeriodScale = "Days"
	AndroidCertificateProfileBaseCertificateValidityPeriodScalemonths AndroidCertificateProfileBaseCertificateValidityPeriodScale = "Months"
	AndroidCertificateProfileBaseCertificateValidityPeriodScaleyears  AndroidCertificateProfileBaseCertificateValidityPeriodScale = "Years"
)

func PossibleValuesForAndroidCertificateProfileBaseCertificateValidityPeriodScale() []string {
	return []string{
		string(AndroidCertificateProfileBaseCertificateValidityPeriodScaledays),
		string(AndroidCertificateProfileBaseCertificateValidityPeriodScalemonths),
		string(AndroidCertificateProfileBaseCertificateValidityPeriodScaleyears),
	}
}

func parseAndroidCertificateProfileBaseCertificateValidityPeriodScale(input string) (*AndroidCertificateProfileBaseCertificateValidityPeriodScale, error) {
	vals := map[string]AndroidCertificateProfileBaseCertificateValidityPeriodScale{
		"days":   AndroidCertificateProfileBaseCertificateValidityPeriodScaledays,
		"months": AndroidCertificateProfileBaseCertificateValidityPeriodScalemonths,
		"years":  AndroidCertificateProfileBaseCertificateValidityPeriodScaleyears,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidCertificateProfileBaseCertificateValidityPeriodScale(input)
	return &out, nil
}
