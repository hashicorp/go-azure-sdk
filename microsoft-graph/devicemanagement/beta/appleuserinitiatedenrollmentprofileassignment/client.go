package appleuserinitiatedenrollmentprofileassignment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppleUserInitiatedEnrollmentProfileAssignmentClient struct {
	Client *msgraph.Client
}

func NewAppleUserInitiatedEnrollmentProfileAssignmentClientWithBaseURI(sdkApi sdkEnv.Api) (*AppleUserInitiatedEnrollmentProfileAssignmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "appleuserinitiatedenrollmentprofileassignment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AppleUserInitiatedEnrollmentProfileAssignmentClient: %+v", err)
	}

	return &AppleUserInitiatedEnrollmentProfileAssignmentClient{
		Client: client,
	}, nil
}
