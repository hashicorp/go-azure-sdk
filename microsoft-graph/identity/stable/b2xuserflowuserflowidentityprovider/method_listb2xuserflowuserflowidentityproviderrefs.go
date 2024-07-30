package b2xuserflowuserflowidentityprovider

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListB2xUserFlowUserFlowIdentityProviderRefsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]string
}

type ListB2xUserFlowUserFlowIdentityProviderRefsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []string
}

type ListB2xUserFlowUserFlowIdentityProviderRefsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListB2xUserFlowUserFlowIdentityProviderRefsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListB2xUserFlowUserFlowIdentityProviderRefs ...
func (c B2xUserFlowUserFlowIdentityProviderClient) ListB2xUserFlowUserFlowIdentityProviderRefs(ctx context.Context, id IdentityB2xUserFlowId) (result ListB2xUserFlowUserFlowIdentityProviderRefsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListB2xUserFlowUserFlowIdentityProviderRefsCustomPager{},
		Path:       fmt.Sprintf("%s/userFlowIdentityProviders/$ref", id.ID()),
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
		Values *[]string `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListB2xUserFlowUserFlowIdentityProviderRefsComplete retrieves all the results into a single object
func (c B2xUserFlowUserFlowIdentityProviderClient) ListB2xUserFlowUserFlowIdentityProviderRefsComplete(ctx context.Context, id IdentityB2xUserFlowId) (result ListB2xUserFlowUserFlowIdentityProviderRefsCompleteResult, err error) {
	items := make([]string, 0)

	resp, err := c.ListB2xUserFlowUserFlowIdentityProviderRefs(ctx, id)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			items = append(items, v)
		}
	}

	result = ListB2xUserFlowUserFlowIdentityProviderRefsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
