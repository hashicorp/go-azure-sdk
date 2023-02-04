package locationbasedcapabilities

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExecuteOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]CapabilityProperties
}

type ExecuteCompleteResult struct {
	Items []CapabilityProperties
}

// Execute ...
func (c LocationBasedCapabilitiesClient) Execute(ctx context.Context, id LocationId) (result ExecuteOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/capabilities", id.ID()),
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
		Values *[]CapabilityProperties `json:"values"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ExecuteComplete retrieves all the results into a single object
func (c LocationBasedCapabilitiesClient) ExecuteComplete(ctx context.Context, id LocationId) (ExecuteCompleteResult, error) {
	return c.ExecuteCompleteMatchingPredicate(ctx, id, CapabilityPropertiesOperationPredicate{})
}

// ExecuteCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c LocationBasedCapabilitiesClient) ExecuteCompleteMatchingPredicate(ctx context.Context, id LocationId, predicate CapabilityPropertiesOperationPredicate) (result ExecuteCompleteResult, err error) {
	items := make([]CapabilityProperties, 0)

	resp, err := c.Execute(ctx, id)
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

	result = ExecuteCompleteResult{
		Items: items,
	}
	return
}
