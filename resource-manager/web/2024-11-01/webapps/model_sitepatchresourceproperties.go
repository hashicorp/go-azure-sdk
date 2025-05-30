package webapps

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SitePatchResourceProperties struct {
	AvailabilityState           *SiteAvailabilityState     `json:"availabilityState,omitempty"`
	ClientAffinityEnabled       *bool                      `json:"clientAffinityEnabled,omitempty"`
	ClientAffinityProxyEnabled  *bool                      `json:"clientAffinityProxyEnabled,omitempty"`
	ClientCertEnabled           *bool                      `json:"clientCertEnabled,omitempty"`
	ClientCertExclusionPaths    *string                    `json:"clientCertExclusionPaths,omitempty"`
	ClientCertMode              *ClientCertMode            `json:"clientCertMode,omitempty"`
	CloningInfo                 *CloningInfo               `json:"cloningInfo,omitempty"`
	ContainerSize               *int64                     `json:"containerSize,omitempty"`
	CustomDomainVerificationId  *string                    `json:"customDomainVerificationId,omitempty"`
	DailyMemoryTimeQuota        *int64                     `json:"dailyMemoryTimeQuota,omitempty"`
	DefaultHostName             *string                    `json:"defaultHostName,omitempty"`
	DnsConfiguration            *SiteDnsConfig             `json:"dnsConfiguration,omitempty"`
	Enabled                     *bool                      `json:"enabled,omitempty"`
	EnabledHostNames            *[]string                  `json:"enabledHostNames,omitempty"`
	HTTPSOnly                   *bool                      `json:"httpsOnly,omitempty"`
	HostNameSslStates           *[]HostNameSslState        `json:"hostNameSslStates,omitempty"`
	HostNames                   *[]string                  `json:"hostNames,omitempty"`
	HostNamesDisabled           *bool                      `json:"hostNamesDisabled,omitempty"`
	HostingEnvironmentProfile   *HostingEnvironmentProfile `json:"hostingEnvironmentProfile,omitempty"`
	HyperV                      *bool                      `json:"hyperV,omitempty"`
	InProgressOperationId       *string                    `json:"inProgressOperationId,omitempty"`
	IsDefaultContainer          *bool                      `json:"isDefaultContainer,omitempty"`
	IsXenon                     *bool                      `json:"isXenon,omitempty"`
	KeyVaultReferenceIdentity   *string                    `json:"keyVaultReferenceIdentity,omitempty"`
	LastModifiedTimeUtc         *string                    `json:"lastModifiedTimeUtc,omitempty"`
	MaxNumberOfWorkers          *int64                     `json:"maxNumberOfWorkers,omitempty"`
	OutboundIPAddresses         *string                    `json:"outboundIpAddresses,omitempty"`
	PossibleOutboundIPAddresses *string                    `json:"possibleOutboundIpAddresses,omitempty"`
	RedundancyMode              *RedundancyMode            `json:"redundancyMode,omitempty"`
	RepositorySiteName          *string                    `json:"repositorySiteName,omitempty"`
	Reserved                    *bool                      `json:"reserved,omitempty"`
	ResourceGroup               *string                    `json:"resourceGroup,omitempty"`
	ScmSiteAlsoStopped          *bool                      `json:"scmSiteAlsoStopped,omitempty"`
	ServerFarmId                *string                    `json:"serverFarmId,omitempty"`
	SiteConfig                  *SiteConfig                `json:"siteConfig,omitempty"`
	SlotSwapStatus              *SlotSwapStatus            `json:"slotSwapStatus,omitempty"`
	State                       *string                    `json:"state,omitempty"`
	StorageAccountRequired      *bool                      `json:"storageAccountRequired,omitempty"`
	SuspendedTill               *string                    `json:"suspendedTill,omitempty"`
	TargetSwapSlot              *string                    `json:"targetSwapSlot,omitempty"`
	TrafficManagerHostNames     *[]string                  `json:"trafficManagerHostNames,omitempty"`
	UsageState                  *UsageState                `json:"usageState,omitempty"`
	VirtualNetworkSubnetId      *string                    `json:"virtualNetworkSubnetId,omitempty"`
}

func (o *SitePatchResourceProperties) GetLastModifiedTimeUtcAsTime() (*time.Time, error) {
	if o.LastModifiedTimeUtc == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastModifiedTimeUtc, "2006-01-02T15:04:05Z07:00")
}

func (o *SitePatchResourceProperties) SetLastModifiedTimeUtcAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastModifiedTimeUtc = &formatted
}

func (o *SitePatchResourceProperties) GetSuspendedTillAsTime() (*time.Time, error) {
	if o.SuspendedTill == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.SuspendedTill, "2006-01-02T15:04:05Z07:00")
}

func (o *SitePatchResourceProperties) SetSuspendedTillAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.SuspendedTill = &formatted
}
