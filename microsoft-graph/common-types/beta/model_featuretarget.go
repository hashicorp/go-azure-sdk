package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FeatureTarget struct {
	Id         *string                  `json:"id,omitempty"`
	ODataType  *string                  `json:"@odata.type,omitempty"`
	TargetType *FeatureTargetTargetType `json:"targetType,omitempty"`
}
