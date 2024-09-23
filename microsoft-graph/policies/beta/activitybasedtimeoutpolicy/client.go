package activitybasedtimeoutpolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActivityBasedTimeoutPolicyClient struct {
	Client *msgraph.Client
}

func NewActivityBasedTimeoutPolicyClientWithBaseURI(sdkApi sdkEnv.Api) (*ActivityBasedTimeoutPolicyClient, error) {
	client, err := msgraph.NewClient(sdkApi, "activitybasedtimeoutpolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ActivityBasedTimeoutPolicyClient: %+v", err)
	}

	return &ActivityBasedTimeoutPolicyClient{
		Client: client,
	}, nil
}
