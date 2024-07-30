package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkLocationDetailNetworkType string

const (
	NetworkLocationDetailNetworkType_Extranet             NetworkLocationDetailNetworkType = "extranet"
	NetworkLocationDetailNetworkType_Intranet             NetworkLocationDetailNetworkType = "intranet"
	NetworkLocationDetailNetworkType_NamedNetwork         NetworkLocationDetailNetworkType = "namedNetwork"
	NetworkLocationDetailNetworkType_Trusted              NetworkLocationDetailNetworkType = "trusted"
	NetworkLocationDetailNetworkType_TrustedNamedLocation NetworkLocationDetailNetworkType = "trustedNamedLocation"
)

func PossibleValuesForNetworkLocationDetailNetworkType() []string {
	return []string{
		string(NetworkLocationDetailNetworkType_Extranet),
		string(NetworkLocationDetailNetworkType_Intranet),
		string(NetworkLocationDetailNetworkType_NamedNetwork),
		string(NetworkLocationDetailNetworkType_Trusted),
		string(NetworkLocationDetailNetworkType_TrustedNamedLocation),
	}
}

func (s *NetworkLocationDetailNetworkType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkLocationDetailNetworkType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkLocationDetailNetworkType(input string) (*NetworkLocationDetailNetworkType, error) {
	vals := map[string]NetworkLocationDetailNetworkType{
		"extranet":             NetworkLocationDetailNetworkType_Extranet,
		"intranet":             NetworkLocationDetailNetworkType_Intranet,
		"namednetwork":         NetworkLocationDetailNetworkType_NamedNetwork,
		"trusted":              NetworkLocationDetailNetworkType_Trusted,
		"trustednamedlocation": NetworkLocationDetailNetworkType_TrustedNamedLocation,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkLocationDetailNetworkType(input)
	return &out, nil
}
