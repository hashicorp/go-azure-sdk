package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AospDeviceOwnerScepCertificateProfileCertificateValidityPeriodScale string

const (
	AospDeviceOwnerScepCertificateProfileCertificateValidityPeriodScaledays   AospDeviceOwnerScepCertificateProfileCertificateValidityPeriodScale = "Days"
	AospDeviceOwnerScepCertificateProfileCertificateValidityPeriodScalemonths AospDeviceOwnerScepCertificateProfileCertificateValidityPeriodScale = "Months"
	AospDeviceOwnerScepCertificateProfileCertificateValidityPeriodScaleyears  AospDeviceOwnerScepCertificateProfileCertificateValidityPeriodScale = "Years"
)

func PossibleValuesForAospDeviceOwnerScepCertificateProfileCertificateValidityPeriodScale() []string {
	return []string{
		string(AospDeviceOwnerScepCertificateProfileCertificateValidityPeriodScaledays),
		string(AospDeviceOwnerScepCertificateProfileCertificateValidityPeriodScalemonths),
		string(AospDeviceOwnerScepCertificateProfileCertificateValidityPeriodScaleyears),
	}
}

func parseAospDeviceOwnerScepCertificateProfileCertificateValidityPeriodScale(input string) (*AospDeviceOwnerScepCertificateProfileCertificateValidityPeriodScale, error) {
	vals := map[string]AospDeviceOwnerScepCertificateProfileCertificateValidityPeriodScale{
		"days":   AospDeviceOwnerScepCertificateProfileCertificateValidityPeriodScaledays,
		"months": AospDeviceOwnerScepCertificateProfileCertificateValidityPeriodScalemonths,
		"years":  AospDeviceOwnerScepCertificateProfileCertificateValidityPeriodScaleyears,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AospDeviceOwnerScepCertificateProfileCertificateValidityPeriodScale(input)
	return &out, nil
}
