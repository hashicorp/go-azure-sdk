package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidPkcsCertificateProfileCertificateValidityPeriodScale string

const (
	AndroidPkcsCertificateProfileCertificateValidityPeriodScaledays   AndroidPkcsCertificateProfileCertificateValidityPeriodScale = "Days"
	AndroidPkcsCertificateProfileCertificateValidityPeriodScalemonths AndroidPkcsCertificateProfileCertificateValidityPeriodScale = "Months"
	AndroidPkcsCertificateProfileCertificateValidityPeriodScaleyears  AndroidPkcsCertificateProfileCertificateValidityPeriodScale = "Years"
)

func PossibleValuesForAndroidPkcsCertificateProfileCertificateValidityPeriodScale() []string {
	return []string{
		string(AndroidPkcsCertificateProfileCertificateValidityPeriodScaledays),
		string(AndroidPkcsCertificateProfileCertificateValidityPeriodScalemonths),
		string(AndroidPkcsCertificateProfileCertificateValidityPeriodScaleyears),
	}
}

func parseAndroidPkcsCertificateProfileCertificateValidityPeriodScale(input string) (*AndroidPkcsCertificateProfileCertificateValidityPeriodScale, error) {
	vals := map[string]AndroidPkcsCertificateProfileCertificateValidityPeriodScale{
		"days":   AndroidPkcsCertificateProfileCertificateValidityPeriodScaledays,
		"months": AndroidPkcsCertificateProfileCertificateValidityPeriodScalemonths,
		"years":  AndroidPkcsCertificateProfileCertificateValidityPeriodScaleyears,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidPkcsCertificateProfileCertificateValidityPeriodScale(input)
	return &out, nil
}
