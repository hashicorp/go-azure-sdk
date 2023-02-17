package configurations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DigestConfig struct {
	ActionGroupResourceId *string            `json:"actionGroupResourceId,omitempty"`
	Categories            *[]Category        `json:"categories,omitempty"`
	Frequency             *int64             `json:"frequency,omitempty"`
	Language              *string            `json:"language,omitempty"`
	Name                  *string            `json:"name,omitempty"`
	State                 *DigestConfigState `json:"state,omitempty"`
}
