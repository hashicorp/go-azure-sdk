package servertrustgroups

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TrustScope string

const (
	TrustScopeGlobalTransactions TrustScope = "GlobalTransactions"
	TrustScopeServiceBroker      TrustScope = "ServiceBroker"
)

func PossibleValuesForTrustScope() []string {
	return []string{
		string(TrustScopeGlobalTransactions),
		string(TrustScopeServiceBroker),
	}
}

func (s *TrustScope) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTrustScope(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTrustScope(input string) (*TrustScope, error) {
	vals := map[string]TrustScope{
		"globaltransactions": TrustScopeGlobalTransactions,
		"servicebroker":      TrustScopeServiceBroker,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TrustScope(input)
	return &out, nil
}
