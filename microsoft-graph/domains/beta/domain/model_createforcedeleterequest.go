package domain

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateForceDeleteRequest struct {
	DisableUserAccounts *bool `json:"disableUserAccounts,omitempty"`
}
