package tokenlifetimepolicy

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListTokenLifetimePolicyRefsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.DirectoryObject
}

type ListTokenLifetimePolicyRefsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.DirectoryObject
}

type ListTokenLifetimePolicyRefsOperationOptions struct {
	Count     *bool
	Filter    *string
	Metadata  *odata.Metadata
	OrderBy   *odata.OrderBy
	RetryFunc client.RequestRetryFunc
	Search    *string
	Skip      *int64
	Top       *int64
}

func DefaultListTokenLifetimePolicyRefsOperationOptions() ListTokenLifetimePolicyRefsOperationOptions {
	return ListTokenLifetimePolicyRefsOperationOptions{}
}

func (o ListTokenLifetimePolicyRefsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListTokenLifetimePolicyRefsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Count != nil {
		out.Count = *o.Count
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
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListTokenLifetimePolicyRefsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListTokenLifetimePolicyRefsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTokenLifetimePolicyRefsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTokenLifetimePolicyRefs - List assigned tokenLifetimePolicies. List the tokenLifetimePolicy objects that are
// assigned to an application. Only one object is returned in the collection because only one tokenLifetimePolicy can be
// assigned to an application.
func (c TokenLifetimePolicyClient) ListTokenLifetimePolicyRefs(ctx context.Context, id stable.ApplicationId, options ListTokenLifetimePolicyRefsOperationOptions) (result ListTokenLifetimePolicyRefsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListTokenLifetimePolicyRefsCustomPager{},
		Path:          fmt.Sprintf("%s/tokenLifetimePolicies/$ref", id.ID()),
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

// ListTokenLifetimePolicyRefsComplete retrieves all the results into a single object
func (c TokenLifetimePolicyClient) ListTokenLifetimePolicyRefsComplete(ctx context.Context, id stable.ApplicationId, options ListTokenLifetimePolicyRefsOperationOptions) (ListTokenLifetimePolicyRefsCompleteResult, error) {
	return c.ListTokenLifetimePolicyRefsCompleteMatchingPredicate(ctx, id, options, DirectoryObjectOperationPredicate{})
}

// ListTokenLifetimePolicyRefsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TokenLifetimePolicyClient) ListTokenLifetimePolicyRefsCompleteMatchingPredicate(ctx context.Context, id stable.ApplicationId, options ListTokenLifetimePolicyRefsOperationOptions, predicate DirectoryObjectOperationPredicate) (result ListTokenLifetimePolicyRefsCompleteResult, err error) {
	items := make([]stable.DirectoryObject, 0)

	resp, err := c.ListTokenLifetimePolicyRefs(ctx, id, options)
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

	result = ListTokenLifetimePolicyRefsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
