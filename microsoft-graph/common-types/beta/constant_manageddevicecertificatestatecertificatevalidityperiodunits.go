package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceCertificateStateCertificateValidityPeriodUnits string

const (
	ManagedDeviceCertificateStateCertificateValidityPeriodUnits_Days   ManagedDeviceCertificateStateCertificateValidityPeriodUnits = "days"
	ManagedDeviceCertificateStateCertificateValidityPeriodUnits_Months ManagedDeviceCertificateStateCertificateValidityPeriodUnits = "months"
	ManagedDeviceCertificateStateCertificateValidityPeriodUnits_Years  ManagedDeviceCertificateStateCertificateValidityPeriodUnits = "years"
)

func PossibleValuesForManagedDeviceCertificateStateCertificateValidityPeriodUnits() []string {
	return []string{
		string(ManagedDeviceCertificateStateCertificateValidityPeriodUnits_Days),
		string(ManagedDeviceCertificateStateCertificateValidityPeriodUnits_Months),
		string(ManagedDeviceCertificateStateCertificateValidityPeriodUnits_Years),
	}
}

func (s *ManagedDeviceCertificateStateCertificateValidityPeriodUnits) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedDeviceCertificateStateCertificateValidityPeriodUnits(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedDeviceCertificateStateCertificateValidityPeriodUnits(input string) (*ManagedDeviceCertificateStateCertificateValidityPeriodUnits, error) {
	vals := map[string]ManagedDeviceCertificateStateCertificateValidityPeriodUnits{
		"days":   ManagedDeviceCertificateStateCertificateValidityPeriodUnits_Days,
		"months": ManagedDeviceCertificateStateCertificateValidityPeriodUnits_Months,
		"years":  ManagedDeviceCertificateStateCertificateValidityPeriodUnits_Years,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceCertificateStateCertificateValidityPeriodUnits(input)
	return &out, nil
}
