package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UsedInsight struct {
	Id                    *string                `json:"id,omitempty"`
	LastUsed              *UsageDetails          `json:"lastUsed,omitempty"`
	ODataType             *string                `json:"@odata.type,omitempty"`
	Resource              *Entity                `json:"resource,omitempty"`
	ResourceReference     *ResourceReference     `json:"resourceReference,omitempty"`
	ResourceVisualization *ResourceVisualization `json:"resourceVisualization,omitempty"`
}
