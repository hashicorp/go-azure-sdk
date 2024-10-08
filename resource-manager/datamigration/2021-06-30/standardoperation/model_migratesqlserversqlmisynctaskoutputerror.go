package standardoperation

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MigrateSqlServerSqlMISyncTaskOutput = MigrateSqlServerSqlMISyncTaskOutputError{}

type MigrateSqlServerSqlMISyncTaskOutputError struct {
	Error *ReportableException `json:"error,omitempty"`

	// Fields inherited from MigrateSqlServerSqlMISyncTaskOutput

	Id         *string `json:"id,omitempty"`
	ResultType string  `json:"resultType"`
}

func (s MigrateSqlServerSqlMISyncTaskOutputError) MigrateSqlServerSqlMISyncTaskOutput() BaseMigrateSqlServerSqlMISyncTaskOutputImpl {
	return BaseMigrateSqlServerSqlMISyncTaskOutputImpl{
		Id:         s.Id,
		ResultType: s.ResultType,
	}
}

var _ json.Marshaler = MigrateSqlServerSqlMISyncTaskOutputError{}

func (s MigrateSqlServerSqlMISyncTaskOutputError) MarshalJSON() ([]byte, error) {
	type wrapper MigrateSqlServerSqlMISyncTaskOutputError
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MigrateSqlServerSqlMISyncTaskOutputError: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MigrateSqlServerSqlMISyncTaskOutputError: %+v", err)
	}

	decoded["resultType"] = "ErrorOutput"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MigrateSqlServerSqlMISyncTaskOutputError: %+v", err)
	}

	return encoded, nil
}
