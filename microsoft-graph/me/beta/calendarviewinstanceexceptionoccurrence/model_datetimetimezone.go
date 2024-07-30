package calendarviewinstanceexceptionoccurrence

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DateTimeTimeZone struct {
	DateTime  *string `json:"dateTime,omitempty"`
	ODataType *string `json:"@odata.type,omitempty"`
	TimeZone  *string `json:"timeZone,omitempty"`
}
