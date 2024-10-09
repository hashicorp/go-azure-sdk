package bookmark

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FileEntityProperties struct {
	AdditionalData    *map[string]interface{} `json:"additionalData,omitempty"`
	Directory         *string                 `json:"directory,omitempty"`
	FileHashEntityIds *[]string               `json:"fileHashEntityIds,omitempty"`
	FileName          *string                 `json:"fileName,omitempty"`
	FriendlyName      *string                 `json:"friendlyName,omitempty"`
	HostEntityId      *string                 `json:"hostEntityId,omitempty"`
}
