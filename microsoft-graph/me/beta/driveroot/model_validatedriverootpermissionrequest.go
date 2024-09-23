package driveroot

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ValidateDriveRootPermissionRequest struct {
	ChallengeToken nullable.Type[string] `json:"challengeToken,omitempty"`
	Password       *string               `json:"password,omitempty"`
}
