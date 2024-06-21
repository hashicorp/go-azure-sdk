package migrations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ValidationDetails struct {
	DbLevelValidationDetails     *[]DbLevelValidationStatus `json:"dbLevelValidationDetails,omitempty"`
	ServerLevelValidationDetails *[]ValidationSummaryItem   `json:"serverLevelValidationDetails,omitempty"`
	Status                       *ValidationState           `json:"status,omitempty"`
	ValidationEndTimeInUtc       *string                    `json:"validationEndTimeInUtc,omitempty"`
	ValidationStartTimeInUtc     *string                    `json:"validationStartTimeInUtc,omitempty"`
}
