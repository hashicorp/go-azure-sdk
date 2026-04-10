package provider

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LinuxJavaContainerSettings struct {
	EndOfLifeDate *string `json:"endOfLifeDate,omitempty"`
	IsAutoUpdate  *bool   `json:"isAutoUpdate,omitempty"`
	IsDeprecated  *bool   `json:"isDeprecated,omitempty"`
	IsEarlyAccess *bool   `json:"isEarlyAccess,omitempty"`
	IsHidden      *bool   `json:"isHidden,omitempty"`
	IsPreview     *bool   `json:"isPreview,omitempty"`
	Java11Runtime *string `json:"java11Runtime,omitempty"`
	Java8Runtime  *string `json:"java8Runtime,omitempty"`
}

func (o *LinuxJavaContainerSettings) GetEndOfLifeDateAsTime() (*time.Time, error) {
	if o.EndOfLifeDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EndOfLifeDate, "2006-01-02T15:04:05Z07:00")
}

func (o *LinuxJavaContainerSettings) SetEndOfLifeDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EndOfLifeDate = &formatted
}
