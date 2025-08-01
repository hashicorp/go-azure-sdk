package ownedobject

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListOwnedObjectApplicationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.Application
}

type ListOwnedObjectApplicationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.Application
}

type ListOwnedObjectApplicationsOperationOptions struct {
	ConsistencyLevel *odata.ConsistencyLevel
	Count            *bool
	Expand           *odata.Expand
	Filter           *string
	Metadata         *odata.Metadata
	OrderBy          *odata.OrderBy
	RetryFunc        client.RequestRetryFunc
	Search           *string
	Select           *[]string
	Skip             *int64
	Top              *int64
}

func DefaultListOwnedObjectApplicationsOperationOptions() ListOwnedObjectApplicationsOperationOptions {
	return ListOwnedObjectApplicationsOperationOptions{}
}

func (o ListOwnedObjectApplicationsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListOwnedObjectApplicationsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.ConsistencyLevel != nil {
		out.ConsistencyLevel = *o.ConsistencyLevel
	}
	if o.Count != nil {
		out.Count = *o.Count
	}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.OrderBy != nil {
		out.OrderBy = *o.OrderBy
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListOwnedObjectApplicationsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListOwnedObjectApplicationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListOwnedObjectApplicationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListOwnedObjectApplications - Get the items of type microsoft.graph.application in the
// microsoft.graph.directoryObject collection
func (c OwnedObjectClient) ListOwnedObjectApplications(ctx context.Context, options ListOwnedObjectApplicationsOperationOptions) (result ListOwnedObjectApplicationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListOwnedObjectApplicationsCustomPager{},
		Path:          "/me/ownedObjects/application",
		RetryFunc:     options.RetryFunc,
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
		Values *[]stable.Application `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListOwnedObjectApplicationsComplete retrieves all the results into a single object
func (c OwnedObjectClient) ListOwnedObjectApplicationsComplete(ctx context.Context, options ListOwnedObjectApplicationsOperationOptions) (ListOwnedObjectApplicationsCompleteResult, error) {
	return c.ListOwnedObjectApplicationsCompleteMatchingPredicate(ctx, options, ApplicationOperationPredicate{})
}

// ListOwnedObjectApplicationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OwnedObjectClient) ListOwnedObjectApplicationsCompleteMatchingPredicate(ctx context.Context, options ListOwnedObjectApplicationsOperationOptions, predicate ApplicationOperationPredicate) (result ListOwnedObjectApplicationsCompleteResult, err error) {
	items := make([]stable.Application, 0)

	resp, err := c.ListOwnedObjectApplications(ctx, options)
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

	result = ListOwnedObjectApplicationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
