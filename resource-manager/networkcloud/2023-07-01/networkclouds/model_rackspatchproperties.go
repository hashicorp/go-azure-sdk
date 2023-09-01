package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RacksPatchProperties struct {
	RackLocation     *string `json:"rackLocation,omitempty"`
	RackSerialNumber *string `json:"rackSerialNumber,omitempty"`
}
