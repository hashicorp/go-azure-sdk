package enterpriseapproleassignmentschedulerequestappscope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnterpriseAppRoleAssignmentScheduleRequestAppScopeClient struct {
	Client *msgraph.Client
}

func NewEnterpriseAppRoleAssignmentScheduleRequestAppScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*EnterpriseAppRoleAssignmentScheduleRequestAppScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "enterpriseapproleassignmentschedulerequestappscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EnterpriseAppRoleAssignmentScheduleRequestAppScopeClient: %+v", err)
	}

	return &EnterpriseAppRoleAssignmentScheduleRequestAppScopeClient{
		Client: client,
	}, nil
}
