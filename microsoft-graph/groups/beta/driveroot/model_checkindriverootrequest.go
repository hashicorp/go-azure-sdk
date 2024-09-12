package driveroot

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CheckinDriveRootRequest struct {
	CheckInAs nullable.Type[string] `json:"checkInAs,omitempty"`
	Comment   nullable.Type[string] `json:"comment,omitempty"`
}
