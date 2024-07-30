package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProfileCardProperty struct {
	Annotations           *[]ProfileCardAnnotation `json:"annotations,omitempty"`
	DirectoryPropertyName *string                  `json:"directoryPropertyName,omitempty"`
	Id                    *string                  `json:"id,omitempty"`
	ODataType             *string                  `json:"@odata.type,omitempty"`
}
