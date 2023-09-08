package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/domains/beta/domain"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/domains/beta/domaindomainnamereference"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/domains/beta/domainfederationconfiguration"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/domains/beta/domainserviceconfigurationrecord"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/domains/beta/domainsharedemaildomaininvitation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/domains/beta/domainverificationdnsrecord"
	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
)

type Client struct {
	Domain                            *domain.DomainClient
	DomainDomainNameReference         *domaindomainnamereference.DomainDomainNameReferenceClient
	DomainFederationConfiguration     *domainfederationconfiguration.DomainFederationConfigurationClient
	DomainServiceConfigurationRecord  *domainserviceconfigurationrecord.DomainServiceConfigurationRecordClient
	DomainSharedEmailDomainInvitation *domainsharedemaildomaininvitation.DomainSharedEmailDomainInvitationClient
	DomainVerificationDnsRecord       *domainverificationdnsrecord.DomainVerificationDnsRecordClient
}

func NewClientWithBaseURI(api skdEnv.Api, configureFunc func(c *msgraph.Client)) (*Client, error) {
	domainClient, err := domain.NewDomainClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Domain client: %+v", err)
	}
	configureFunc(domainClient.Client)

	domainDomainNameReferenceClient, err := domaindomainnamereference.NewDomainDomainNameReferenceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DomainDomainNameReference client: %+v", err)
	}
	configureFunc(domainDomainNameReferenceClient.Client)

	domainFederationConfigurationClient, err := domainfederationconfiguration.NewDomainFederationConfigurationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DomainFederationConfiguration client: %+v", err)
	}
	configureFunc(domainFederationConfigurationClient.Client)

	domainServiceConfigurationRecordClient, err := domainserviceconfigurationrecord.NewDomainServiceConfigurationRecordClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DomainServiceConfigurationRecord client: %+v", err)
	}
	configureFunc(domainServiceConfigurationRecordClient.Client)

	domainSharedEmailDomainInvitationClient, err := domainsharedemaildomaininvitation.NewDomainSharedEmailDomainInvitationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DomainSharedEmailDomainInvitation client: %+v", err)
	}
	configureFunc(domainSharedEmailDomainInvitationClient.Client)

	domainVerificationDnsRecordClient, err := domainverificationdnsrecord.NewDomainVerificationDnsRecordClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DomainVerificationDnsRecord client: %+v", err)
	}
	configureFunc(domainVerificationDnsRecordClient.Client)

	return &Client{
		Domain:                            domainClient,
		DomainDomainNameReference:         domainDomainNameReferenceClient,
		DomainFederationConfiguration:     domainFederationConfigurationClient,
		DomainServiceConfigurationRecord:  domainServiceConfigurationRecordClient,
		DomainSharedEmailDomainInvitation: domainSharedEmailDomainInvitationClient,
		DomainVerificationDnsRecord:       domainVerificationDnsRecordClient,
	}, nil
}
