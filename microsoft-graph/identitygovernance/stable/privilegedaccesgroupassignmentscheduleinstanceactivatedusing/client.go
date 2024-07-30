package privilegedaccesgroupassignmentscheduleinstanceactivatedusing

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccesGroupAssignmentScheduleInstanceActivatedUsingClient struct {
	Client *msgraph.Client
}

func NewPrivilegedAccesGroupAssignmentScheduleInstanceActivatedUsingClientWithBaseURI(sdkApi sdkEnv.Api) (*PrivilegedAccesGroupAssignmentScheduleInstanceActivatedUsingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "privilegedaccesgroupassignmentscheduleinstanceactivatedusing", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PrivilegedAccesGroupAssignmentScheduleInstanceActivatedUsingClient: %+v", err)
	}

	return &PrivilegedAccesGroupAssignmentScheduleInstanceActivatedUsingClient{
		Client: client,
	}, nil
}
