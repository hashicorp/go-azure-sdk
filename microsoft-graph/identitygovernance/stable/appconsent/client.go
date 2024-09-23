package appconsent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppConsentClient struct {
	Client *msgraph.Client
}

func NewAppConsentClientWithBaseURI(sdkApi sdkEnv.Api) (*AppConsentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "appconsent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AppConsentClient: %+v", err)
	}

	return &AppConsentClient{
		Client: client,
	}, nil
}
