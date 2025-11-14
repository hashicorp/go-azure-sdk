package openapis

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Reseller struct {
	ResellerDescription *string `json:"resellerDescription,omitempty"`
	ResellerId          *string `json:"resellerId,omitempty"`
}
