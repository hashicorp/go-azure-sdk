package privilegedaccessgroupeligibilityschedulerequestprincipal

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessGroupEligibilityScheduleRequestPrincipalClient struct {
	Client *msgraph.Client
}

func NewPrivilegedAccessGroupEligibilityScheduleRequestPrincipalClientWithBaseURI(sdkApi sdkEnv.Api) (*PrivilegedAccessGroupEligibilityScheduleRequestPrincipalClient, error) {
	client, err := msgraph.NewClient(sdkApi, "privilegedaccessgroupeligibilityschedulerequestprincipal", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PrivilegedAccessGroupEligibilityScheduleRequestPrincipalClient: %+v", err)
	}

	return &PrivilegedAccessGroupEligibilityScheduleRequestPrincipalClient{
		Client: client,
	}, nil
}
