package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventMessageResponseType string

const (
	EventMessageResponseType_Exception      EventMessageResponseType = "exception"
	EventMessageResponseType_Occurrence     EventMessageResponseType = "occurrence"
	EventMessageResponseType_SeriesMaster   EventMessageResponseType = "seriesMaster"
	EventMessageResponseType_SingleInstance EventMessageResponseType = "singleInstance"
)

func PossibleValuesForEventMessageResponseType() []string {
	return []string{
		string(EventMessageResponseType_Exception),
		string(EventMessageResponseType_Occurrence),
		string(EventMessageResponseType_SeriesMaster),
		string(EventMessageResponseType_SingleInstance),
	}
}

func (s *EventMessageResponseType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEventMessageResponseType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEventMessageResponseType(input string) (*EventMessageResponseType, error) {
	vals := map[string]EventMessageResponseType{
		"exception":      EventMessageResponseType_Exception,
		"occurrence":     EventMessageResponseType_Occurrence,
		"seriesmaster":   EventMessageResponseType_SeriesMaster,
		"singleinstance": EventMessageResponseType_SingleInstance,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EventMessageResponseType(input)
	return &out, nil
}
