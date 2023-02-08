package dscnodeconfiguration

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DscNodeConfiguration struct {
	Configuration    *DscConfigurationAssociationProperty `json:"configuration,omitempty"`
	CreationTime     *string                              `json:"creationTime,omitempty"`
	Id               *string                              `json:"id,omitempty"`
	LastModifiedTime *string                              `json:"lastModifiedTime,omitempty"`
	Name             *string                              `json:"name,omitempty"`
	Type             *string                              `json:"type,omitempty"`
}

func (o *DscNodeConfiguration) GetCreationTimeAsTime() (*time.Time, error) {
	if o.CreationTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CreationTime, "2006-01-02T15:04:05Z07:00")
}

func (o *DscNodeConfiguration) SetCreationTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.CreationTime = &formatted
}

func (o *DscNodeConfiguration) GetLastModifiedTimeAsTime() (*time.Time, error) {
	if o.LastModifiedTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastModifiedTime, "2006-01-02T15:04:05Z07:00")
}

func (o *DscNodeConfiguration) SetLastModifiedTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastModifiedTime = &formatted
}
