package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfilePkcsCertificateProfileCertificateValidityPeriodScale string

const (
	AndroidWorkProfilePkcsCertificateProfileCertificateValidityPeriodScale_Days   AndroidWorkProfilePkcsCertificateProfileCertificateValidityPeriodScale = "days"
	AndroidWorkProfilePkcsCertificateProfileCertificateValidityPeriodScale_Months AndroidWorkProfilePkcsCertificateProfileCertificateValidityPeriodScale = "months"
	AndroidWorkProfilePkcsCertificateProfileCertificateValidityPeriodScale_Years  AndroidWorkProfilePkcsCertificateProfileCertificateValidityPeriodScale = "years"
)

func PossibleValuesForAndroidWorkProfilePkcsCertificateProfileCertificateValidityPeriodScale() []string {
	return []string{
		string(AndroidWorkProfilePkcsCertificateProfileCertificateValidityPeriodScale_Days),
		string(AndroidWorkProfilePkcsCertificateProfileCertificateValidityPeriodScale_Months),
		string(AndroidWorkProfilePkcsCertificateProfileCertificateValidityPeriodScale_Years),
	}
}

func (s *AndroidWorkProfilePkcsCertificateProfileCertificateValidityPeriodScale) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidWorkProfilePkcsCertificateProfileCertificateValidityPeriodScale(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidWorkProfilePkcsCertificateProfileCertificateValidityPeriodScale(input string) (*AndroidWorkProfilePkcsCertificateProfileCertificateValidityPeriodScale, error) {
	vals := map[string]AndroidWorkProfilePkcsCertificateProfileCertificateValidityPeriodScale{
		"days":   AndroidWorkProfilePkcsCertificateProfileCertificateValidityPeriodScale_Days,
		"months": AndroidWorkProfilePkcsCertificateProfileCertificateValidityPeriodScale_Months,
		"years":  AndroidWorkProfilePkcsCertificateProfileCertificateValidityPeriodScale_Years,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfilePkcsCertificateProfileCertificateValidityPeriodScale(input)
	return &out, nil
}
