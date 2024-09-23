package termsofuseagreementfilelocalization

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TermsOfUseAgreementFileLocalizationClient struct {
	Client *msgraph.Client
}

func NewTermsOfUseAgreementFileLocalizationClientWithBaseURI(sdkApi sdkEnv.Api) (*TermsOfUseAgreementFileLocalizationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "termsofuseagreementfilelocalization", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TermsOfUseAgreementFileLocalizationClient: %+v", err)
	}

	return &TermsOfUseAgreementFileLocalizationClient{
		Client: client,
	}, nil
}
