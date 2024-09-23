package privilegedaccessgroupassignmentschedulerequesttargetschedule

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessGroupAssignmentScheduleRequestTargetScheduleClient struct {
	Client *msgraph.Client
}

func NewPrivilegedAccessGroupAssignmentScheduleRequestTargetScheduleClientWithBaseURI(sdkApi sdkEnv.Api) (*PrivilegedAccessGroupAssignmentScheduleRequestTargetScheduleClient, error) {
	client, err := msgraph.NewClient(sdkApi, "privilegedaccessgroupassignmentschedulerequesttargetschedule", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PrivilegedAccessGroupAssignmentScheduleRequestTargetScheduleClient: %+v", err)
	}

	return &PrivilegedAccessGroupAssignmentScheduleRequestTargetScheduleClient{
		Client: client,
	}, nil
}
