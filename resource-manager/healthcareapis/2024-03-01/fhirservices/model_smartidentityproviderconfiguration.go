package fhirservices

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SmartIdentityProviderConfiguration struct {
	Applications *[]SmartIdentityProviderApplication `json:"applications,omitempty"`
	Authority    *string                             `json:"authority,omitempty"`
}
