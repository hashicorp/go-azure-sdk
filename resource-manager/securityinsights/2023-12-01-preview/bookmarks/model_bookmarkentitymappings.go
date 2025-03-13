package bookmarks

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookmarkEntityMappings struct {
	EntityType    *string               `json:"entityType,omitempty"`
	FieldMappings *[]EntityFieldMapping `json:"fieldMappings,omitempty"`
}
