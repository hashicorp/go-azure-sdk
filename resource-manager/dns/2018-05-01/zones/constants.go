package zones

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ZoneType string

const (
	ZoneTypePrivate ZoneType = "Private"
	ZoneTypePublic  ZoneType = "Public"
)

func PossibleValuesForZoneType() []string {
	return []string{
		string(ZoneTypePrivate),
		string(ZoneTypePublic),
	}
}

func (s *ZoneType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForZoneType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = ZoneType(decoded)
	return nil
}
