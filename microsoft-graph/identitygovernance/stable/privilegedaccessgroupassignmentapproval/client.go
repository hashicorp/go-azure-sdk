package privilegedaccessgroupassignmentapproval

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessGroupAssignmentApprovalClient struct {
	Client *msgraph.Client
}

func NewPrivilegedAccessGroupAssignmentApprovalClientWithBaseURI(sdkApi sdkEnv.Api) (*PrivilegedAccessGroupAssignmentApprovalClient, error) {
	client, err := msgraph.NewClient(sdkApi, "privilegedaccessgroupassignmentapproval", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PrivilegedAccessGroupAssignmentApprovalClient: %+v", err)
	}

	return &PrivilegedAccessGroupAssignmentApprovalClient{
		Client: client,
	}, nil
}
