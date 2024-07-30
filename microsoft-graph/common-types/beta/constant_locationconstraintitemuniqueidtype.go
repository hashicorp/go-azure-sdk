package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LocationConstraintItemUniqueIdType string

const (
	LocationConstraintItemUniqueIdType_Bing          LocationConstraintItemUniqueIdType = "bing"
	LocationConstraintItemUniqueIdType_Directory     LocationConstraintItemUniqueIdType = "directory"
	LocationConstraintItemUniqueIdType_LocationStore LocationConstraintItemUniqueIdType = "locationStore"
	LocationConstraintItemUniqueIdType_Private       LocationConstraintItemUniqueIdType = "private"
	LocationConstraintItemUniqueIdType_Unknown       LocationConstraintItemUniqueIdType = "unknown"
)

func PossibleValuesForLocationConstraintItemUniqueIdType() []string {
	return []string{
		string(LocationConstraintItemUniqueIdType_Bing),
		string(LocationConstraintItemUniqueIdType_Directory),
		string(LocationConstraintItemUniqueIdType_LocationStore),
		string(LocationConstraintItemUniqueIdType_Private),
		string(LocationConstraintItemUniqueIdType_Unknown),
	}
}

func (s *LocationConstraintItemUniqueIdType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLocationConstraintItemUniqueIdType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLocationConstraintItemUniqueIdType(input string) (*LocationConstraintItemUniqueIdType, error) {
	vals := map[string]LocationConstraintItemUniqueIdType{
		"bing":          LocationConstraintItemUniqueIdType_Bing,
		"directory":     LocationConstraintItemUniqueIdType_Directory,
		"locationstore": LocationConstraintItemUniqueIdType_LocationStore,
		"private":       LocationConstraintItemUniqueIdType_Private,
		"unknown":       LocationConstraintItemUniqueIdType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LocationConstraintItemUniqueIdType(input)
	return &out, nil
}
