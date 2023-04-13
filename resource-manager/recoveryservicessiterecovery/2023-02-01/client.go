package v2023_02_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-02-01/migrationrecoverypoints"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-02-01/operations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-02-01/recoverypoints"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-02-01/replicationalertsettings"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-02-01/replicationappliances"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-02-01/replicationeligibilityresults"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-02-01/replicationevents"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-02-01/replicationfabrics"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-02-01/replicationjobs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-02-01/replicationlogicalnetworks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-02-01/replicationmigrationitems"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-02-01/replicationnetworkmappings"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-02-01/replicationnetworks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-02-01/replicationpolicies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-02-01/replicationprotectableitems"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-02-01/replicationprotecteditems"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-02-01/replicationprotectioncontainermappings"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-02-01/replicationprotectioncontainers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-02-01/replicationprotectionintents"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-02-01/replicationrecoveryplans"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-02-01/replicationrecoveryservicesproviders"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-02-01/replicationstorageclassificationmappings"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-02-01/replicationstorageclassifications"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-02-01/replicationvaulthealth"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-02-01/replicationvaultsetting"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-02-01/replicationvcenters"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-02-01/supportedoperatingsystems"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-02-01/targetcomputesizes"
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

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	migrationRecoveryPointsClient := migrationrecoverypoints.NewMigrationRecoveryPointsClientWithBaseURI(endpoint)
	configureAuthFunc(&migrationRecoveryPointsClient.Client)

	operationsClient := operations.NewOperationsClientWithBaseURI(endpoint)
	configureAuthFunc(&operationsClient.Client)

	recoveryPointsClient := recoverypoints.NewRecoveryPointsClientWithBaseURI(endpoint)
	configureAuthFunc(&recoveryPointsClient.Client)

	replicationAlertSettingsClient := replicationalertsettings.NewReplicationAlertSettingsClientWithBaseURI(endpoint)
	configureAuthFunc(&replicationAlertSettingsClient.Client)

	replicationAppliancesClient := replicationappliances.NewReplicationAppliancesClientWithBaseURI(endpoint)
	configureAuthFunc(&replicationAppliancesClient.Client)

	replicationEligibilityResultsClient := replicationeligibilityresults.NewReplicationEligibilityResultsClientWithBaseURI(endpoint)
	configureAuthFunc(&replicationEligibilityResultsClient.Client)

	replicationEventsClient := replicationevents.NewReplicationEventsClientWithBaseURI(endpoint)
	configureAuthFunc(&replicationEventsClient.Client)

	replicationFabricsClient := replicationfabrics.NewReplicationFabricsClientWithBaseURI(endpoint)
	configureAuthFunc(&replicationFabricsClient.Client)

	replicationJobsClient := replicationjobs.NewReplicationJobsClientWithBaseURI(endpoint)
	configureAuthFunc(&replicationJobsClient.Client)

	replicationLogicalNetworksClient := replicationlogicalnetworks.NewReplicationLogicalNetworksClientWithBaseURI(endpoint)
	configureAuthFunc(&replicationLogicalNetworksClient.Client)

	replicationMigrationItemsClient := replicationmigrationitems.NewReplicationMigrationItemsClientWithBaseURI(endpoint)
	configureAuthFunc(&replicationMigrationItemsClient.Client)

	replicationNetworkMappingsClient := replicationnetworkmappings.NewReplicationNetworkMappingsClientWithBaseURI(endpoint)
	configureAuthFunc(&replicationNetworkMappingsClient.Client)

	replicationNetworksClient := replicationnetworks.NewReplicationNetworksClientWithBaseURI(endpoint)
	configureAuthFunc(&replicationNetworksClient.Client)

	replicationPoliciesClient := replicationpolicies.NewReplicationPoliciesClientWithBaseURI(endpoint)
	configureAuthFunc(&replicationPoliciesClient.Client)

	replicationProtectableItemsClient := replicationprotectableitems.NewReplicationProtectableItemsClientWithBaseURI(endpoint)
	configureAuthFunc(&replicationProtectableItemsClient.Client)

	replicationProtectedItemsClient := replicationprotecteditems.NewReplicationProtectedItemsClientWithBaseURI(endpoint)
	configureAuthFunc(&replicationProtectedItemsClient.Client)

	replicationProtectionContainerMappingsClient := replicationprotectioncontainermappings.NewReplicationProtectionContainerMappingsClientWithBaseURI(endpoint)
	configureAuthFunc(&replicationProtectionContainerMappingsClient.Client)

	replicationProtectionContainersClient := replicationprotectioncontainers.NewReplicationProtectionContainersClientWithBaseURI(endpoint)
	configureAuthFunc(&replicationProtectionContainersClient.Client)

	replicationProtectionIntentsClient := replicationprotectionintents.NewReplicationProtectionIntentsClientWithBaseURI(endpoint)
	configureAuthFunc(&replicationProtectionIntentsClient.Client)

	replicationRecoveryPlansClient := replicationrecoveryplans.NewReplicationRecoveryPlansClientWithBaseURI(endpoint)
	configureAuthFunc(&replicationRecoveryPlansClient.Client)

	replicationRecoveryServicesProvidersClient := replicationrecoveryservicesproviders.NewReplicationRecoveryServicesProvidersClientWithBaseURI(endpoint)
	configureAuthFunc(&replicationRecoveryServicesProvidersClient.Client)

	replicationStorageClassificationMappingsClient := replicationstorageclassificationmappings.NewReplicationStorageClassificationMappingsClientWithBaseURI(endpoint)
	configureAuthFunc(&replicationStorageClassificationMappingsClient.Client)

	replicationStorageClassificationsClient := replicationstorageclassifications.NewReplicationStorageClassificationsClientWithBaseURI(endpoint)
	configureAuthFunc(&replicationStorageClassificationsClient.Client)

	replicationVaultHealthClient := replicationvaulthealth.NewReplicationVaultHealthClientWithBaseURI(endpoint)
	configureAuthFunc(&replicationVaultHealthClient.Client)

	replicationVaultSettingClient := replicationvaultsetting.NewReplicationVaultSettingClientWithBaseURI(endpoint)
	configureAuthFunc(&replicationVaultSettingClient.Client)

	replicationvCentersClient := replicationvcenters.NewReplicationvCentersClientWithBaseURI(endpoint)
	configureAuthFunc(&replicationvCentersClient.Client)

	supportedOperatingSystemsClient := supportedoperatingsystems.NewSupportedOperatingSystemsClientWithBaseURI(endpoint)
	configureAuthFunc(&supportedOperatingSystemsClient.Client)

	targetComputeSizesClient := targetcomputesizes.NewTargetComputeSizesClientWithBaseURI(endpoint)
	configureAuthFunc(&targetComputeSizesClient.Client)

	return Client{
		MigrationRecoveryPoints:                  &migrationRecoveryPointsClient,
		Operations:                               &operationsClient,
		RecoveryPoints:                           &recoveryPointsClient,
		ReplicationAlertSettings:                 &replicationAlertSettingsClient,
		ReplicationAppliances:                    &replicationAppliancesClient,
		ReplicationEligibilityResults:            &replicationEligibilityResultsClient,
		ReplicationEvents:                        &replicationEventsClient,
		ReplicationFabrics:                       &replicationFabricsClient,
		ReplicationJobs:                          &replicationJobsClient,
		ReplicationLogicalNetworks:               &replicationLogicalNetworksClient,
		ReplicationMigrationItems:                &replicationMigrationItemsClient,
		ReplicationNetworkMappings:               &replicationNetworkMappingsClient,
		ReplicationNetworks:                      &replicationNetworksClient,
		ReplicationPolicies:                      &replicationPoliciesClient,
		ReplicationProtectableItems:              &replicationProtectableItemsClient,
		ReplicationProtectedItems:                &replicationProtectedItemsClient,
		ReplicationProtectionContainerMappings:   &replicationProtectionContainerMappingsClient,
		ReplicationProtectionContainers:          &replicationProtectionContainersClient,
		ReplicationProtectionIntents:             &replicationProtectionIntentsClient,
		ReplicationRecoveryPlans:                 &replicationRecoveryPlansClient,
		ReplicationRecoveryServicesProviders:     &replicationRecoveryServicesProvidersClient,
		ReplicationStorageClassificationMappings: &replicationStorageClassificationMappingsClient,
		ReplicationStorageClassifications:        &replicationStorageClassificationsClient,
		ReplicationVaultHealth:                   &replicationVaultHealthClient,
		ReplicationVaultSetting:                  &replicationVaultSettingClient,
		ReplicationvCenters:                      &replicationvCentersClient,
		SupportedOperatingSystems:                &supportedOperatingSystemsClient,
		TargetComputeSizes:                       &targetComputeSizesClient,
	}
}
