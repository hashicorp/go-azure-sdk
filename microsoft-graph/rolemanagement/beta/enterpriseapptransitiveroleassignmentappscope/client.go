package enterpriseapptransitiveroleassignmentappscope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnterpriseAppTransitiveRoleAssignmentAppScopeClient struct {
	Client *msgraph.Client
}

func NewEnterpriseAppTransitiveRoleAssignmentAppScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*EnterpriseAppTransitiveRoleAssignmentAppScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "enterpriseapptransitiveroleassignmentappscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EnterpriseAppTransitiveRoleAssignmentAppScopeClient: %+v", err)
	}

	return &EnterpriseAppTransitiveRoleAssignmentAppScopeClient{
		Client: client,
	}, nil
}
