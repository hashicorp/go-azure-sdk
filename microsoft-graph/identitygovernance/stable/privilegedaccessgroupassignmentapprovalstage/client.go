package privilegedaccessgroupassignmentapprovalstage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessGroupAssignmentApprovalStageClient struct {
	Client *msgraph.Client
}

func NewPrivilegedAccessGroupAssignmentApprovalStageClientWithBaseURI(sdkApi sdkEnv.Api) (*PrivilegedAccessGroupAssignmentApprovalStageClient, error) {
	client, err := msgraph.NewClient(sdkApi, "privilegedaccessgroupassignmentapprovalstage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PrivilegedAccessGroupAssignmentApprovalStageClient: %+v", err)
	}

	return &PrivilegedAccessGroupAssignmentApprovalStageClient{
		Client: client,
	}, nil
}
