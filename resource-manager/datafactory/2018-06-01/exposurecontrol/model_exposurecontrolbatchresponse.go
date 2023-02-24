package exposurecontrol

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExposureControlBatchResponse struct {
	ExposureControlResponses []ExposureControlResponse `json:"exposureControlResponses"`
}
