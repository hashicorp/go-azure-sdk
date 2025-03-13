package netappresource

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FilePathAvailabilityRequest struct {
	AvailabilityZone *string `json:"availabilityZone,omitempty"`
	Name             string  `json:"name"`
	SubnetId         string  `json:"subnetId"`
}
