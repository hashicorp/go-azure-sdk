package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsAddLogRequest struct {
	LogInformation *string `json:"logInformation,omitempty"`
	ODataType      *string `json:"@odata.type,omitempty"`
}
