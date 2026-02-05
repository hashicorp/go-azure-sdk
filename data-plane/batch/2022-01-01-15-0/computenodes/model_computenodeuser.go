package computenodes

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComputeNodeUser struct {
	ExpiryTime   *string `json:"expiryTime,omitempty"`
	IsAdmin      *bool   `json:"isAdmin,omitempty"`
	Name         string  `json:"name"`
	Password     *string `json:"password,omitempty"`
	SshPublicKey *string `json:"sshPublicKey,omitempty"`
}

func (o *ComputeNodeUser) GetExpiryTimeAsTime() (*time.Time, error) {
	if o.ExpiryTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ExpiryTime, "2006-01-02T15:04:05Z07:00")
}

func (o *ComputeNodeUser) SetExpiryTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ExpiryTime = &formatted
}
