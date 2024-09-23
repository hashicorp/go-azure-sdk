package privilegedaccessgroupassignmentscheduleinstancegroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessGroupAssignmentScheduleInstanceGroupClient struct {
	Client *msgraph.Client
}

func NewPrivilegedAccessGroupAssignmentScheduleInstanceGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*PrivilegedAccessGroupAssignmentScheduleInstanceGroupClient, error) {
	client, err := msgraph.NewClient(sdkApi, "privilegedaccessgroupassignmentscheduleinstancegroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PrivilegedAccessGroupAssignmentScheduleInstanceGroupClient: %+v", err)
	}

	return &PrivilegedAccessGroupAssignmentScheduleInstanceGroupClient{
		Client: client,
	}, nil
}
