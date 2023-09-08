package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceCertificateStateCertificateValidityPeriodUnits string

const (
	ManagedDeviceCertificateStateCertificateValidityPeriodUnitsdays   ManagedDeviceCertificateStateCertificateValidityPeriodUnits = "Days"
	ManagedDeviceCertificateStateCertificateValidityPeriodUnitsmonths ManagedDeviceCertificateStateCertificateValidityPeriodUnits = "Months"
	ManagedDeviceCertificateStateCertificateValidityPeriodUnitsyears  ManagedDeviceCertificateStateCertificateValidityPeriodUnits = "Years"
)

func PossibleValuesForManagedDeviceCertificateStateCertificateValidityPeriodUnits() []string {
	return []string{
		string(ManagedDeviceCertificateStateCertificateValidityPeriodUnitsdays),
		string(ManagedDeviceCertificateStateCertificateValidityPeriodUnitsmonths),
		string(ManagedDeviceCertificateStateCertificateValidityPeriodUnitsyears),
	}
}

func parseManagedDeviceCertificateStateCertificateValidityPeriodUnits(input string) (*ManagedDeviceCertificateStateCertificateValidityPeriodUnits, error) {
	vals := map[string]ManagedDeviceCertificateStateCertificateValidityPeriodUnits{
		"days":   ManagedDeviceCertificateStateCertificateValidityPeriodUnitsdays,
		"months": ManagedDeviceCertificateStateCertificateValidityPeriodUnitsmonths,
		"years":  ManagedDeviceCertificateStateCertificateValidityPeriodUnitsyears,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceCertificateStateCertificateValidityPeriodUnits(input)
	return &out, nil
}
