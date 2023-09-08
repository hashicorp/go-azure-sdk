package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSScepCertificateProfileCertificateValidityPeriodScale string

const (
	MacOSScepCertificateProfileCertificateValidityPeriodScaledays   MacOSScepCertificateProfileCertificateValidityPeriodScale = "Days"
	MacOSScepCertificateProfileCertificateValidityPeriodScalemonths MacOSScepCertificateProfileCertificateValidityPeriodScale = "Months"
	MacOSScepCertificateProfileCertificateValidityPeriodScaleyears  MacOSScepCertificateProfileCertificateValidityPeriodScale = "Years"
)

func PossibleValuesForMacOSScepCertificateProfileCertificateValidityPeriodScale() []string {
	return []string{
		string(MacOSScepCertificateProfileCertificateValidityPeriodScaledays),
		string(MacOSScepCertificateProfileCertificateValidityPeriodScalemonths),
		string(MacOSScepCertificateProfileCertificateValidityPeriodScaleyears),
	}
}

func parseMacOSScepCertificateProfileCertificateValidityPeriodScale(input string) (*MacOSScepCertificateProfileCertificateValidityPeriodScale, error) {
	vals := map[string]MacOSScepCertificateProfileCertificateValidityPeriodScale{
		"days":   MacOSScepCertificateProfileCertificateValidityPeriodScaledays,
		"months": MacOSScepCertificateProfileCertificateValidityPeriodScalemonths,
		"years":  MacOSScepCertificateProfileCertificateValidityPeriodScaleyears,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSScepCertificateProfileCertificateValidityPeriodScale(input)
	return &out, nil
}
