package v2023_06_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-06-01/migrationrecoverypoints"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-06-01/operations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-06-01/recoverypoints"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-06-01/replicationalertsettings"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-06-01/replicationappliances"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-06-01/replicationeligibilityresults"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-06-01/replicationevents"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-06-01/replicationfabrics"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-06-01/replicationjobs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-06-01/replicationlogicalnetworks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-06-01/replicationmigrationitems"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-06-01/replicationnetworkmappings"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-06-01/replicationnetworks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-06-01/replicationpolicies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-06-01/replicationprotectableitems"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-06-01/replicationprotecteditems"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-06-01/replicationprotectioncontainermappings"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-06-01/replicationprotectioncontainers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-06-01/replicationprotectionintents"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-06-01/replicationrecoveryplans"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-06-01/replicationrecoveryservicesproviders"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-06-01/replicationstorageclassificationmappings"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-06-01/replicationstorageclassifications"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-06-01/replicationvaulthealth"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-06-01/replicationvaultsetting"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-06-01/replicationvcenters"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-06-01/supportedoperatingsystems"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-06-01/targetcomputesizes"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	MigrationRecoveryPoints                  *migrationrecoverypoints.MigrationRecoveryPointsClient
	Operations                               *operations.OperationsClient
	RecoveryPoints                           *recoverypoints.RecoveryPointsClient
	ReplicationAlertSettings                 *replicationalertsettings.ReplicationAlertSettingsClient
	ReplicationAppliances                    *replicationappliances.ReplicationAppliancesClient
	ReplicationEligibilityResults            *replicationeligibilityresults.ReplicationEligibilityResultsClient
	ReplicationEvents                        *replicationevents.ReplicationEventsClient
	ReplicationFabrics                       *replicationfabrics.ReplicationFabricsClient
	ReplicationJobs                          *replicationjobs.ReplicationJobsClient
	ReplicationLogicalNetworks               *replicationlogicalnetworks.ReplicationLogicalNetworksClient
	ReplicationMigrationItems                *replicationmigrationitems.ReplicationMigrationItemsClient
	ReplicationNetworkMappings               *replicationnetworkmappings.ReplicationNetworkMappingsClient
	ReplicationNetworks                      *replicationnetworks.ReplicationNetworksClient
	ReplicationPolicies                      *replicationpolicies.ReplicationPoliciesClient
	ReplicationProtectableItems              *replicationprotectableitems.ReplicationProtectableItemsClient
	ReplicationProtectedItems                *replicationprotecteditems.ReplicationProtectedItemsClient
	ReplicationProtectionContainerMappings   *replicationprotectioncontainermappings.ReplicationProtectionContainerMappingsClient
	ReplicationProtectionContainers          *replicationprotectioncontainers.ReplicationProtectionContainersClient
	ReplicationProtectionIntents             *replicationprotectionintents.ReplicationProtectionIntentsClient
	ReplicationRecoveryPlans                 *replicationrecoveryplans.ReplicationRecoveryPlansClient
	ReplicationRecoveryServicesProviders     *replicationrecoveryservicesproviders.ReplicationRecoveryServicesProvidersClient
	ReplicationStorageClassificationMappings *replicationstorageclassificationmappings.ReplicationStorageClassificationMappingsClient
	ReplicationStorageClassifications        *replicationstorageclassifications.ReplicationStorageClassificationsClient
	ReplicationVaultHealth                   *replicationvaulthealth.ReplicationVaultHealthClient
	ReplicationVaultSetting                  *replicationvaultsetting.ReplicationVaultSettingClient
	ReplicationvCenters                      *replicationvcenters.ReplicationvCentersClient
	SupportedOperatingSystems                *supportedoperatingsystems.SupportedOperatingSystemsClient
	TargetComputeSizes                       *targetcomputesizes.TargetComputeSizesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	migrationRecoveryPointsClient, err := migrationrecoverypoints.NewMigrationRecoveryPointsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MigrationRecoveryPoints client: %+v", err)
	}
	configureFunc(migrationRecoveryPointsClient.Client)

	operationsClient, err := operations.NewOperationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Operations client: %+v", err)
	}
	configureFunc(operationsClient.Client)

	recoveryPointsClient, err := recoverypoints.NewRecoveryPointsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RecoveryPoints client: %+v", err)
	}
	configureFunc(recoveryPointsClient.Client)

	replicationAlertSettingsClient, err := replicationalertsettings.NewReplicationAlertSettingsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ReplicationAlertSettings client: %+v", err)
	}
	configureFunc(replicationAlertSettingsClient.Client)

	replicationAppliancesClient, err := replicationappliances.NewReplicationAppliancesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ReplicationAppliances client: %+v", err)
	}
	configureFunc(replicationAppliancesClient.Client)

	replicationEligibilityResultsClient, err := replicationeligibilityresults.NewReplicationEligibilityResultsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ReplicationEligibilityResults client: %+v", err)
	}
	configureFunc(replicationEligibilityResultsClient.Client)

	replicationEventsClient, err := replicationevents.NewReplicationEventsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ReplicationEvents client: %+v", err)
	}
	configureFunc(replicationEventsClient.Client)

	replicationFabricsClient, err := replicationfabrics.NewReplicationFabricsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ReplicationFabrics client: %+v", err)
	}
	configureFunc(replicationFabricsClient.Client)

	replicationJobsClient, err := replicationjobs.NewReplicationJobsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ReplicationJobs client: %+v", err)
	}
	configureFunc(replicationJobsClient.Client)

	replicationLogicalNetworksClient, err := replicationlogicalnetworks.NewReplicationLogicalNetworksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ReplicationLogicalNetworks client: %+v", err)
	}
	configureFunc(replicationLogicalNetworksClient.Client)

	replicationMigrationItemsClient, err := replicationmigrationitems.NewReplicationMigrationItemsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ReplicationMigrationItems client: %+v", err)
	}
	configureFunc(replicationMigrationItemsClient.Client)

	replicationNetworkMappingsClient, err := replicationnetworkmappings.NewReplicationNetworkMappingsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ReplicationNetworkMappings client: %+v", err)
	}
	configureFunc(replicationNetworkMappingsClient.Client)

	replicationNetworksClient, err := replicationnetworks.NewReplicationNetworksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ReplicationNetworks client: %+v", err)
	}
	configureFunc(replicationNetworksClient.Client)

	replicationPoliciesClient, err := replicationpolicies.NewReplicationPoliciesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ReplicationPolicies client: %+v", err)
	}
	configureFunc(replicationPoliciesClient.Client)

	replicationProtectableItemsClient, err := replicationprotectableitems.NewReplicationProtectableItemsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ReplicationProtectableItems client: %+v", err)
	}
	configureFunc(replicationProtectableItemsClient.Client)

	replicationProtectedItemsClient, err := replicationprotecteditems.NewReplicationProtectedItemsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ReplicationProtectedItems client: %+v", err)
	}
	configureFunc(replicationProtectedItemsClient.Client)

	replicationProtectionContainerMappingsClient, err := replicationprotectioncontainermappings.NewReplicationProtectionContainerMappingsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ReplicationProtectionContainerMappings client: %+v", err)
	}
	configureFunc(replicationProtectionContainerMappingsClient.Client)

	replicationProtectionContainersClient, err := replicationprotectioncontainers.NewReplicationProtectionContainersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ReplicationProtectionContainers client: %+v", err)
	}
	configureFunc(replicationProtectionContainersClient.Client)

	replicationProtectionIntentsClient, err := replicationprotectionintents.NewReplicationProtectionIntentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ReplicationProtectionIntents client: %+v", err)
	}
	configureFunc(replicationProtectionIntentsClient.Client)

	replicationRecoveryPlansClient, err := replicationrecoveryplans.NewReplicationRecoveryPlansClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ReplicationRecoveryPlans client: %+v", err)
	}
	configureFunc(replicationRecoveryPlansClient.Client)

	replicationRecoveryServicesProvidersClient, err := replicationrecoveryservicesproviders.NewReplicationRecoveryServicesProvidersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ReplicationRecoveryServicesProviders client: %+v", err)
	}
	configureFunc(replicationRecoveryServicesProvidersClient.Client)

	replicationStorageClassificationMappingsClient, err := replicationstorageclassificationmappings.NewReplicationStorageClassificationMappingsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ReplicationStorageClassificationMappings client: %+v", err)
	}
	configureFunc(replicationStorageClassificationMappingsClient.Client)

	replicationStorageClassificationsClient, err := replicationstorageclassifications.NewReplicationStorageClassificationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ReplicationStorageClassifications client: %+v", err)
	}
	configureFunc(replicationStorageClassificationsClient.Client)

	replicationVaultHealthClient, err := replicationvaulthealth.NewReplicationVaultHealthClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ReplicationVaultHealth client: %+v", err)
	}
	configureFunc(replicationVaultHealthClient.Client)

	replicationVaultSettingClient, err := replicationvaultsetting.NewReplicationVaultSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ReplicationVaultSetting client: %+v", err)
	}
	configureFunc(replicationVaultSettingClient.Client)

	replicationvCentersClient, err := replicationvcenters.NewReplicationvCentersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ReplicationvCenters client: %+v", err)
	}
	configureFunc(replicationvCentersClient.Client)

	supportedOperatingSystemsClient, err := supportedoperatingsystems.NewSupportedOperatingSystemsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SupportedOperatingSystems client: %+v", err)
	}
	configureFunc(supportedOperatingSystemsClient.Client)

	targetComputeSizesClient, err := targetcomputesizes.NewTargetComputeSizesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TargetComputeSizes client: %+v", err)
	}
	configureFunc(targetComputeSizesClient.Client)

	return &Client{
		MigrationRecoveryPoints:                  migrationRecoveryPointsClient,
		Operations:                               operationsClient,
		RecoveryPoints:                           recoveryPointsClient,
		ReplicationAlertSettings:                 replicationAlertSettingsClient,
		ReplicationAppliances:                    replicationAppliancesClient,
		ReplicationEligibilityResults:            replicationEligibilityResultsClient,
		ReplicationEvents:                        replicationEventsClient,
		ReplicationFabrics:                       replicationFabricsClient,
		ReplicationJobs:                          replicationJobsClient,
		ReplicationLogicalNetworks:               replicationLogicalNetworksClient,
		ReplicationMigrationItems:                replicationMigrationItemsClient,
		ReplicationNetworkMappings:               replicationNetworkMappingsClient,
		ReplicationNetworks:                      replicationNetworksClient,
		ReplicationPolicies:                      replicationPoliciesClient,
		ReplicationProtectableItems:              replicationProtectableItemsClient,
		ReplicationProtectedItems:                replicationProtectedItemsClient,
		ReplicationProtectionContainerMappings:   replicationProtectionContainerMappingsClient,
		ReplicationProtectionContainers:          replicationProtectionContainersClient,
		ReplicationProtectionIntents:             replicationProtectionIntentsClient,
		ReplicationRecoveryPlans:                 replicationRecoveryPlansClient,
		ReplicationRecoveryServicesProviders:     replicationRecoveryServicesProvidersClient,
		ReplicationStorageClassificationMappings: replicationStorageClassificationMappingsClient,
		ReplicationStorageClassifications:        replicationStorageClassificationsClient,
		ReplicationVaultHealth:                   replicationVaultHealthClient,
		ReplicationVaultSetting:                  replicationVaultSettingClient,
		ReplicationvCenters:                      replicationvCentersClient,
		SupportedOperatingSystems:                supportedOperatingSystemsClient,
		TargetComputeSizes:                       targetComputeSizesClient,
	}, nil
}
