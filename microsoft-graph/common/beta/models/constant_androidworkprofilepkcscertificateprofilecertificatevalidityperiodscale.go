package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfilePkcsCertificateProfileCertificateValidityPeriodScale string

const (
	AndroidWorkProfilePkcsCertificateProfileCertificateValidityPeriodScaledays   AndroidWorkProfilePkcsCertificateProfileCertificateValidityPeriodScale = "Days"
	AndroidWorkProfilePkcsCertificateProfileCertificateValidityPeriodScalemonths AndroidWorkProfilePkcsCertificateProfileCertificateValidityPeriodScale = "Months"
	AndroidWorkProfilePkcsCertificateProfileCertificateValidityPeriodScaleyears  AndroidWorkProfilePkcsCertificateProfileCertificateValidityPeriodScale = "Years"
)

func PossibleValuesForAndroidWorkProfilePkcsCertificateProfileCertificateValidityPeriodScale() []string {
	return []string{
		string(AndroidWorkProfilePkcsCertificateProfileCertificateValidityPeriodScaledays),
		string(AndroidWorkProfilePkcsCertificateProfileCertificateValidityPeriodScalemonths),
		string(AndroidWorkProfilePkcsCertificateProfileCertificateValidityPeriodScaleyears),
	}
}

func parseAndroidWorkProfilePkcsCertificateProfileCertificateValidityPeriodScale(input string) (*AndroidWorkProfilePkcsCertificateProfileCertificateValidityPeriodScale, error) {
	vals := map[string]AndroidWorkProfilePkcsCertificateProfileCertificateValidityPeriodScale{
		"days":   AndroidWorkProfilePkcsCertificateProfileCertificateValidityPeriodScaledays,
		"months": AndroidWorkProfilePkcsCertificateProfileCertificateValidityPeriodScalemonths,
		"years":  AndroidWorkProfilePkcsCertificateProfileCertificateValidityPeriodScaleyears,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfilePkcsCertificateProfileCertificateValidityPeriodScale(input)
	return &out, nil
}
