package cloudhsmclusterbackupcreate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudHSMClusterBackupCreateClient struct {
	Client *resourcemanager.Client
}

func NewCloudHSMClusterBackupCreateClientWithBaseURI(sdkApi sdkEnv.Api) (*CloudHSMClusterBackupCreateClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "cloudhsmclusterbackupcreate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CloudHSMClusterBackupCreateClient: %+v", err)
	}

	return &CloudHSMClusterBackupCreateClient{
		Client: client,
	}, nil
}
