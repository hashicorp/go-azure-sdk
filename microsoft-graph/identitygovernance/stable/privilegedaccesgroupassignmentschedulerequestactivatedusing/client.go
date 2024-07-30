package privilegedaccesgroupassignmentschedulerequestactivatedusing

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccesGroupAssignmentScheduleRequestActivatedUsingClient struct {
	Client *msgraph.Client
}

func NewPrivilegedAccesGroupAssignmentScheduleRequestActivatedUsingClientWithBaseURI(sdkApi sdkEnv.Api) (*PrivilegedAccesGroupAssignmentScheduleRequestActivatedUsingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "privilegedaccesgroupassignmentschedulerequestactivatedusing", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PrivilegedAccesGroupAssignmentScheduleRequestActivatedUsingClient: %+v", err)
	}

	return &PrivilegedAccesGroupAssignmentScheduleRequestActivatedUsingClient{
		Client: client,
	}, nil
}
