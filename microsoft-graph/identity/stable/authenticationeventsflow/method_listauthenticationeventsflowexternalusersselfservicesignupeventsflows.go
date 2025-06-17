package authenticationeventsflow

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

type ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ExternalUsersSelfServiceSignUpEventsFlow
}

type ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ExternalUsersSelfServiceSignUpEventsFlow
}

type ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowsOperationOptions struct {
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

func DefaultListAuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowsOperationOptions() ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowsOperationOptions {
	return ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowsOperationOptions{}
}

func (o ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowsOperationOptions) ToOData() *odata.Query {
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

func (o ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlows - Get the items of type
// microsoft.graph.externalUsersSelfServiceSignUpEventsFlow in the microsoft.graph.authenticationEventsFlow collection
func (c AuthenticationEventsFlowClient) ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlows(ctx context.Context, options ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowsOperationOptions) (result ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowsCustomPager{},
		Path:          "/identity/authenticationEventsFlows/externalUsersSelfServiceSignUpEventsFlow",
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
		Values *[]stable.ExternalUsersSelfServiceSignUpEventsFlow `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowsComplete retrieves all the results into a single object
func (c AuthenticationEventsFlowClient) ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowsComplete(ctx context.Context, options ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowsOperationOptions) (ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowsCompleteResult, error) {
	return c.ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowsCompleteMatchingPredicate(ctx, options, ExternalUsersSelfServiceSignUpEventsFlowOperationPredicate{})
}

// ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AuthenticationEventsFlowClient) ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowsCompleteMatchingPredicate(ctx context.Context, options ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowsOperationOptions, predicate ExternalUsersSelfServiceSignUpEventsFlowOperationPredicate) (result ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowsCompleteResult, err error) {
	items := make([]stable.ExternalUsersSelfServiceSignUpEventsFlow, 0)

	resp, err := c.ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlows(ctx, options)
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

	result = ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
