package feature

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Feature struct {
	DataType    *FeatureDataType   `json:"dataType,omitempty"`
	Description *string            `json:"description,omitempty"`
	FeatureName *string            `json:"featureName,omitempty"`
	Properties  *map[string]string `json:"properties,omitempty"`
	Tags        *map[string]string `json:"tags,omitempty"`
}
