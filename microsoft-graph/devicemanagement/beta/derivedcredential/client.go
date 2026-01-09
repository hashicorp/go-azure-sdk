package derivedcredential

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DerivedCredentialClient struct {
	Client *msgraph.Client
}

func NewDerivedCredentialClientWithBaseURI(sdkApi sdkEnv.Api) (*DerivedCredentialClient, error) {
	client, err := msgraph.NewClient(sdkApi, "derivedcredential", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DerivedCredentialClient: %+v", err)
	}

	return &DerivedCredentialClient{
		Client: client,
	}, nil
}
