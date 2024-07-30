package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessIpSecEncryption string

const (
	NetworkaccessIpSecEncryption_GcmAes128 NetworkaccessIpSecEncryption = "gcmAes128"
	NetworkaccessIpSecEncryption_GcmAes192 NetworkaccessIpSecEncryption = "gcmAes192"
	NetworkaccessIpSecEncryption_GcmAes256 NetworkaccessIpSecEncryption = "gcmAes256"
	NetworkaccessIpSecEncryption_None      NetworkaccessIpSecEncryption = "none"
)

func PossibleValuesForNetworkaccessIpSecEncryption() []string {
	return []string{
		string(NetworkaccessIpSecEncryption_GcmAes128),
		string(NetworkaccessIpSecEncryption_GcmAes192),
		string(NetworkaccessIpSecEncryption_GcmAes256),
		string(NetworkaccessIpSecEncryption_None),
	}
}

func (s *NetworkaccessIpSecEncryption) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessIpSecEncryption(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessIpSecEncryption(input string) (*NetworkaccessIpSecEncryption, error) {
	vals := map[string]NetworkaccessIpSecEncryption{
		"gcmaes128": NetworkaccessIpSecEncryption_GcmAes128,
		"gcmaes192": NetworkaccessIpSecEncryption_GcmAes192,
		"gcmaes256": NetworkaccessIpSecEncryption_GcmAes256,
		"none":      NetworkaccessIpSecEncryption_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessIpSecEncryption(input)
	return &out, nil
}
