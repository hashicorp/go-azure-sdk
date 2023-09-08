package directoryadministrativeunitmember

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

type ListDirectoryAdministrativeUnitByIdMemberRefsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.StringCollectionResponse
}

type ListDirectoryAdministrativeUnitByIdMemberRefsCompleteResult struct {
	Items []models.StringCollectionResponse
}

// ListDirectoryAdministrativeUnitByIdMemberRefs ...
func (c DirectoryAdministrativeUnitMemberClient) ListDirectoryAdministrativeUnitByIdMemberRefs(ctx context.Context, id DirectoryAdministrativeUnitId) (result ListDirectoryAdministrativeUnitByIdMemberRefsOperationResponse, err error) {
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

// ListDirectoryAdministrativeUnitByIdMemberRefsComplete retrieves all the results into a single object
func (c DirectoryAdministrativeUnitMemberClient) ListDirectoryAdministrativeUnitByIdMemberRefsComplete(ctx context.Context, id DirectoryAdministrativeUnitId) (ListDirectoryAdministrativeUnitByIdMemberRefsCompleteResult, error) {
	return c.ListDirectoryAdministrativeUnitByIdMemberRefsCompleteMatchingPredicate(ctx, id, models.StringCollectionResponseOperationPredicate{})
}

// ListDirectoryAdministrativeUnitByIdMemberRefsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DirectoryAdministrativeUnitMemberClient) ListDirectoryAdministrativeUnitByIdMemberRefsCompleteMatchingPredicate(ctx context.Context, id DirectoryAdministrativeUnitId, predicate models.StringCollectionResponseOperationPredicate) (result ListDirectoryAdministrativeUnitByIdMemberRefsCompleteResult, err error) {
	items := make([]models.StringCollectionResponse, 0)

	resp, err := c.ListDirectoryAdministrativeUnitByIdMemberRefs(ctx, id)
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

	result = ListDirectoryAdministrativeUnitByIdMemberRefsCompleteResult{
		Items: items,
	}
	return
}
