package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LastSignIn struct {
	InteractiveDateTime    *string `json:"interactiveDateTime,omitempty"`
	NonInteractiveDateTime *string `json:"nonInteractiveDateTime,omitempty"`
	ODataType              *string `json:"@odata.type,omitempty"`
}
