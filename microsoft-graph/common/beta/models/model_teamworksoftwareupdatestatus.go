package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkSoftwareUpdateStatus struct {
	AvailableVersion  *string                                        `json:"availableVersion,omitempty"`
	CurrentVersion    *string                                        `json:"currentVersion,omitempty"`
	ODataType         *string                                        `json:"@odata.type,omitempty"`
	SoftwareFreshness *TeamworkSoftwareUpdateStatusSoftwareFreshness `json:"softwareFreshness,omitempty"`
}
