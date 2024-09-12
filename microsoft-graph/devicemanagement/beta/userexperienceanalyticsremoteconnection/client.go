package userexperienceanalyticsremoteconnection

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsRemoteConnectionClient struct {
	Client *msgraph.Client
}

func NewUserExperienceAnalyticsRemoteConnectionClientWithBaseURI(sdkApi sdkEnv.Api) (*UserExperienceAnalyticsRemoteConnectionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userexperienceanalyticsremoteconnection", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserExperienceAnalyticsRemoteConnectionClient: %+v", err)
	}

	return &UserExperienceAnalyticsRemoteConnectionClient{
		Client: client,
	}, nil
}
