package enterpriseapproleassignmentscheduleinstanceprincipal

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnterpriseAppRoleAssignmentScheduleInstancePrincipalClient struct {
	Client *msgraph.Client
}

func NewEnterpriseAppRoleAssignmentScheduleInstancePrincipalClientWithBaseURI(sdkApi sdkEnv.Api) (*EnterpriseAppRoleAssignmentScheduleInstancePrincipalClient, error) {
	client, err := msgraph.NewClient(sdkApi, "enterpriseapproleassignmentscheduleinstanceprincipal", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EnterpriseAppRoleAssignmentScheduleInstancePrincipalClient: %+v", err)
	}

	return &EnterpriseAppRoleAssignmentScheduleInstancePrincipalClient{
		Client: client,
	}, nil
}
