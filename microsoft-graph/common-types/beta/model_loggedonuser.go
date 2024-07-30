package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LoggedOnUser struct {
	LastLogOnDateTime *string `json:"lastLogOnDateTime,omitempty"`
	ODataType         *string `json:"@odata.type,omitempty"`
	UserId            *string `json:"userId,omitempty"`
}
