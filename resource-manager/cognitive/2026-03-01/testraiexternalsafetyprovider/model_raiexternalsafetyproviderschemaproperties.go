package testraiexternalsafetyprovider

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RaiExternalSafetyProviderSchemaProperties struct {
	CreatedAt       *string `json:"createdAt,omitempty"`
	KeyVaultUri     *string `json:"keyVaultUri,omitempty"`
	LastModifiedAt  *string `json:"lastModifiedAt,omitempty"`
	ManagedIdentity *string `json:"managedIdentity,omitempty"`
	Mode            *string `json:"mode,omitempty"`
	ProviderId      *string `json:"providerId,omitempty"`
	ProviderName    *string `json:"providerName,omitempty"`
	SecretName      *string `json:"secretName,omitempty"`
	Url             *string `json:"url,omitempty"`
}

func (o *RaiExternalSafetyProviderSchemaProperties) GetCreatedAtAsTime() (*time.Time, error) {
	if o.CreatedAt == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CreatedAt, "2006-01-02T15:04:05Z07:00")
}

func (o *RaiExternalSafetyProviderSchemaProperties) SetCreatedAtAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.CreatedAt = &formatted
}

func (o *RaiExternalSafetyProviderSchemaProperties) GetLastModifiedAtAsTime() (*time.Time, error) {
	if o.LastModifiedAt == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastModifiedAt, "2006-01-02T15:04:05Z07:00")
}

func (o *RaiExternalSafetyProviderSchemaProperties) SetLastModifiedAtAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastModifiedAt = &formatted
}
