package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type X509CertificateIssuerHintsConfigurationState string

const (
	X509CertificateIssuerHintsConfigurationState_Disabled X509CertificateIssuerHintsConfigurationState = "disabled"
	X509CertificateIssuerHintsConfigurationState_Enabled  X509CertificateIssuerHintsConfigurationState = "enabled"
)

func PossibleValuesForX509CertificateIssuerHintsConfigurationState() []string {
	return []string{
		string(X509CertificateIssuerHintsConfigurationState_Disabled),
		string(X509CertificateIssuerHintsConfigurationState_Enabled),
	}
}

func (s *X509CertificateIssuerHintsConfigurationState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseX509CertificateIssuerHintsConfigurationState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseX509CertificateIssuerHintsConfigurationState(input string) (*X509CertificateIssuerHintsConfigurationState, error) {
	vals := map[string]X509CertificateIssuerHintsConfigurationState{
		"disabled": X509CertificateIssuerHintsConfigurationState_Disabled,
		"enabled":  X509CertificateIssuerHintsConfigurationState_Enabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := X509CertificateIssuerHintsConfigurationState(input)
	return &out, nil
}
