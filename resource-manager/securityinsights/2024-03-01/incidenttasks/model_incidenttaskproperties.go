package incidenttasks

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IncidentTaskProperties struct {
	CreatedBy           *ClientInfo        `json:"createdBy,omitempty"`
	CreatedTimeUtc      *string            `json:"createdTimeUtc,omitempty"`
	Description         *string            `json:"description,omitempty"`
	LastModifiedBy      *ClientInfo        `json:"lastModifiedBy,omitempty"`
	LastModifiedTimeUtc *string            `json:"lastModifiedTimeUtc,omitempty"`
	Status              IncidentTaskStatus `json:"status"`
	Title               string             `json:"title"`
}
