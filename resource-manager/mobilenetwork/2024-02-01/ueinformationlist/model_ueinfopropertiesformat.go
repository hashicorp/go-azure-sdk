package ueinformationlist

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UeInfoPropertiesFormat struct {
	LastReadAt    *string      `json:"lastReadAt,omitempty"`
	RatType       RatType      `json:"ratType"`
	UeIPAddresses *[]DnnIPPair `json:"ueIpAddresses,omitempty"`
	UeState       UeState      `json:"ueState"`
}
