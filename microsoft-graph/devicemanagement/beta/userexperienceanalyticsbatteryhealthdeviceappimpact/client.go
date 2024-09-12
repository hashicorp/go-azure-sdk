package userexperienceanalyticsbatteryhealthdeviceappimpact

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsBatteryHealthDeviceAppImpactClient struct {
	Client *msgraph.Client
}

func NewUserExperienceAnalyticsBatteryHealthDeviceAppImpactClientWithBaseURI(sdkApi sdkEnv.Api) (*UserExperienceAnalyticsBatteryHealthDeviceAppImpactClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userexperienceanalyticsbatteryhealthdeviceappimpact", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserExperienceAnalyticsBatteryHealthDeviceAppImpactClient: %+v", err)
	}

	return &UserExperienceAnalyticsBatteryHealthDeviceAppImpactClient{
		Client: client,
	}, nil
}
