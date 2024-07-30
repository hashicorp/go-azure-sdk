package privilegedaccesgroupeligibilityschedulerequest

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccesGroupEligibilityScheduleRequestClient struct {
	Client *msgraph.Client
}

func NewPrivilegedAccesGroupEligibilityScheduleRequestClientWithBaseURI(sdkApi sdkEnv.Api) (*PrivilegedAccesGroupEligibilityScheduleRequestClient, error) {
	client, err := msgraph.NewClient(sdkApi, "privilegedaccesgroupeligibilityschedulerequest", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PrivilegedAccesGroupEligibilityScheduleRequestClient: %+v", err)
	}

	return &PrivilegedAccesGroupEligibilityScheduleRequestClient{
		Client: client,
	}, nil
}
