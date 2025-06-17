package authenticationeventsflowexternalusersselfservicesignupeventsflowonattributecollectiononattributecollectionexternalusersselfservicesignupattribute

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

type ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAttributeCollectionExternalAttributesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.IdentityUserFlowAttribute
}

type ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAttributeCollectionExternalAttributesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.IdentityUserFlowAttribute
}

type ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAttributeCollectionExternalAttributesOperationOptions struct {
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

func DefaultListAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAttributeCollectionExternalAttributesOperationOptions() ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAttributeCollectionExternalAttributesOperationOptions {
	return ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAttributeCollectionExternalAttributesOperationOptions{}
}

func (o ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAttributeCollectionExternalAttributesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAttributeCollectionExternalAttributesOperationOptions) ToOData() *odata.Query {
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

func (o ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAttributeCollectionExternalAttributesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAttributeCollectionExternalAttributesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAttributeCollectionExternalAttributesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAttributeCollectionExternalAttributes - List attributes
// (of a user flow). Get an identityUserFlowAttribute collection associated with an external identities self-service
// user flow represented by an externalUsersSelfServiceSignupEventsFlow object. These attributes are collected from the
// user during the authentication experience defined by the user flow.
func (c AuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionOnAttributeCollectionExternalUsersSelfServiceSignUpAttributeClient) ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAttributeCollectionExternalAttributes(ctx context.Context, id stable.IdentityAuthenticationEventsFlowId, options ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAttributeCollectionExternalAttributesOperationOptions) (result ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAttributeCollectionExternalAttributesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAttributeCollectionExternalAttributesCustomPager{},
		Path:          fmt.Sprintf("%s/externalUsersSelfServiceSignUpEventsFlow/onAttributeCollection/onAttributeCollectionExternalUsersSelfServiceSignUp/attributes", id.ID()),
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

	temp := make([]stable.IdentityUserFlowAttribute, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := stable.UnmarshalIdentityUserFlowAttributeImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for stable.IdentityUserFlowAttribute (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAttributeCollectionExternalAttributesComplete retrieves all the results into a single object
func (c AuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionOnAttributeCollectionExternalUsersSelfServiceSignUpAttributeClient) ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAttributeCollectionExternalAttributesComplete(ctx context.Context, id stable.IdentityAuthenticationEventsFlowId, options ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAttributeCollectionExternalAttributesOperationOptions) (ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAttributeCollectionExternalAttributesCompleteResult, error) {
	return c.ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAttributeCollectionExternalAttributesCompleteMatchingPredicate(ctx, id, options, IdentityUserFlowAttributeOperationPredicate{})
}

// ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAttributeCollectionExternalAttributesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionOnAttributeCollectionExternalUsersSelfServiceSignUpAttributeClient) ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAttributeCollectionExternalAttributesCompleteMatchingPredicate(ctx context.Context, id stable.IdentityAuthenticationEventsFlowId, options ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAttributeCollectionExternalAttributesOperationOptions, predicate IdentityUserFlowAttributeOperationPredicate) (result ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAttributeCollectionExternalAttributesCompleteResult, err error) {
	items := make([]stable.IdentityUserFlowAttribute, 0)

	resp, err := c.ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAttributeCollectionExternalAttributes(ctx, id, options)
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

	result = ListAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAttributeCollectionExternalAttributesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
