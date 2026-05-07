package openapis

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RemediationsListDeploymentsAtResourceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]RemediationDeployment
}

type RemediationsListDeploymentsAtResourceCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []RemediationDeployment
}

type RemediationsListDeploymentsAtResourceOperationOptions struct {
	Top *int64
}

func DefaultRemediationsListDeploymentsAtResourceOperationOptions() RemediationsListDeploymentsAtResourceOperationOptions {
	return RemediationsListDeploymentsAtResourceOperationOptions{}
}

func (o RemediationsListDeploymentsAtResourceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o RemediationsListDeploymentsAtResourceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o RemediationsListDeploymentsAtResourceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

type RemediationsListDeploymentsAtResourceCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *RemediationsListDeploymentsAtResourceCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// RemediationsListDeploymentsAtResource ...
func (c OpenapisClient) RemediationsListDeploymentsAtResource(ctx context.Context, id ScopedRemediationId, options RemediationsListDeploymentsAtResourceOperationOptions) (result RemediationsListDeploymentsAtResourceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &RemediationsListDeploymentsAtResourceCustomPager{},
		Path:          fmt.Sprintf("%s/listDeployments", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]RemediationDeployment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// RemediationsListDeploymentsAtResourceComplete retrieves all the results into a single object
func (c OpenapisClient) RemediationsListDeploymentsAtResourceComplete(ctx context.Context, id ScopedRemediationId, options RemediationsListDeploymentsAtResourceOperationOptions) (RemediationsListDeploymentsAtResourceCompleteResult, error) {
	return c.RemediationsListDeploymentsAtResourceCompleteMatchingPredicate(ctx, id, options, RemediationDeploymentOperationPredicate{})
}

// RemediationsListDeploymentsAtResourceCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OpenapisClient) RemediationsListDeploymentsAtResourceCompleteMatchingPredicate(ctx context.Context, id ScopedRemediationId, options RemediationsListDeploymentsAtResourceOperationOptions, predicate RemediationDeploymentOperationPredicate) (result RemediationsListDeploymentsAtResourceCompleteResult, err error) {
	items := make([]RemediationDeployment, 0)

	resp, err := c.RemediationsListDeploymentsAtResource(ctx, id, options)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = RemediationsListDeploymentsAtResourceCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
