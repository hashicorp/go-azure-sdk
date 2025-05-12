package edgenodes

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CidrIPAddress struct {
	BaseIPAddress *string `json:"baseIpAddress,omitempty"`
	PrefixLength  *int64  `json:"prefixLength,omitempty"`
}
