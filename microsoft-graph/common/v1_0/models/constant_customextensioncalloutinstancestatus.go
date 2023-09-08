package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomExtensionCalloutInstanceStatus string

const (
	CustomExtensionCalloutInstanceStatuscallbackReceived   CustomExtensionCalloutInstanceStatus = "CallbackReceived"
	CustomExtensionCalloutInstanceStatuscallbackTimedOut   CustomExtensionCalloutInstanceStatus = "CallbackTimedOut"
	CustomExtensionCalloutInstanceStatuscalloutFailed      CustomExtensionCalloutInstanceStatus = "CalloutFailed"
	CustomExtensionCalloutInstanceStatuscalloutSent        CustomExtensionCalloutInstanceStatus = "CalloutSent"
	CustomExtensionCalloutInstanceStatuswaitingForCallback CustomExtensionCalloutInstanceStatus = "WaitingForCallback"
)

func PossibleValuesForCustomExtensionCalloutInstanceStatus() []string {
	return []string{
		string(CustomExtensionCalloutInstanceStatuscallbackReceived),
		string(CustomExtensionCalloutInstanceStatuscallbackTimedOut),
		string(CustomExtensionCalloutInstanceStatuscalloutFailed),
		string(CustomExtensionCalloutInstanceStatuscalloutSent),
		string(CustomExtensionCalloutInstanceStatuswaitingForCallback),
	}
}

func parseCustomExtensionCalloutInstanceStatus(input string) (*CustomExtensionCalloutInstanceStatus, error) {
	vals := map[string]CustomExtensionCalloutInstanceStatus{
		"callbackreceived":   CustomExtensionCalloutInstanceStatuscallbackReceived,
		"callbacktimedout":   CustomExtensionCalloutInstanceStatuscallbackTimedOut,
		"calloutfailed":      CustomExtensionCalloutInstanceStatuscalloutFailed,
		"calloutsent":        CustomExtensionCalloutInstanceStatuscalloutSent,
		"waitingforcallback": CustomExtensionCalloutInstanceStatuswaitingForCallback,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CustomExtensionCalloutInstanceStatus(input)
	return &out, nil
}
