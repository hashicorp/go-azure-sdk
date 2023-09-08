package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCConnectionSettings struct {
	EnableSingleSignOn *bool   `json:"enableSingleSignOn,omitempty"`
	ODataType          *string `json:"@odata.type,omitempty"`
}
