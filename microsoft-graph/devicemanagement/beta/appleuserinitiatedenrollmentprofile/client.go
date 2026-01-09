package appleuserinitiatedenrollmentprofile

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppleUserInitiatedEnrollmentProfileClient struct {
	Client *msgraph.Client
}

func NewAppleUserInitiatedEnrollmentProfileClientWithBaseURI(sdkApi sdkEnv.Api) (*AppleUserInitiatedEnrollmentProfileClient, error) {
	client, err := msgraph.NewClient(sdkApi, "appleuserinitiatedenrollmentprofile", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AppleUserInitiatedEnrollmentProfileClient: %+v", err)
	}

	return &AppleUserInitiatedEnrollmentProfileClient{
		Client: client,
	}, nil
}
