package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesDeploymentStateReasonValue string

const (
	WindowsUpdatesDeploymentStateReasonValuefaultedByContentOutdated WindowsUpdatesDeploymentStateReasonValue = "FaultedByContentOutdated"
	WindowsUpdatesDeploymentStateReasonValueofferingByRequest        WindowsUpdatesDeploymentStateReasonValue = "OfferingByRequest"
	WindowsUpdatesDeploymentStateReasonValuepausedByMonitoring       WindowsUpdatesDeploymentStateReasonValue = "PausedByMonitoring"
	WindowsUpdatesDeploymentStateReasonValuepausedByRequest          WindowsUpdatesDeploymentStateReasonValue = "PausedByRequest"
	WindowsUpdatesDeploymentStateReasonValuescheduledByOfferWindow   WindowsUpdatesDeploymentStateReasonValue = "ScheduledByOfferWindow"
)

func PossibleValuesForWindowsUpdatesDeploymentStateReasonValue() []string {
	return []string{
		string(WindowsUpdatesDeploymentStateReasonValuefaultedByContentOutdated),
		string(WindowsUpdatesDeploymentStateReasonValueofferingByRequest),
		string(WindowsUpdatesDeploymentStateReasonValuepausedByMonitoring),
		string(WindowsUpdatesDeploymentStateReasonValuepausedByRequest),
		string(WindowsUpdatesDeploymentStateReasonValuescheduledByOfferWindow),
	}
}

func parseWindowsUpdatesDeploymentStateReasonValue(input string) (*WindowsUpdatesDeploymentStateReasonValue, error) {
	vals := map[string]WindowsUpdatesDeploymentStateReasonValue{
		"faultedbycontentoutdated": WindowsUpdatesDeploymentStateReasonValuefaultedByContentOutdated,
		"offeringbyrequest":        WindowsUpdatesDeploymentStateReasonValueofferingByRequest,
		"pausedbymonitoring":       WindowsUpdatesDeploymentStateReasonValuepausedByMonitoring,
		"pausedbyrequest":          WindowsUpdatesDeploymentStateReasonValuepausedByRequest,
		"scheduledbyofferwindow":   WindowsUpdatesDeploymentStateReasonValuescheduledByOfferWindow,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUpdatesDeploymentStateReasonValue(input)
	return &out, nil
}
