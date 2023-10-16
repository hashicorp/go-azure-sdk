package workbooksapis

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkbookProperties struct {
	Category         string         `json:"category"`
	Kind             SharedTypeKind `json:"kind"`
	Name             string         `json:"name"`
	SerializedData   string         `json:"serializedData"`
	SourceResourceId *string        `json:"sourceResourceId,omitempty"`
	Tags             *[]string      `json:"tags,omitempty"`
	TimeModified     *string        `json:"timeModified,omitempty"`
	UserId           string         `json:"userId"`
	Version          *string        `json:"version,omitempty"`
	WorkbookId       string         `json:"workbookId"`
}
