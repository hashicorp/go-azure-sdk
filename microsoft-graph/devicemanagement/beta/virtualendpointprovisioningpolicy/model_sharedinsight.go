package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharedInsight struct {
	Id                    *string                `json:"id,omitempty"`
	LastShared            *SharingDetail         `json:"lastShared,omitempty"`
	LastSharedMethod      *Entity                `json:"lastSharedMethod,omitempty"`
	ODataType             *string                `json:"@odata.type,omitempty"`
	Resource              *Entity                `json:"resource,omitempty"`
	ResourceReference     *ResourceReference     `json:"resourceReference,omitempty"`
	ResourceVisualization *ResourceVisualization `json:"resourceVisualization,omitempty"`
	SharingHistory        *[]SharingDetail       `json:"sharingHistory,omitempty"`
}
