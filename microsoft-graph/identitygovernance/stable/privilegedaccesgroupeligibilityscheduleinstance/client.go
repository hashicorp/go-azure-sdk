package privilegedaccesgroupeligibilityscheduleinstance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccesGroupEligibilityScheduleInstanceClient struct {
	Client *msgraph.Client
}

func NewPrivilegedAccesGroupEligibilityScheduleInstanceClientWithBaseURI(sdkApi sdkEnv.Api) (*PrivilegedAccesGroupEligibilityScheduleInstanceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "privilegedaccesgroupeligibilityscheduleinstance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PrivilegedAccesGroupEligibilityScheduleInstanceClient: %+v", err)
	}

	return &PrivilegedAccesGroupEligibilityScheduleInstanceClient{
		Client: client,
	}, nil
}
