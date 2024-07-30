package conditionalaccesauthenticationcontextclassreference

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccesAuthenticationContextClassReferenceClient struct {
	Client *msgraph.Client
}

func NewConditionalAccesAuthenticationContextClassReferenceClientWithBaseURI(sdkApi sdkEnv.Api) (*ConditionalAccesAuthenticationContextClassReferenceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "conditionalaccesauthenticationcontextclassreference", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ConditionalAccesAuthenticationContextClassReferenceClient: %+v", err)
	}

	return &ConditionalAccesAuthenticationContextClassReferenceClient{
		Client: client,
	}, nil
}
