package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCBulkRemoteActionResult struct {
	FailedDeviceIds       *[]string `json:"failedDeviceIds,omitempty"`
	NotFoundDeviceIds     *[]string `json:"notFoundDeviceIds,omitempty"`
	NotSupportedDeviceIds *[]string `json:"notSupportedDeviceIds,omitempty"`
	ODataType             *string   `json:"@odata.type,omitempty"`
	SuccessfulDeviceIds   *[]string `json:"successfulDeviceIds,omitempty"`
}
