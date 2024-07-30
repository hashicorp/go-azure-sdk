package privilegedaccesgroupeligibilityschedulerequestprincipal

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccesGroupEligibilityScheduleRequestPrincipalClient struct {
	Client *msgraph.Client
}

func NewPrivilegedAccesGroupEligibilityScheduleRequestPrincipalClientWithBaseURI(sdkApi sdkEnv.Api) (*PrivilegedAccesGroupEligibilityScheduleRequestPrincipalClient, error) {
	client, err := msgraph.NewClient(sdkApi, "privilegedaccesgroupeligibilityschedulerequestprincipal", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PrivilegedAccesGroupEligibilityScheduleRequestPrincipalClient: %+v", err)
	}

	return &PrivilegedAccesGroupEligibilityScheduleRequestPrincipalClient{
		Client: client,
	}, nil
}
