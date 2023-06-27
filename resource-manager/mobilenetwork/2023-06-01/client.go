package v2023_06_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2023-06-01/attacheddatanetwork"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2023-06-01/attacheddatanetworks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2023-06-01/datanetwork"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2023-06-01/datanetworks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2023-06-01/diagnosticspackages"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2023-06-01/mobilenetwork"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2023-06-01/mobilenetworks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2023-06-01/packetcaptures"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2023-06-01/packetcorecontrolplane"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2023-06-01/packetcorecontrolplanecollectdiagnosticspackage"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2023-06-01/packetcorecontrolplanereinstall"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2023-06-01/packetcorecontrolplanerollback"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2023-06-01/packetcorecontrolplanes"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2023-06-01/packetcorecontrolplaneversion"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2023-06-01/packetcoredataplane"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2023-06-01/packetcoredataplanes"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2023-06-01/service"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2023-06-01/services"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2023-06-01/sim"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2023-06-01/simgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2023-06-01/simgroups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2023-06-01/simpolicies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2023-06-01/simpolicy"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2023-06-01/sims"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2023-06-01/site"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2023-06-01/sites"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2023-06-01/slice"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2023-06-01/slices"
)

type Client struct {
	AttachedDataNetwork                             *attacheddatanetwork.AttachedDataNetworkClient
	AttachedDataNetworks                            *attacheddatanetworks.AttachedDataNetworksClient
	DataNetwork                                     *datanetwork.DataNetworkClient
	DataNetworks                                    *datanetworks.DataNetworksClient
	DiagnosticsPackages                             *diagnosticspackages.DiagnosticsPackagesClient
	MobileNetwork                                   *mobilenetwork.MobileNetworkClient
	MobileNetworks                                  *mobilenetworks.MobileNetworksClient
	PacketCaptures                                  *packetcaptures.PacketCapturesClient
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

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	attachedDataNetworkClient := attacheddatanetwork.NewAttachedDataNetworkClientWithBaseURI(endpoint)
	configureAuthFunc(&attachedDataNetworkClient.Client)

	attachedDataNetworksClient := attacheddatanetworks.NewAttachedDataNetworksClientWithBaseURI(endpoint)
	configureAuthFunc(&attachedDataNetworksClient.Client)

	dataNetworkClient := datanetwork.NewDataNetworkClientWithBaseURI(endpoint)
	configureAuthFunc(&dataNetworkClient.Client)

	dataNetworksClient := datanetworks.NewDataNetworksClientWithBaseURI(endpoint)
	configureAuthFunc(&dataNetworksClient.Client)

	diagnosticsPackagesClient := diagnosticspackages.NewDiagnosticsPackagesClientWithBaseURI(endpoint)
	configureAuthFunc(&diagnosticsPackagesClient.Client)

	mobileNetworkClient := mobilenetwork.NewMobileNetworkClientWithBaseURI(endpoint)
	configureAuthFunc(&mobileNetworkClient.Client)

	mobileNetworksClient := mobilenetworks.NewMobileNetworksClientWithBaseURI(endpoint)
	configureAuthFunc(&mobileNetworksClient.Client)

	packetCapturesClient := packetcaptures.NewPacketCapturesClientWithBaseURI(endpoint)
	configureAuthFunc(&packetCapturesClient.Client)

	packetCoreControlPlaneClient := packetcorecontrolplane.NewPacketCoreControlPlaneClientWithBaseURI(endpoint)
	configureAuthFunc(&packetCoreControlPlaneClient.Client)

	packetCoreControlPlaneCollectDiagnosticsPackageClient := packetcorecontrolplanecollectdiagnosticspackage.NewPacketCoreControlPlaneCollectDiagnosticsPackageClientWithBaseURI(endpoint)
	configureAuthFunc(&packetCoreControlPlaneCollectDiagnosticsPackageClient.Client)

	packetCoreControlPlaneReinstallClient := packetcorecontrolplanereinstall.NewPacketCoreControlPlaneReinstallClientWithBaseURI(endpoint)
	configureAuthFunc(&packetCoreControlPlaneReinstallClient.Client)

	packetCoreControlPlaneRollbackClient := packetcorecontrolplanerollback.NewPacketCoreControlPlaneRollbackClientWithBaseURI(endpoint)
	configureAuthFunc(&packetCoreControlPlaneRollbackClient.Client)

	packetCoreControlPlaneVersionClient := packetcorecontrolplaneversion.NewPacketCoreControlPlaneVersionClientWithBaseURI(endpoint)
	configureAuthFunc(&packetCoreControlPlaneVersionClient.Client)

	packetCoreControlPlanesClient := packetcorecontrolplanes.NewPacketCoreControlPlanesClientWithBaseURI(endpoint)
	configureAuthFunc(&packetCoreControlPlanesClient.Client)

	packetCoreDataPlaneClient := packetcoredataplane.NewPacketCoreDataPlaneClientWithBaseURI(endpoint)
	configureAuthFunc(&packetCoreDataPlaneClient.Client)

	packetCoreDataPlanesClient := packetcoredataplanes.NewPacketCoreDataPlanesClientWithBaseURI(endpoint)
	configureAuthFunc(&packetCoreDataPlanesClient.Client)

	sIMClient := sim.NewSIMClientWithBaseURI(endpoint)
	configureAuthFunc(&sIMClient.Client)

	sIMGroupClient := simgroup.NewSIMGroupClientWithBaseURI(endpoint)
	configureAuthFunc(&sIMGroupClient.Client)

	sIMGroupsClient := simgroups.NewSIMGroupsClientWithBaseURI(endpoint)
	configureAuthFunc(&sIMGroupsClient.Client)

	sIMPoliciesClient := simpolicies.NewSIMPoliciesClientWithBaseURI(endpoint)
	configureAuthFunc(&sIMPoliciesClient.Client)

	sIMPolicyClient := simpolicy.NewSIMPolicyClientWithBaseURI(endpoint)
	configureAuthFunc(&sIMPolicyClient.Client)

	sIMsClient := sims.NewSIMsClientWithBaseURI(endpoint)
	configureAuthFunc(&sIMsClient.Client)

	serviceClient := service.NewServiceClientWithBaseURI(endpoint)
	configureAuthFunc(&serviceClient.Client)

	servicesClient := services.NewServicesClientWithBaseURI(endpoint)
	configureAuthFunc(&servicesClient.Client)

	siteClient := site.NewSiteClientWithBaseURI(endpoint)
	configureAuthFunc(&siteClient.Client)

	sitesClient := sites.NewSitesClientWithBaseURI(endpoint)
	configureAuthFunc(&sitesClient.Client)

	sliceClient := slice.NewSliceClientWithBaseURI(endpoint)
	configureAuthFunc(&sliceClient.Client)

	slicesClient := slices.NewSlicesClientWithBaseURI(endpoint)
	configureAuthFunc(&slicesClient.Client)

	return Client{
		AttachedDataNetwork:    &attachedDataNetworkClient,
		AttachedDataNetworks:   &attachedDataNetworksClient,
		DataNetwork:            &dataNetworkClient,
		DataNetworks:           &dataNetworksClient,
		DiagnosticsPackages:    &diagnosticsPackagesClient,
		MobileNetwork:          &mobileNetworkClient,
		MobileNetworks:         &mobileNetworksClient,
		PacketCaptures:         &packetCapturesClient,
		PacketCoreControlPlane: &packetCoreControlPlaneClient,
		PacketCoreControlPlaneCollectDiagnosticsPackage: &packetCoreControlPlaneCollectDiagnosticsPackageClient,
		PacketCoreControlPlaneReinstall:                 &packetCoreControlPlaneReinstallClient,
		PacketCoreControlPlaneRollback:                  &packetCoreControlPlaneRollbackClient,
		PacketCoreControlPlaneVersion:                   &packetCoreControlPlaneVersionClient,
		PacketCoreControlPlanes:                         &packetCoreControlPlanesClient,
		PacketCoreDataPlane:                             &packetCoreDataPlaneClient,
		PacketCoreDataPlanes:                            &packetCoreDataPlanesClient,
		SIM:                                             &sIMClient,
		SIMGroup:                                        &sIMGroupClient,
		SIMGroups:                                       &sIMGroupsClient,
		SIMPolicies:                                     &sIMPoliciesClient,
		SIMPolicy:                                       &sIMPolicyClient,
		SIMs:                                            &sIMsClient,
		Service:                                         &serviceClient,
		Services:                                        &servicesClient,
		Site:                                            &siteClient,
		Sites:                                           &sitesClient,
		Slice:                                           &sliceClient,
		Slices:                                          &slicesClient,
	}
}
