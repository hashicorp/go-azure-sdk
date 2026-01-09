package directoryroleeligibilityschedulerequesttargetschedule

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryRoleEligibilityScheduleRequestTargetScheduleClient struct {
	Client *msgraph.Client
}

func NewDirectoryRoleEligibilityScheduleRequestTargetScheduleClientWithBaseURI(sdkApi sdkEnv.Api) (*DirectoryRoleEligibilityScheduleRequestTargetScheduleClient, error) {
	client, err := msgraph.NewClient(sdkApi, "directoryroleeligibilityschedulerequesttargetschedule", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryRoleEligibilityScheduleRequestTargetScheduleClient: %+v", err)
	}

	return &DirectoryRoleEligibilityScheduleRequestTargetScheduleClient{
		Client: client,
	}, nil
}
