package termsofuse

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TermsOfUseClient struct {
	Client *msgraph.Client
}

func NewTermsOfUseClientWithBaseURI(sdkApi sdkEnv.Api) (*TermsOfUseClient, error) {
	client, err := msgraph.NewClient(sdkApi, "termsofuse", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TermsOfUseClient: %+v", err)
	}

	return &TermsOfUseClient{
		Client: client,
	}, nil
}
