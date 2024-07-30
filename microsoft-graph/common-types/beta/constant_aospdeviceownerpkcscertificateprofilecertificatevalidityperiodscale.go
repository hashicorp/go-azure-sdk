package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AospDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScale string

const (
	AospDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScale_Days   AospDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScale = "days"
	AospDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScale_Months AospDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScale = "months"
	AospDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScale_Years  AospDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScale = "years"
)

func PossibleValuesForAospDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScale() []string {
	return []string{
		string(AospDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScale_Days),
		string(AospDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScale_Months),
		string(AospDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScale_Years),
	}
}

func (s *AospDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScale) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAospDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScale(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAospDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScale(input string) (*AospDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScale, error) {
	vals := map[string]AospDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScale{
		"days":   AospDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScale_Days,
		"months": AospDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScale_Months,
		"years":  AospDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScale_Years,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AospDeviceOwnerPkcsCertificateProfileCertificateValidityPeriodScale(input)
	return &out, nil
}
