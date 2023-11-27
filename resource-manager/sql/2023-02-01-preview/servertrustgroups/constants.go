package servertrustgroups

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TrustScopes string

const (
	TrustScopesGlobalTransactions TrustScopes = "GlobalTransactions"
	TrustScopesServiceBroker      TrustScopes = "ServiceBroker"
)

func PossibleValuesForTrustScopes() []string {
	return []string{
		string(TrustScopesGlobalTransactions),
		string(TrustScopesServiceBroker),
	}
}

func (s *TrustScopes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTrustScopes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTrustScopes(input string) (*TrustScopes, error) {
	vals := map[string]TrustScopes{
		"globaltransactions": TrustScopesGlobalTransactions,
		"servicebroker":      TrustScopesServiceBroker,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TrustScopes(input)
	return &out, nil
}
