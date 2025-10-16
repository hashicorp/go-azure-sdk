package namespacediscoveredassets

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NamespaceDiscoveredManagementGroup struct {
	Actions                      *[]NamespaceDiscoveredManagementAction `json:"actions,omitempty"`
	DataSource                   *string                                `json:"dataSource,omitempty"`
	DefaultTimeoutInSeconds      *int64                                 `json:"defaultTimeoutInSeconds,omitempty"`
	DefaultTopic                 *string                                `json:"defaultTopic,omitempty"`
	LastUpdatedOn                *string                                `json:"lastUpdatedOn,omitempty"`
	ManagementGroupConfiguration *string                                `json:"managementGroupConfiguration,omitempty"`
	Name                         string                                 `json:"name"`
	TypeRef                      *string                                `json:"typeRef,omitempty"`
}

func (o *NamespaceDiscoveredManagementGroup) GetLastUpdatedOnAsTime() (*time.Time, error) {
	if o.LastUpdatedOn == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastUpdatedOn, "2006-01-02T15:04:05Z07:00")
}

func (o *NamespaceDiscoveredManagementGroup) SetLastUpdatedOnAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastUpdatedOn = &formatted
}
