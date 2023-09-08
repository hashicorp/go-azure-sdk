package me

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateMeChangePasswordRequest struct {
	CurrentPassword *string `json:"currentPassword,omitempty"`
	NewPassword     *string `json:"newPassword,omitempty"`
}
