package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerImportedPFXCertificateProfileCertificateValidityPeriodScale string

const (
	AndroidDeviceOwnerImportedPFXCertificateProfileCertificateValidityPeriodScaledays   AndroidDeviceOwnerImportedPFXCertificateProfileCertificateValidityPeriodScale = "Days"
	AndroidDeviceOwnerImportedPFXCertificateProfileCertificateValidityPeriodScalemonths AndroidDeviceOwnerImportedPFXCertificateProfileCertificateValidityPeriodScale = "Months"
	AndroidDeviceOwnerImportedPFXCertificateProfileCertificateValidityPeriodScaleyears  AndroidDeviceOwnerImportedPFXCertificateProfileCertificateValidityPeriodScale = "Years"
)

func PossibleValuesForAndroidDeviceOwnerImportedPFXCertificateProfileCertificateValidityPeriodScale() []string {
	return []string{
		string(AndroidDeviceOwnerImportedPFXCertificateProfileCertificateValidityPeriodScaledays),
		string(AndroidDeviceOwnerImportedPFXCertificateProfileCertificateValidityPeriodScalemonths),
		string(AndroidDeviceOwnerImportedPFXCertificateProfileCertificateValidityPeriodScaleyears),
	}
}

func parseAndroidDeviceOwnerImportedPFXCertificateProfileCertificateValidityPeriodScale(input string) (*AndroidDeviceOwnerImportedPFXCertificateProfileCertificateValidityPeriodScale, error) {
	vals := map[string]AndroidDeviceOwnerImportedPFXCertificateProfileCertificateValidityPeriodScale{
		"days":   AndroidDeviceOwnerImportedPFXCertificateProfileCertificateValidityPeriodScaledays,
		"months": AndroidDeviceOwnerImportedPFXCertificateProfileCertificateValidityPeriodScalemonths,
		"years":  AndroidDeviceOwnerImportedPFXCertificateProfileCertificateValidityPeriodScaleyears,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerImportedPFXCertificateProfileCertificateValidityPeriodScale(input)
	return &out, nil
}
