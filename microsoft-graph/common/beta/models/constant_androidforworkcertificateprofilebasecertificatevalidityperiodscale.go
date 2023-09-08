package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkCertificateProfileBaseCertificateValidityPeriodScale string

const (
	AndroidForWorkCertificateProfileBaseCertificateValidityPeriodScaledays   AndroidForWorkCertificateProfileBaseCertificateValidityPeriodScale = "Days"
	AndroidForWorkCertificateProfileBaseCertificateValidityPeriodScalemonths AndroidForWorkCertificateProfileBaseCertificateValidityPeriodScale = "Months"
	AndroidForWorkCertificateProfileBaseCertificateValidityPeriodScaleyears  AndroidForWorkCertificateProfileBaseCertificateValidityPeriodScale = "Years"
)

func PossibleValuesForAndroidForWorkCertificateProfileBaseCertificateValidityPeriodScale() []string {
	return []string{
		string(AndroidForWorkCertificateProfileBaseCertificateValidityPeriodScaledays),
		string(AndroidForWorkCertificateProfileBaseCertificateValidityPeriodScalemonths),
		string(AndroidForWorkCertificateProfileBaseCertificateValidityPeriodScaleyears),
	}
}

func parseAndroidForWorkCertificateProfileBaseCertificateValidityPeriodScale(input string) (*AndroidForWorkCertificateProfileBaseCertificateValidityPeriodScale, error) {
	vals := map[string]AndroidForWorkCertificateProfileBaseCertificateValidityPeriodScale{
		"days":   AndroidForWorkCertificateProfileBaseCertificateValidityPeriodScaledays,
		"months": AndroidForWorkCertificateProfileBaseCertificateValidityPeriodScalemonths,
		"years":  AndroidForWorkCertificateProfileBaseCertificateValidityPeriodScaleyears,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkCertificateProfileBaseCertificateValidityPeriodScale(input)
	return &out, nil
}
