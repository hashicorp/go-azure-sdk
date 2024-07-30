package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10ImportedPFXCertificateProfileCertificateValidityPeriodScale string

const (
	Windows10ImportedPFXCertificateProfileCertificateValidityPeriodScale_Days   Windows10ImportedPFXCertificateProfileCertificateValidityPeriodScale = "days"
	Windows10ImportedPFXCertificateProfileCertificateValidityPeriodScale_Months Windows10ImportedPFXCertificateProfileCertificateValidityPeriodScale = "months"
	Windows10ImportedPFXCertificateProfileCertificateValidityPeriodScale_Years  Windows10ImportedPFXCertificateProfileCertificateValidityPeriodScale = "years"
)

func PossibleValuesForWindows10ImportedPFXCertificateProfileCertificateValidityPeriodScale() []string {
	return []string{
		string(Windows10ImportedPFXCertificateProfileCertificateValidityPeriodScale_Days),
		string(Windows10ImportedPFXCertificateProfileCertificateValidityPeriodScale_Months),
		string(Windows10ImportedPFXCertificateProfileCertificateValidityPeriodScale_Years),
	}
}

func (s *Windows10ImportedPFXCertificateProfileCertificateValidityPeriodScale) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10ImportedPFXCertificateProfileCertificateValidityPeriodScale(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10ImportedPFXCertificateProfileCertificateValidityPeriodScale(input string) (*Windows10ImportedPFXCertificateProfileCertificateValidityPeriodScale, error) {
	vals := map[string]Windows10ImportedPFXCertificateProfileCertificateValidityPeriodScale{
		"days":   Windows10ImportedPFXCertificateProfileCertificateValidityPeriodScale_Days,
		"months": Windows10ImportedPFXCertificateProfileCertificateValidityPeriodScale_Months,
		"years":  Windows10ImportedPFXCertificateProfileCertificateValidityPeriodScale_Years,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10ImportedPFXCertificateProfileCertificateValidityPeriodScale(input)
	return &out, nil
}
