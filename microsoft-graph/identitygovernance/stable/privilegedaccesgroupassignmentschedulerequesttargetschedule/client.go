package privilegedaccesgroupassignmentschedulerequesttargetschedule

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccesGroupAssignmentScheduleRequestTargetScheduleClient struct {
	Client *msgraph.Client
}

func NewPrivilegedAccesGroupAssignmentScheduleRequestTargetScheduleClientWithBaseURI(sdkApi sdkEnv.Api) (*PrivilegedAccesGroupAssignmentScheduleRequestTargetScheduleClient, error) {
	client, err := msgraph.NewClient(sdkApi, "privilegedaccesgroupassignmentschedulerequesttargetschedule", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PrivilegedAccesGroupAssignmentScheduleRequestTargetScheduleClient: %+v", err)
	}

	return &PrivilegedAccesGroupAssignmentScheduleRequestTargetScheduleClient{
		Client: client,
	}, nil
}
