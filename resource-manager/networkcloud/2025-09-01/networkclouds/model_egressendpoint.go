package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EgressEndpoint struct {
	Category  string               `json:"category"`
	Endpoints []EndpointDependency `json:"endpoints"`
}
