package jitnetworkaccesspolicies

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JitNetworkAccessRequest struct {
	Justification   *string                                 `json:"justification,omitempty"`
	Requestor       string                                  `json:"requestor"`
	StartTimeUtc    string                                  `json:"startTimeUtc"`
	VirtualMachines []JitNetworkAccessRequestVirtualMachine `json:"virtualMachines"`
}

func (o *JitNetworkAccessRequest) GetStartTimeUtcAsTime() (*time.Time, error) {
	return dates.ParseAsFormat(&o.StartTimeUtc, "2006-01-02T15:04:05Z07:00")
}

func (o *JitNetworkAccessRequest) SetStartTimeUtcAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.StartTimeUtc = formatted
}
