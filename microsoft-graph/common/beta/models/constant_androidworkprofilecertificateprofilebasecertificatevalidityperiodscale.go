package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileCertificateProfileBaseCertificateValidityPeriodScale string

const (
	AndroidWorkProfileCertificateProfileBaseCertificateValidityPeriodScaledays   AndroidWorkProfileCertificateProfileBaseCertificateValidityPeriodScale = "Days"
	AndroidWorkProfileCertificateProfileBaseCertificateValidityPeriodScalemonths AndroidWorkProfileCertificateProfileBaseCertificateValidityPeriodScale = "Months"
	AndroidWorkProfileCertificateProfileBaseCertificateValidityPeriodScaleyears  AndroidWorkProfileCertificateProfileBaseCertificateValidityPeriodScale = "Years"
)

func PossibleValuesForAndroidWorkProfileCertificateProfileBaseCertificateValidityPeriodScale() []string {
	return []string{
		string(AndroidWorkProfileCertificateProfileBaseCertificateValidityPeriodScaledays),
		string(AndroidWorkProfileCertificateProfileBaseCertificateValidityPeriodScalemonths),
		string(AndroidWorkProfileCertificateProfileBaseCertificateValidityPeriodScaleyears),
	}
}

func parseAndroidWorkProfileCertificateProfileBaseCertificateValidityPeriodScale(input string) (*AndroidWorkProfileCertificateProfileBaseCertificateValidityPeriodScale, error) {
	vals := map[string]AndroidWorkProfileCertificateProfileBaseCertificateValidityPeriodScale{
		"days":   AndroidWorkProfileCertificateProfileBaseCertificateValidityPeriodScaledays,
		"months": AndroidWorkProfileCertificateProfileBaseCertificateValidityPeriodScalemonths,
		"years":  AndroidWorkProfileCertificateProfileBaseCertificateValidityPeriodScaleyears,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileCertificateProfileBaseCertificateValidityPeriodScale(input)
	return &out, nil
}
