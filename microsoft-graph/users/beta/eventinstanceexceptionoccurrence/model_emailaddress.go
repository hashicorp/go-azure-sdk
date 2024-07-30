package eventinstanceexceptionoccurrence

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmailAddress struct {
	Address   *string `json:"address,omitempty"`
	Name      *string `json:"name,omitempty"`
	ODataType *string `json:"@odata.type,omitempty"`
}
