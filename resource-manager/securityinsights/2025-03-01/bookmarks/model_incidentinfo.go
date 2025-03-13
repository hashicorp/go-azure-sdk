package bookmarks

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IncidentInfo struct {
	IncidentId   *string           `json:"incidentId,omitempty"`
	RelationName *string           `json:"relationName,omitempty"`
	Severity     *IncidentSeverity `json:"severity,omitempty"`
	Title        *string           `json:"title,omitempty"`
}
