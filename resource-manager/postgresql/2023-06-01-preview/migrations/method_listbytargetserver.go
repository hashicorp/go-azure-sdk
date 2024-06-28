package migrations

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByTargetServerOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]MigrationResource
}

type ListByTargetServerCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []MigrationResource
}

type ListByTargetServerOperationOptions struct {
	MigrationListFilter *MigrationListFilter
}

func DefaultListByTargetServerOperationOptions() ListByTargetServerOperationOptions {
	return ListByTargetServerOperationOptions{}
}

func (o ListByTargetServerOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListByTargetServerOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ListByTargetServerOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.MigrationListFilter != nil {
		out.Append("migrationListFilter", fmt.Sprintf("%v", *o.MigrationListFilter))
	}
	return &out
}

type ListByTargetServerCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListByTargetServerCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListByTargetServer ...
func (c MigrationsClient) ListByTargetServer(ctx context.Context, id FlexibleServerId, options ListByTargetServerOperationOptions) (result ListByTargetServerOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Pager:         &ListByTargetServerCustomPager{},
		Path:          fmt.Sprintf("%s/migrations", id.ID()),
		OptionsObject: options,
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
		Values *[]MigrationResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByTargetServerComplete retrieves all the results into a single object
func (c MigrationsClient) ListByTargetServerComplete(ctx context.Context, id FlexibleServerId, options ListByTargetServerOperationOptions) (ListByTargetServerCompleteResult, error) {
	return c.ListByTargetServerCompleteMatchingPredicate(ctx, id, options, MigrationResourceOperationPredicate{})
}

// ListByTargetServerCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MigrationsClient) ListByTargetServerCompleteMatchingPredicate(ctx context.Context, id FlexibleServerId, options ListByTargetServerOperationOptions, predicate MigrationResourceOperationPredicate) (result ListByTargetServerCompleteResult, err error) {
	items := make([]MigrationResource, 0)

	resp, err := c.ListByTargetServer(ctx, id, options)
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

	result = ListByTargetServerCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
