package sitelogsconfigoperationgroup

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LogLevel string

const (
	LogLevelError       LogLevel = "Error"
	LogLevelInformation LogLevel = "Information"
	LogLevelOff         LogLevel = "Off"
	LogLevelVerbose     LogLevel = "Verbose"
	LogLevelWarning     LogLevel = "Warning"
)

func PossibleValuesForLogLevel() []string {
	return []string{
		string(LogLevelError),
		string(LogLevelInformation),
		string(LogLevelOff),
		string(LogLevelVerbose),
		string(LogLevelWarning),
	}
}

func (s *LogLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLogLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLogLevel(input string) (*LogLevel, error) {
	vals := map[string]LogLevel{
		"error":       LogLevelError,
		"information": LogLevelInformation,
		"off":         LogLevelOff,
		"verbose":     LogLevelVerbose,
		"warning":     LogLevelWarning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LogLevel(input)
	return &out, nil
}
