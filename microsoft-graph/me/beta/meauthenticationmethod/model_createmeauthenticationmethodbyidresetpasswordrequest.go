package meauthenticationmethod

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateMeAuthenticationMethodByIdResetPasswordRequest struct {
	NewPassword *string `json:"newPassword,omitempty"`
}
