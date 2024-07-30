package privilegedaccesgroupassignmentschedulerequestprincipal

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccesGroupAssignmentScheduleRequestPrincipalClient struct {
	Client *msgraph.Client
}

func NewPrivilegedAccesGroupAssignmentScheduleRequestPrincipalClientWithBaseURI(sdkApi sdkEnv.Api) (*PrivilegedAccesGroupAssignmentScheduleRequestPrincipalClient, error) {
	client, err := msgraph.NewClient(sdkApi, "privilegedaccesgroupassignmentschedulerequestprincipal", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PrivilegedAccesGroupAssignmentScheduleRequestPrincipalClient: %+v", err)
	}

	return &PrivilegedAccesGroupAssignmentScheduleRequestPrincipalClient{
		Client: client,
	}, nil
}
