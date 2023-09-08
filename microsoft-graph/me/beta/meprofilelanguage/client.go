package meprofilelanguage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeProfileLanguageClient struct {
	Client *msgraph.Client
}

func NewMeProfileLanguageClientWithBaseURI(api sdkEnv.Api) (*MeProfileLanguageClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meprofilelanguage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeProfileLanguageClient: %+v", err)
	}

	return &MeProfileLanguageClient{
		Client: client,
	}, nil
}
