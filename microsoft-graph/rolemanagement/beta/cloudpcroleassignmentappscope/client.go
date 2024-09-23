package cloudpcroleassignmentappscope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCRoleAssignmentAppScopeClient struct {
	Client *msgraph.Client
}

func NewCloudPCRoleAssignmentAppScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*CloudPCRoleAssignmentAppScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "cloudpcroleassignmentappscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CloudPCRoleAssignmentAppScopeClient: %+v", err)
	}

	return &CloudPCRoleAssignmentAppScopeClient{
		Client: client,
	}, nil
}
