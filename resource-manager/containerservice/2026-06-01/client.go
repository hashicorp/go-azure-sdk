package v2026_06_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/containerservice/2026-06-01/agentpools"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerservice/2026-06-01/autoupgradeprofileoperations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerservice/2026-06-01/autoupgradeprofiles"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerservice/2026-06-01/fleetmanagednamespaces"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerservice/2026-06-01/fleetmembers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerservice/2026-06-01/fleets"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerservice/2026-06-01/fleetupdatestrategies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerservice/2026-06-01/gates"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerservice/2026-06-01/identitybindings"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerservice/2026-06-01/machines"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerservice/2026-06-01/maintenanceconfigurations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerservice/2026-06-01/managedclusters"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerservice/2026-06-01/managednamespaces"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerservice/2026-06-01/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerservice/2026-06-01/privatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerservice/2026-06-01/resolveprivatelinkserviceid"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerservice/2026-06-01/snapshots"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerservice/2026-06-01/trustedaccess"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerservice/2026-06-01/updateruns"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AgentPools                   *agentpools.AgentPoolsClient
	AutoUpgradeProfileOperations *autoupgradeprofileoperations.AutoUpgradeProfileOperationsClient
	AutoUpgradeProfiles          *autoupgradeprofiles.AutoUpgradeProfilesClient
	FleetManagedNamespaces       *fleetmanagednamespaces.FleetManagedNamespacesClient
	FleetMembers                 *fleetmembers.FleetMembersClient
	FleetUpdateStrategies        *fleetupdatestrategies.FleetUpdateStrategiesClient
	Fleets                       *fleets.FleetsClient
	Gates                        *gates.GatesClient
	IdentityBindings             *identitybindings.IdentityBindingsClient
	Machines                     *machines.MachinesClient
	MaintenanceConfigurations    *maintenanceconfigurations.MaintenanceConfigurationsClient
	ManagedClusters              *managedclusters.ManagedClustersClient
	ManagedNamespaces            *managednamespaces.ManagedNamespacesClient
	PrivateEndpointConnections   *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources         *privatelinkresources.PrivateLinkResourcesClient
	ResolvePrivateLinkServiceId  *resolveprivatelinkserviceid.ResolvePrivateLinkServiceIdClient
	Snapshots                    *snapshots.SnapshotsClient
	TrustedAccess                *trustedaccess.TrustedAccessClient
	UpdateRuns                   *updateruns.UpdateRunsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	agentPoolsClient, err := agentpools.NewAgentPoolsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AgentPools client: %+v", err)
	}
	configureFunc(agentPoolsClient.Client)

	autoUpgradeProfileOperationsClient, err := autoupgradeprofileoperations.NewAutoUpgradeProfileOperationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AutoUpgradeProfileOperations client: %+v", err)
	}
	configureFunc(autoUpgradeProfileOperationsClient.Client)

	autoUpgradeProfilesClient, err := autoupgradeprofiles.NewAutoUpgradeProfilesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AutoUpgradeProfiles client: %+v", err)
	}
	configureFunc(autoUpgradeProfilesClient.Client)

	fleetManagedNamespacesClient, err := fleetmanagednamespaces.NewFleetManagedNamespacesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building FleetManagedNamespaces client: %+v", err)
	}
	configureFunc(fleetManagedNamespacesClient.Client)

	fleetMembersClient, err := fleetmembers.NewFleetMembersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building FleetMembers client: %+v", err)
	}
	configureFunc(fleetMembersClient.Client)

	fleetUpdateStrategiesClient, err := fleetupdatestrategies.NewFleetUpdateStrategiesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building FleetUpdateStrategies client: %+v", err)
	}
	configureFunc(fleetUpdateStrategiesClient.Client)

	fleetsClient, err := fleets.NewFleetsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Fleets client: %+v", err)
	}
	configureFunc(fleetsClient.Client)

	gatesClient, err := gates.NewGatesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Gates client: %+v", err)
	}
	configureFunc(gatesClient.Client)

	identityBindingsClient, err := identitybindings.NewIdentityBindingsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IdentityBindings client: %+v", err)
	}
	configureFunc(identityBindingsClient.Client)

	machinesClient, err := machines.NewMachinesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Machines client: %+v", err)
	}
	configureFunc(machinesClient.Client)

	maintenanceConfigurationsClient, err := maintenanceconfigurations.NewMaintenanceConfigurationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MaintenanceConfigurations client: %+v", err)
	}
	configureFunc(maintenanceConfigurationsClient.Client)

	managedClustersClient, err := managedclusters.NewManagedClustersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedClusters client: %+v", err)
	}
	configureFunc(managedClustersClient.Client)

	managedNamespacesClient, err := managednamespaces.NewManagedNamespacesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedNamespaces client: %+v", err)
	}
	configureFunc(managedNamespacesClient.Client)

	privateEndpointConnectionsClient, err := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateEndpointConnections client: %+v", err)
	}
	configureFunc(privateEndpointConnectionsClient.Client)

	privateLinkResourcesClient, err := privatelinkresources.NewPrivateLinkResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateLinkResources client: %+v", err)
	}
	configureFunc(privateLinkResourcesClient.Client)

	resolvePrivateLinkServiceIdClient, err := resolveprivatelinkserviceid.NewResolvePrivateLinkServiceIdClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ResolvePrivateLinkServiceId client: %+v", err)
	}
	configureFunc(resolvePrivateLinkServiceIdClient.Client)

	snapshotsClient, err := snapshots.NewSnapshotsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Snapshots client: %+v", err)
	}
	configureFunc(snapshotsClient.Client)

	trustedAccessClient, err := trustedaccess.NewTrustedAccessClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TrustedAccess client: %+v", err)
	}
	configureFunc(trustedAccessClient.Client)

	updateRunsClient, err := updateruns.NewUpdateRunsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UpdateRuns client: %+v", err)
	}
	configureFunc(updateRunsClient.Client)

	return &Client{
		AgentPools:                   agentPoolsClient,
		AutoUpgradeProfileOperations: autoUpgradeProfileOperationsClient,
		AutoUpgradeProfiles:          autoUpgradeProfilesClient,
		FleetManagedNamespaces:       fleetManagedNamespacesClient,
		FleetMembers:                 fleetMembersClient,
		FleetUpdateStrategies:        fleetUpdateStrategiesClient,
		Fleets:                       fleetsClient,
		Gates:                        gatesClient,
		IdentityBindings:             identityBindingsClient,
		Machines:                     machinesClient,
		MaintenanceConfigurations:    maintenanceConfigurationsClient,
		ManagedClusters:              managedClustersClient,
		ManagedNamespaces:            managedNamespacesClient,
		PrivateEndpointConnections:   privateEndpointConnectionsClient,
		PrivateLinkResources:         privateLinkResourcesClient,
		ResolvePrivateLinkServiceId:  resolvePrivateLinkServiceIdClient,
		Snapshots:                    snapshotsClient,
		TrustedAccess:                trustedAccessClient,
		UpdateRuns:                   updateRunsClient,
	}, nil
}
