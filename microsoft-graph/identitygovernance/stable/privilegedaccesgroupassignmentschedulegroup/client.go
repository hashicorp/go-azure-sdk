package privilegedaccesgroupassignmentschedulegroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccesGroupAssignmentScheduleGroupClient struct {
	Client *msgraph.Client
}

func NewPrivilegedAccesGroupAssignmentScheduleGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*PrivilegedAccesGroupAssignmentScheduleGroupClient, error) {
	client, err := msgraph.NewClient(sdkApi, "privilegedaccesgroupassignmentschedulegroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PrivilegedAccesGroupAssignmentScheduleGroupClient: %+v", err)
	}

	return &PrivilegedAccesGroupAssignmentScheduleGroupClient{
		Client: client,
	}, nil
}
