package entities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IconDefinition struct {
	CustomData *string `json:"customData,omitempty"`
	IconName   string  `json:"iconName"`
}
