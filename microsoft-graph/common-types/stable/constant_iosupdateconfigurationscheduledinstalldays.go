package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosUpdateConfigurationScheduledInstallDays string

const (
	IosUpdateConfigurationScheduledInstallDays_Friday    IosUpdateConfigurationScheduledInstallDays = "friday"
	IosUpdateConfigurationScheduledInstallDays_Monday    IosUpdateConfigurationScheduledInstallDays = "monday"
	IosUpdateConfigurationScheduledInstallDays_Saturday  IosUpdateConfigurationScheduledInstallDays = "saturday"
	IosUpdateConfigurationScheduledInstallDays_Sunday    IosUpdateConfigurationScheduledInstallDays = "sunday"
	IosUpdateConfigurationScheduledInstallDays_Thursday  IosUpdateConfigurationScheduledInstallDays = "thursday"
	IosUpdateConfigurationScheduledInstallDays_Tuesday   IosUpdateConfigurationScheduledInstallDays = "tuesday"
	IosUpdateConfigurationScheduledInstallDays_Wednesday IosUpdateConfigurationScheduledInstallDays = "wednesday"
)

func PossibleValuesForIosUpdateConfigurationScheduledInstallDays() []string {
	return []string{
		string(IosUpdateConfigurationScheduledInstallDays_Friday),
		string(IosUpdateConfigurationScheduledInstallDays_Monday),
		string(IosUpdateConfigurationScheduledInstallDays_Saturday),
		string(IosUpdateConfigurationScheduledInstallDays_Sunday),
		string(IosUpdateConfigurationScheduledInstallDays_Thursday),
		string(IosUpdateConfigurationScheduledInstallDays_Tuesday),
		string(IosUpdateConfigurationScheduledInstallDays_Wednesday),
	}
}

func (s *IosUpdateConfigurationScheduledInstallDays) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosUpdateConfigurationScheduledInstallDays(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosUpdateConfigurationScheduledInstallDays(input string) (*IosUpdateConfigurationScheduledInstallDays, error) {
	vals := map[string]IosUpdateConfigurationScheduledInstallDays{
		"friday":    IosUpdateConfigurationScheduledInstallDays_Friday,
		"monday":    IosUpdateConfigurationScheduledInstallDays_Monday,
		"saturday":  IosUpdateConfigurationScheduledInstallDays_Saturday,
		"sunday":    IosUpdateConfigurationScheduledInstallDays_Sunday,
		"thursday":  IosUpdateConfigurationScheduledInstallDays_Thursday,
		"tuesday":   IosUpdateConfigurationScheduledInstallDays_Tuesday,
		"wednesday": IosUpdateConfigurationScheduledInstallDays_Wednesday,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosUpdateConfigurationScheduledInstallDays(input)
	return &out, nil
}
