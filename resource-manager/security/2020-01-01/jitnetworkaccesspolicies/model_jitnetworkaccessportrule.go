package jitnetworkaccesspolicies

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JitNetworkAccessPortRule struct {
	AllowedSourceAddressPrefix   *string   `json:"allowedSourceAddressPrefix,omitempty"`
	AllowedSourceAddressPrefixes *[]string `json:"allowedSourceAddressPrefixes,omitempty"`
	MaxRequestAccessDuration     string    `json:"maxRequestAccessDuration"`
	Number                       int64     `json:"number"`
	Protocol                     Protocol  `json:"protocol"`
}
