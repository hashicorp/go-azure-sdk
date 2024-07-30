package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WinGetAppInstallTimeSettings struct {
	DeadlineDateTime *string `json:"deadlineDateTime,omitempty"`
	ODataType        *string `json:"@odata.type,omitempty"`
	UseLocalTime     *bool   `json:"useLocalTime,omitempty"`
}
