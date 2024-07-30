package privilegedaccesgroupassignmentscheduleinstanceprincipal

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccesGroupAssignmentScheduleInstancePrincipalClient struct {
	Client *msgraph.Client
}

func NewPrivilegedAccesGroupAssignmentScheduleInstancePrincipalClientWithBaseURI(sdkApi sdkEnv.Api) (*PrivilegedAccesGroupAssignmentScheduleInstancePrincipalClient, error) {
	client, err := msgraph.NewClient(sdkApi, "privilegedaccesgroupassignmentscheduleinstanceprincipal", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PrivilegedAccesGroupAssignmentScheduleInstancePrincipalClient: %+v", err)
	}

	return &PrivilegedAccesGroupAssignmentScheduleInstancePrincipalClient{
		Client: client,
	}, nil
}
