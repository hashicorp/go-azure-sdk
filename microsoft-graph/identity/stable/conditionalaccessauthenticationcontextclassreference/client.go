package conditionalaccessauthenticationcontextclassreference

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessAuthenticationContextClassReferenceClient struct {
	Client *msgraph.Client
}

func NewConditionalAccessAuthenticationContextClassReferenceClientWithBaseURI(sdkApi sdkEnv.Api) (*ConditionalAccessAuthenticationContextClassReferenceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "conditionalaccessauthenticationcontextclassreference", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ConditionalAccessAuthenticationContextClassReferenceClient: %+v", err)
	}

	return &ConditionalAccessAuthenticationContextClassReferenceClient{
		Client: client,
	}, nil
}
