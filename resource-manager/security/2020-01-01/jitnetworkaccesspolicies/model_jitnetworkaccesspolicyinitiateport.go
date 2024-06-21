package jitnetworkaccesspolicies

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JitNetworkAccessPolicyInitiatePort struct {
	AllowedSourceAddressPrefix *string `json:"allowedSourceAddressPrefix,omitempty"`
	EndTimeUtc                 string  `json:"endTimeUtc"`
	Number                     int64   `json:"number"`
}
