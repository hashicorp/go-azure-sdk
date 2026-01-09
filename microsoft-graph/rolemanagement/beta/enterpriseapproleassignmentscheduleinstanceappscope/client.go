package enterpriseapproleassignmentscheduleinstanceappscope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnterpriseAppRoleAssignmentScheduleInstanceAppScopeClient struct {
	Client *msgraph.Client
}

func NewEnterpriseAppRoleAssignmentScheduleInstanceAppScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*EnterpriseAppRoleAssignmentScheduleInstanceAppScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "enterpriseapproleassignmentscheduleinstanceappscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EnterpriseAppRoleAssignmentScheduleInstanceAppScopeClient: %+v", err)
	}

	return &EnterpriseAppRoleAssignmentScheduleInstanceAppScopeClient{
		Client: client,
	}, nil
}
