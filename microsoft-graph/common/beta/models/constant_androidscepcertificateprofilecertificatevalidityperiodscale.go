package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidScepCertificateProfileCertificateValidityPeriodScale string

const (
	AndroidScepCertificateProfileCertificateValidityPeriodScaledays   AndroidScepCertificateProfileCertificateValidityPeriodScale = "Days"
	AndroidScepCertificateProfileCertificateValidityPeriodScalemonths AndroidScepCertificateProfileCertificateValidityPeriodScale = "Months"
	AndroidScepCertificateProfileCertificateValidityPeriodScaleyears  AndroidScepCertificateProfileCertificateValidityPeriodScale = "Years"
)

func PossibleValuesForAndroidScepCertificateProfileCertificateValidityPeriodScale() []string {
	return []string{
		string(AndroidScepCertificateProfileCertificateValidityPeriodScaledays),
		string(AndroidScepCertificateProfileCertificateValidityPeriodScalemonths),
		string(AndroidScepCertificateProfileCertificateValidityPeriodScaleyears),
	}
}

func parseAndroidScepCertificateProfileCertificateValidityPeriodScale(input string) (*AndroidScepCertificateProfileCertificateValidityPeriodScale, error) {
	vals := map[string]AndroidScepCertificateProfileCertificateValidityPeriodScale{
		"days":   AndroidScepCertificateProfileCertificateValidityPeriodScaledays,
		"months": AndroidScepCertificateProfileCertificateValidityPeriodScalemonths,
		"years":  AndroidScepCertificateProfileCertificateValidityPeriodScaleyears,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidScepCertificateProfileCertificateValidityPeriodScale(input)
	return &out, nil
}
