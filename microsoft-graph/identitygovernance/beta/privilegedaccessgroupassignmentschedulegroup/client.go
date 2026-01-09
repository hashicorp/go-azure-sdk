package privilegedaccessgroupassignmentschedulegroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessGroupAssignmentScheduleGroupClient struct {
	Client *msgraph.Client
}

func NewPrivilegedAccessGroupAssignmentScheduleGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*PrivilegedAccessGroupAssignmentScheduleGroupClient, error) {
	client, err := msgraph.NewClient(sdkApi, "privilegedaccessgroupassignmentschedulegroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PrivilegedAccessGroupAssignmentScheduleGroupClient: %+v", err)
	}

	return &PrivilegedAccessGroupAssignmentScheduleGroupClient{
		Client: client,
	}, nil
}
