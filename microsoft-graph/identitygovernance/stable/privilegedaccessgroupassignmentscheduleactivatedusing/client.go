package privilegedaccessgroupassignmentscheduleactivatedusing

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessGroupAssignmentScheduleActivatedUsingClient struct {
	Client *msgraph.Client
}

func NewPrivilegedAccessGroupAssignmentScheduleActivatedUsingClientWithBaseURI(sdkApi sdkEnv.Api) (*PrivilegedAccessGroupAssignmentScheduleActivatedUsingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "privilegedaccessgroupassignmentscheduleactivatedusing", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PrivilegedAccessGroupAssignmentScheduleActivatedUsingClient: %+v", err)
	}

	return &PrivilegedAccessGroupAssignmentScheduleActivatedUsingClient{
		Client: client,
	}, nil
}
