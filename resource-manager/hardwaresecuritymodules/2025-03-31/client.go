package v2025_03_31

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/hardwaresecuritymodules/2025-03-31/cloudhsmclusterbackupcreate"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hardwaresecuritymodules/2025-03-31/cloudhsmclustercreateorupdate"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hardwaresecuritymodules/2025-03-31/cloudhsmclusterdelete"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hardwaresecuritymodules/2025-03-31/cloudhsmclusterget"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hardwaresecuritymodules/2025-03-31/cloudhsmclustergetbyresourcegroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hardwaresecuritymodules/2025-03-31/cloudhsmclustergetbysubscription"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hardwaresecuritymodules/2025-03-31/cloudhsmclusterprivateendpointconnectionsget"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hardwaresecuritymodules/2025-03-31/cloudhsmclusterprivateendpointcreate"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hardwaresecuritymodules/2025-03-31/cloudhsmclusterprivateendpointdelete"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hardwaresecuritymodules/2025-03-31/cloudhsmclusterprivateendpointget"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hardwaresecuritymodules/2025-03-31/cloudhsmclusterprivatelinkget"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hardwaresecuritymodules/2025-03-31/cloudhsmclustersbackupoperationstatusget"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hardwaresecuritymodules/2025-03-31/cloudhsmclustersrestoreoperationcreate"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hardwaresecuritymodules/2025-03-31/cloudhsmclusterupdate"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hardwaresecuritymodules/2025-03-31/cloudhsmclustervalidatebackupproperties"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hardwaresecuritymodules/2025-03-31/dedicatedhsms"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hardwaresecuritymodules/2025-03-31/returnsrestoreoperationstatus"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hardwaresecuritymodules/2025-03-31/validatecloudhsmclusterrestoreproperties"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	CloudHSMClusterBackupCreate                  *cloudhsmclusterbackupcreate.CloudHSMClusterBackupCreateClient
	CloudHSMClusterCreateOrUpdate                *cloudhsmclustercreateorupdate.CloudHSMClusterCreateOrUpdateClient
	CloudHSMClusterDelete                        *cloudhsmclusterdelete.CloudHSMClusterDeleteClient
	CloudHSMClusterGet                           *cloudhsmclusterget.CloudHSMClusterGetClient
	CloudHSMClusterGetByResourceGroup            *cloudhsmclustergetbyresourcegroup.CloudHSMClusterGetByResourceGroupClient
	CloudHSMClusterGetBySubscription             *cloudhsmclustergetbysubscription.CloudHSMClusterGetBySubscriptionClient
	CloudHSMClusterPrivateEndpointConnectionsGet *cloudhsmclusterprivateendpointconnectionsget.CloudHSMClusterPrivateEndpointConnectionsGetClient
	CloudHSMClusterPrivateEndpointCreate         *cloudhsmclusterprivateendpointcreate.CloudHSMClusterPrivateEndpointCreateClient
	CloudHSMClusterPrivateEndpointDelete         *cloudhsmclusterprivateendpointdelete.CloudHSMClusterPrivateEndpointDeleteClient
	CloudHSMClusterPrivateEndpointGet            *cloudhsmclusterprivateendpointget.CloudHSMClusterPrivateEndpointGetClient
	CloudHSMClusterPrivateLinkGet                *cloudhsmclusterprivatelinkget.CloudHSMClusterPrivateLinkGetClient
	CloudHSMClusterUpdate                        *cloudhsmclusterupdate.CloudHSMClusterUpdateClient
	CloudHSMClusterValidateBackupProperties      *cloudhsmclustervalidatebackupproperties.CloudHSMClusterValidateBackupPropertiesClient
	CloudHSMClustersBackupOperationStatusGet     *cloudhsmclustersbackupoperationstatusget.CloudHSMClustersBackupOperationStatusGetClient
	CloudHSMClustersRestoreOperationCreate       *cloudhsmclustersrestoreoperationcreate.CloudHSMClustersRestoreOperationCreateClient
	DedicatedHsms                                *dedicatedhsms.DedicatedHsmsClient
	ReturnsRestoreOperationStatus                *returnsrestoreoperationstatus.ReturnsRestoreOperationStatusClient
	ValidateCloudHSMClusterRestoreProperties     *validatecloudhsmclusterrestoreproperties.ValidateCloudHSMClusterRestorePropertiesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	cloudHSMClusterBackupCreateClient, err := cloudhsmclusterbackupcreate.NewCloudHSMClusterBackupCreateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CloudHSMClusterBackupCreate client: %+v", err)
	}
	configureFunc(cloudHSMClusterBackupCreateClient.Client)

	cloudHSMClusterCreateOrUpdateClient, err := cloudhsmclustercreateorupdate.NewCloudHSMClusterCreateOrUpdateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CloudHSMClusterCreateOrUpdate client: %+v", err)
	}
	configureFunc(cloudHSMClusterCreateOrUpdateClient.Client)

	cloudHSMClusterDeleteClient, err := cloudhsmclusterdelete.NewCloudHSMClusterDeleteClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CloudHSMClusterDelete client: %+v", err)
	}
	configureFunc(cloudHSMClusterDeleteClient.Client)

	cloudHSMClusterGetByResourceGroupClient, err := cloudhsmclustergetbyresourcegroup.NewCloudHSMClusterGetByResourceGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CloudHSMClusterGetByResourceGroup client: %+v", err)
	}
	configureFunc(cloudHSMClusterGetByResourceGroupClient.Client)

	cloudHSMClusterGetBySubscriptionClient, err := cloudhsmclustergetbysubscription.NewCloudHSMClusterGetBySubscriptionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CloudHSMClusterGetBySubscription client: %+v", err)
	}
	configureFunc(cloudHSMClusterGetBySubscriptionClient.Client)

	cloudHSMClusterGetClient, err := cloudhsmclusterget.NewCloudHSMClusterGetClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CloudHSMClusterGet client: %+v", err)
	}
	configureFunc(cloudHSMClusterGetClient.Client)

	cloudHSMClusterPrivateEndpointConnectionsGetClient, err := cloudhsmclusterprivateendpointconnectionsget.NewCloudHSMClusterPrivateEndpointConnectionsGetClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CloudHSMClusterPrivateEndpointConnectionsGet client: %+v", err)
	}
	configureFunc(cloudHSMClusterPrivateEndpointConnectionsGetClient.Client)

	cloudHSMClusterPrivateEndpointCreateClient, err := cloudhsmclusterprivateendpointcreate.NewCloudHSMClusterPrivateEndpointCreateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CloudHSMClusterPrivateEndpointCreate client: %+v", err)
	}
	configureFunc(cloudHSMClusterPrivateEndpointCreateClient.Client)

	cloudHSMClusterPrivateEndpointDeleteClient, err := cloudhsmclusterprivateendpointdelete.NewCloudHSMClusterPrivateEndpointDeleteClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CloudHSMClusterPrivateEndpointDelete client: %+v", err)
	}
	configureFunc(cloudHSMClusterPrivateEndpointDeleteClient.Client)

	cloudHSMClusterPrivateEndpointGetClient, err := cloudhsmclusterprivateendpointget.NewCloudHSMClusterPrivateEndpointGetClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CloudHSMClusterPrivateEndpointGet client: %+v", err)
	}
	configureFunc(cloudHSMClusterPrivateEndpointGetClient.Client)

	cloudHSMClusterPrivateLinkGetClient, err := cloudhsmclusterprivatelinkget.NewCloudHSMClusterPrivateLinkGetClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CloudHSMClusterPrivateLinkGet client: %+v", err)
	}
	configureFunc(cloudHSMClusterPrivateLinkGetClient.Client)

	cloudHSMClusterUpdateClient, err := cloudhsmclusterupdate.NewCloudHSMClusterUpdateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CloudHSMClusterUpdate client: %+v", err)
	}
	configureFunc(cloudHSMClusterUpdateClient.Client)

	cloudHSMClusterValidateBackupPropertiesClient, err := cloudhsmclustervalidatebackupproperties.NewCloudHSMClusterValidateBackupPropertiesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CloudHSMClusterValidateBackupProperties client: %+v", err)
	}
	configureFunc(cloudHSMClusterValidateBackupPropertiesClient.Client)

	cloudHSMClustersBackupOperationStatusGetClient, err := cloudhsmclustersbackupoperationstatusget.NewCloudHSMClustersBackupOperationStatusGetClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CloudHSMClustersBackupOperationStatusGet client: %+v", err)
	}
	configureFunc(cloudHSMClustersBackupOperationStatusGetClient.Client)

	cloudHSMClustersRestoreOperationCreateClient, err := cloudhsmclustersrestoreoperationcreate.NewCloudHSMClustersRestoreOperationCreateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CloudHSMClustersRestoreOperationCreate client: %+v", err)
	}
	configureFunc(cloudHSMClustersRestoreOperationCreateClient.Client)

	dedicatedHsmsClient, err := dedicatedhsms.NewDedicatedHsmsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DedicatedHsms client: %+v", err)
	}
	configureFunc(dedicatedHsmsClient.Client)

	returnsRestoreOperationStatusClient, err := returnsrestoreoperationstatus.NewReturnsRestoreOperationStatusClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ReturnsRestoreOperationStatus client: %+v", err)
	}
	configureFunc(returnsRestoreOperationStatusClient.Client)

	validateCloudHSMClusterRestorePropertiesClient, err := validatecloudhsmclusterrestoreproperties.NewValidateCloudHSMClusterRestorePropertiesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ValidateCloudHSMClusterRestoreProperties client: %+v", err)
	}
	configureFunc(validateCloudHSMClusterRestorePropertiesClient.Client)

	return &Client{
		CloudHSMClusterBackupCreate:                  cloudHSMClusterBackupCreateClient,
		CloudHSMClusterCreateOrUpdate:                cloudHSMClusterCreateOrUpdateClient,
		CloudHSMClusterDelete:                        cloudHSMClusterDeleteClient,
		CloudHSMClusterGet:                           cloudHSMClusterGetClient,
		CloudHSMClusterGetByResourceGroup:            cloudHSMClusterGetByResourceGroupClient,
		CloudHSMClusterGetBySubscription:             cloudHSMClusterGetBySubscriptionClient,
		CloudHSMClusterPrivateEndpointConnectionsGet: cloudHSMClusterPrivateEndpointConnectionsGetClient,
		CloudHSMClusterPrivateEndpointCreate:         cloudHSMClusterPrivateEndpointCreateClient,
		CloudHSMClusterPrivateEndpointDelete:         cloudHSMClusterPrivateEndpointDeleteClient,
		CloudHSMClusterPrivateEndpointGet:            cloudHSMClusterPrivateEndpointGetClient,
		CloudHSMClusterPrivateLinkGet:                cloudHSMClusterPrivateLinkGetClient,
		CloudHSMClusterUpdate:                        cloudHSMClusterUpdateClient,
		CloudHSMClusterValidateBackupProperties:      cloudHSMClusterValidateBackupPropertiesClient,
		CloudHSMClustersBackupOperationStatusGet:     cloudHSMClustersBackupOperationStatusGetClient,
		CloudHSMClustersRestoreOperationCreate:       cloudHSMClustersRestoreOperationCreateClient,
		DedicatedHsms:                                dedicatedHsmsClient,
		ReturnsRestoreOperationStatus:                returnsRestoreOperationStatusClient,
		ValidateCloudHSMClusterRestoreProperties:     validateCloudHSMClusterRestorePropertiesClient,
	}, nil
}
