package appconsentrequest

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppConsentRequestClient struct {
	Client *msgraph.Client
}

func NewAppConsentRequestClientWithBaseURI(sdkApi sdkEnv.Api) (*AppConsentRequestClient, error) {
	client, err := msgraph.NewClient(sdkApi, "appconsentrequest", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AppConsentRequestClient: %+v", err)
	}

	return &AppConsentRequestClient{
		Client: client,
	}, nil
}
