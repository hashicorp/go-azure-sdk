package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesMonitoringRule struct {
	Action    *WindowsUpdatesMonitoringRuleAction `json:"action,omitempty"`
	ODataType *string                             `json:"@odata.type,omitempty"`
	Signal    *WindowsUpdatesMonitoringRuleSignal `json:"signal,omitempty"`
	Threshold *int64                              `json:"threshold,omitempty"`
}
