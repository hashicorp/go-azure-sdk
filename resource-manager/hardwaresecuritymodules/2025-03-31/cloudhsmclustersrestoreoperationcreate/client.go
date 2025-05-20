package cloudhsmclustersrestoreoperationcreate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudHSMClustersRestoreOperationCreateClient struct {
	Client *resourcemanager.Client
}

func NewCloudHSMClustersRestoreOperationCreateClientWithBaseURI(sdkApi sdkEnv.Api) (*CloudHSMClustersRestoreOperationCreateClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "cloudhsmclustersrestoreoperationcreate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CloudHSMClustersRestoreOperationCreateClient: %+v", err)
	}

	return &CloudHSMClustersRestoreOperationCreateClient{
		Client: client,
	}, nil
}
