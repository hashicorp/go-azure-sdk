package privilegedaccessgroupassignmentschedulerequestactivatedusing

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessGroupAssignmentScheduleRequestActivatedUsingClient struct {
	Client *msgraph.Client
}

func NewPrivilegedAccessGroupAssignmentScheduleRequestActivatedUsingClientWithBaseURI(sdkApi sdkEnv.Api) (*PrivilegedAccessGroupAssignmentScheduleRequestActivatedUsingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "privilegedaccessgroupassignmentschedulerequestactivatedusing", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PrivilegedAccessGroupAssignmentScheduleRequestActivatedUsingClient: %+v", err)
	}

	return &PrivilegedAccessGroupAssignmentScheduleRequestActivatedUsingClient{
		Client: client,
	}, nil
}
