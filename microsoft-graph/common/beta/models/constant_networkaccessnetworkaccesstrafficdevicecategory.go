package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessNetworkAccessTrafficDeviceCategory string

const (
	NetworkaccessNetworkAccessTrafficDeviceCategorybranch NetworkaccessNetworkAccessTrafficDeviceCategory = "Branch"
	NetworkaccessNetworkAccessTrafficDeviceCategoryclient NetworkaccessNetworkAccessTrafficDeviceCategory = "Client"
)

func PossibleValuesForNetworkaccessNetworkAccessTrafficDeviceCategory() []string {
	return []string{
		string(NetworkaccessNetworkAccessTrafficDeviceCategorybranch),
		string(NetworkaccessNetworkAccessTrafficDeviceCategoryclient),
	}
}

func parseNetworkaccessNetworkAccessTrafficDeviceCategory(input string) (*NetworkaccessNetworkAccessTrafficDeviceCategory, error) {
	vals := map[string]NetworkaccessNetworkAccessTrafficDeviceCategory{
		"branch": NetworkaccessNetworkAccessTrafficDeviceCategorybranch,
		"client": NetworkaccessNetworkAccessTrafficDeviceCategoryclient,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessNetworkAccessTrafficDeviceCategory(input)
	return &out, nil
}
