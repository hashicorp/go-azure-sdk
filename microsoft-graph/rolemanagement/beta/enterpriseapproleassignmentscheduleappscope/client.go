package enterpriseapproleassignmentscheduleappscope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnterpriseAppRoleAssignmentScheduleAppScopeClient struct {
	Client *msgraph.Client
}

func NewEnterpriseAppRoleAssignmentScheduleAppScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*EnterpriseAppRoleAssignmentScheduleAppScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "enterpriseapproleassignmentscheduleappscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EnterpriseAppRoleAssignmentScheduleAppScopeClient: %+v", err)
	}

	return &EnterpriseAppRoleAssignmentScheduleAppScopeClient{
		Client: client,
	}, nil
}
