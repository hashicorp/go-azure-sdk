package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharePointOneDriveOptions struct {
	IncludeContent *SharePointOneDriveOptionsIncludeContent `json:"includeContent,omitempty"`
	ODataType      *string                                  `json:"@odata.type,omitempty"`
}
