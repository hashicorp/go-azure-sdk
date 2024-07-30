package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationSynchronizationError struct {
	EntryType            *string `json:"entryType,omitempty"`
	ErrorCode            *string `json:"errorCode,omitempty"`
	ErrorMessage         *string `json:"errorMessage,omitempty"`
	Id                   *string `json:"id,omitempty"`
	JoiningValue         *string `json:"joiningValue,omitempty"`
	ODataType            *string `json:"@odata.type,omitempty"`
	RecordedDateTime     *string `json:"recordedDateTime,omitempty"`
	ReportableIdentifier *string `json:"reportableIdentifier,omitempty"`
}
