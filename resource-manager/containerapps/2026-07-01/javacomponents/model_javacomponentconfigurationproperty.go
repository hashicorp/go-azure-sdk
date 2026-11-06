package javacomponents

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JavaComponentConfigurationProperty struct {
	PropertyName *string `json:"propertyName,omitempty"`
	Value        *string `json:"value,omitempty"`
}
