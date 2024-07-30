package lifecycleworkflowdeleteditemworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Photo struct {
	CameraMake          *string  `json:"cameraMake,omitempty"`
	CameraModel         *string  `json:"cameraModel,omitempty"`
	ExposureDenominator *float64 `json:"exposureDenominator,omitempty"`
	ExposureNumerator   *float64 `json:"exposureNumerator,omitempty"`
	FNumber             *float64 `json:"fNumber,omitempty"`
	FocalLength         *float64 `json:"focalLength,omitempty"`
	Iso                 *int64   `json:"iso,omitempty"`
	ODataType           *string  `json:"@odata.type,omitempty"`
	Orientation         *int64   `json:"orientation,omitempty"`
	TakenDateTime       *string  `json:"takenDateTime,omitempty"`
}
