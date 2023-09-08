package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerScepCertificateProfileCertificateValidityPeriodScale string

const (
	AndroidDeviceOwnerScepCertificateProfileCertificateValidityPeriodScaledays   AndroidDeviceOwnerScepCertificateProfileCertificateValidityPeriodScale = "Days"
	AndroidDeviceOwnerScepCertificateProfileCertificateValidityPeriodScalemonths AndroidDeviceOwnerScepCertificateProfileCertificateValidityPeriodScale = "Months"
	AndroidDeviceOwnerScepCertificateProfileCertificateValidityPeriodScaleyears  AndroidDeviceOwnerScepCertificateProfileCertificateValidityPeriodScale = "Years"
)

func PossibleValuesForAndroidDeviceOwnerScepCertificateProfileCertificateValidityPeriodScale() []string {
	return []string{
		string(AndroidDeviceOwnerScepCertificateProfileCertificateValidityPeriodScaledays),
		string(AndroidDeviceOwnerScepCertificateProfileCertificateValidityPeriodScalemonths),
		string(AndroidDeviceOwnerScepCertificateProfileCertificateValidityPeriodScaleyears),
	}
}

func parseAndroidDeviceOwnerScepCertificateProfileCertificateValidityPeriodScale(input string) (*AndroidDeviceOwnerScepCertificateProfileCertificateValidityPeriodScale, error) {
	vals := map[string]AndroidDeviceOwnerScepCertificateProfileCertificateValidityPeriodScale{
		"days":   AndroidDeviceOwnerScepCertificateProfileCertificateValidityPeriodScaledays,
		"months": AndroidDeviceOwnerScepCertificateProfileCertificateValidityPeriodScalemonths,
		"years":  AndroidDeviceOwnerScepCertificateProfileCertificateValidityPeriodScaleyears,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerScepCertificateProfileCertificateValidityPeriodScale(input)
	return &out, nil
}
