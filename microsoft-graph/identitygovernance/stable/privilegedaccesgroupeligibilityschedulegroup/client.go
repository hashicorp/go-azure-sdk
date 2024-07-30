package privilegedaccesgroupeligibilityschedulegroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccesGroupEligibilityScheduleGroupClient struct {
	Client *msgraph.Client
}

func NewPrivilegedAccesGroupEligibilityScheduleGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*PrivilegedAccesGroupEligibilityScheduleGroupClient, error) {
	client, err := msgraph.NewClient(sdkApi, "privilegedaccesgroupeligibilityschedulegroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PrivilegedAccesGroupEligibilityScheduleGroupClient: %+v", err)
	}

	return &PrivilegedAccesGroupEligibilityScheduleGroupClient{
		Client: client,
	}, nil
}
