package activesessionhostconfiguration

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActiveSessionHostConfigurationProperties struct {
	AvailabilityZones            *[]int64                       `json:"availabilityZones,omitempty"`
	BootDiagnosticsInfo          *BootDiagnosticsInfoProperties `json:"bootDiagnosticsInfo,omitempty"`
	CustomConfigurationScriptURL *string                        `json:"customConfigurationScriptUrl,omitempty"`
	DiskInfo                     DiskInfoProperties             `json:"diskInfo"`
	DomainInfo                   DomainInfoProperties           `json:"domainInfo"`
	FriendlyName                 *string                        `json:"friendlyName,omitempty"`
	ImageInfo                    ImageInfoProperties            `json:"imageInfo"`
	NetworkInfo                  NetworkInfoProperties          `json:"networkInfo"`
	SecurityInfo                 *SecurityInfoProperties        `json:"securityInfo,omitempty"`
	VMAdminCredentials           KeyVaultCredentialsProperties  `json:"vmAdminCredentials"`
	VMLocation                   *string                        `json:"vmLocation,omitempty"`
	VMNamePrefix                 string                         `json:"vmNamePrefix"`
	VMResourceGroup              *string                        `json:"vmResourceGroup,omitempty"`
	VMSizeId                     string                         `json:"vmSizeId"`
	VMTags                       *map[string]string             `json:"vmTags,omitempty"`
	Version                      *string                        `json:"version,omitempty"`
}

func (o *ActiveSessionHostConfigurationProperties) GetVersionAsTime() (*time.Time, error) {
	if o.Version == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.Version, "2006-01-02T15:04:05Z07:00")
}

func (o *ActiveSessionHostConfigurationProperties) SetVersionAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.Version = &formatted
}
