package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerPkcsCertificateProfileCertificateAccessType string

const (
	AndroidDeviceOwnerPkcsCertificateProfileCertificateAccessType_SpecificApps AndroidDeviceOwnerPkcsCertificateProfileCertificateAccessType = "specificApps"
	AndroidDeviceOwnerPkcsCertificateProfileCertificateAccessType_UserApproval AndroidDeviceOwnerPkcsCertificateProfileCertificateAccessType = "userApproval"
)

func PossibleValuesForAndroidDeviceOwnerPkcsCertificateProfileCertificateAccessType() []string {
	return []string{
		string(AndroidDeviceOwnerPkcsCertificateProfileCertificateAccessType_SpecificApps),
		string(AndroidDeviceOwnerPkcsCertificateProfileCertificateAccessType_UserApproval),
	}
}

func (s *AndroidDeviceOwnerPkcsCertificateProfileCertificateAccessType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerPkcsCertificateProfileCertificateAccessType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerPkcsCertificateProfileCertificateAccessType(input string) (*AndroidDeviceOwnerPkcsCertificateProfileCertificateAccessType, error) {
	vals := map[string]AndroidDeviceOwnerPkcsCertificateProfileCertificateAccessType{
		"specificapps": AndroidDeviceOwnerPkcsCertificateProfileCertificateAccessType_SpecificApps,
		"userapproval": AndroidDeviceOwnerPkcsCertificateProfileCertificateAccessType_UserApproval,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerPkcsCertificateProfileCertificateAccessType(input)
	return &out, nil
}
