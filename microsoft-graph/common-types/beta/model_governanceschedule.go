package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GovernanceSchedule struct {
	Duration      *string `json:"duration,omitempty"`
	EndDateTime   *string `json:"endDateTime,omitempty"`
	ODataType     *string `json:"@odata.type,omitempty"`
	StartDateTime *string `json:"startDateTime,omitempty"`
	Type          *string `json:"type,omitempty"`
}
