package networkclouds

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConsoleProperties struct {
	DetailedStatus         *ConsoleDetailedStatus    `json:"detailedStatus,omitempty"`
	DetailedStatusMessage  *string                   `json:"detailedStatusMessage,omitempty"`
	Enabled                ConsoleEnabled            `json:"enabled"`
	Expiration             *string                   `json:"expiration,omitempty"`
	PrivateLinkServiceId   *string                   `json:"privateLinkServiceId,omitempty"`
	ProvisioningState      *ConsoleProvisioningState `json:"provisioningState,omitempty"`
	SshPublicKey           SshPublicKey              `json:"sshPublicKey"`
	VirtualMachineAccessId *string                   `json:"virtualMachineAccessId,omitempty"`
}

func (o *ConsoleProperties) GetExpirationAsTime() (*time.Time, error) {
	if o.Expiration == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.Expiration, "2006-01-02T15:04:05Z07:00")
}

func (o *ConsoleProperties) SetExpirationAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.Expiration = &formatted
}
