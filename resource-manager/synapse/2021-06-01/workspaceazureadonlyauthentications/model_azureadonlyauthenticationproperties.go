package workspaceazureadonlyauthentications

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureADOnlyAuthenticationProperties struct {
	AzureADOnlyAuthentication bool        `json:"azureADOnlyAuthentication"`
	CreationDate              *string     `json:"creationDate,omitempty"`
	State                     *StateValue `json:"state,omitempty"`
}

func (o *AzureADOnlyAuthenticationProperties) GetCreationDateAsTime() (*time.Time, error) {
	if o.CreationDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CreationDate, "2006-01-02T15:04:05Z07:00")
}
