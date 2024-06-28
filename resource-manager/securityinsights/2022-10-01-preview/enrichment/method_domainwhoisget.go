package enrichment

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DomainWhoisGetOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *EnrichmentDomainWhois
}

type DomainWhoisGetOperationOptions struct {
	Domain *string
}

func DefaultDomainWhoisGetOperationOptions() DomainWhoisGetOperationOptions {
	return DomainWhoisGetOperationOptions{}
}

func (o DomainWhoisGetOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o DomainWhoisGetOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o DomainWhoisGetOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Domain != nil {
		out.Append("domain", fmt.Sprintf("%v", *o.Domain))
	}
	return &out
}

// DomainWhoisGet ...
func (c EnrichmentClient) DomainWhoisGet(ctx context.Context, id commonids.ResourceGroupId, options DomainWhoisGetOperationOptions) (result DomainWhoisGetOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/providers/Microsoft.SecurityInsights/enrichment/domain/whois", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var model EnrichmentDomainWhois
	result.Model = &model

	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
