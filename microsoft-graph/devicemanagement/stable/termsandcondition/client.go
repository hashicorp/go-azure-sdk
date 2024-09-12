package termsandcondition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TermsAndConditionClient struct {
	Client *msgraph.Client
}

func NewTermsAndConditionClientWithBaseURI(sdkApi sdkEnv.Api) (*TermsAndConditionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "termsandcondition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TermsAndConditionClient: %+v", err)
	}

	return &TermsAndConditionClient{
		Client: client,
	}, nil
}
