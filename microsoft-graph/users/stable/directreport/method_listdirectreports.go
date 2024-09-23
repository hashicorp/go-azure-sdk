package directreport

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListDirectReportsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.DirectoryObject
}

type ListDirectReportsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.DirectoryObject
}

type ListDirectReportsOperationOptions struct {
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

func DefaultListDirectReportsOperationOptions() ListDirectReportsOperationOptions {
	return ListDirectReportsOperationOptions{}
}

func (o ListDirectReportsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDirectReportsOperationOptions) ToOData() *odata.Query {
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

func (o ListDirectReportsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDirectReportsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDirectReportsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDirectReports - Get directReports from users. The users and contacts that report to the user. (The users and
// contacts that have their manager property set to this user.) Read-only. Nullable. Supports $expand.
func (c DirectReportClient) ListDirectReports(ctx context.Context, id stable.UserId, options ListDirectReportsOperationOptions) (result ListDirectReportsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDirectReportsCustomPager{},
		Path:          fmt.Sprintf("%s/directReports", id.ID()),
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
		Values *[]json.RawMessage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	temp := make([]stable.DirectoryObject, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := stable.UnmarshalDirectoryObjectImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for stable.DirectoryObject (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListDirectReportsComplete retrieves all the results into a single object
func (c DirectReportClient) ListDirectReportsComplete(ctx context.Context, id stable.UserId, options ListDirectReportsOperationOptions) (ListDirectReportsCompleteResult, error) {
	return c.ListDirectReportsCompleteMatchingPredicate(ctx, id, options, DirectoryObjectOperationPredicate{})
}

// ListDirectReportsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DirectReportClient) ListDirectReportsCompleteMatchingPredicate(ctx context.Context, id stable.UserId, options ListDirectReportsOperationOptions, predicate DirectoryObjectOperationPredicate) (result ListDirectReportsCompleteResult, err error) {
	items := make([]stable.DirectoryObject, 0)

	resp, err := c.ListDirectReports(ctx, id, options)
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

	result = ListDirectReportsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
