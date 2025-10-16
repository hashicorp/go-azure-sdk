package networkclouds

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BmcKeySetProperties struct {
	AzureGroupId          string                      `json:"azureGroupId"`
	DetailedStatus        *BmcKeySetDetailedStatus    `json:"detailedStatus,omitempty"`
	DetailedStatusMessage *string                     `json:"detailedStatusMessage,omitempty"`
	Expiration            string                      `json:"expiration"`
	LastValidation        *string                     `json:"lastValidation,omitempty"`
	PrivilegeLevel        BmcKeySetPrivilegeLevel     `json:"privilegeLevel"`
	ProvisioningState     *BmcKeySetProvisioningState `json:"provisioningState,omitempty"`
	UserList              []KeySetUser                `json:"userList"`
	UserListStatus        *[]KeySetUserStatus         `json:"userListStatus,omitempty"`
}

func (o *BmcKeySetProperties) GetExpirationAsTime() (*time.Time, error) {
	return dates.ParseAsFormat(&o.Expiration, "2006-01-02T15:04:05Z07:00")
}

func (o *BmcKeySetProperties) SetExpirationAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.Expiration = formatted
}

func (o *BmcKeySetProperties) GetLastValidationAsTime() (*time.Time, error) {
	if o.LastValidation == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastValidation, "2006-01-02T15:04:05Z07:00")
}

func (o *BmcKeySetProperties) SetLastValidationAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastValidation = &formatted
}
