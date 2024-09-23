package privilegedaccessgroupassignmentschedulerequestprincipal

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessGroupAssignmentScheduleRequestPrincipalClient struct {
	Client *msgraph.Client
}

func NewPrivilegedAccessGroupAssignmentScheduleRequestPrincipalClientWithBaseURI(sdkApi sdkEnv.Api) (*PrivilegedAccessGroupAssignmentScheduleRequestPrincipalClient, error) {
	client, err := msgraph.NewClient(sdkApi, "privilegedaccessgroupassignmentschedulerequestprincipal", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PrivilegedAccessGroupAssignmentScheduleRequestPrincipalClient: %+v", err)
	}

	return &PrivilegedAccessGroupAssignmentScheduleRequestPrincipalClient{
		Client: client,
	}, nil
}
