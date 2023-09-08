package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScale string

const (
	AndroidDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScaledays   AndroidDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScale = "Days"
	AndroidDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScalemonths AndroidDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScale = "Months"
	AndroidDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScaleyears  AndroidDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScale = "Years"
)

func PossibleValuesForAndroidDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScale() []string {
	return []string{
		string(AndroidDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScaledays),
		string(AndroidDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScalemonths),
		string(AndroidDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScaleyears),
	}
}

func parseAndroidDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScale(input string) (*AndroidDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScale, error) {
	vals := map[string]AndroidDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScale{
		"days":   AndroidDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScaledays,
		"months": AndroidDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScalemonths,
		"years":  AndroidDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScaleyears,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScale(input)
	return &out, nil
}
