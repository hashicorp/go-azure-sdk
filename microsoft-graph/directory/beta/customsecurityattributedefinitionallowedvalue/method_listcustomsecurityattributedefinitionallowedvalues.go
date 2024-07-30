package customsecurityattributedefinitionallowedvalue

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

type ListCustomSecurityAttributeDefinitionAllowedValuesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AllowedValue
}

type ListCustomSecurityAttributeDefinitionAllowedValuesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AllowedValue
}

type ListCustomSecurityAttributeDefinitionAllowedValuesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCustomSecurityAttributeDefinitionAllowedValuesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCustomSecurityAttributeDefinitionAllowedValues ...
func (c CustomSecurityAttributeDefinitionAllowedValueClient) ListCustomSecurityAttributeDefinitionAllowedValues(ctx context.Context, id DirectoryCustomSecurityAttributeDefinitionId) (result ListCustomSecurityAttributeDefinitionAllowedValuesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListCustomSecurityAttributeDefinitionAllowedValuesCustomPager{},
		Path:       fmt.Sprintf("%s/allowedValues", id.ID()),
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
		Values *[]beta.AllowedValue `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListCustomSecurityAttributeDefinitionAllowedValuesComplete retrieves all the results into a single object
func (c CustomSecurityAttributeDefinitionAllowedValueClient) ListCustomSecurityAttributeDefinitionAllowedValuesComplete(ctx context.Context, id DirectoryCustomSecurityAttributeDefinitionId) (ListCustomSecurityAttributeDefinitionAllowedValuesCompleteResult, error) {
	return c.ListCustomSecurityAttributeDefinitionAllowedValuesCompleteMatchingPredicate(ctx, id, AllowedValueOperationPredicate{})
}

// ListCustomSecurityAttributeDefinitionAllowedValuesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CustomSecurityAttributeDefinitionAllowedValueClient) ListCustomSecurityAttributeDefinitionAllowedValuesCompleteMatchingPredicate(ctx context.Context, id DirectoryCustomSecurityAttributeDefinitionId, predicate AllowedValueOperationPredicate) (result ListCustomSecurityAttributeDefinitionAllowedValuesCompleteResult, err error) {
	items := make([]beta.AllowedValue, 0)

	resp, err := c.ListCustomSecurityAttributeDefinitionAllowedValues(ctx, id)
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

	result = ListCustomSecurityAttributeDefinitionAllowedValuesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
