package administrativeunitmember

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/beta/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListAdministrativeUnitByIdMemberRefsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.StringCollectionResponse
}

type ListAdministrativeUnitByIdMemberRefsCompleteResult struct {
	Items []models.StringCollectionResponse
}

// ListAdministrativeUnitByIdMemberRefs ...
func (c AdministrativeUnitMemberClient) ListAdministrativeUnitByIdMemberRefs(ctx context.Context, id AdministrativeUnitId) (result ListAdministrativeUnitByIdMemberRefsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/members/$ref", id.ID()),
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
		Values *[]models.StringCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAdministrativeUnitByIdMemberRefsComplete retrieves all the results into a single object
func (c AdministrativeUnitMemberClient) ListAdministrativeUnitByIdMemberRefsComplete(ctx context.Context, id AdministrativeUnitId) (ListAdministrativeUnitByIdMemberRefsCompleteResult, error) {
	return c.ListAdministrativeUnitByIdMemberRefsCompleteMatchingPredicate(ctx, id, models.StringCollectionResponseOperationPredicate{})
}

// ListAdministrativeUnitByIdMemberRefsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AdministrativeUnitMemberClient) ListAdministrativeUnitByIdMemberRefsCompleteMatchingPredicate(ctx context.Context, id AdministrativeUnitId, predicate models.StringCollectionResponseOperationPredicate) (result ListAdministrativeUnitByIdMemberRefsCompleteResult, err error) {
	items := make([]models.StringCollectionResponse, 0)

	resp, err := c.ListAdministrativeUnitByIdMemberRefs(ctx, id)
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

	result = ListAdministrativeUnitByIdMemberRefsCompleteResult{
		Items: items,
	}
	return
}
