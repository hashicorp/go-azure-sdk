package wcfrelaysops

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Relaytype string

const (
	RelaytypeHTTP   Relaytype = "Http"
	RelaytypeNetTcp Relaytype = "NetTcp"
)

func PossibleValuesForRelaytype() []string {
	return []string{
		string(RelaytypeHTTP),
		string(RelaytypeNetTcp),
	}
}

func (s *Relaytype) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRelaytype(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRelaytype(input string) (*Relaytype, error) {
	vals := map[string]Relaytype{
		"http":   RelaytypeHTTP,
		"nettcp": RelaytypeNetTcp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Relaytype(input)
	return &out, nil
}
