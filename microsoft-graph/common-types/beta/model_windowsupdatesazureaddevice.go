package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesAzureADDevice struct {
	Enrollments *[]WindowsUpdatesUpdatableAssetEnrollment `json:"enrollments,omitempty"`
	Errors      *[]WindowsUpdatesUpdatableAssetError      `json:"errors,omitempty"`
	Id          *string                                   `json:"id,omitempty"`
	ODataType   *string                                   `json:"@odata.type,omitempty"`
}
