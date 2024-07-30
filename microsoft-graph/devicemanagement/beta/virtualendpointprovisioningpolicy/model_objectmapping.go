package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ObjectMapping struct {
	AttributeMappings *[]AttributeMapping           `json:"attributeMappings,omitempty"`
	Enabled           *bool                         `json:"enabled,omitempty"`
	FlowTypes         *ObjectMappingFlowTypes       `json:"flowTypes,omitempty"`
	Metadata          *[]ObjectMappingMetadataEntry `json:"metadata,omitempty"`
	Name              *string                       `json:"name,omitempty"`
	ODataType         *string                       `json:"@odata.type,omitempty"`
	Scope             *Filter                       `json:"scope,omitempty"`
	SourceObjectName  *string                       `json:"sourceObjectName,omitempty"`
	TargetObjectName  *string                       `json:"targetObjectName,omitempty"`
}
