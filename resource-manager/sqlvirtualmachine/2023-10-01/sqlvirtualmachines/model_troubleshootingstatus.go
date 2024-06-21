package sqlvirtualmachines

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TroubleshootingStatus struct {
	EndTimeUtc              *string                              `json:"endTimeUtc,omitempty"`
	LastTriggerTimeUtc      *string                              `json:"lastTriggerTimeUtc,omitempty"`
	Properties              *TroubleshootingAdditionalProperties `json:"properties,omitempty"`
	RootCause               *string                              `json:"rootCause,omitempty"`
	StartTimeUtc            *string                              `json:"startTimeUtc,omitempty"`
	TroubleshootingScenario *TroubleshootingScenario             `json:"troubleshootingScenario,omitempty"`
}
