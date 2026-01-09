package cloudpcroleassignmentdirectoryscope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCRoleAssignmentDirectoryScopeClient struct {
	Client *msgraph.Client
}

func NewCloudPCRoleAssignmentDirectoryScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*CloudPCRoleAssignmentDirectoryScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "cloudpcroleassignmentdirectoryscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CloudPCRoleAssignmentDirectoryScopeClient: %+v", err)
	}

	return &CloudPCRoleAssignmentDirectoryScopeClient{
		Client: client,
	}, nil
}
