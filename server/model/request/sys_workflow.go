package request

import "server/model"

type WorkflowProcessSearch struct {
	model.WorkflowProcess
	PageInfo
}
