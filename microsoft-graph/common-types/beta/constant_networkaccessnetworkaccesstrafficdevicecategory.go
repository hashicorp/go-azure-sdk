package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessNetworkAccessTrafficDeviceCategory string

const (
	NetworkaccessNetworkAccessTrafficDeviceCategory_Branch NetworkaccessNetworkAccessTrafficDeviceCategory = "branch"
	NetworkaccessNetworkAccessTrafficDeviceCategory_Client NetworkaccessNetworkAccessTrafficDeviceCategory = "client"
)

func PossibleValuesForNetworkaccessNetworkAccessTrafficDeviceCategory() []string {
	return []string{
		string(NetworkaccessNetworkAccessTrafficDeviceCategory_Branch),
		string(NetworkaccessNetworkAccessTrafficDeviceCategory_Client),
	}
}

func (s *NetworkaccessNetworkAccessTrafficDeviceCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessNetworkAccessTrafficDeviceCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessNetworkAccessTrafficDeviceCategory(input string) (*NetworkaccessNetworkAccessTrafficDeviceCategory, error) {
	vals := map[string]NetworkaccessNetworkAccessTrafficDeviceCategory{
		"branch": NetworkaccessNetworkAccessTrafficDeviceCategory_Branch,
		"client": NetworkaccessNetworkAccessTrafficDeviceCategory_Client,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessNetworkAccessTrafficDeviceCategory(input)
	return &out, nil
}
