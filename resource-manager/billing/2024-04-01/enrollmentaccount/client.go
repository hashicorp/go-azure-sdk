package enrollmentaccount

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnrollmentAccountClient struct {
	Client *resourcemanager.Client
}

func NewEnrollmentAccountClientWithBaseURI(sdkApi sdkEnv.Api) (*EnrollmentAccountClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "enrollmentaccount", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EnrollmentAccountClient: %+v", err)
	}

	return &EnrollmentAccountClient{
		Client: client,
	}, nil
}
