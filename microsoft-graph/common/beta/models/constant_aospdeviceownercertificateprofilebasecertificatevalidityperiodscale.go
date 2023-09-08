package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AospDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScale string

const (
	AospDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScaledays   AospDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScale = "Days"
	AospDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScalemonths AospDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScale = "Months"
	AospDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScaleyears  AospDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScale = "Years"
)

func PossibleValuesForAospDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScale() []string {
	return []string{
		string(AospDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScaledays),
		string(AospDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScalemonths),
		string(AospDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScaleyears),
	}
}

func parseAospDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScale(input string) (*AospDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScale, error) {
	vals := map[string]AospDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScale{
		"days":   AospDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScaledays,
		"months": AospDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScalemonths,
		"years":  AospDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScaleyears,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AospDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScale(input)
	return &out, nil
}
