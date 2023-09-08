package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileScepCertificateProfileCertificateValidityPeriodScale string

const (
	AndroidWorkProfileScepCertificateProfileCertificateValidityPeriodScaledays   AndroidWorkProfileScepCertificateProfileCertificateValidityPeriodScale = "Days"
	AndroidWorkProfileScepCertificateProfileCertificateValidityPeriodScalemonths AndroidWorkProfileScepCertificateProfileCertificateValidityPeriodScale = "Months"
	AndroidWorkProfileScepCertificateProfileCertificateValidityPeriodScaleyears  AndroidWorkProfileScepCertificateProfileCertificateValidityPeriodScale = "Years"
)

func PossibleValuesForAndroidWorkProfileScepCertificateProfileCertificateValidityPeriodScale() []string {
	return []string{
		string(AndroidWorkProfileScepCertificateProfileCertificateValidityPeriodScaledays),
		string(AndroidWorkProfileScepCertificateProfileCertificateValidityPeriodScalemonths),
		string(AndroidWorkProfileScepCertificateProfileCertificateValidityPeriodScaleyears),
	}
}

func parseAndroidWorkProfileScepCertificateProfileCertificateValidityPeriodScale(input string) (*AndroidWorkProfileScepCertificateProfileCertificateValidityPeriodScale, error) {
	vals := map[string]AndroidWorkProfileScepCertificateProfileCertificateValidityPeriodScale{
		"days":   AndroidWorkProfileScepCertificateProfileCertificateValidityPeriodScaledays,
		"months": AndroidWorkProfileScepCertificateProfileCertificateValidityPeriodScalemonths,
		"years":  AndroidWorkProfileScepCertificateProfileCertificateValidityPeriodScaleyears,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileScepCertificateProfileCertificateValidityPeriodScale(input)
	return &out, nil
}
