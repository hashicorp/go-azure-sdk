package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows81SCEPCertificateProfileCertificateValidityPeriodScale string

const (
	Windows81SCEPCertificateProfileCertificateValidityPeriodScale_Days   Windows81SCEPCertificateProfileCertificateValidityPeriodScale = "days"
	Windows81SCEPCertificateProfileCertificateValidityPeriodScale_Months Windows81SCEPCertificateProfileCertificateValidityPeriodScale = "months"
	Windows81SCEPCertificateProfileCertificateValidityPeriodScale_Years  Windows81SCEPCertificateProfileCertificateValidityPeriodScale = "years"
)

func PossibleValuesForWindows81SCEPCertificateProfileCertificateValidityPeriodScale() []string {
	return []string{
		string(Windows81SCEPCertificateProfileCertificateValidityPeriodScale_Days),
		string(Windows81SCEPCertificateProfileCertificateValidityPeriodScale_Months),
		string(Windows81SCEPCertificateProfileCertificateValidityPeriodScale_Years),
	}
}

func (s *Windows81SCEPCertificateProfileCertificateValidityPeriodScale) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows81SCEPCertificateProfileCertificateValidityPeriodScale(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows81SCEPCertificateProfileCertificateValidityPeriodScale(input string) (*Windows81SCEPCertificateProfileCertificateValidityPeriodScale, error) {
	vals := map[string]Windows81SCEPCertificateProfileCertificateValidityPeriodScale{
		"days":   Windows81SCEPCertificateProfileCertificateValidityPeriodScale_Days,
		"months": Windows81SCEPCertificateProfileCertificateValidityPeriodScale_Months,
		"years":  Windows81SCEPCertificateProfileCertificateValidityPeriodScale_Years,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows81SCEPCertificateProfileCertificateValidityPeriodScale(input)
	return &out, nil
}
