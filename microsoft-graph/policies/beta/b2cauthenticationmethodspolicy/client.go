package b2cauthenticationmethodspolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type B2cAuthenticationMethodsPolicyClient struct {
	Client *msgraph.Client
}

func NewB2cAuthenticationMethodsPolicyClientWithBaseURI(sdkApi sdkEnv.Api) (*B2cAuthenticationMethodsPolicyClient, error) {
	client, err := msgraph.NewClient(sdkApi, "b2cauthenticationmethodspolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating B2cAuthenticationMethodsPolicyClient: %+v", err)
	}

	return &B2cAuthenticationMethodsPolicyClient{
		Client: client,
	}, nil
}
