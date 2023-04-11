package jitnetworkaccesspolicies

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JitNetworkAccessPolicy struct {
	Id         *string                          `json:"id,omitempty"`
	Kind       *string                          `json:"kind,omitempty"`
	Location   *string                          `json:"location,omitempty"`
	Name       *string                          `json:"name,omitempty"`
	Properties JitNetworkAccessPolicyProperties `json:"properties"`
	Type       *string                          `json:"type,omitempty"`
}
