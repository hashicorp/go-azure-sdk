package portalrevision

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PortalRevisionContractProperties struct {
	CreatedDateTime   *string               `json:"createdDateTime,omitempty"`
	Description       *string               `json:"description,omitempty"`
	IsCurrent         *bool                 `json:"isCurrent,omitempty"`
	ProvisioningState *string               `json:"provisioningState,omitempty"`
	Status            *PortalRevisionStatus `json:"status,omitempty"`
	StatusDetails     *string               `json:"statusDetails,omitempty"`
	UpdatedDateTime   *string               `json:"updatedDateTime,omitempty"`
}

func (o *PortalRevisionContractProperties) GetCreatedDateTimeAsTime() (*time.Time, error) {
	if o.CreatedDateTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CreatedDateTime, "2006-01-02T15:04:05Z07:00")
}

func (o *PortalRevisionContractProperties) SetCreatedDateTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.CreatedDateTime = &formatted
}

func (o *PortalRevisionContractProperties) GetUpdatedDateTimeAsTime() (*time.Time, error) {
	if o.UpdatedDateTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.UpdatedDateTime, "2006-01-02T15:04:05Z07:00")
}

func (o *PortalRevisionContractProperties) SetUpdatedDateTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.UpdatedDateTime = &formatted
}
