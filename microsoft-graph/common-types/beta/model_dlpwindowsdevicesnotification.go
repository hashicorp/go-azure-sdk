package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DlpWindowsDevicesNotification struct {
	Author        *string `json:"author,omitempty"`
	ContentName   *string `json:"contentName,omitempty"`
	LastModfiedBy *string `json:"lastModfiedBy,omitempty"`
	ODataType     *string `json:"@odata.type,omitempty"`
}
