package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkHardwareConfiguration struct {
	Compute        *TeamworkPeripheral `json:"compute,omitempty"`
	HdmiIngest     *TeamworkPeripheral `json:"hdmiIngest,omitempty"`
	ODataType      *string             `json:"@odata.type,omitempty"`
	ProcessorModel *string             `json:"processorModel,omitempty"`
}
