package appconsentappconsentrequest

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppConsentAppConsentRequestClient struct {
	Client *msgraph.Client
}

func NewAppConsentAppConsentRequestClientWithBaseURI(sdkApi sdkEnv.Api) (*AppConsentAppConsentRequestClient, error) {
	client, err := msgraph.NewClient(sdkApi, "appconsentappconsentrequest", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AppConsentAppConsentRequestClient: %+v", err)
	}

	return &AppConsentAppConsentRequestClient{
		Client: client,
	}, nil
}
