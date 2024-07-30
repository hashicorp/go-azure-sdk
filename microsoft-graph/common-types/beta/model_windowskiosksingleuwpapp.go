package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsKioskSingleUWPApp struct {
	ODataType *string             `json:"@odata.type,omitempty"`
	UwpApp    *WindowsKioskUWPApp `json:"uwpApp,omitempty"`
}
