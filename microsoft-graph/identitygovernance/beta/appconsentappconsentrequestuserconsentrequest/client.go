package appconsentappconsentrequestuserconsentrequest

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppConsentAppConsentRequestUserConsentRequestClient struct {
	Client *msgraph.Client
}

func NewAppConsentAppConsentRequestUserConsentRequestClientWithBaseURI(sdkApi sdkEnv.Api) (*AppConsentAppConsentRequestUserConsentRequestClient, error) {
	client, err := msgraph.NewClient(sdkApi, "appconsentappconsentrequestuserconsentrequest", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AppConsentAppConsentRequestUserConsentRequestClient: %+v", err)
	}

	return &AppConsentAppConsentRequestUserConsentRequestClient{
		Client: client,
	}, nil
}
