package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AospDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScale string

const (
	AospDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScaledays   AospDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScale = "Days"
	AospDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScalemonths AospDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScale = "Months"
	AospDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScaleyears  AospDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScale = "Years"
)

func PossibleValuesForAospDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScale() []string {
	return []string{
		string(AospDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScaledays),
		string(AospDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScalemonths),
		string(AospDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScaleyears),
	}
}

func parseAospDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScale(input string) (*AospDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScale, error) {
	vals := map[string]AospDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScale{
		"days":   AospDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScaledays,
		"months": AospDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScalemonths,
		"years":  AospDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScaleyears,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AospDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScale(input)
	return &out, nil
}
