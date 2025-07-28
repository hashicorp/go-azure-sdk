package sourcecontrols

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServicePrincipal struct {
	AppId               *string `json:"appId,omitempty"`
	CredentialsExpireOn *string `json:"credentialsExpireOn,omitempty"`
	Id                  *string `json:"id,omitempty"`
	TenantId            *string `json:"tenantId,omitempty"`
}

func (o *ServicePrincipal) GetCredentialsExpireOnAsTime() (*time.Time, error) {
	if o.CredentialsExpireOn == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CredentialsExpireOn, "2006-01-02T15:04:05Z07:00")
}

func (o *ServicePrincipal) SetCredentialsExpireOnAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.CredentialsExpireOn = &formatted
}
