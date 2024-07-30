package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdgeSearchEngineEdgeSearchEngineType string

const (
	EdgeSearchEngineEdgeSearchEngineType_Bing    EdgeSearchEngineEdgeSearchEngineType = "bing"
	EdgeSearchEngineEdgeSearchEngineType_Default EdgeSearchEngineEdgeSearchEngineType = "default"
)

func PossibleValuesForEdgeSearchEngineEdgeSearchEngineType() []string {
	return []string{
		string(EdgeSearchEngineEdgeSearchEngineType_Bing),
		string(EdgeSearchEngineEdgeSearchEngineType_Default),
	}
}

func (s *EdgeSearchEngineEdgeSearchEngineType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEdgeSearchEngineEdgeSearchEngineType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEdgeSearchEngineEdgeSearchEngineType(input string) (*EdgeSearchEngineEdgeSearchEngineType, error) {
	vals := map[string]EdgeSearchEngineEdgeSearchEngineType{
		"bing":    EdgeSearchEngineEdgeSearchEngineType_Bing,
		"default": EdgeSearchEngineEdgeSearchEngineType_Default,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdgeSearchEngineEdgeSearchEngineType(input)
	return &out, nil
}
