package provider

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsJavaContainerSettings struct {
	EndOfLifeDate        *string `json:"endOfLifeDate,omitempty"`
	IsAutoUpdate         *bool   `json:"isAutoUpdate,omitempty"`
	IsDeprecated         *bool   `json:"isDeprecated,omitempty"`
	IsEarlyAccess        *bool   `json:"isEarlyAccess,omitempty"`
	IsHidden             *bool   `json:"isHidden,omitempty"`
	IsPreview            *bool   `json:"isPreview,omitempty"`
	JavaContainer        *string `json:"javaContainer,omitempty"`
	JavaContainerVersion *string `json:"javaContainerVersion,omitempty"`
}
