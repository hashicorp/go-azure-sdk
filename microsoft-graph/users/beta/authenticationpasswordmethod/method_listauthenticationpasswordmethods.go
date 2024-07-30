package authenticationpasswordmethod

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

type ListAuthenticationPasswordMethodsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.PasswordAuthenticationMethod
}

type ListAuthenticationPasswordMethodsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.PasswordAuthenticationMethod
}

type ListAuthenticationPasswordMethodsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAuthenticationPasswordMethodsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAuthenticationPasswordMethods ...
func (c AuthenticationPasswordMethodClient) ListAuthenticationPasswordMethods(ctx context.Context, id UserId) (result ListAuthenticationPasswordMethodsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListAuthenticationPasswordMethodsCustomPager{},
		Path:       fmt.Sprintf("%s/authentication/passwordMethods", id.ID()),
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
		Values *[]beta.PasswordAuthenticationMethod `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAuthenticationPasswordMethodsComplete retrieves all the results into a single object
func (c AuthenticationPasswordMethodClient) ListAuthenticationPasswordMethodsComplete(ctx context.Context, id UserId) (ListAuthenticationPasswordMethodsCompleteResult, error) {
	return c.ListAuthenticationPasswordMethodsCompleteMatchingPredicate(ctx, id, PasswordAuthenticationMethodOperationPredicate{})
}

// ListAuthenticationPasswordMethodsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AuthenticationPasswordMethodClient) ListAuthenticationPasswordMethodsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate PasswordAuthenticationMethodOperationPredicate) (result ListAuthenticationPasswordMethodsCompleteResult, err error) {
	items := make([]beta.PasswordAuthenticationMethod, 0)

	resp, err := c.ListAuthenticationPasswordMethods(ctx, id)
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

	result = ListAuthenticationPasswordMethodsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
