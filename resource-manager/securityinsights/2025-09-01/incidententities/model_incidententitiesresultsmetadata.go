package incidententities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IncidentEntitiesResultsMetadata struct {
	Count      int64          `json:"count"`
	EntityKind EntityKindEnum `json:"entityKind"`
}
