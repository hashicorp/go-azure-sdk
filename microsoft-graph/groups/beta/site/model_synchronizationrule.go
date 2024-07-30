package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SynchronizationRule struct {
	ContainerFilter     *ContainerFilter            `json:"containerFilter,omitempty"`
	Editable            *bool                       `json:"editable,omitempty"`
	GroupFilter         *GroupFilter                `json:"groupFilter,omitempty"`
	Id                  *string                     `json:"id,omitempty"`
	Metadata            *[]StringKeyStringValuePair `json:"metadata,omitempty"`
	Name                *string                     `json:"name,omitempty"`
	ODataType           *string                     `json:"@odata.type,omitempty"`
	ObjectMappings      *[]ObjectMapping            `json:"objectMappings,omitempty"`
	Priority            *int64                      `json:"priority,omitempty"`
	SourceDirectoryName *string                     `json:"sourceDirectoryName,omitempty"`
	TargetDirectoryName *string                     `json:"targetDirectoryName,omitempty"`
}
