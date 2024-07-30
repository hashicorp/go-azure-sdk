package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SynchronizationSchema struct {
	Directories          *[]DirectoryDefinition `json:"directories,omitempty"`
	Id                   *string                `json:"id,omitempty"`
	ODataType            *string                `json:"@odata.type,omitempty"`
	SynchronizationRules *[]SynchronizationRule `json:"synchronizationRules,omitempty"`
	Version              *string                `json:"version,omitempty"`
}
