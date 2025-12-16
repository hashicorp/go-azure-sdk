package elastics

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

type ElasticVersionsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ElasticVersionListFormat
}

type ElasticVersionsListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ElasticVersionListFormat
}

type ElasticVersionsListOperationOptions struct {
	Region *string
}

func DefaultElasticVersionsListOperationOptions() ElasticVersionsListOperationOptions {
	return ElasticVersionsListOperationOptions{}
}

func (o ElasticVersionsListOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ElasticVersionsListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ElasticVersionsListOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Region != nil {
		out.Append("region", fmt.Sprintf("%v", *o.Region))
	}
	return &out
}

type ElasticVersionsListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ElasticVersionsListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ElasticVersionsList ...
func (c ElasticsClient) ElasticVersionsList(ctx context.Context, id commonids.SubscriptionId, options ElasticVersionsListOperationOptions) (result ElasticVersionsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ElasticVersionsListCustomPager{},
		Path:          fmt.Sprintf("%s/providers/Microsoft.Elastic/elasticVersions", id.ID()),
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
		Values *[]ElasticVersionListFormat `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ElasticVersionsListComplete retrieves all the results into a single object
func (c ElasticsClient) ElasticVersionsListComplete(ctx context.Context, id commonids.SubscriptionId, options ElasticVersionsListOperationOptions) (ElasticVersionsListCompleteResult, error) {
	return c.ElasticVersionsListCompleteMatchingPredicate(ctx, id, options, ElasticVersionListFormatOperationPredicate{})
}

// ElasticVersionsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ElasticsClient) ElasticVersionsListCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, options ElasticVersionsListOperationOptions, predicate ElasticVersionListFormatOperationPredicate) (result ElasticVersionsListCompleteResult, err error) {
	items := make([]ElasticVersionListFormat, 0)

	resp, err := c.ElasticVersionsList(ctx, id, options)
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

	result = ElasticVersionsListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
