package jitnetworkaccesspolicies

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JitNetworkAccessPolicyInitiatePort struct {
	AllowedSourceAddressPrefix *string `json:"allowedSourceAddressPrefix,omitempty"`
	EndTimeUtc                 string  `json:"endTimeUtc"`
	Number                     int64   `json:"number"`
}

func (o *JitNetworkAccessPolicyInitiatePort) GetEndTimeUtcAsTime() (*time.Time, error) {
	return dates.ParseAsFormat(&o.EndTimeUtc, "2006-01-02T15:04:05Z07:00")
}

func (o *JitNetworkAccessPolicyInitiatePort) SetEndTimeUtcAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EndTimeUtc = formatted
}
