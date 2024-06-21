package workbooksapis

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkbookProperties struct {
	Category       string    `json:"category"`
	Description    *string   `json:"description,omitempty"`
	DisplayName    string    `json:"displayName"`
	Revision       *string   `json:"revision,omitempty"`
	SerializedData string    `json:"serializedData"`
	SourceId       *string   `json:"sourceId,omitempty"`
	StorageUri     *string   `json:"storageUri,omitempty"`
	Tags           *[]string `json:"tags,omitempty"`
	TimeModified   *string   `json:"timeModified,omitempty"`
	UserId         *string   `json:"userId,omitempty"`
	Version        *string   `json:"version,omitempty"`
}
