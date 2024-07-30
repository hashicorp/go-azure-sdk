package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosikEv2VpnConfigurationServerCertificateType string

const (
	IosikEv2VpnConfigurationServerCertificateType_Ecdsa256 IosikEv2VpnConfigurationServerCertificateType = "ecdsa256"
	IosikEv2VpnConfigurationServerCertificateType_Ecdsa384 IosikEv2VpnConfigurationServerCertificateType = "ecdsa384"
	IosikEv2VpnConfigurationServerCertificateType_Ecdsa521 IosikEv2VpnConfigurationServerCertificateType = "ecdsa521"
	IosikEv2VpnConfigurationServerCertificateType_Rsa      IosikEv2VpnConfigurationServerCertificateType = "rsa"
)

func PossibleValuesForIosikEv2VpnConfigurationServerCertificateType() []string {
	return []string{
		string(IosikEv2VpnConfigurationServerCertificateType_Ecdsa256),
		string(IosikEv2VpnConfigurationServerCertificateType_Ecdsa384),
		string(IosikEv2VpnConfigurationServerCertificateType_Ecdsa521),
		string(IosikEv2VpnConfigurationServerCertificateType_Rsa),
	}
}

func (s *IosikEv2VpnConfigurationServerCertificateType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosikEv2VpnConfigurationServerCertificateType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosikEv2VpnConfigurationServerCertificateType(input string) (*IosikEv2VpnConfigurationServerCertificateType, error) {
	vals := map[string]IosikEv2VpnConfigurationServerCertificateType{
		"ecdsa256": IosikEv2VpnConfigurationServerCertificateType_Ecdsa256,
		"ecdsa384": IosikEv2VpnConfigurationServerCertificateType_Ecdsa384,
		"ecdsa521": IosikEv2VpnConfigurationServerCertificateType_Ecdsa521,
		"rsa":      IosikEv2VpnConfigurationServerCertificateType_Rsa,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosikEv2VpnConfigurationServerCertificateType(input)
	return &out, nil
}
