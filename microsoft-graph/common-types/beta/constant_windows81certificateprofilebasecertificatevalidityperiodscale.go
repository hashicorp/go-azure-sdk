package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows81CertificateProfileBaseCertificateValidityPeriodScale string

const (
	Windows81CertificateProfileBaseCertificateValidityPeriodScale_Days   Windows81CertificateProfileBaseCertificateValidityPeriodScale = "days"
	Windows81CertificateProfileBaseCertificateValidityPeriodScale_Months Windows81CertificateProfileBaseCertificateValidityPeriodScale = "months"
	Windows81CertificateProfileBaseCertificateValidityPeriodScale_Years  Windows81CertificateProfileBaseCertificateValidityPeriodScale = "years"
)

func PossibleValuesForWindows81CertificateProfileBaseCertificateValidityPeriodScale() []string {
	return []string{
		string(Windows81CertificateProfileBaseCertificateValidityPeriodScale_Days),
		string(Windows81CertificateProfileBaseCertificateValidityPeriodScale_Months),
		string(Windows81CertificateProfileBaseCertificateValidityPeriodScale_Years),
	}
}

func (s *Windows81CertificateProfileBaseCertificateValidityPeriodScale) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows81CertificateProfileBaseCertificateValidityPeriodScale(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows81CertificateProfileBaseCertificateValidityPeriodScale(input string) (*Windows81CertificateProfileBaseCertificateValidityPeriodScale, error) {
	vals := map[string]Windows81CertificateProfileBaseCertificateValidityPeriodScale{
		"days":   Windows81CertificateProfileBaseCertificateValidityPeriodScale_Days,
		"months": Windows81CertificateProfileBaseCertificateValidityPeriodScale_Months,
		"years":  Windows81CertificateProfileBaseCertificateValidityPeriodScale_Years,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows81CertificateProfileBaseCertificateValidityPeriodScale(input)
	return &out, nil
}
