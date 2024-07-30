package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessIpSecIntegrity string

const (
	NetworkaccessIpSecIntegrity_GcmAes128 NetworkaccessIpSecIntegrity = "gcmAes128"
	NetworkaccessIpSecIntegrity_GcmAes192 NetworkaccessIpSecIntegrity = "gcmAes192"
	NetworkaccessIpSecIntegrity_GcmAes256 NetworkaccessIpSecIntegrity = "gcmAes256"
	NetworkaccessIpSecIntegrity_Sha256    NetworkaccessIpSecIntegrity = "sha256"
)

func PossibleValuesForNetworkaccessIpSecIntegrity() []string {
	return []string{
		string(NetworkaccessIpSecIntegrity_GcmAes128),
		string(NetworkaccessIpSecIntegrity_GcmAes192),
		string(NetworkaccessIpSecIntegrity_GcmAes256),
		string(NetworkaccessIpSecIntegrity_Sha256),
	}
}

func (s *NetworkaccessIpSecIntegrity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessIpSecIntegrity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessIpSecIntegrity(input string) (*NetworkaccessIpSecIntegrity, error) {
	vals := map[string]NetworkaccessIpSecIntegrity{
		"gcmaes128": NetworkaccessIpSecIntegrity_GcmAes128,
		"gcmaes192": NetworkaccessIpSecIntegrity_GcmAes192,
		"gcmaes256": NetworkaccessIpSecIntegrity_GcmAes256,
		"sha256":    NetworkaccessIpSecIntegrity_Sha256,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessIpSecIntegrity(input)
	return &out, nil
}
