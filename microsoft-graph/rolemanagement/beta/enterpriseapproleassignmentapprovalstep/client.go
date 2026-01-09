package enterpriseapproleassignmentapprovalstep

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnterpriseAppRoleAssignmentApprovalStepClient struct {
	Client *msgraph.Client
}

func NewEnterpriseAppRoleAssignmentApprovalStepClientWithBaseURI(sdkApi sdkEnv.Api) (*EnterpriseAppRoleAssignmentApprovalStepClient, error) {
	client, err := msgraph.NewClient(sdkApi, "enterpriseapproleassignmentapprovalstep", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EnterpriseAppRoleAssignmentApprovalStepClient: %+v", err)
	}

	return &EnterpriseAppRoleAssignmentApprovalStepClient{
		Client: client,
	}, nil
}
