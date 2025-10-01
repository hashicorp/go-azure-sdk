package bms

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SQLDataDirectoryMapping struct {
	MappingType       *SQLDataDirectoryType `json:"mappingType,omitempty"`
	SourceLogicalName *string               `json:"sourceLogicalName,omitempty"`
	SourcePath        *string               `json:"sourcePath,omitempty"`
	TargetPath        *string               `json:"targetPath,omitempty"`
}
