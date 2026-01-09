package authenticationeventsflowexternalusersselfservicesignupeventsflowonauthenticationmethodloadstartonauthenticationmethodloadstartexternalusersselfservicesignupidentityprovider

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

type ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAuthenticationMethodLoadStartExternalIdentityProviderRefsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.DirectoryObject
}

type ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAuthenticationMethodLoadStartExternalIdentityProviderRefsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.DirectoryObject
}

type ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAuthenticationMethodLoadStartExternalIdentityProviderRefsOperationOptions struct {
	Count     *bool
	Filter    *string
	Metadata  *odata.Metadata
	OrderBy   *odata.OrderBy
	RetryFunc client.RequestRetryFunc
	Search    *string
	Skip      *int64
	Top       *int64
}

func DefaultListAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAuthenticationMethodLoadStartExternalIdentityProviderRefsOperationOptions() ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAuthenticationMethodLoadStartExternalIdentityProviderRefsOperationOptions {
	return ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAuthenticationMethodLoadStartExternalIdentityProviderRefsOperationOptions{}
}

func (o ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAuthenticationMethodLoadStartExternalIdentityProviderRefsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAuthenticationMethodLoadStartExternalIdentityProviderRefsOperationOptions) ToOData() *odata.Query {
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

func (o ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAuthenticationMethodLoadStartExternalIdentityProviderRefsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAuthenticationMethodLoadStartExternalIdentityProviderRefsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAuthenticationMethodLoadStartExternalIdentityProviderRefsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAuthenticationMethodLoadStartExternalIdentityProviderRefs
// - List identityProviders (in a user flow). Get the identity providers that are defined for an external identities
// self-service sign up user flow that's represented by an externalUsersSelfServiceSignupEventsFlow object type.
func (c AuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpIdentityProviderClient) ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAuthenticationMethodLoadStartExternalIdentityProviderRefs(ctx context.Context, id stable.IdentityAuthenticationEventsFlowId, options ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAuthenticationMethodLoadStartExternalIdentityProviderRefsOperationOptions) (result ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAuthenticationMethodLoadStartExternalIdentityProviderRefsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAuthenticationMethodLoadStartExternalIdentityProviderRefsCustomPager{},
		Path:          fmt.Sprintf("%s/externalUsersSelfServiceSignUpEventsFlow/onAuthenticationMethodLoadStart/onAuthenticationMethodLoadStartExternalUsersSelfServiceSignUp/identityProviders/$ref", id.ID()),
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

// ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAuthenticationMethodLoadStartExternalIdentityProviderRefsComplete retrieves all the results into a single object
func (c AuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpIdentityProviderClient) ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAuthenticationMethodLoadStartExternalIdentityProviderRefsComplete(ctx context.Context, id stable.IdentityAuthenticationEventsFlowId, options ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAuthenticationMethodLoadStartExternalIdentityProviderRefsOperationOptions) (ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAuthenticationMethodLoadStartExternalIdentityProviderRefsCompleteResult, error) {
	return c.ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAuthenticationMethodLoadStartExternalIdentityProviderRefsCompleteMatchingPredicate(ctx, id, options, DirectoryObjectOperationPredicate{})
}

// ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAuthenticationMethodLoadStartExternalIdentityProviderRefsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartOnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUpIdentityProviderClient) ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAuthenticationMethodLoadStartExternalIdentityProviderRefsCompleteMatchingPredicate(ctx context.Context, id stable.IdentityAuthenticationEventsFlowId, options ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAuthenticationMethodLoadStartExternalIdentityProviderRefsOperationOptions, predicate DirectoryObjectOperationPredicate) (result ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAuthenticationMethodLoadStartExternalIdentityProviderRefsCompleteResult, err error) {
	items := make([]stable.DirectoryObject, 0)

	resp, err := c.ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAuthenticationMethodLoadStartExternalIdentityProviderRefs(ctx, id, options)
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

	result = ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAuthenticationMethodLoadStartExternalIdentityProviderRefsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
