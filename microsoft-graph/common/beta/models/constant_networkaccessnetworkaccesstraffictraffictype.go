package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessNetworkAccessTrafficTrafficType string

const (
	NetworkaccessNetworkAccessTrafficTrafficTypeall          NetworkaccessNetworkAccessTrafficTrafficType = "All"
	NetworkaccessNetworkAccessTrafficTrafficTypeinternet     NetworkaccessNetworkAccessTrafficTrafficType = "Internet"
	NetworkaccessNetworkAccessTrafficTrafficTypemicrosoft365 NetworkaccessNetworkAccessTrafficTrafficType = "Microsoft365"
	NetworkaccessNetworkAccessTrafficTrafficTypeprivate      NetworkaccessNetworkAccessTrafficTrafficType = "Private"
)

func PossibleValuesForNetworkaccessNetworkAccessTrafficTrafficType() []string {
	return []string{
		string(NetworkaccessNetworkAccessTrafficTrafficTypeall),
		string(NetworkaccessNetworkAccessTrafficTrafficTypeinternet),
		string(NetworkaccessNetworkAccessTrafficTrafficTypemicrosoft365),
		string(NetworkaccessNetworkAccessTrafficTrafficTypeprivate),
	}
}

func parseNetworkaccessNetworkAccessTrafficTrafficType(input string) (*NetworkaccessNetworkAccessTrafficTrafficType, error) {
	vals := map[string]NetworkaccessNetworkAccessTrafficTrafficType{
		"all":          NetworkaccessNetworkAccessTrafficTrafficTypeall,
		"internet":     NetworkaccessNetworkAccessTrafficTrafficTypeinternet,
		"microsoft365": NetworkaccessNetworkAccessTrafficTrafficTypemicrosoft365,
		"private":      NetworkaccessNetworkAccessTrafficTrafficTypeprivate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessNetworkAccessTrafficTrafficType(input)
	return &out, nil
}
