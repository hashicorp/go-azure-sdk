package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerScepCertificateProfileCertificateAccessType string

const (
	AndroidDeviceOwnerScepCertificateProfileCertificateAccessType_SpecificApps AndroidDeviceOwnerScepCertificateProfileCertificateAccessType = "specificApps"
	AndroidDeviceOwnerScepCertificateProfileCertificateAccessType_UserApproval AndroidDeviceOwnerScepCertificateProfileCertificateAccessType = "userApproval"
)

func PossibleValuesForAndroidDeviceOwnerScepCertificateProfileCertificateAccessType() []string {
	return []string{
		string(AndroidDeviceOwnerScepCertificateProfileCertificateAccessType_SpecificApps),
		string(AndroidDeviceOwnerScepCertificateProfileCertificateAccessType_UserApproval),
	}
}

func (s *AndroidDeviceOwnerScepCertificateProfileCertificateAccessType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerScepCertificateProfileCertificateAccessType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerScepCertificateProfileCertificateAccessType(input string) (*AndroidDeviceOwnerScepCertificateProfileCertificateAccessType, error) {
	vals := map[string]AndroidDeviceOwnerScepCertificateProfileCertificateAccessType{
		"specificapps": AndroidDeviceOwnerScepCertificateProfileCertificateAccessType_SpecificApps,
		"userapproval": AndroidDeviceOwnerScepCertificateProfileCertificateAccessType_UserApproval,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerScepCertificateProfileCertificateAccessType(input)
	return &out, nil
}
