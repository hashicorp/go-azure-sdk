package privilegedaccesgroupassignmentapproval

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccesGroupAssignmentApprovalClient struct {
	Client *msgraph.Client
}

func NewPrivilegedAccesGroupAssignmentApprovalClientWithBaseURI(sdkApi sdkEnv.Api) (*PrivilegedAccesGroupAssignmentApprovalClient, error) {
	client, err := msgraph.NewClient(sdkApi, "privilegedaccesgroupassignmentapproval", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PrivilegedAccesGroupAssignmentApprovalClient: %+v", err)
	}

	return &PrivilegedAccesGroupAssignmentApprovalClient{
		Client: client,
	}, nil
}
