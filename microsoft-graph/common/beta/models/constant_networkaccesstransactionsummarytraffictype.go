package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessTransactionSummaryTrafficType string

const (
	NetworkaccessTransactionSummaryTrafficTypeall          NetworkaccessTransactionSummaryTrafficType = "All"
	NetworkaccessTransactionSummaryTrafficTypeinternet     NetworkaccessTransactionSummaryTrafficType = "Internet"
	NetworkaccessTransactionSummaryTrafficTypemicrosoft365 NetworkaccessTransactionSummaryTrafficType = "Microsoft365"
	NetworkaccessTransactionSummaryTrafficTypeprivate      NetworkaccessTransactionSummaryTrafficType = "Private"
)

func PossibleValuesForNetworkaccessTransactionSummaryTrafficType() []string {
	return []string{
		string(NetworkaccessTransactionSummaryTrafficTypeall),
		string(NetworkaccessTransactionSummaryTrafficTypeinternet),
		string(NetworkaccessTransactionSummaryTrafficTypemicrosoft365),
		string(NetworkaccessTransactionSummaryTrafficTypeprivate),
	}
}

func parseNetworkaccessTransactionSummaryTrafficType(input string) (*NetworkaccessTransactionSummaryTrafficType, error) {
	vals := map[string]NetworkaccessTransactionSummaryTrafficType{
		"all":          NetworkaccessTransactionSummaryTrafficTypeall,
		"internet":     NetworkaccessTransactionSummaryTrafficTypeinternet,
		"microsoft365": NetworkaccessTransactionSummaryTrafficTypemicrosoft365,
		"private":      NetworkaccessTransactionSummaryTrafficTypeprivate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessTransactionSummaryTrafficType(input)
	return &out, nil
}
