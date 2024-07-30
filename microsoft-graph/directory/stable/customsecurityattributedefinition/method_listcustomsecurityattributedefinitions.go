package customsecurityattributedefinition

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

type ListCustomSecurityAttributeDefinitionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.CustomSecurityAttributeDefinition
}

type ListCustomSecurityAttributeDefinitionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.CustomSecurityAttributeDefinition
}

type ListCustomSecurityAttributeDefinitionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCustomSecurityAttributeDefinitionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCustomSecurityAttributeDefinitions ...
func (c CustomSecurityAttributeDefinitionClient) ListCustomSecurityAttributeDefinitions(ctx context.Context) (result ListCustomSecurityAttributeDefinitionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListCustomSecurityAttributeDefinitionsCustomPager{},
		Path:       "/directory/customSecurityAttributeDefinitions",
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
		Values *[]stable.CustomSecurityAttributeDefinition `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListCustomSecurityAttributeDefinitionsComplete retrieves all the results into a single object
func (c CustomSecurityAttributeDefinitionClient) ListCustomSecurityAttributeDefinitionsComplete(ctx context.Context) (ListCustomSecurityAttributeDefinitionsCompleteResult, error) {
	return c.ListCustomSecurityAttributeDefinitionsCompleteMatchingPredicate(ctx, CustomSecurityAttributeDefinitionOperationPredicate{})
}

// ListCustomSecurityAttributeDefinitionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CustomSecurityAttributeDefinitionClient) ListCustomSecurityAttributeDefinitionsCompleteMatchingPredicate(ctx context.Context, predicate CustomSecurityAttributeDefinitionOperationPredicate) (result ListCustomSecurityAttributeDefinitionsCompleteResult, err error) {
	items := make([]stable.CustomSecurityAttributeDefinition, 0)

	resp, err := c.ListCustomSecurityAttributeDefinitions(ctx)
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

	result = ListCustomSecurityAttributeDefinitionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
