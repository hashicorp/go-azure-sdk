package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventMessageType string

const (
	EventMessageType_Exception      EventMessageType = "exception"
	EventMessageType_Occurrence     EventMessageType = "occurrence"
	EventMessageType_SeriesMaster   EventMessageType = "seriesMaster"
	EventMessageType_SingleInstance EventMessageType = "singleInstance"
)

func PossibleValuesForEventMessageType() []string {
	return []string{
		string(EventMessageType_Exception),
		string(EventMessageType_Occurrence),
		string(EventMessageType_SeriesMaster),
		string(EventMessageType_SingleInstance),
	}
}

func (s *EventMessageType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEventMessageType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEventMessageType(input string) (*EventMessageType, error) {
	vals := map[string]EventMessageType{
		"exception":      EventMessageType_Exception,
		"occurrence":     EventMessageType_Occurrence,
		"seriesmaster":   EventMessageType_SeriesMaster,
		"singleinstance": EventMessageType_SingleInstance,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EventMessageType(input)
	return &out, nil
}
