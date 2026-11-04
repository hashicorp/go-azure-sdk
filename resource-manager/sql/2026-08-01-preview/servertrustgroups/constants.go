package servertrustgroups

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServerTrustGroupPropertiesTrustScopesItem string

const (
	ServerTrustGroupPropertiesTrustScopesItemGlobalTransactions ServerTrustGroupPropertiesTrustScopesItem = "GlobalTransactions"
	ServerTrustGroupPropertiesTrustScopesItemServiceBroker      ServerTrustGroupPropertiesTrustScopesItem = "ServiceBroker"
)

func PossibleValuesForServerTrustGroupPropertiesTrustScopesItem() []string {
	return []string{
		string(ServerTrustGroupPropertiesTrustScopesItemGlobalTransactions),
		string(ServerTrustGroupPropertiesTrustScopesItemServiceBroker),
	}
}

func (s *ServerTrustGroupPropertiesTrustScopesItem) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseServerTrustGroupPropertiesTrustScopesItem(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseServerTrustGroupPropertiesTrustScopesItem(input string) (*ServerTrustGroupPropertiesTrustScopesItem, error) {
	vals := map[string]ServerTrustGroupPropertiesTrustScopesItem{
		"globaltransactions": ServerTrustGroupPropertiesTrustScopesItemGlobalTransactions,
		"servicebroker":      ServerTrustGroupPropertiesTrustScopesItemServiceBroker,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServerTrustGroupPropertiesTrustScopesItem(input)
	return &out, nil
}
