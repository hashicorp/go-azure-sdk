package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosPkcsCertificateProfileCertificateValidityPeriodScale string

const (
	IosPkcsCertificateProfileCertificateValidityPeriodScale_Days   IosPkcsCertificateProfileCertificateValidityPeriodScale = "days"
	IosPkcsCertificateProfileCertificateValidityPeriodScale_Months IosPkcsCertificateProfileCertificateValidityPeriodScale = "months"
	IosPkcsCertificateProfileCertificateValidityPeriodScale_Years  IosPkcsCertificateProfileCertificateValidityPeriodScale = "years"
)

func PossibleValuesForIosPkcsCertificateProfileCertificateValidityPeriodScale() []string {
	return []string{
		string(IosPkcsCertificateProfileCertificateValidityPeriodScale_Days),
		string(IosPkcsCertificateProfileCertificateValidityPeriodScale_Months),
		string(IosPkcsCertificateProfileCertificateValidityPeriodScale_Years),
	}
}

func (s *IosPkcsCertificateProfileCertificateValidityPeriodScale) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosPkcsCertificateProfileCertificateValidityPeriodScale(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosPkcsCertificateProfileCertificateValidityPeriodScale(input string) (*IosPkcsCertificateProfileCertificateValidityPeriodScale, error) {
	vals := map[string]IosPkcsCertificateProfileCertificateValidityPeriodScale{
		"days":   IosPkcsCertificateProfileCertificateValidityPeriodScale_Days,
		"months": IosPkcsCertificateProfileCertificateValidityPeriodScale_Months,
		"years":  IosPkcsCertificateProfileCertificateValidityPeriodScale_Years,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosPkcsCertificateProfileCertificateValidityPeriodScale(input)
	return &out, nil
}
