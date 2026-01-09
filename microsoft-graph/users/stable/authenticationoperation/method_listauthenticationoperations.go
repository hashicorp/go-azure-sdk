package authenticationoperation

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

type ListAuthenticationOperationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.LongRunningOperation
}

type ListAuthenticationOperationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.LongRunningOperation
}

type ListAuthenticationOperationsOperationOptions struct {
	Count     *bool
	Expand    *odata.Expand
	Filter    *string
	Metadata  *odata.Metadata
	OrderBy   *odata.OrderBy
	RetryFunc client.RequestRetryFunc
	Search    *string
	Select    *[]string
	Skip      *int64
	Top       *int64
}

func DefaultListAuthenticationOperationsOperationOptions() ListAuthenticationOperationsOperationOptions {
	return ListAuthenticationOperationsOperationOptions{}
}

func (o ListAuthenticationOperationsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListAuthenticationOperationsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
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

func (o ListAuthenticationOperationsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListAuthenticationOperationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAuthenticationOperationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAuthenticationOperations - Get longRunningOperation. Read the properties and relationships of a
// longRunningOperation object. This API allows you to retrieve the details and status of the following long-running
// Microsoft Graph API operations. The possible states of the long-running operation are notStarted, running, succeeded,
// failed, unknownFutureValue where succeeded and failed are terminal states.
func (c AuthenticationOperationClient) ListAuthenticationOperations(ctx context.Context, id stable.UserId, options ListAuthenticationOperationsOperationOptions) (result ListAuthenticationOperationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListAuthenticationOperationsCustomPager{},
		Path:          fmt.Sprintf("%s/authentication/operations", id.ID()),
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

	temp := make([]stable.LongRunningOperation, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := stable.UnmarshalLongRunningOperationImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for stable.LongRunningOperation (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListAuthenticationOperationsComplete retrieves all the results into a single object
func (c AuthenticationOperationClient) ListAuthenticationOperationsComplete(ctx context.Context, id stable.UserId, options ListAuthenticationOperationsOperationOptions) (ListAuthenticationOperationsCompleteResult, error) {
	return c.ListAuthenticationOperationsCompleteMatchingPredicate(ctx, id, options, LongRunningOperationOperationPredicate{})
}

// ListAuthenticationOperationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AuthenticationOperationClient) ListAuthenticationOperationsCompleteMatchingPredicate(ctx context.Context, id stable.UserId, options ListAuthenticationOperationsOperationOptions, predicate LongRunningOperationOperationPredicate) (result ListAuthenticationOperationsCompleteResult, err error) {
	items := make([]stable.LongRunningOperation, 0)

	resp, err := c.ListAuthenticationOperations(ctx, id, options)
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

	result = ListAuthenticationOperationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
