package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeliveryOptimizationGroupIdSourceOptions struct {
	GroupIdSourceOption *DeliveryOptimizationGroupIdSourceOptionsGroupIdSourceOption `json:"groupIdSourceOption,omitempty"`
	ODataType           *string                                                      `json:"@odata.type,omitempty"`
}
