package privilegedaccessgroupassignmentapprovalstep

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessGroupAssignmentApprovalStepClient struct {
	Client *msgraph.Client
}

func NewPrivilegedAccessGroupAssignmentApprovalStepClientWithBaseURI(sdkApi sdkEnv.Api) (*PrivilegedAccessGroupAssignmentApprovalStepClient, error) {
	client, err := msgraph.NewClient(sdkApi, "privilegedaccessgroupassignmentapprovalstep", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PrivilegedAccessGroupAssignmentApprovalStepClient: %+v", err)
	}

	return &PrivilegedAccessGroupAssignmentApprovalStepClient{
		Client: client,
	}, nil
}
