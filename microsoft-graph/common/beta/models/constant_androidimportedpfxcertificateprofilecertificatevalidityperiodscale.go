package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidImportedPFXCertificateProfileCertificateValidityPeriodScale string

const (
	AndroidImportedPFXCertificateProfileCertificateValidityPeriodScaledays   AndroidImportedPFXCertificateProfileCertificateValidityPeriodScale = "Days"
	AndroidImportedPFXCertificateProfileCertificateValidityPeriodScalemonths AndroidImportedPFXCertificateProfileCertificateValidityPeriodScale = "Months"
	AndroidImportedPFXCertificateProfileCertificateValidityPeriodScaleyears  AndroidImportedPFXCertificateProfileCertificateValidityPeriodScale = "Years"
)

func PossibleValuesForAndroidImportedPFXCertificateProfileCertificateValidityPeriodScale() []string {
	return []string{
		string(AndroidImportedPFXCertificateProfileCertificateValidityPeriodScaledays),
		string(AndroidImportedPFXCertificateProfileCertificateValidityPeriodScalemonths),
		string(AndroidImportedPFXCertificateProfileCertificateValidityPeriodScaleyears),
	}
}

func parseAndroidImportedPFXCertificateProfileCertificateValidityPeriodScale(input string) (*AndroidImportedPFXCertificateProfileCertificateValidityPeriodScale, error) {
	vals := map[string]AndroidImportedPFXCertificateProfileCertificateValidityPeriodScale{
		"days":   AndroidImportedPFXCertificateProfileCertificateValidityPeriodScaledays,
		"months": AndroidImportedPFXCertificateProfileCertificateValidityPeriodScalemonths,
		"years":  AndroidImportedPFXCertificateProfileCertificateValidityPeriodScaleyears,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidImportedPFXCertificateProfileCertificateValidityPeriodScale(input)
	return &out, nil
}
