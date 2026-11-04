package restorabledroppeddatabases

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

type ListByServerOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]RestorableDroppedDatabase
}

type ListByServerCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []RestorableDroppedDatabase
}

type ListByServerOperationOptions struct {
	Top *int64
}

func DefaultListByServerOperationOptions() ListByServerOperationOptions {
	return ListByServerOperationOptions{}
}

func (o ListByServerOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListByServerOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ListByServerOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

type ListByServerCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListByServerCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListByServer ...
func (c RestorableDroppedDatabasesClient) ListByServer(ctx context.Context, id commonids.SqlServerId, options ListByServerOperationOptions) (result ListByServerOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListByServerCustomPager{},
		Path:          fmt.Sprintf("%s/restorableDroppedDatabases", id.ID()),
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
		Values *[]RestorableDroppedDatabase `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByServerComplete retrieves all the results into a single object
func (c RestorableDroppedDatabasesClient) ListByServerComplete(ctx context.Context, id commonids.SqlServerId, options ListByServerOperationOptions) (ListByServerCompleteResult, error) {
	return c.ListByServerCompleteMatchingPredicate(ctx, id, options, RestorableDroppedDatabaseOperationPredicate{})
}

// ListByServerCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RestorableDroppedDatabasesClient) ListByServerCompleteMatchingPredicate(ctx context.Context, id commonids.SqlServerId, options ListByServerOperationOptions, predicate RestorableDroppedDatabaseOperationPredicate) (result ListByServerCompleteResult, err error) {
	items := make([]RestorableDroppedDatabase, 0)

	resp, err := c.ListByServer(ctx, id, options)
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

	result = ListByServerCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
