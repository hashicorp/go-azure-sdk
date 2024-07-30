package authenticationmethoduserregistrationdetail

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListAuthenticationMethodUserRegistrationDetailsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UserRegistrationDetails
}

type ListAuthenticationMethodUserRegistrationDetailsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UserRegistrationDetails
}

type ListAuthenticationMethodUserRegistrationDetailsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAuthenticationMethodUserRegistrationDetailsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAuthenticationMethodUserRegistrationDetails ...
func (c AuthenticationMethodUserRegistrationDetailClient) ListAuthenticationMethodUserRegistrationDetails(ctx context.Context) (result ListAuthenticationMethodUserRegistrationDetailsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListAuthenticationMethodUserRegistrationDetailsCustomPager{},
		Path:       "/reports/authenticationMethods/userRegistrationDetails",
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
		Values *[]beta.UserRegistrationDetails `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAuthenticationMethodUserRegistrationDetailsComplete retrieves all the results into a single object
func (c AuthenticationMethodUserRegistrationDetailClient) ListAuthenticationMethodUserRegistrationDetailsComplete(ctx context.Context) (ListAuthenticationMethodUserRegistrationDetailsCompleteResult, error) {
	return c.ListAuthenticationMethodUserRegistrationDetailsCompleteMatchingPredicate(ctx, UserRegistrationDetailsOperationPredicate{})
}

// ListAuthenticationMethodUserRegistrationDetailsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AuthenticationMethodUserRegistrationDetailClient) ListAuthenticationMethodUserRegistrationDetailsCompleteMatchingPredicate(ctx context.Context, predicate UserRegistrationDetailsOperationPredicate) (result ListAuthenticationMethodUserRegistrationDetailsCompleteResult, err error) {
	items := make([]beta.UserRegistrationDetails, 0)

	resp, err := c.ListAuthenticationMethodUserRegistrationDetails(ctx)
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

	result = ListAuthenticationMethodUserRegistrationDetailsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
