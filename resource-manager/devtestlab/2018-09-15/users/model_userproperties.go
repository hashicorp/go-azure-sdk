package users

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserProperties struct {
	CreatedDate       *string          `json:"createdDate,omitempty"`
	Identity          *UserIdentity    `json:"identity,omitempty"`
	ProvisioningState *string          `json:"provisioningState,omitempty"`
	SecretStore       *UserSecretStore `json:"secretStore,omitempty"`
	UniqueIdentifier  *string          `json:"uniqueIdentifier,omitempty"`
}

func (o *UserProperties) GetCreatedDateAsTime() (*time.Time, error) {
	if o.CreatedDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CreatedDate, "2006-01-02T15:04:05Z07:00")
}
