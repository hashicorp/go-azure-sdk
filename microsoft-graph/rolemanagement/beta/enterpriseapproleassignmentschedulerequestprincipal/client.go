package enterpriseapproleassignmentschedulerequestprincipal

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnterpriseAppRoleAssignmentScheduleRequestPrincipalClient struct {
	Client *msgraph.Client
}

func NewEnterpriseAppRoleAssignmentScheduleRequestPrincipalClientWithBaseURI(sdkApi sdkEnv.Api) (*EnterpriseAppRoleAssignmentScheduleRequestPrincipalClient, error) {
	client, err := msgraph.NewClient(sdkApi, "enterpriseapproleassignmentschedulerequestprincipal", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EnterpriseAppRoleAssignmentScheduleRequestPrincipalClient: %+v", err)
	}

	return &EnterpriseAppRoleAssignmentScheduleRequestPrincipalClient{
		Client: client,
	}, nil
}
