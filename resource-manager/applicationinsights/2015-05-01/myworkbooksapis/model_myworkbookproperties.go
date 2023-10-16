package myworkbooksapis

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MyWorkbookProperties struct {
	Category       string    `json:"category"`
	DisplayName    string    `json:"displayName"`
	SerializedData string    `json:"serializedData"`
	SourceId       *string   `json:"sourceId,omitempty"`
	Tags           *[]string `json:"tags,omitempty"`
	TimeModified   *string   `json:"timeModified,omitempty"`
	UserId         *string   `json:"userId,omitempty"`
	Version        *string   `json:"version,omitempty"`
}
