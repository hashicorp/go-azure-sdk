package allowedconnections

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AllowedConnectionsResourceProperties struct {
	CalculatedDateTime   *string                `json:"calculatedDateTime,omitempty"`
	ConnectableResources *[]ConnectableResource `json:"connectableResources,omitempty"`
}
