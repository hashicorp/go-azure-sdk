package authenticationemailmethod

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

type ListAuthenticationEmailMethodsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.EmailAuthenticationMethod
}

type ListAuthenticationEmailMethodsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.EmailAuthenticationMethod
}

type ListAuthenticationEmailMethodsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAuthenticationEmailMethodsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAuthenticationEmailMethods ...
func (c AuthenticationEmailMethodClient) ListAuthenticationEmailMethods(ctx context.Context, id UserId) (result ListAuthenticationEmailMethodsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListAuthenticationEmailMethodsCustomPager{},
		Path:       fmt.Sprintf("%s/authentication/emailMethods", id.ID()),
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
		Values *[]beta.EmailAuthenticationMethod `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAuthenticationEmailMethodsComplete retrieves all the results into a single object
func (c AuthenticationEmailMethodClient) ListAuthenticationEmailMethodsComplete(ctx context.Context, id UserId) (ListAuthenticationEmailMethodsCompleteResult, error) {
	return c.ListAuthenticationEmailMethodsCompleteMatchingPredicate(ctx, id, EmailAuthenticationMethodOperationPredicate{})
}

// ListAuthenticationEmailMethodsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AuthenticationEmailMethodClient) ListAuthenticationEmailMethodsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate EmailAuthenticationMethodOperationPredicate) (result ListAuthenticationEmailMethodsCompleteResult, err error) {
	items := make([]beta.EmailAuthenticationMethod, 0)

	resp, err := c.ListAuthenticationEmailMethods(ctx, id)
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

	result = ListAuthenticationEmailMethodsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
