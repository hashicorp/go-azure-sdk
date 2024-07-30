package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RequestSchedule struct {
	Expiration    *ExpirationPattern   `json:"expiration,omitempty"`
	ODataType     *string              `json:"@odata.type,omitempty"`
	Recurrence    *PatternedRecurrence `json:"recurrence,omitempty"`
	StartDateTime *string              `json:"startDateTime,omitempty"`
}
