package privilegedaccesgroupassignmentschedulerequestgroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccesGroupAssignmentScheduleRequestGroupClient struct {
	Client *msgraph.Client
}

func NewPrivilegedAccesGroupAssignmentScheduleRequestGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*PrivilegedAccesGroupAssignmentScheduleRequestGroupClient, error) {
	client, err := msgraph.NewClient(sdkApi, "privilegedaccesgroupassignmentschedulerequestgroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PrivilegedAccesGroupAssignmentScheduleRequestGroupClient: %+v", err)
	}

	return &PrivilegedAccesGroupAssignmentScheduleRequestGroupClient{
		Client: client,
	}, nil
}
