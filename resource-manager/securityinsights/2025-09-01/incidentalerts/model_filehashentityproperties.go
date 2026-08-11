package incidentalerts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FileHashEntityProperties struct {
	AdditionalData *map[string]interface{} `json:"additionalData,omitempty"`
	Algorithm      *FileHashAlgorithm      `json:"algorithm,omitempty"`
	FriendlyName   *string                 `json:"friendlyName,omitempty"`
	HashValue      *string                 `json:"hashValue,omitempty"`
}
