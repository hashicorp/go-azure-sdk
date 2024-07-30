package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10CertificateProfileBaseCertificateValidityPeriodScale string

const (
	Windows10CertificateProfileBaseCertificateValidityPeriodScale_Days   Windows10CertificateProfileBaseCertificateValidityPeriodScale = "days"
	Windows10CertificateProfileBaseCertificateValidityPeriodScale_Months Windows10CertificateProfileBaseCertificateValidityPeriodScale = "months"
	Windows10CertificateProfileBaseCertificateValidityPeriodScale_Years  Windows10CertificateProfileBaseCertificateValidityPeriodScale = "years"
)

func PossibleValuesForWindows10CertificateProfileBaseCertificateValidityPeriodScale() []string {
	return []string{
		string(Windows10CertificateProfileBaseCertificateValidityPeriodScale_Days),
		string(Windows10CertificateProfileBaseCertificateValidityPeriodScale_Months),
		string(Windows10CertificateProfileBaseCertificateValidityPeriodScale_Years),
	}
}

func (s *Windows10CertificateProfileBaseCertificateValidityPeriodScale) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10CertificateProfileBaseCertificateValidityPeriodScale(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10CertificateProfileBaseCertificateValidityPeriodScale(input string) (*Windows10CertificateProfileBaseCertificateValidityPeriodScale, error) {
	vals := map[string]Windows10CertificateProfileBaseCertificateValidityPeriodScale{
		"days":   Windows10CertificateProfileBaseCertificateValidityPeriodScale_Days,
		"months": Windows10CertificateProfileBaseCertificateValidityPeriodScale_Months,
		"years":  Windows10CertificateProfileBaseCertificateValidityPeriodScale_Years,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10CertificateProfileBaseCertificateValidityPeriodScale(input)
	return &out, nil
}
