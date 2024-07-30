package serviceprincipal

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

type GetServicePrincipalsAvailableExtensionPropertiesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ExtensionProperty
}

type GetServicePrincipalsAvailableExtensionPropertiesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ExtensionProperty
}

type GetServicePrincipalsAvailableExtensionPropertiesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *GetServicePrincipalsAvailableExtensionPropertiesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// GetServicePrincipalsAvailableExtensionProperties ...
func (c ServicePrincipalClient) GetServicePrincipalsAvailableExtensionProperties(ctx context.Context, input GetServicePrincipalsAvailableExtensionPropertiesRequest) (result GetServicePrincipalsAvailableExtensionPropertiesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Pager:      &GetServicePrincipalsAvailableExtensionPropertiesCustomPager{},
		Path:       "/servicePrincipals/getAvailableExtensionProperties",
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
		Values *[]stable.ExtensionProperty `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// GetServicePrincipalsAvailableExtensionPropertiesComplete retrieves all the results into a single object
func (c ServicePrincipalClient) GetServicePrincipalsAvailableExtensionPropertiesComplete(ctx context.Context, input GetServicePrincipalsAvailableExtensionPropertiesRequest) (GetServicePrincipalsAvailableExtensionPropertiesCompleteResult, error) {
	return c.GetServicePrincipalsAvailableExtensionPropertiesCompleteMatchingPredicate(ctx, input, ExtensionPropertyOperationPredicate{})
}

// GetServicePrincipalsAvailableExtensionPropertiesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ServicePrincipalClient) GetServicePrincipalsAvailableExtensionPropertiesCompleteMatchingPredicate(ctx context.Context, input GetServicePrincipalsAvailableExtensionPropertiesRequest, predicate ExtensionPropertyOperationPredicate) (result GetServicePrincipalsAvailableExtensionPropertiesCompleteResult, err error) {
	items := make([]stable.ExtensionProperty, 0)

	resp, err := c.GetServicePrincipalsAvailableExtensionProperties(ctx, input)
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

	result = GetServicePrincipalsAvailableExtensionPropertiesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
