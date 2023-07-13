package v2022_11_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2022-11-01/attacheddatanetwork"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2022-11-01/attacheddatanetworks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2022-11-01/datanetwork"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2022-11-01/datanetworks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2022-11-01/mobilenetwork"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2022-11-01/mobilenetworks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2022-11-01/packetcorecontrolplane"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2022-11-01/packetcorecontrolplanecollectdiagnosticspackage"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2022-11-01/packetcorecontrolplanereinstall"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2022-11-01/packetcorecontrolplanerollback"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2022-11-01/packetcorecontrolplanes"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2022-11-01/packetcorecontrolplaneversion"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2022-11-01/packetcoredataplane"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2022-11-01/packetcoredataplanes"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2022-11-01/service"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2022-11-01/services"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2022-11-01/sim"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2022-11-01/simgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2022-11-01/simgroups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2022-11-01/simpolicies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2022-11-01/simpolicy"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2022-11-01/sims"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2022-11-01/site"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2022-11-01/sites"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2022-11-01/slice"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2022-11-01/slices"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AttachedDataNetwork                             *attacheddatanetwork.AttachedDataNetworkClient
	AttachedDataNetworks                            *attacheddatanetworks.AttachedDataNetworksClient
	DataNetwork                                     *datanetwork.DataNetworkClient
	DataNetworks                                    *datanetworks.DataNetworksClient
	MobileNetwork                                   *mobilenetwork.MobileNetworkClient
	MobileNetworks                                  *mobilenetworks.MobileNetworksClient
	PacketCoreControlPlane                          *packetcorecontrolplane.PacketCoreControlPlaneClient
	PacketCoreControlPlaneCollectDiagnosticsPackage *packetcorecontrolplanecollectdiagnosticspackage.PacketCoreControlPlaneCollectDiagnosticsPackageClient
	PacketCoreControlPlaneReinstall                 *packetcorecontrolplanereinstall.PacketCoreControlPlaneReinstallClient
	PacketCoreControlPlaneRollback                  *packetcorecontrolplanerollback.PacketCoreControlPlaneRollbackClient
	PacketCoreControlPlaneVersion                   *packetcorecontrolplaneversion.PacketCoreControlPlaneVersionClient
	PacketCoreControlPlanes                         *packetcorecontrolplanes.PacketCoreControlPlanesClient
	PacketCoreDataPlane                             *packetcoredataplane.PacketCoreDataPlaneClient
	PacketCoreDataPlanes                            *packetcoredataplanes.PacketCoreDataPlanesClient
	SIM                                             *sim.SIMClient
	SIMGroup                                        *simgroup.SIMGroupClient
	SIMGroups                                       *simgroups.SIMGroupsClient
	SIMPolicies                                     *simpolicies.SIMPoliciesClient
	SIMPolicy                                       *simpolicy.SIMPolicyClient
	SIMs                                            *sims.SIMsClient
	Service                                         *service.ServiceClient
	Services                                        *services.ServicesClient
	Site                                            *site.SiteClient
	Sites                                           *sites.SitesClient
	Slice                                           *slice.SliceClient
	Slices                                          *slices.SlicesClient
}

func NewClientWithBaseURI(api environments.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	attachedDataNetworkClient, err := attacheddatanetwork.NewAttachedDataNetworkClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building AttachedDataNetwork client: %+v", err)
	}
	configureFunc(attachedDataNetworkClient.Client)

	attachedDataNetworksClient, err := attacheddatanetworks.NewAttachedDataNetworksClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building AttachedDataNetworks client: %+v", err)
	}
	configureFunc(attachedDataNetworksClient.Client)

	dataNetworkClient, err := datanetwork.NewDataNetworkClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building DataNetwork client: %+v", err)
	}
	configureFunc(dataNetworkClient.Client)

	dataNetworksClient, err := datanetworks.NewDataNetworksClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building DataNetworks client: %+v", err)
	}
	configureFunc(dataNetworksClient.Client)

	mobileNetworkClient, err := mobilenetwork.NewMobileNetworkClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building MobileNetwork client: %+v", err)
	}
	configureFunc(mobileNetworkClient.Client)

	mobileNetworksClient, err := mobilenetworks.NewMobileNetworksClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building MobileNetworks client: %+v", err)
	}
	configureFunc(mobileNetworksClient.Client)

	packetCoreControlPlaneClient, err := packetcorecontrolplane.NewPacketCoreControlPlaneClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building PacketCoreControlPlane client: %+v", err)
	}
	configureFunc(packetCoreControlPlaneClient.Client)

	packetCoreControlPlaneCollectDiagnosticsPackageClient, err := packetcorecontrolplanecollectdiagnosticspackage.NewPacketCoreControlPlaneCollectDiagnosticsPackageClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building PacketCoreControlPlaneCollectDiagnosticsPackage client: %+v", err)
	}
	configureFunc(packetCoreControlPlaneCollectDiagnosticsPackageClient.Client)

	packetCoreControlPlaneReinstallClient, err := packetcorecontrolplanereinstall.NewPacketCoreControlPlaneReinstallClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building PacketCoreControlPlaneReinstall client: %+v", err)
	}
	configureFunc(packetCoreControlPlaneReinstallClient.Client)

	packetCoreControlPlaneRollbackClient, err := packetcorecontrolplanerollback.NewPacketCoreControlPlaneRollbackClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building PacketCoreControlPlaneRollback client: %+v", err)
	}
	configureFunc(packetCoreControlPlaneRollbackClient.Client)

	packetCoreControlPlaneVersionClient, err := packetcorecontrolplaneversion.NewPacketCoreControlPlaneVersionClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building PacketCoreControlPlaneVersion client: %+v", err)
	}
	configureFunc(packetCoreControlPlaneVersionClient.Client)

	packetCoreControlPlanesClient, err := packetcorecontrolplanes.NewPacketCoreControlPlanesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building PacketCoreControlPlanes client: %+v", err)
	}
	configureFunc(packetCoreControlPlanesClient.Client)

	packetCoreDataPlaneClient, err := packetcoredataplane.NewPacketCoreDataPlaneClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building PacketCoreDataPlane client: %+v", err)
	}
	configureFunc(packetCoreDataPlaneClient.Client)

	packetCoreDataPlanesClient, err := packetcoredataplanes.NewPacketCoreDataPlanesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building PacketCoreDataPlanes client: %+v", err)
	}
	configureFunc(packetCoreDataPlanesClient.Client)

	sIMClient, err := sim.NewSIMClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building SIM client: %+v", err)
	}
	configureFunc(sIMClient.Client)

	sIMGroupClient, err := simgroup.NewSIMGroupClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building SIMGroup client: %+v", err)
	}
	configureFunc(sIMGroupClient.Client)

	sIMGroupsClient, err := simgroups.NewSIMGroupsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building SIMGroups client: %+v", err)
	}
	configureFunc(sIMGroupsClient.Client)

	sIMPoliciesClient, err := simpolicies.NewSIMPoliciesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building SIMPolicies client: %+v", err)
	}
	configureFunc(sIMPoliciesClient.Client)

	sIMPolicyClient, err := simpolicy.NewSIMPolicyClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building SIMPolicy client: %+v", err)
	}
	configureFunc(sIMPolicyClient.Client)

	sIMsClient, err := sims.NewSIMsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building SIMs client: %+v", err)
	}
	configureFunc(sIMsClient.Client)

	serviceClient, err := service.NewServiceClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Service client: %+v", err)
	}
	configureFunc(serviceClient.Client)

	servicesClient, err := services.NewServicesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Services client: %+v", err)
	}
	configureFunc(servicesClient.Client)

	siteClient, err := site.NewSiteClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Site client: %+v", err)
	}
	configureFunc(siteClient.Client)

	sitesClient, err := sites.NewSitesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Sites client: %+v", err)
	}
	configureFunc(sitesClient.Client)

	sliceClient, err := slice.NewSliceClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Slice client: %+v", err)
	}
	configureFunc(sliceClient.Client)

	slicesClient, err := slices.NewSlicesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Slices client: %+v", err)
	}
	configureFunc(slicesClient.Client)

	return &Client{
		AttachedDataNetwork:    attachedDataNetworkClient,
		AttachedDataNetworks:   attachedDataNetworksClient,
		DataNetwork:            dataNetworkClient,
		DataNetworks:           dataNetworksClient,
		MobileNetwork:          mobileNetworkClient,
		MobileNetworks:         mobileNetworksClient,
		PacketCoreControlPlane: packetCoreControlPlaneClient,
		PacketCoreControlPlaneCollectDiagnosticsPackage: packetCoreControlPlaneCollectDiagnosticsPackageClient,
		PacketCoreControlPlaneReinstall:                 packetCoreControlPlaneReinstallClient,
		PacketCoreControlPlaneRollback:                  packetCoreControlPlaneRollbackClient,
		PacketCoreControlPlaneVersion:                   packetCoreControlPlaneVersionClient,
		PacketCoreControlPlanes:                         packetCoreControlPlanesClient,
		PacketCoreDataPlane:                             packetCoreDataPlaneClient,
		PacketCoreDataPlanes:                            packetCoreDataPlanesClient,
		SIM:                                             sIMClient,
		SIMGroup:                                        sIMGroupClient,
		SIMGroups:                                       sIMGroupsClient,
		SIMPolicies:                                     sIMPoliciesClient,
		SIMPolicy:                                       sIMPolicyClient,
		SIMs:                                            sIMsClient,
		Service:                                         serviceClient,
		Services:                                        servicesClient,
		Site:                                            siteClient,
		Sites:                                           sitesClient,
		Slice:                                           sliceClient,
		Slices:                                          slicesClient,
	}, nil
}
