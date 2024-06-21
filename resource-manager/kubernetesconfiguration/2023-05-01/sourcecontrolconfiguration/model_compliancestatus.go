package sourcecontrolconfiguration

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComplianceStatus struct {
	ComplianceState   *ComplianceStateType `json:"complianceState,omitempty"`
	LastConfigApplied *string              `json:"lastConfigApplied,omitempty"`
	Message           *string              `json:"message,omitempty"`
	MessageLevel      *MessageLevelType    `json:"messageLevel,omitempty"`
}
