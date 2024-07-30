package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessTransactionSummaryTrafficType string

const (
	NetworkaccessTransactionSummaryTrafficType_All          NetworkaccessTransactionSummaryTrafficType = "all"
	NetworkaccessTransactionSummaryTrafficType_Internet     NetworkaccessTransactionSummaryTrafficType = "internet"
	NetworkaccessTransactionSummaryTrafficType_Microsoft365 NetworkaccessTransactionSummaryTrafficType = "microsoft365"
	NetworkaccessTransactionSummaryTrafficType_Private      NetworkaccessTransactionSummaryTrafficType = "private"
)

func PossibleValuesForNetworkaccessTransactionSummaryTrafficType() []string {
	return []string{
		string(NetworkaccessTransactionSummaryTrafficType_All),
		string(NetworkaccessTransactionSummaryTrafficType_Internet),
		string(NetworkaccessTransactionSummaryTrafficType_Microsoft365),
		string(NetworkaccessTransactionSummaryTrafficType_Private),
	}
}

func (s *NetworkaccessTransactionSummaryTrafficType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessTransactionSummaryTrafficType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessTransactionSummaryTrafficType(input string) (*NetworkaccessTransactionSummaryTrafficType, error) {
	vals := map[string]NetworkaccessTransactionSummaryTrafficType{
		"all":          NetworkaccessTransactionSummaryTrafficType_All,
		"internet":     NetworkaccessTransactionSummaryTrafficType_Internet,
		"microsoft365": NetworkaccessTransactionSummaryTrafficType_Microsoft365,
		"private":      NetworkaccessTransactionSummaryTrafficType_Private,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessTransactionSummaryTrafficType(input)
	return &out, nil
}
