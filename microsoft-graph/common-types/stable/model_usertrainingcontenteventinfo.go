package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserTrainingContentEventInfo struct {
	Browser                 *string  `json:"browser,omitempty"`
	ContentDateTime         *string  `json:"contentDateTime,omitempty"`
	IpAddress               *string  `json:"ipAddress,omitempty"`
	ODataType               *string  `json:"@odata.type,omitempty"`
	OsPlatformDeviceDetails *string  `json:"osPlatformDeviceDetails,omitempty"`
	PotentialScoreImpact    *float64 `json:"potentialScoreImpact,omitempty"`
}
