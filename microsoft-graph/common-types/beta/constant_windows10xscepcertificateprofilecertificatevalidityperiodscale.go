package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10XSCEPCertificateProfileCertificateValidityPeriodScale string

const (
	Windows10XSCEPCertificateProfileCertificateValidityPeriodScale_Days   Windows10XSCEPCertificateProfileCertificateValidityPeriodScale = "days"
	Windows10XSCEPCertificateProfileCertificateValidityPeriodScale_Months Windows10XSCEPCertificateProfileCertificateValidityPeriodScale = "months"
	Windows10XSCEPCertificateProfileCertificateValidityPeriodScale_Years  Windows10XSCEPCertificateProfileCertificateValidityPeriodScale = "years"
)

func PossibleValuesForWindows10XSCEPCertificateProfileCertificateValidityPeriodScale() []string {
	return []string{
		string(Windows10XSCEPCertificateProfileCertificateValidityPeriodScale_Days),
		string(Windows10XSCEPCertificateProfileCertificateValidityPeriodScale_Months),
		string(Windows10XSCEPCertificateProfileCertificateValidityPeriodScale_Years),
	}
}

func (s *Windows10XSCEPCertificateProfileCertificateValidityPeriodScale) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10XSCEPCertificateProfileCertificateValidityPeriodScale(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10XSCEPCertificateProfileCertificateValidityPeriodScale(input string) (*Windows10XSCEPCertificateProfileCertificateValidityPeriodScale, error) {
	vals := map[string]Windows10XSCEPCertificateProfileCertificateValidityPeriodScale{
		"days":   Windows10XSCEPCertificateProfileCertificateValidityPeriodScale_Days,
		"months": Windows10XSCEPCertificateProfileCertificateValidityPeriodScale_Months,
		"years":  Windows10XSCEPCertificateProfileCertificateValidityPeriodScale_Years,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10XSCEPCertificateProfileCertificateValidityPeriodScale(input)
	return &out, nil
}
