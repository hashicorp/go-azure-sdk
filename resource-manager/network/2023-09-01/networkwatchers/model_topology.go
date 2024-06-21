package networkwatchers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Topology struct {
	CreatedDateTime *string             `json:"createdDateTime,omitempty"`
	Id              *string             `json:"id,omitempty"`
	LastModified    *string             `json:"lastModified,omitempty"`
	Resources       *[]TopologyResource `json:"resources,omitempty"`
}
