package jitnetworkaccesspolicies

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JitNetworkAccessRequestPort struct {
	AllowedSourceAddressPrefix   *string      `json:"allowedSourceAddressPrefix,omitempty"`
	AllowedSourceAddressPrefixes *[]string    `json:"allowedSourceAddressPrefixes,omitempty"`
	EndTimeUtc                   string       `json:"endTimeUtc"`
	MappedPort                   *int64       `json:"mappedPort,omitempty"`
	Number                       int64        `json:"number"`
	Status                       Status       `json:"status"`
	StatusReason                 StatusReason `json:"statusReason"`
}
