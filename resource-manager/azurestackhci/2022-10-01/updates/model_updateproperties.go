package updates

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateProperties struct {
	AdditionalProperties  *string                `json:"additionalProperties,omitempty"`
	AvailabilityType      *AvailabilityType      `json:"availabilityType,omitempty"`
	Description           *string                `json:"description,omitempty"`
	DisplayName           *string                `json:"displayName,omitempty"`
	InstalledDate         *string                `json:"installedDate,omitempty"`
	PackagePath           *string                `json:"packagePath,omitempty"`
	PackageSizeInMb       *float64               `json:"packageSizeInMb,omitempty"`
	PackageType           *string                `json:"packageType,omitempty"`
	Prerequisites         *[]UpdatePrerequisite  `json:"prerequisites,omitempty"`
	ProvisioningState     *ProvisioningState     `json:"provisioningState,omitempty"`
	Publisher             *string                `json:"publisher,omitempty"`
	ReleaseLink           *string                `json:"releaseLink,omitempty"`
	State                 *State                 `json:"state,omitempty"`
	UpdateStateProperties *UpdateStateProperties `json:"updateStateProperties"`
	Version               *string                `json:"version,omitempty"`
}

func (o *UpdateProperties) GetInstalledDateAsTime() (*time.Time, error) {
	if o.InstalledDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.InstalledDate, "2006-01-02T15:04:05Z07:00")
}

func (o *UpdateProperties) SetInstalledDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.InstalledDate = &formatted
}
