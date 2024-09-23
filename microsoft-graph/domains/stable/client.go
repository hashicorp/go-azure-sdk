package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/domains/stable/domain"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/domains/stable/domainnamereference"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/domains/stable/federationconfiguration"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/domains/stable/serviceconfigurationrecord"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/domains/stable/verificationdnsrecord"
	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Domain                     *domain.DomainClient
	DomainNameReference        *domainnamereference.DomainNameReferenceClient
	FederationConfiguration    *federationconfiguration.FederationConfigurationClient
	ServiceConfigurationRecord *serviceconfigurationrecord.ServiceConfigurationRecordClient
	VerificationDnsRecord      *verificationdnsrecord.VerificationDnsRecordClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *msgraph.Client)) (*Client, error) {
	domainClient, err := domain.NewDomainClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Domain client: %+v", err)
	}
	configureFunc(domainClient.Client)

	domainNameReferenceClient, err := domainnamereference.NewDomainNameReferenceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DomainNameReference client: %+v", err)
	}
	configureFunc(domainNameReferenceClient.Client)

	federationConfigurationClient, err := federationconfiguration.NewFederationConfigurationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building FederationConfiguration client: %+v", err)
	}
	configureFunc(federationConfigurationClient.Client)

	serviceConfigurationRecordClient, err := serviceconfigurationrecord.NewServiceConfigurationRecordClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServiceConfigurationRecord client: %+v", err)
	}
	configureFunc(serviceConfigurationRecordClient.Client)

	verificationDnsRecordClient, err := verificationdnsrecord.NewVerificationDnsRecordClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VerificationDnsRecord client: %+v", err)
	}
	configureFunc(verificationDnsRecordClient.Client)

	return &Client{
		Domain:                     domainClient,
		DomainNameReference:        domainNameReferenceClient,
		FederationConfiguration:    federationConfigurationClient,
		ServiceConfigurationRecord: serviceConfigurationRecordClient,
		VerificationDnsRecord:      verificationDnsRecordClient,
	}, nil
}
