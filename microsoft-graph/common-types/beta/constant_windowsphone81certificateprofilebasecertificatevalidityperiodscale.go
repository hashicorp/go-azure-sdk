package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPhone81CertificateProfileBaseCertificateValidityPeriodScale string

const (
	WindowsPhone81CertificateProfileBaseCertificateValidityPeriodScale_Days   WindowsPhone81CertificateProfileBaseCertificateValidityPeriodScale = "days"
	WindowsPhone81CertificateProfileBaseCertificateValidityPeriodScale_Months WindowsPhone81CertificateProfileBaseCertificateValidityPeriodScale = "months"
	WindowsPhone81CertificateProfileBaseCertificateValidityPeriodScale_Years  WindowsPhone81CertificateProfileBaseCertificateValidityPeriodScale = "years"
)

func PossibleValuesForWindowsPhone81CertificateProfileBaseCertificateValidityPeriodScale() []string {
	return []string{
		string(WindowsPhone81CertificateProfileBaseCertificateValidityPeriodScale_Days),
		string(WindowsPhone81CertificateProfileBaseCertificateValidityPeriodScale_Months),
		string(WindowsPhone81CertificateProfileBaseCertificateValidityPeriodScale_Years),
	}
}

func (s *WindowsPhone81CertificateProfileBaseCertificateValidityPeriodScale) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsPhone81CertificateProfileBaseCertificateValidityPeriodScale(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsPhone81CertificateProfileBaseCertificateValidityPeriodScale(input string) (*WindowsPhone81CertificateProfileBaseCertificateValidityPeriodScale, error) {
	vals := map[string]WindowsPhone81CertificateProfileBaseCertificateValidityPeriodScale{
		"days":   WindowsPhone81CertificateProfileBaseCertificateValidityPeriodScale_Days,
		"months": WindowsPhone81CertificateProfileBaseCertificateValidityPeriodScale_Months,
		"years":  WindowsPhone81CertificateProfileBaseCertificateValidityPeriodScale_Years,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPhone81CertificateProfileBaseCertificateValidityPeriodScale(input)
	return &out, nil
}
