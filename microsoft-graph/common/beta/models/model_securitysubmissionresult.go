package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecuritySubmissionResult struct {
	Category           *SecuritySubmissionResultCategory           `json:"category,omitempty"`
	Detail             *SecuritySubmissionResultDetail             `json:"detail,omitempty"`
	DetectedFiles      *[]SecuritySubmissionDetectedFile           `json:"detectedFiles,omitempty"`
	DetectedUrls       *[]string                                   `json:"detectedUrls,omitempty"`
	ODataType          *string                                     `json:"@odata.type,omitempty"`
	UserMailboxSetting *SecuritySubmissionResultUserMailboxSetting `json:"userMailboxSetting,omitempty"`
}
