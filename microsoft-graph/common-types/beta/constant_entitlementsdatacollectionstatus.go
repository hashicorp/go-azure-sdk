package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementsDataCollectionStatus string

const (
	EntitlementsDataCollectionStatus_Offline EntitlementsDataCollectionStatus = "offline"
	EntitlementsDataCollectionStatus_Online  EntitlementsDataCollectionStatus = "online"
)

func PossibleValuesForEntitlementsDataCollectionStatus() []string {
	return []string{
		string(EntitlementsDataCollectionStatus_Offline),
		string(EntitlementsDataCollectionStatus_Online),
	}
}

func (s *EntitlementsDataCollectionStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEntitlementsDataCollectionStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEntitlementsDataCollectionStatus(input string) (*EntitlementsDataCollectionStatus, error) {
	vals := map[string]EntitlementsDataCollectionStatus{
		"offline": EntitlementsDataCollectionStatus_Offline,
		"online":  EntitlementsDataCollectionStatus_Online,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EntitlementsDataCollectionStatus(input)
	return &out, nil
}
