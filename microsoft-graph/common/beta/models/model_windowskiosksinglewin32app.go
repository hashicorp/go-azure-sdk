package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsKioskSingleWin32App struct {
	ODataType *string               `json:"@odata.type,omitempty"`
	Win32App  *WindowsKioskWin32App `json:"win32App,omitempty"`
}
