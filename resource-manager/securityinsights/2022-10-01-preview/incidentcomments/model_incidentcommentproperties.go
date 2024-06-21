package incidentcomments

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IncidentCommentProperties struct {
	Author              *ClientInfo `json:"author,omitempty"`
	CreatedTimeUtc      *string     `json:"createdTimeUtc,omitempty"`
	LastModifiedTimeUtc *string     `json:"lastModifiedTimeUtc,omitempty"`
	Message             string      `json:"message"`
}
