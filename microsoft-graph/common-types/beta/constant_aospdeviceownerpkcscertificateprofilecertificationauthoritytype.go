package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AospDeviceOwnerPkcsCertificateProfileCertificationAuthorityType string

const (
	AospDeviceOwnerPkcsCertificateProfileCertificationAuthorityType_DigiCert      AospDeviceOwnerPkcsCertificateProfileCertificationAuthorityType = "digiCert"
	AospDeviceOwnerPkcsCertificateProfileCertificationAuthorityType_Microsoft     AospDeviceOwnerPkcsCertificateProfileCertificationAuthorityType = "microsoft"
	AospDeviceOwnerPkcsCertificateProfileCertificationAuthorityType_NotConfigured AospDeviceOwnerPkcsCertificateProfileCertificationAuthorityType = "notConfigured"
)

func PossibleValuesForAospDeviceOwnerPkcsCertificateProfileCertificationAuthorityType() []string {
	return []string{
		string(AospDeviceOwnerPkcsCertificateProfileCertificationAuthorityType_DigiCert),
		string(AospDeviceOwnerPkcsCertificateProfileCertificationAuthorityType_Microsoft),
		string(AospDeviceOwnerPkcsCertificateProfileCertificationAuthorityType_NotConfigured),
	}
}

func (s *AospDeviceOwnerPkcsCertificateProfileCertificationAuthorityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAospDeviceOwnerPkcsCertificateProfileCertificationAuthorityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAospDeviceOwnerPkcsCertificateProfileCertificationAuthorityType(input string) (*AospDeviceOwnerPkcsCertificateProfileCertificationAuthorityType, error) {
	vals := map[string]AospDeviceOwnerPkcsCertificateProfileCertificationAuthorityType{
		"digicert":      AospDeviceOwnerPkcsCertificateProfileCertificationAuthorityType_DigiCert,
		"microsoft":     AospDeviceOwnerPkcsCertificateProfileCertificationAuthorityType_Microsoft,
		"notconfigured": AospDeviceOwnerPkcsCertificateProfileCertificationAuthorityType_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AospDeviceOwnerPkcsCertificateProfileCertificationAuthorityType(input)
	return &out, nil
}
