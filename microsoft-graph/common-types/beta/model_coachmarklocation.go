package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CoachmarkLocation struct {
	Length    *int64                 `json:"length,omitempty"`
	ODataType *string                `json:"@odata.type,omitempty"`
	Offset    *int64                 `json:"offset,omitempty"`
	Type      *CoachmarkLocationType `json:"type,omitempty"`
}
