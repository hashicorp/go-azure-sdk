package enterpriseapproleassignmentscheduleprincipal

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnterpriseAppRoleAssignmentSchedulePrincipalClient struct {
	Client *msgraph.Client
}

func NewEnterpriseAppRoleAssignmentSchedulePrincipalClientWithBaseURI(sdkApi sdkEnv.Api) (*EnterpriseAppRoleAssignmentSchedulePrincipalClient, error) {
	client, err := msgraph.NewClient(sdkApi, "enterpriseapproleassignmentscheduleprincipal", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EnterpriseAppRoleAssignmentSchedulePrincipalClient: %+v", err)
	}

	return &EnterpriseAppRoleAssignmentSchedulePrincipalClient{
		Client: client,
	}, nil
}
