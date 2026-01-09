package privilegedaccessgroupassignmentscheduleinstanceprincipal

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessGroupAssignmentScheduleInstancePrincipalClient struct {
	Client *msgraph.Client
}

func NewPrivilegedAccessGroupAssignmentScheduleInstancePrincipalClientWithBaseURI(sdkApi sdkEnv.Api) (*PrivilegedAccessGroupAssignmentScheduleInstancePrincipalClient, error) {
	client, err := msgraph.NewClient(sdkApi, "privilegedaccessgroupassignmentscheduleinstanceprincipal", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PrivilegedAccessGroupAssignmentScheduleInstancePrincipalClient: %+v", err)
	}

	return &PrivilegedAccessGroupAssignmentScheduleInstancePrincipalClient{
		Client: client,
	}, nil
}
