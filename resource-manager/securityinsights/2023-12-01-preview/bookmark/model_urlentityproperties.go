package bookmark

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type URLEntityProperties struct {
	AdditionalData *interface{} `json:"additionalData,omitempty"`
	FriendlyName   *string      `json:"friendlyName,omitempty"`
	Url            *string      `json:"url,omitempty"`
}
