package profilelanguage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProfileLanguageClient struct {
	Client *msgraph.Client
}

func NewProfileLanguageClientWithBaseURI(sdkApi sdkEnv.Api) (*ProfileLanguageClient, error) {
	client, err := msgraph.NewClient(sdkApi, "profilelanguage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ProfileLanguageClient: %+v", err)
	}

	return &ProfileLanguageClient{
		Client: client,
	}, nil
}
