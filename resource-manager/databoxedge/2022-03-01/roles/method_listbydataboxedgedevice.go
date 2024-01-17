package roles

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

type ListByDataBoxEdgeDeviceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Role
}

type ListByDataBoxEdgeDeviceCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Role
}

// ListByDataBoxEdgeDevice ...
func (c RolesClient) ListByDataBoxEdgeDevice(ctx context.Context, id DataBoxEdgeDeviceId) (result ListByDataBoxEdgeDeviceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/roles", id.ID()),
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

	temp := make([]Role, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := unmarshalRoleImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for Role (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListByDataBoxEdgeDeviceComplete retrieves all the results into a single object
func (c RolesClient) ListByDataBoxEdgeDeviceComplete(ctx context.Context, id DataBoxEdgeDeviceId) (ListByDataBoxEdgeDeviceCompleteResult, error) {
	return c.ListByDataBoxEdgeDeviceCompleteMatchingPredicate(ctx, id, RoleOperationPredicate{})
}

// ListByDataBoxEdgeDeviceCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RolesClient) ListByDataBoxEdgeDeviceCompleteMatchingPredicate(ctx context.Context, id DataBoxEdgeDeviceId, predicate RoleOperationPredicate) (result ListByDataBoxEdgeDeviceCompleteResult, err error) {
	items := make([]Role, 0)

	resp, err := c.ListByDataBoxEdgeDevice(ctx, id)
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

	result = ListByDataBoxEdgeDeviceCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
