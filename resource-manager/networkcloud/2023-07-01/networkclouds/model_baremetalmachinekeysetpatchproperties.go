package networkclouds

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BareMetalMachineKeySetPatchProperties struct {
	Expiration       *string       `json:"expiration,omitempty"`
	JumpHostsAllowed *[]string     `json:"jumpHostsAllowed,omitempty"`
	UserList         *[]KeySetUser `json:"userList,omitempty"`
}

func (o *BareMetalMachineKeySetPatchProperties) GetExpirationAsTime() (*time.Time, error) {
	if o.Expiration == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.Expiration, "2006-01-02T15:04:05Z07:00")
}

func (o *BareMetalMachineKeySetPatchProperties) SetExpirationAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.Expiration = &formatted
}
