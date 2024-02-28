package externalsecuritysolutions

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByHomeRegionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ExternalSecuritySolution
}

type ListByHomeRegionCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ExternalSecuritySolution
}

// ListByHomeRegion ...
func (c ExternalSecuritySolutionsClient) ListByHomeRegion(ctx context.Context, id LocationId) (result ListByHomeRegionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/externalSecuritySolutions", id.ID()),
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

	temp := make([]ExternalSecuritySolution, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := unmarshalExternalSecuritySolutionImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for ExternalSecuritySolution (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListByHomeRegionComplete retrieves all the results into a single object
func (c ExternalSecuritySolutionsClient) ListByHomeRegionComplete(ctx context.Context, id LocationId) (ListByHomeRegionCompleteResult, error) {
	return c.ListByHomeRegionCompleteMatchingPredicate(ctx, id, ExternalSecuritySolutionOperationPredicate{})
}

// ListByHomeRegionCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ExternalSecuritySolutionsClient) ListByHomeRegionCompleteMatchingPredicate(ctx context.Context, id LocationId, predicate ExternalSecuritySolutionOperationPredicate) (result ListByHomeRegionCompleteResult, err error) {
	items := make([]ExternalSecuritySolution, 0)

	resp, err := c.ListByHomeRegion(ctx, id)
	if err != nil {
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

	result = ListByHomeRegionCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
