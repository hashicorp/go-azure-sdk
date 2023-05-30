package changedatacapture

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataMapperMapping struct {
	AttributeMappingInfo      *MapperAttributeMappings   `json:"attributeMappingInfo,omitempty"`
	SourceConnectionReference *MapperConnectionReference `json:"sourceConnectionReference,omitempty"`
	SourceDenormalizeInfo     *interface{}               `json:"sourceDenormalizeInfo,omitempty"`
	SourceEntityName          *string                    `json:"sourceEntityName,omitempty"`
	TargetEntityName          *string                    `json:"targetEntityName,omitempty"`
}
