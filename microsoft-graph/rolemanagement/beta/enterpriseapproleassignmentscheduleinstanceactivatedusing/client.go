package enterpriseapproleassignmentscheduleinstanceactivatedusing

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnterpriseAppRoleAssignmentScheduleInstanceActivatedUsingClient struct {
	Client *msgraph.Client
}

func NewEnterpriseAppRoleAssignmentScheduleInstanceActivatedUsingClientWithBaseURI(sdkApi sdkEnv.Api) (*EnterpriseAppRoleAssignmentScheduleInstanceActivatedUsingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "enterpriseapproleassignmentscheduleinstanceactivatedusing", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EnterpriseAppRoleAssignmentScheduleInstanceActivatedUsingClient: %+v", err)
	}

	return &EnterpriseAppRoleAssignmentScheduleInstanceActivatedUsingClient{
		Client: client,
	}, nil
}
