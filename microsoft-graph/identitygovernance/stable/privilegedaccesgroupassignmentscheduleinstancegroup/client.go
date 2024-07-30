package privilegedaccesgroupassignmentscheduleinstancegroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccesGroupAssignmentScheduleInstanceGroupClient struct {
	Client *msgraph.Client
}

func NewPrivilegedAccesGroupAssignmentScheduleInstanceGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*PrivilegedAccesGroupAssignmentScheduleInstanceGroupClient, error) {
	client, err := msgraph.NewClient(sdkApi, "privilegedaccesgroupassignmentscheduleinstancegroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PrivilegedAccesGroupAssignmentScheduleInstanceGroupClient: %+v", err)
	}

	return &PrivilegedAccesGroupAssignmentScheduleInstanceGroupClient{
		Client: client,
	}, nil
}
