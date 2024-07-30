package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventMessageRequestType string

const (
	EventMessageRequestType_Exception      EventMessageRequestType = "exception"
	EventMessageRequestType_Occurrence     EventMessageRequestType = "occurrence"
	EventMessageRequestType_SeriesMaster   EventMessageRequestType = "seriesMaster"
	EventMessageRequestType_SingleInstance EventMessageRequestType = "singleInstance"
)

func PossibleValuesForEventMessageRequestType() []string {
	return []string{
		string(EventMessageRequestType_Exception),
		string(EventMessageRequestType_Occurrence),
		string(EventMessageRequestType_SeriesMaster),
		string(EventMessageRequestType_SingleInstance),
	}
}

func (s *EventMessageRequestType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEventMessageRequestType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEventMessageRequestType(input string) (*EventMessageRequestType, error) {
	vals := map[string]EventMessageRequestType{
		"exception":      EventMessageRequestType_Exception,
		"occurrence":     EventMessageRequestType_Occurrence,
		"seriesmaster":   EventMessageRequestType_SeriesMaster,
		"singleinstance": EventMessageRequestType_SingleInstance,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EventMessageRequestType(input)
	return &out, nil
}
