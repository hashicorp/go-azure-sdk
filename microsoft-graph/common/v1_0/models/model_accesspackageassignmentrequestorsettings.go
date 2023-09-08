package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageAssignmentRequestorSettings struct {
	AllowCustomAssignmentSchedule          *bool         `json:"allowCustomAssignmentSchedule,omitempty"`
	EnableOnBehalfRequestorsToAddAccess    *bool         `json:"enableOnBehalfRequestorsToAddAccess,omitempty"`
	EnableOnBehalfRequestorsToRemoveAccess *bool         `json:"enableOnBehalfRequestorsToRemoveAccess,omitempty"`
	EnableOnBehalfRequestorsToUpdateAccess *bool         `json:"enableOnBehalfRequestorsToUpdateAccess,omitempty"`
	EnableTargetsToSelfAddAccess           *bool         `json:"enableTargetsToSelfAddAccess,omitempty"`
	EnableTargetsToSelfRemoveAccess        *bool         `json:"enableTargetsToSelfRemoveAccess,omitempty"`
	EnableTargetsToSelfUpdateAccess        *bool         `json:"enableTargetsToSelfUpdateAccess,omitempty"`
	ODataType                              *string       `json:"@odata.type,omitempty"`
	OnBehalfRequestors                     *[]SubjectSet `json:"onBehalfRequestors,omitempty"`
}
