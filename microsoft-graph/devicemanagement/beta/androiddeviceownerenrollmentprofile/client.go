package androiddeviceownerenrollmentprofile

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerEnrollmentProfileClient struct {
	Client *msgraph.Client
}

func NewAndroidDeviceOwnerEnrollmentProfileClientWithBaseURI(sdkApi sdkEnv.Api) (*AndroidDeviceOwnerEnrollmentProfileClient, error) {
	client, err := msgraph.NewClient(sdkApi, "androiddeviceownerenrollmentprofile", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AndroidDeviceOwnerEnrollmentProfileClient: %+v", err)
	}

	return &AndroidDeviceOwnerEnrollmentProfileClient{
		Client: client,
	}, nil
}
