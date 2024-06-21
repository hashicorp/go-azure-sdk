package provider

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LinuxJavaContainerSettings struct {
	EndOfLifeDate *string `json:"endOfLifeDate,omitempty"`
	IsAutoUpdate  *bool   `json:"isAutoUpdate,omitempty"`
	IsDeprecated  *bool   `json:"isDeprecated,omitempty"`
	IsEarlyAccess *bool   `json:"isEarlyAccess,omitempty"`
	IsHidden      *bool   `json:"isHidden,omitempty"`
	IsPreview     *bool   `json:"isPreview,omitempty"`
	Java11Runtime *string `json:"java11Runtime,omitempty"`
	Java8Runtime  *string `json:"java8Runtime,omitempty"`
}
