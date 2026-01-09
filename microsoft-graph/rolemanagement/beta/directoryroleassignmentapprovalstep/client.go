package directoryroleassignmentapprovalstep

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryRoleAssignmentApprovalStepClient struct {
	Client *msgraph.Client
}

func NewDirectoryRoleAssignmentApprovalStepClientWithBaseURI(sdkApi sdkEnv.Api) (*DirectoryRoleAssignmentApprovalStepClient, error) {
	client, err := msgraph.NewClient(sdkApi, "directoryroleassignmentapprovalstep", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryRoleAssignmentApprovalStepClient: %+v", err)
	}

	return &DirectoryRoleAssignmentApprovalStepClient{
		Client: client,
	}, nil
}
