package enterpriseapproleassignmentappscope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnterpriseAppRoleAssignmentAppScopeClient struct {
	Client *msgraph.Client
}

func NewEnterpriseAppRoleAssignmentAppScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*EnterpriseAppRoleAssignmentAppScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "enterpriseapproleassignmentappscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EnterpriseAppRoleAssignmentAppScopeClient: %+v", err)
	}

	return &EnterpriseAppRoleAssignmentAppScopeClient{
		Client: client,
	}, nil
}
