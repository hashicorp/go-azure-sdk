package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSCertificateProfileBaseCertificateValidityPeriodScale string

const (
	MacOSCertificateProfileBaseCertificateValidityPeriodScale_Days   MacOSCertificateProfileBaseCertificateValidityPeriodScale = "days"
	MacOSCertificateProfileBaseCertificateValidityPeriodScale_Months MacOSCertificateProfileBaseCertificateValidityPeriodScale = "months"
	MacOSCertificateProfileBaseCertificateValidityPeriodScale_Years  MacOSCertificateProfileBaseCertificateValidityPeriodScale = "years"
)

func PossibleValuesForMacOSCertificateProfileBaseCertificateValidityPeriodScale() []string {
	return []string{
		string(MacOSCertificateProfileBaseCertificateValidityPeriodScale_Days),
		string(MacOSCertificateProfileBaseCertificateValidityPeriodScale_Months),
		string(MacOSCertificateProfileBaseCertificateValidityPeriodScale_Years),
	}
}

func (s *MacOSCertificateProfileBaseCertificateValidityPeriodScale) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSCertificateProfileBaseCertificateValidityPeriodScale(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSCertificateProfileBaseCertificateValidityPeriodScale(input string) (*MacOSCertificateProfileBaseCertificateValidityPeriodScale, error) {
	vals := map[string]MacOSCertificateProfileBaseCertificateValidityPeriodScale{
		"days":   MacOSCertificateProfileBaseCertificateValidityPeriodScale_Days,
		"months": MacOSCertificateProfileBaseCertificateValidityPeriodScale_Months,
		"years":  MacOSCertificateProfileBaseCertificateValidityPeriodScale_Years,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSCertificateProfileBaseCertificateValidityPeriodScale(input)
	return &out, nil
}
