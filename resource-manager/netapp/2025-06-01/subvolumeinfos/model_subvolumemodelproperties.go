package subvolumeinfos

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubvolumeModelProperties struct {
	AccessedTimeStamp *string `json:"accessedTimeStamp,omitempty"`
	BytesUsed         *int64  `json:"bytesUsed,omitempty"`
	ChangedTimeStamp  *string `json:"changedTimeStamp,omitempty"`
	CreationTimeStamp *string `json:"creationTimeStamp,omitempty"`
	ModifiedTimeStamp *string `json:"modifiedTimeStamp,omitempty"`
	ParentPath        *string `json:"parentPath,omitempty"`
	Path              *string `json:"path,omitempty"`
	Permissions       *string `json:"permissions,omitempty"`
	ProvisioningState *string `json:"provisioningState,omitempty"`
	Size              *int64  `json:"size,omitempty"`
}

func (o *SubvolumeModelProperties) GetAccessedTimeStampAsTime() (*time.Time, error) {
	if o.AccessedTimeStamp == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.AccessedTimeStamp, "2006-01-02T15:04:05Z07:00")
}

func (o *SubvolumeModelProperties) SetAccessedTimeStampAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.AccessedTimeStamp = &formatted
}

func (o *SubvolumeModelProperties) GetChangedTimeStampAsTime() (*time.Time, error) {
	if o.ChangedTimeStamp == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ChangedTimeStamp, "2006-01-02T15:04:05Z07:00")
}

func (o *SubvolumeModelProperties) SetChangedTimeStampAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ChangedTimeStamp = &formatted
}

func (o *SubvolumeModelProperties) GetCreationTimeStampAsTime() (*time.Time, error) {
	if o.CreationTimeStamp == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CreationTimeStamp, "2006-01-02T15:04:05Z07:00")
}

func (o *SubvolumeModelProperties) SetCreationTimeStampAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.CreationTimeStamp = &formatted
}

func (o *SubvolumeModelProperties) GetModifiedTimeStampAsTime() (*time.Time, error) {
	if o.ModifiedTimeStamp == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ModifiedTimeStamp, "2006-01-02T15:04:05Z07:00")
}

func (o *SubvolumeModelProperties) SetModifiedTimeStampAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ModifiedTimeStamp = &formatted
}
