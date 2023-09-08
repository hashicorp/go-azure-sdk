package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkScepCertificateProfileCertificateValidityPeriodScale string

const (
	AndroidForWorkScepCertificateProfileCertificateValidityPeriodScaledays   AndroidForWorkScepCertificateProfileCertificateValidityPeriodScale = "Days"
	AndroidForWorkScepCertificateProfileCertificateValidityPeriodScalemonths AndroidForWorkScepCertificateProfileCertificateValidityPeriodScale = "Months"
	AndroidForWorkScepCertificateProfileCertificateValidityPeriodScaleyears  AndroidForWorkScepCertificateProfileCertificateValidityPeriodScale = "Years"
)

func PossibleValuesForAndroidForWorkScepCertificateProfileCertificateValidityPeriodScale() []string {
	return []string{
		string(AndroidForWorkScepCertificateProfileCertificateValidityPeriodScaledays),
		string(AndroidForWorkScepCertificateProfileCertificateValidityPeriodScalemonths),
		string(AndroidForWorkScepCertificateProfileCertificateValidityPeriodScaleyears),
	}
}

func parseAndroidForWorkScepCertificateProfileCertificateValidityPeriodScale(input string) (*AndroidForWorkScepCertificateProfileCertificateValidityPeriodScale, error) {
	vals := map[string]AndroidForWorkScepCertificateProfileCertificateValidityPeriodScale{
		"days":   AndroidForWorkScepCertificateProfileCertificateValidityPeriodScaledays,
		"months": AndroidForWorkScepCertificateProfileCertificateValidityPeriodScalemonths,
		"years":  AndroidForWorkScepCertificateProfileCertificateValidityPeriodScaleyears,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkScepCertificateProfileCertificateValidityPeriodScale(input)
	return &out, nil
}
