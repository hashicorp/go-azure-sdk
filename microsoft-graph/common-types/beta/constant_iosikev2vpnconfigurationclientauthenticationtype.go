package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosikEv2VpnConfigurationClientAuthenticationType string

const (
	IosikEv2VpnConfigurationClientAuthenticationType_DeviceAuthentication IosikEv2VpnConfigurationClientAuthenticationType = "deviceAuthentication"
	IosikEv2VpnConfigurationClientAuthenticationType_UserAuthentication   IosikEv2VpnConfigurationClientAuthenticationType = "userAuthentication"
)

func PossibleValuesForIosikEv2VpnConfigurationClientAuthenticationType() []string {
	return []string{
		string(IosikEv2VpnConfigurationClientAuthenticationType_DeviceAuthentication),
		string(IosikEv2VpnConfigurationClientAuthenticationType_UserAuthentication),
	}
}

func (s *IosikEv2VpnConfigurationClientAuthenticationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosikEv2VpnConfigurationClientAuthenticationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosikEv2VpnConfigurationClientAuthenticationType(input string) (*IosikEv2VpnConfigurationClientAuthenticationType, error) {
	vals := map[string]IosikEv2VpnConfigurationClientAuthenticationType{
		"deviceauthentication": IosikEv2VpnConfigurationClientAuthenticationType_DeviceAuthentication,
		"userauthentication":   IosikEv2VpnConfigurationClientAuthenticationType_UserAuthentication,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosikEv2VpnConfigurationClientAuthenticationType(input)
	return &out, nil
}
