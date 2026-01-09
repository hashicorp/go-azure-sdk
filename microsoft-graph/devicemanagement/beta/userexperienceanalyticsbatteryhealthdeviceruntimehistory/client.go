package userexperienceanalyticsbatteryhealthdeviceruntimehistory

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryClient struct {
	Client *msgraph.Client
}

func NewUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryClientWithBaseURI(sdkApi sdkEnv.Api) (*UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userexperienceanalyticsbatteryhealthdeviceruntimehistory", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryClient: %+v", err)
	}

	return &UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryClient{
		Client: client,
	}, nil
}
