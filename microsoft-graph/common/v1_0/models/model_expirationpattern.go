package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExpirationPattern struct {
	Duration    *string                `json:"duration,omitempty"`
	EndDateTime *string                `json:"endDateTime,omitempty"`
	ODataType   *string                `json:"@odata.type,omitempty"`
	Type        *ExpirationPatternType `json:"type,omitempty"`
}
