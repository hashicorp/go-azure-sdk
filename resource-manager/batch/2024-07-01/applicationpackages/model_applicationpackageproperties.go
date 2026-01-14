package applicationpackages

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationPackageProperties struct {
	Format             *string       `json:"format,omitempty"`
	LastActivationTime *string       `json:"lastActivationTime,omitempty"`
	State              *PackageState `json:"state,omitempty"`
	StorageURL         *string       `json:"storageUrl,omitempty"`
	StorageURLExpiry   *string       `json:"storageUrlExpiry,omitempty"`
}

func (o *ApplicationPackageProperties) GetLastActivationTimeAsTime() (*time.Time, error) {
	if o.LastActivationTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastActivationTime, "2006-01-02T15:04:05Z07:00")
}

func (o *ApplicationPackageProperties) SetLastActivationTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastActivationTime = &formatted
}

func (o *ApplicationPackageProperties) GetStorageURLExpiryAsTime() (*time.Time, error) {
	if o.StorageURLExpiry == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.StorageURLExpiry, "2006-01-02T15:04:05Z07:00")
}

func (o *ApplicationPackageProperties) SetStorageURLExpiryAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.StorageURLExpiry = &formatted
}
