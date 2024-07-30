package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesMonitoringRuleAction string

const (
	WindowsUpdatesMonitoringRuleAction_AlertError      WindowsUpdatesMonitoringRuleAction = "alertError"
	WindowsUpdatesMonitoringRuleAction_OfferFallback   WindowsUpdatesMonitoringRuleAction = "offerFallback"
	WindowsUpdatesMonitoringRuleAction_PauseDeployment WindowsUpdatesMonitoringRuleAction = "pauseDeployment"
)

func PossibleValuesForWindowsUpdatesMonitoringRuleAction() []string {
	return []string{
		string(WindowsUpdatesMonitoringRuleAction_AlertError),
		string(WindowsUpdatesMonitoringRuleAction_OfferFallback),
		string(WindowsUpdatesMonitoringRuleAction_PauseDeployment),
	}
}

func (s *WindowsUpdatesMonitoringRuleAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsUpdatesMonitoringRuleAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsUpdatesMonitoringRuleAction(input string) (*WindowsUpdatesMonitoringRuleAction, error) {
	vals := map[string]WindowsUpdatesMonitoringRuleAction{
		"alerterror":      WindowsUpdatesMonitoringRuleAction_AlertError,
		"offerfallback":   WindowsUpdatesMonitoringRuleAction_OfferFallback,
		"pausedeployment": WindowsUpdatesMonitoringRuleAction_PauseDeployment,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUpdatesMonitoringRuleAction(input)
	return &out, nil
}
