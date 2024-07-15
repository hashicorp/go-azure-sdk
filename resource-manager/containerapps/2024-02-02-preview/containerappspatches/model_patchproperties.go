package containerappspatches

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PatchProperties struct {
	CreatedAt            *string           `json:"createdAt,omitempty"`
	LastModifiedAt       *string           `json:"lastModifiedAt,omitempty"`
	PatchApplyStatus     *PatchApplyStatus `json:"patchApplyStatus,omitempty"`
	PatchDetails         *[]PatchDetails   `json:"patchDetails,omitempty"`
	TargetContainerAppId *string           `json:"targetContainerAppId,omitempty"`
	TargetEnvironmentId  *string           `json:"targetEnvironmentId,omitempty"`
	TargetRevisionId     *string           `json:"targetRevisionId,omitempty"`
}

func (o *PatchProperties) GetCreatedAtAsTime() (*time.Time, error) {
	if o.CreatedAt == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CreatedAt, "2006-01-02T15:04:05Z07:00")
}

func (o *PatchProperties) SetCreatedAtAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.CreatedAt = &formatted
}

func (o *PatchProperties) GetLastModifiedAtAsTime() (*time.Time, error) {
	if o.LastModifiedAt == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastModifiedAt, "2006-01-02T15:04:05Z07:00")
}

func (o *PatchProperties) SetLastModifiedAtAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastModifiedAt = &formatted
}
