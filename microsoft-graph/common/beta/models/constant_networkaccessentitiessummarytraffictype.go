package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessEntitiesSummaryTrafficType string

const (
	NetworkaccessEntitiesSummaryTrafficTypeall          NetworkaccessEntitiesSummaryTrafficType = "All"
	NetworkaccessEntitiesSummaryTrafficTypeinternet     NetworkaccessEntitiesSummaryTrafficType = "Internet"
	NetworkaccessEntitiesSummaryTrafficTypemicrosoft365 NetworkaccessEntitiesSummaryTrafficType = "Microsoft365"
	NetworkaccessEntitiesSummaryTrafficTypeprivate      NetworkaccessEntitiesSummaryTrafficType = "Private"
)

func PossibleValuesForNetworkaccessEntitiesSummaryTrafficType() []string {
	return []string{
		string(NetworkaccessEntitiesSummaryTrafficTypeall),
		string(NetworkaccessEntitiesSummaryTrafficTypeinternet),
		string(NetworkaccessEntitiesSummaryTrafficTypemicrosoft365),
		string(NetworkaccessEntitiesSummaryTrafficTypeprivate),
	}
}

func parseNetworkaccessEntitiesSummaryTrafficType(input string) (*NetworkaccessEntitiesSummaryTrafficType, error) {
	vals := map[string]NetworkaccessEntitiesSummaryTrafficType{
		"all":          NetworkaccessEntitiesSummaryTrafficTypeall,
		"internet":     NetworkaccessEntitiesSummaryTrafficTypeinternet,
		"microsoft365": NetworkaccessEntitiesSummaryTrafficTypemicrosoft365,
		"private":      NetworkaccessEntitiesSummaryTrafficTypeprivate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessEntitiesSummaryTrafficType(input)
	return &out, nil
}
