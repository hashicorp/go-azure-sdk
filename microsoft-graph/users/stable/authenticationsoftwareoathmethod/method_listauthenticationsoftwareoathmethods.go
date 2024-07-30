package authenticationsoftwareoathmethod

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

type ListAuthenticationSoftwareOathMethodsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.SoftwareOathAuthenticationMethod
}

type ListAuthenticationSoftwareOathMethodsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.SoftwareOathAuthenticationMethod
}

type ListAuthenticationSoftwareOathMethodsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAuthenticationSoftwareOathMethodsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAuthenticationSoftwareOathMethods ...
func (c AuthenticationSoftwareOathMethodClient) ListAuthenticationSoftwareOathMethods(ctx context.Context, id UserId) (result ListAuthenticationSoftwareOathMethodsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListAuthenticationSoftwareOathMethodsCustomPager{},
		Path:       fmt.Sprintf("%s/authentication/softwareOathMethods", id.ID()),
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
		Values *[]stable.SoftwareOathAuthenticationMethod `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAuthenticationSoftwareOathMethodsComplete retrieves all the results into a single object
func (c AuthenticationSoftwareOathMethodClient) ListAuthenticationSoftwareOathMethodsComplete(ctx context.Context, id UserId) (ListAuthenticationSoftwareOathMethodsCompleteResult, error) {
	return c.ListAuthenticationSoftwareOathMethodsCompleteMatchingPredicate(ctx, id, SoftwareOathAuthenticationMethodOperationPredicate{})
}

// ListAuthenticationSoftwareOathMethodsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AuthenticationSoftwareOathMethodClient) ListAuthenticationSoftwareOathMethodsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate SoftwareOathAuthenticationMethodOperationPredicate) (result ListAuthenticationSoftwareOathMethodsCompleteResult, err error) {
	items := make([]stable.SoftwareOathAuthenticationMethod, 0)

	resp, err := c.ListAuthenticationSoftwareOathMethods(ctx, id)
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

	result = ListAuthenticationSoftwareOathMethodsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
