package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10PkcsCertificateProfileCertificateValidityPeriodScale string

const (
	Windows10PkcsCertificateProfileCertificateValidityPeriodScale_Days   Windows10PkcsCertificateProfileCertificateValidityPeriodScale = "days"
	Windows10PkcsCertificateProfileCertificateValidityPeriodScale_Months Windows10PkcsCertificateProfileCertificateValidityPeriodScale = "months"
	Windows10PkcsCertificateProfileCertificateValidityPeriodScale_Years  Windows10PkcsCertificateProfileCertificateValidityPeriodScale = "years"
)

func PossibleValuesForWindows10PkcsCertificateProfileCertificateValidityPeriodScale() []string {
	return []string{
		string(Windows10PkcsCertificateProfileCertificateValidityPeriodScale_Days),
		string(Windows10PkcsCertificateProfileCertificateValidityPeriodScale_Months),
		string(Windows10PkcsCertificateProfileCertificateValidityPeriodScale_Years),
	}
}

func (s *Windows10PkcsCertificateProfileCertificateValidityPeriodScale) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10PkcsCertificateProfileCertificateValidityPeriodScale(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10PkcsCertificateProfileCertificateValidityPeriodScale(input string) (*Windows10PkcsCertificateProfileCertificateValidityPeriodScale, error) {
	vals := map[string]Windows10PkcsCertificateProfileCertificateValidityPeriodScale{
		"days":   Windows10PkcsCertificateProfileCertificateValidityPeriodScale_Days,
		"months": Windows10PkcsCertificateProfileCertificateValidityPeriodScale_Months,
		"years":  Windows10PkcsCertificateProfileCertificateValidityPeriodScale_Years,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10PkcsCertificateProfileCertificateValidityPeriodScale(input)
	return &out, nil
}
