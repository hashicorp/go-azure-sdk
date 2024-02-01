package sqlpoolsdatamaskingpolicies

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataMaskingPolicy struct {
	Id         *string                      `json:"id,omitempty"`
	Kind       *string                      `json:"kind,omitempty"`
	Location   *string                      `json:"location,omitempty"`
	ManagedBy  *string                      `json:"managedBy,omitempty"`
	Name       *string                      `json:"name,omitempty"`
	Properties *DataMaskingPolicyProperties `json:"properties,omitempty"`
	Type       *string                      `json:"type,omitempty"`
}
