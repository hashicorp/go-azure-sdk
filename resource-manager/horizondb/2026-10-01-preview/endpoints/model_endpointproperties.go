package endpoints

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EndpointProperties struct {
	AssociatedResourceId     string        `json:"associatedResourceId"`
	FullyQualifiedDomainName string        `json:"fullyQualifiedDomainName"`
	Role                     EndpointRole  `json:"role"`
	Scope                    EndpointScope `json:"scope"`
}
