package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkPkcsCertificateProfileCertificateValidityPeriodScale string

const (
	AndroidForWorkPkcsCertificateProfileCertificateValidityPeriodScaledays   AndroidForWorkPkcsCertificateProfileCertificateValidityPeriodScale = "Days"
	AndroidForWorkPkcsCertificateProfileCertificateValidityPeriodScalemonths AndroidForWorkPkcsCertificateProfileCertificateValidityPeriodScale = "Months"
	AndroidForWorkPkcsCertificateProfileCertificateValidityPeriodScaleyears  AndroidForWorkPkcsCertificateProfileCertificateValidityPeriodScale = "Years"
)

func PossibleValuesForAndroidForWorkPkcsCertificateProfileCertificateValidityPeriodScale() []string {
	return []string{
		string(AndroidForWorkPkcsCertificateProfileCertificateValidityPeriodScaledays),
		string(AndroidForWorkPkcsCertificateProfileCertificateValidityPeriodScalemonths),
		string(AndroidForWorkPkcsCertificateProfileCertificateValidityPeriodScaleyears),
	}
}

func parseAndroidForWorkPkcsCertificateProfileCertificateValidityPeriodScale(input string) (*AndroidForWorkPkcsCertificateProfileCertificateValidityPeriodScale, error) {
	vals := map[string]AndroidForWorkPkcsCertificateProfileCertificateValidityPeriodScale{
		"days":   AndroidForWorkPkcsCertificateProfileCertificateValidityPeriodScaledays,
		"months": AndroidForWorkPkcsCertificateProfileCertificateValidityPeriodScalemonths,
		"years":  AndroidForWorkPkcsCertificateProfileCertificateValidityPeriodScaleyears,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkPkcsCertificateProfileCertificateValidityPeriodScale(input)
	return &out, nil
}
