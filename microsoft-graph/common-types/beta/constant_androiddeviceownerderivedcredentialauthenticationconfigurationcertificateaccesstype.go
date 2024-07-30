package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerDerivedCredentialAuthenticationConfigurationCertificateAccessType string

const (
	AndroidDeviceOwnerDerivedCredentialAuthenticationConfigurationCertificateAccessType_SpecificApps AndroidDeviceOwnerDerivedCredentialAuthenticationConfigurationCertificateAccessType = "specificApps"
	AndroidDeviceOwnerDerivedCredentialAuthenticationConfigurationCertificateAccessType_UserApproval AndroidDeviceOwnerDerivedCredentialAuthenticationConfigurationCertificateAccessType = "userApproval"
)

func PossibleValuesForAndroidDeviceOwnerDerivedCredentialAuthenticationConfigurationCertificateAccessType() []string {
	return []string{
		string(AndroidDeviceOwnerDerivedCredentialAuthenticationConfigurationCertificateAccessType_SpecificApps),
		string(AndroidDeviceOwnerDerivedCredentialAuthenticationConfigurationCertificateAccessType_UserApproval),
	}
}

func (s *AndroidDeviceOwnerDerivedCredentialAuthenticationConfigurationCertificateAccessType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerDerivedCredentialAuthenticationConfigurationCertificateAccessType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerDerivedCredentialAuthenticationConfigurationCertificateAccessType(input string) (*AndroidDeviceOwnerDerivedCredentialAuthenticationConfigurationCertificateAccessType, error) {
	vals := map[string]AndroidDeviceOwnerDerivedCredentialAuthenticationConfigurationCertificateAccessType{
		"specificapps": AndroidDeviceOwnerDerivedCredentialAuthenticationConfigurationCertificateAccessType_SpecificApps,
		"userapproval": AndroidDeviceOwnerDerivedCredentialAuthenticationConfigurationCertificateAccessType_UserApproval,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerDerivedCredentialAuthenticationConfigurationCertificateAccessType(input)
	return &out, nil
}
