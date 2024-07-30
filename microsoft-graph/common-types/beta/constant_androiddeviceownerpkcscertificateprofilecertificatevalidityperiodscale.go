package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScale string

const (
	AndroidDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScale_Days   AndroidDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScale = "days"
	AndroidDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScale_Months AndroidDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScale = "months"
	AndroidDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScale_Years  AndroidDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScale = "years"
)

func PossibleValuesForAndroidDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScale() []string {
	return []string{
		string(AndroidDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScale_Days),
		string(AndroidDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScale_Months),
		string(AndroidDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScale_Years),
	}
}

func (s *AndroidDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScale) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScale(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScale(input string) (*AndroidDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScale, error) {
	vals := map[string]AndroidDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScale{
		"days":   AndroidDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScale_Days,
		"months": AndroidDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScale_Months,
		"years":  AndroidDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScale_Years,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScale(input)
	return &out, nil
}
