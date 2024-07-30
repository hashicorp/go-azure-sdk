package enterpriseapproleassignmentscheduleinstancedirectoryscope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnterpriseAppRoleAssignmentScheduleInstanceDirectoryScopeClient struct {
	Client *msgraph.Client
}

func NewEnterpriseAppRoleAssignmentScheduleInstanceDirectoryScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*EnterpriseAppRoleAssignmentScheduleInstanceDirectoryScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "enterpriseapproleassignmentscheduleinstancedirectoryscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EnterpriseAppRoleAssignmentScheduleInstanceDirectoryScopeClient: %+v", err)
	}

	return &EnterpriseAppRoleAssignmentScheduleInstanceDirectoryScopeClient{
		Client: client,
	}, nil
}
