package enterpriseapproleassignmentprincipal

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnterpriseAppRoleAssignmentPrincipalClient struct {
	Client *msgraph.Client
}

func NewEnterpriseAppRoleAssignmentPrincipalClientWithBaseURI(sdkApi sdkEnv.Api) (*EnterpriseAppRoleAssignmentPrincipalClient, error) {
	client, err := msgraph.NewClient(sdkApi, "enterpriseapproleassignmentprincipal", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EnterpriseAppRoleAssignmentPrincipalClient: %+v", err)
	}

	return &EnterpriseAppRoleAssignmentPrincipalClient{
		Client: client,
	}, nil
}
