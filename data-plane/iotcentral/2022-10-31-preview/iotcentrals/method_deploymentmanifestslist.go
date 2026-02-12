package iotcentrals

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeploymentManifestsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]DeploymentManifest
}

type DeploymentManifestsListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []DeploymentManifest
}

type DeploymentManifestsListOperationOptions struct {
	Filter      *string
	Maxpagesize *int64
}

func DefaultDeploymentManifestsListOperationOptions() DeploymentManifestsListOperationOptions {
	return DeploymentManifestsListOperationOptions{}
}

func (o DeploymentManifestsListOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o DeploymentManifestsListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DeploymentManifestsListOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.Maxpagesize != nil {
		out.Append("maxpagesize", fmt.Sprintf("%v", *o.Maxpagesize))
	}
	return &out
}

type DeploymentManifestsListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *DeploymentManifestsListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// DeploymentManifestsList ...
func (c IotcentralsClient) DeploymentManifestsList(ctx context.Context, options DeploymentManifestsListOperationOptions) (result DeploymentManifestsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &DeploymentManifestsListCustomPager{},
		Path:          "/deploymentManifests",
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
		Values *[]DeploymentManifest `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// DeploymentManifestsListComplete retrieves all the results into a single object
func (c IotcentralsClient) DeploymentManifestsListComplete(ctx context.Context, options DeploymentManifestsListOperationOptions) (DeploymentManifestsListCompleteResult, error) {
	return c.DeploymentManifestsListCompleteMatchingPredicate(ctx, options, DeploymentManifestOperationPredicate{})
}

// DeploymentManifestsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c IotcentralsClient) DeploymentManifestsListCompleteMatchingPredicate(ctx context.Context, options DeploymentManifestsListOperationOptions, predicate DeploymentManifestOperationPredicate) (result DeploymentManifestsListCompleteResult, err error) {
	items := make([]DeploymentManifest, 0)

	resp, err := c.DeploymentManifestsList(ctx, options)
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

	result = DeploymentManifestsListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
