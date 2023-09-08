package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkCameraConfiguration struct {
	Cameras                    *[]TeamworkPeripheral               `json:"cameras,omitempty"`
	ContentCameraConfiguration *TeamworkContentCameraConfiguration `json:"contentCameraConfiguration,omitempty"`
	DefaultContentCamera       *TeamworkPeripheral                 `json:"defaultContentCamera,omitempty"`
	ODataType                  *string                             `json:"@odata.type,omitempty"`
}
