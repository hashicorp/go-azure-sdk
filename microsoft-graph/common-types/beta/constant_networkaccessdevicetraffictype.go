package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessDeviceTrafficType string

const (
	NetworkaccessDeviceTrafficType_All          NetworkaccessDeviceTrafficType = "all"
	NetworkaccessDeviceTrafficType_Internet     NetworkaccessDeviceTrafficType = "internet"
	NetworkaccessDeviceTrafficType_Microsoft365 NetworkaccessDeviceTrafficType = "microsoft365"
	NetworkaccessDeviceTrafficType_Private      NetworkaccessDeviceTrafficType = "private"
)

func PossibleValuesForNetworkaccessDeviceTrafficType() []string {
	return []string{
		string(NetworkaccessDeviceTrafficType_All),
		string(NetworkaccessDeviceTrafficType_Internet),
		string(NetworkaccessDeviceTrafficType_Microsoft365),
		string(NetworkaccessDeviceTrafficType_Private),
	}
}

func (s *NetworkaccessDeviceTrafficType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessDeviceTrafficType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessDeviceTrafficType(input string) (*NetworkaccessDeviceTrafficType, error) {
	vals := map[string]NetworkaccessDeviceTrafficType{
		"all":          NetworkaccessDeviceTrafficType_All,
		"internet":     NetworkaccessDeviceTrafficType_Internet,
		"microsoft365": NetworkaccessDeviceTrafficType_Microsoft365,
		"private":      NetworkaccessDeviceTrafficType_Private,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessDeviceTrafficType(input)
	return &out, nil
}
