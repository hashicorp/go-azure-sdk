package driveitem

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CheckinDriveItemRequest struct {
	CheckInAs nullable.Type[string] `json:"checkInAs,omitempty"`
	Comment   nullable.Type[string] `json:"comment,omitempty"`
}
