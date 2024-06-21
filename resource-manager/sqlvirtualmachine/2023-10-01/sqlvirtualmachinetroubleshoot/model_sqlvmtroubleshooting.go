package sqlvirtualmachinetroubleshoot

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlVMTroubleshooting struct {
	EndTimeUtc               *string                              `json:"endTimeUtc,omitempty"`
	Properties               *TroubleshootingAdditionalProperties `json:"properties,omitempty"`
	StartTimeUtc             *string                              `json:"startTimeUtc,omitempty"`
	TroubleshootingScenario  *TroubleshootingScenario             `json:"troubleshootingScenario,omitempty"`
	VirtualMachineResourceId *string                              `json:"virtualMachineResourceId,omitempty"`
}
