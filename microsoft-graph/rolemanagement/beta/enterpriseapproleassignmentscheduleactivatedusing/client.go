package enterpriseapproleassignmentscheduleactivatedusing

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnterpriseAppRoleAssignmentScheduleActivatedUsingClient struct {
	Client *msgraph.Client
}

func NewEnterpriseAppRoleAssignmentScheduleActivatedUsingClientWithBaseURI(sdkApi sdkEnv.Api) (*EnterpriseAppRoleAssignmentScheduleActivatedUsingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "enterpriseapproleassignmentscheduleactivatedusing", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EnterpriseAppRoleAssignmentScheduleActivatedUsingClient: %+v", err)
	}

	return &EnterpriseAppRoleAssignmentScheduleActivatedUsingClient{
		Client: client,
	}, nil
}
