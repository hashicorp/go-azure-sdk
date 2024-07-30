package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerPkcsCertificateProfileCertificationAuthorityType string

const (
	AndroidDeviceOwnerPkcsCertificateProfileCertificationAuthorityType_DigiCert      AndroidDeviceOwnerPkcsCertificateProfileCertificationAuthorityType = "digiCert"
	AndroidDeviceOwnerPkcsCertificateProfileCertificationAuthorityType_Microsoft     AndroidDeviceOwnerPkcsCertificateProfileCertificationAuthorityType = "microsoft"
	AndroidDeviceOwnerPkcsCertificateProfileCertificationAuthorityType_NotConfigured AndroidDeviceOwnerPkcsCertificateProfileCertificationAuthorityType = "notConfigured"
)

func PossibleValuesForAndroidDeviceOwnerPkcsCertificateProfileCertificationAuthorityType() []string {
	return []string{
		string(AndroidDeviceOwnerPkcsCertificateProfileCertificationAuthorityType_DigiCert),
		string(AndroidDeviceOwnerPkcsCertificateProfileCertificationAuthorityType_Microsoft),
		string(AndroidDeviceOwnerPkcsCertificateProfileCertificationAuthorityType_NotConfigured),
	}
}

func (s *AndroidDeviceOwnerPkcsCertificateProfileCertificationAuthorityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerPkcsCertificateProfileCertificationAuthorityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerPkcsCertificateProfileCertificationAuthorityType(input string) (*AndroidDeviceOwnerPkcsCertificateProfileCertificationAuthorityType, error) {
	vals := map[string]AndroidDeviceOwnerPkcsCertificateProfileCertificationAuthorityType{
		"digicert":      AndroidDeviceOwnerPkcsCertificateProfileCertificationAuthorityType_DigiCert,
		"microsoft":     AndroidDeviceOwnerPkcsCertificateProfileCertificationAuthorityType_Microsoft,
		"notconfigured": AndroidDeviceOwnerPkcsCertificateProfileCertificationAuthorityType_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerPkcsCertificateProfileCertificationAuthorityType(input)
	return &out, nil
}
