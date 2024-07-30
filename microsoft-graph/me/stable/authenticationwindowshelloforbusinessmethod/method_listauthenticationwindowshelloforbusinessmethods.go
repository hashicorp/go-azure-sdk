package authenticationwindowshelloforbusinessmethod

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

type ListAuthenticationWindowsHelloForBusinessMethodsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.WindowsHelloForBusinessAuthenticationMethod
}

type ListAuthenticationWindowsHelloForBusinessMethodsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.WindowsHelloForBusinessAuthenticationMethod
}

type ListAuthenticationWindowsHelloForBusinessMethodsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAuthenticationWindowsHelloForBusinessMethodsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAuthenticationWindowsHelloForBusinessMethods ...
func (c AuthenticationWindowsHelloForBusinessMethodClient) ListAuthenticationWindowsHelloForBusinessMethods(ctx context.Context) (result ListAuthenticationWindowsHelloForBusinessMethodsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListAuthenticationWindowsHelloForBusinessMethodsCustomPager{},
		Path:       "/me/authentication/windowsHelloForBusinessMethods",
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
		Values *[]stable.WindowsHelloForBusinessAuthenticationMethod `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAuthenticationWindowsHelloForBusinessMethodsComplete retrieves all the results into a single object
func (c AuthenticationWindowsHelloForBusinessMethodClient) ListAuthenticationWindowsHelloForBusinessMethodsComplete(ctx context.Context) (ListAuthenticationWindowsHelloForBusinessMethodsCompleteResult, error) {
	return c.ListAuthenticationWindowsHelloForBusinessMethodsCompleteMatchingPredicate(ctx, WindowsHelloForBusinessAuthenticationMethodOperationPredicate{})
}

// ListAuthenticationWindowsHelloForBusinessMethodsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AuthenticationWindowsHelloForBusinessMethodClient) ListAuthenticationWindowsHelloForBusinessMethodsCompleteMatchingPredicate(ctx context.Context, predicate WindowsHelloForBusinessAuthenticationMethodOperationPredicate) (result ListAuthenticationWindowsHelloForBusinessMethodsCompleteResult, err error) {
	items := make([]stable.WindowsHelloForBusinessAuthenticationMethod, 0)

	resp, err := c.ListAuthenticationWindowsHelloForBusinessMethods(ctx)
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

	result = ListAuthenticationWindowsHelloForBusinessMethodsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
