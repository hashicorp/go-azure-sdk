package dscnode

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DscNode struct {
	AccountId         *string                                       `json:"accountId,omitempty"`
	Etag              *string                                       `json:"etag,omitempty"`
	ExtensionHandler  *[]DscNodeExtensionHandlerAssociationProperty `json:"extensionHandler,omitempty"`
	IP                *string                                       `json:"ip,omitempty"`
	Id                *string                                       `json:"id,omitempty"`
	LastSeen          *string                                       `json:"lastSeen,omitempty"`
	Name              *string                                       `json:"name,omitempty"`
	NodeConfiguration *DscNodeConfigurationAssociationProperty      `json:"nodeConfiguration,omitempty"`
	NodeId            *string                                       `json:"nodeId,omitempty"`
	RegistrationTime  *string                                       `json:"registrationTime,omitempty"`
	Status            *string                                       `json:"status,omitempty"`
	Type              *string                                       `json:"type,omitempty"`
}

func (o *DscNode) GetLastSeenAsTime() (*time.Time, error) {
	if o.LastSeen == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastSeen, "2006-01-02T15:04:05Z07:00")
}

func (o *DscNode) SetLastSeenAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastSeen = &formatted
}

func (o *DscNode) GetRegistrationTimeAsTime() (*time.Time, error) {
	if o.RegistrationTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.RegistrationTime, "2006-01-02T15:04:05Z07:00")
}

func (o *DscNode) SetRegistrationTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.RegistrationTime = &formatted
}
