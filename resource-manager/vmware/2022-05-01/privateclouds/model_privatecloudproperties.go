package privateclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateCloudProperties struct {
	Availability                 *AvailabilityProperties        `json:"availability"`
	Circuit                      *Circuit                       `json:"circuit"`
	Encryption                   *Encryption                    `json:"encryption"`
	Endpoints                    *Endpoints                     `json:"endpoints"`
	ExternalCloudLinks           *[]string                      `json:"externalCloudLinks,omitempty"`
	IdentitySources              *[]IdentitySource              `json:"identitySources,omitempty"`
	Internet                     *InternetEnum                  `json:"internet,omitempty"`
	ManagementCluster            CommonClusterProperties        `json:"managementCluster"`
	ManagementNetwork            *string                        `json:"managementNetwork,omitempty"`
	NetworkBlock                 string                         `json:"networkBlock"`
	NsxPublicIPQuotaRaised       *NsxPublicIPQuotaRaisedEnum    `json:"nsxPublicIpQuotaRaised,omitempty"`
	NsxtCertificateThumbprint    *string                        `json:"nsxtCertificateThumbprint,omitempty"`
	NsxtPassword                 *string                        `json:"nsxtPassword,omitempty"`
	ProvisioningNetwork          *string                        `json:"provisioningNetwork,omitempty"`
	ProvisioningState            *PrivateCloudProvisioningState `json:"provisioningState,omitempty"`
	SecondaryCircuit             *Circuit                       `json:"secondaryCircuit"`
	VcenterCertificateThumbprint *string                        `json:"vcenterCertificateThumbprint,omitempty"`
	VcenterPassword              *string                        `json:"vcenterPassword,omitempty"`
	VmotionNetwork               *string                        `json:"vmotionNetwork,omitempty"`
}
