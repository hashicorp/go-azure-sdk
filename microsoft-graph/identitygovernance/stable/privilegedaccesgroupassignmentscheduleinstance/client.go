package privilegedaccesgroupassignmentscheduleinstance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccesGroupAssignmentScheduleInstanceClient struct {
	Client *msgraph.Client
}

func NewPrivilegedAccesGroupAssignmentScheduleInstanceClientWithBaseURI(sdkApi sdkEnv.Api) (*PrivilegedAccesGroupAssignmentScheduleInstanceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "privilegedaccesgroupassignmentscheduleinstance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PrivilegedAccesGroupAssignmentScheduleInstanceClient: %+v", err)
	}

	return &PrivilegedAccesGroupAssignmentScheduleInstanceClient{
		Client: client,
	}, nil
}
