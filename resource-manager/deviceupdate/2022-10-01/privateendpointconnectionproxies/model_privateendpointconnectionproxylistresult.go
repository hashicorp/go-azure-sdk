package privateendpointconnectionproxies

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateEndpointConnectionProxyListResult struct {
	NextLink *string                           `json:"nextLink,omitempty"`
	Value    *[]PrivateEndpointConnectionProxy `json:"value,omitempty"`
}
