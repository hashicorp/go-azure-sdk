package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoomBookingType string

const (
	RoomBookingType_Reserved RoomBookingType = "reserved"
	RoomBookingType_Standard RoomBookingType = "standard"
	RoomBookingType_Unknown  RoomBookingType = "unknown"
)

func PossibleValuesForRoomBookingType() []string {
	return []string{
		string(RoomBookingType_Reserved),
		string(RoomBookingType_Standard),
		string(RoomBookingType_Unknown),
	}
}

func (s *RoomBookingType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRoomBookingType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRoomBookingType(input string) (*RoomBookingType, error) {
	vals := map[string]RoomBookingType{
		"reserved": RoomBookingType_Reserved,
		"standard": RoomBookingType_Standard,
		"unknown":  RoomBookingType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RoomBookingType(input)
	return &out, nil
}
