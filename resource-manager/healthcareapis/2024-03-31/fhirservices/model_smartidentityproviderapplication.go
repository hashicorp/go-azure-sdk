package fhirservices

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SmartIdentityProviderApplication struct {
	AllowedDataActions *[]SmartDataActions `json:"allowedDataActions,omitempty"`
	Audience           *string             `json:"audience,omitempty"`
	ClientId           *string             `json:"clientId,omitempty"`
}
