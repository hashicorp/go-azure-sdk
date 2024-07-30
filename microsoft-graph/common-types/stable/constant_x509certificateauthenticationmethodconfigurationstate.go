package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type X509CertificateAuthenticationMethodConfigurationState string

const (
	X509CertificateAuthenticationMethodConfigurationState_Disabled X509CertificateAuthenticationMethodConfigurationState = "disabled"
	X509CertificateAuthenticationMethodConfigurationState_Enabled  X509CertificateAuthenticationMethodConfigurationState = "enabled"
)

func PossibleValuesForX509CertificateAuthenticationMethodConfigurationState() []string {
	return []string{
		string(X509CertificateAuthenticationMethodConfigurationState_Disabled),
		string(X509CertificateAuthenticationMethodConfigurationState_Enabled),
	}
}

func (s *X509CertificateAuthenticationMethodConfigurationState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseX509CertificateAuthenticationMethodConfigurationState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseX509CertificateAuthenticationMethodConfigurationState(input string) (*X509CertificateAuthenticationMethodConfigurationState, error) {
	vals := map[string]X509CertificateAuthenticationMethodConfigurationState{
		"disabled": X509CertificateAuthenticationMethodConfigurationState_Disabled,
		"enabled":  X509CertificateAuthenticationMethodConfigurationState_Enabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := X509CertificateAuthenticationMethodConfigurationState(input)
	return &out, nil
}
