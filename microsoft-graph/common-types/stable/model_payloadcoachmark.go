package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PayloadCoachmark struct {
	CoachmarkLocation *CoachmarkLocation `json:"coachmarkLocation,omitempty"`
	Description       *string            `json:"description,omitempty"`
	Indicator         *string            `json:"indicator,omitempty"`
	IsValid           *bool              `json:"isValid,omitempty"`
	Language          *string            `json:"language,omitempty"`
	ODataType         *string            `json:"@odata.type,omitempty"`
	Order             *string            `json:"order,omitempty"`
}
