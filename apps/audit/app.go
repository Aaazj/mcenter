package audit

import "fmt"

func NewDescriptAuditRequestWithId(id string) *DescribeAuditRequest {
	return &DescribeAuditRequest{Userid: id}
}

func NewDefaultAudit() *OperateLog {
	return &OperateLog{}
}

// Validate 校验
func (req *DescribeAuditRequest) Validate() error {
	if req.Userid == "" {
		return fmt.Errorf("audit user_id is required")
	}

	return nil
}
