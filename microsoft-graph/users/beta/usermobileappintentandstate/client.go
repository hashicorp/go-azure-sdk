package usermobileappintentandstate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserMobileAppIntentAndStateClient struct {
	Client *msgraph.Client
}

func NewUserMobileAppIntentAndStateClientWithBaseURI(api sdkEnv.Api) (*UserMobileAppIntentAndStateClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usermobileappintentandstate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserMobileAppIntentAndStateClient: %+v", err)
	}

	return &UserMobileAppIntentAndStateClient{
		Client: client,
	}, nil
}
