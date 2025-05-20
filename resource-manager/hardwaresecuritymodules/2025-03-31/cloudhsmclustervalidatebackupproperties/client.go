package cloudhsmclustervalidatebackupproperties

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudHSMClusterValidateBackupPropertiesClient struct {
	Client *resourcemanager.Client
}

func NewCloudHSMClusterValidateBackupPropertiesClientWithBaseURI(sdkApi sdkEnv.Api) (*CloudHSMClusterValidateBackupPropertiesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "cloudhsmclustervalidatebackupproperties", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CloudHSMClusterValidateBackupPropertiesClient: %+v", err)
	}

	return &CloudHSMClusterValidateBackupPropertiesClient{
		Client: client,
	}, nil
}
