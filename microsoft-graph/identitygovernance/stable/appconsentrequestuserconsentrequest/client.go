package appconsentrequestuserconsentrequest

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppConsentRequestUserConsentRequestClient struct {
	Client *msgraph.Client
}

func NewAppConsentRequestUserConsentRequestClientWithBaseURI(sdkApi sdkEnv.Api) (*AppConsentRequestUserConsentRequestClient, error) {
	client, err := msgraph.NewClient(sdkApi, "appconsentrequestuserconsentrequest", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AppConsentRequestUserConsentRequestClient: %+v", err)
	}

	return &AppConsentRequestUserConsentRequestClient{
		Client: client,
	}, nil
}
