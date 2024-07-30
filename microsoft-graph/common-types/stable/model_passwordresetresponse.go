package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PasswordResetResponse struct {
	NewPassword *string `json:"newPassword,omitempty"`
	ODataType   *string `json:"@odata.type,omitempty"`
}
