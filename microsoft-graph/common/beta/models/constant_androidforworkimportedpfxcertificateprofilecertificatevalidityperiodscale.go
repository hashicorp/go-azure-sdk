package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkImportedPFXCertificateProfileCertificateValidityPeriodScale string

const (
	AndroidForWorkImportedPFXCertificateProfileCertificateValidityPeriodScaledays   AndroidForWorkImportedPFXCertificateProfileCertificateValidityPeriodScale = "Days"
	AndroidForWorkImportedPFXCertificateProfileCertificateValidityPeriodScalemonths AndroidForWorkImportedPFXCertificateProfileCertificateValidityPeriodScale = "Months"
	AndroidForWorkImportedPFXCertificateProfileCertificateValidityPeriodScaleyears  AndroidForWorkImportedPFXCertificateProfileCertificateValidityPeriodScale = "Years"
)

func PossibleValuesForAndroidForWorkImportedPFXCertificateProfileCertificateValidityPeriodScale() []string {
	return []string{
		string(AndroidForWorkImportedPFXCertificateProfileCertificateValidityPeriodScaledays),
		string(AndroidForWorkImportedPFXCertificateProfileCertificateValidityPeriodScalemonths),
		string(AndroidForWorkImportedPFXCertificateProfileCertificateValidityPeriodScaleyears),
	}
}

func parseAndroidForWorkImportedPFXCertificateProfileCertificateValidityPeriodScale(input string) (*AndroidForWorkImportedPFXCertificateProfileCertificateValidityPeriodScale, error) {
	vals := map[string]AndroidForWorkImportedPFXCertificateProfileCertificateValidityPeriodScale{
		"days":   AndroidForWorkImportedPFXCertificateProfileCertificateValidityPeriodScaledays,
		"months": AndroidForWorkImportedPFXCertificateProfileCertificateValidityPeriodScalemonths,
		"years":  AndroidForWorkImportedPFXCertificateProfileCertificateValidityPeriodScaleyears,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkImportedPFXCertificateProfileCertificateValidityPeriodScale(input)
	return &out, nil
}
