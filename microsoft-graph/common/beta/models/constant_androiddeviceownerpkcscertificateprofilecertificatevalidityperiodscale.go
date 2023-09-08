package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScale string

const (
	AndroidDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScaledays   AndroidDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScale = "Days"
	AndroidDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScalemonths AndroidDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScale = "Months"
	AndroidDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScaleyears  AndroidDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScale = "Years"
)

func PossibleValuesForAndroidDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScale() []string {
	return []string{
		string(AndroidDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScaledays),
		string(AndroidDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScalemonths),
		string(AndroidDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScaleyears),
	}
}

func parseAndroidDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScale(input string) (*AndroidDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScale, error) {
	vals := map[string]AndroidDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScale{
		"days":   AndroidDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScaledays,
		"months": AndroidDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScalemonths,
		"years":  AndroidDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScaleyears,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScale(input)
	return &out, nil
}
