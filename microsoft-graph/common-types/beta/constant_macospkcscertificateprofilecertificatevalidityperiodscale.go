package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSPkcsCertificateProfileCertificateValidityPeriodScale string

const (
	MacOSPkcsCertificateProfileCertificateValidityPeriodScale_Days   MacOSPkcsCertificateProfileCertificateValidityPeriodScale = "days"
	MacOSPkcsCertificateProfileCertificateValidityPeriodScale_Months MacOSPkcsCertificateProfileCertificateValidityPeriodScale = "months"
	MacOSPkcsCertificateProfileCertificateValidityPeriodScale_Years  MacOSPkcsCertificateProfileCertificateValidityPeriodScale = "years"
)

func PossibleValuesForMacOSPkcsCertificateProfileCertificateValidityPeriodScale() []string {
	return []string{
		string(MacOSPkcsCertificateProfileCertificateValidityPeriodScale_Days),
		string(MacOSPkcsCertificateProfileCertificateValidityPeriodScale_Months),
		string(MacOSPkcsCertificateProfileCertificateValidityPeriodScale_Years),
	}
}

func (s *MacOSPkcsCertificateProfileCertificateValidityPeriodScale) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSPkcsCertificateProfileCertificateValidityPeriodScale(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSPkcsCertificateProfileCertificateValidityPeriodScale(input string) (*MacOSPkcsCertificateProfileCertificateValidityPeriodScale, error) {
	vals := map[string]MacOSPkcsCertificateProfileCertificateValidityPeriodScale{
		"days":   MacOSPkcsCertificateProfileCertificateValidityPeriodScale_Days,
		"months": MacOSPkcsCertificateProfileCertificateValidityPeriodScale_Months,
		"years":  MacOSPkcsCertificateProfileCertificateValidityPeriodScale_Years,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSPkcsCertificateProfileCertificateValidityPeriodScale(input)
	return &out, nil
}
