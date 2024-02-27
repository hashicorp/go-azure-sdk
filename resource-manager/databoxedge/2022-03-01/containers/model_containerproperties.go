package containers

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContainerProperties struct {
	ContainerStatus *ContainerStatus         `json:"containerStatus,omitempty"`
	CreatedDateTime *string                  `json:"createdDateTime,omitempty"`
	DataFormat      AzureContainerDataFormat `json:"dataFormat"`
	RefreshDetails  *RefreshDetails          `json:"refreshDetails,omitempty"`
}

func (o *ContainerProperties) GetCreatedDateTimeAsTime() (*time.Time, error) {
	if o.CreatedDateTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CreatedDateTime, "2006-01-02T15:04:05Z07:00")
}
