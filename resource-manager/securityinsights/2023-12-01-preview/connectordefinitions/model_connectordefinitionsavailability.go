package connectordefinitions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectorDefinitionsAvailability struct {
	IsPreview *bool  `json:"isPreview,omitempty"`
	Status    *int64 `json:"status,omitempty"`
}
