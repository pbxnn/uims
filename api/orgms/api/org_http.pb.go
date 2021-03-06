// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.4

package api

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type CompanyHTTPServer interface {
	BatchCreateCompany(context.Context, *CreateCompanyReq) (*BatchCreateCompanyReply, error)
	CreateCompany(context.Context, *CreateCompanyReq) (*CreateCompanyReply, error)
	DeleteCompany(context.Context, *DelCompanyReq) (*DelCompanyReply, error)
	GetCompany(context.Context, *GetCompanyReq) (*GetCompanyReply, error)
	GetCompanyList(context.Context, *GetCompanyListReq) (*GetCompanyListReply, error)
	OrderCompany(context.Context, *OrderCompanyReq) (*OrderCompanyReply, error)
	UpdateCompany(context.Context, *UpdateCompanyReq) (*UpdateCompanyReply, error)
}

func RegisterCompanyHTTPServer(s *http.Server, srv CompanyHTTPServer) {
	r := s.Route("/")
	r.POST("/uims/org/company", _Company_CreateCompany0_HTTP_Handler(srv))
	r.POST("/uims/org/company/batch", _Company_BatchCreateCompany0_HTTP_Handler(srv))
	r.GET("/uims/org/company", _Company_GetCompanyList0_HTTP_Handler(srv))
	r.GET("/uims/org/company/{company_id}", _Company_GetCompany0_HTTP_Handler(srv))
	r.PUT("/uims/org/company", _Company_UpdateCompany0_HTTP_Handler(srv))
	r.DELETE("/uims/org/company/{company_id}", _Company_DeleteCompany0_HTTP_Handler(srv))
	r.PATCH("/uims/org/company", _Company_OrderCompany0_HTTP_Handler(srv))
}

func _Company_CreateCompany0_HTTP_Handler(srv CompanyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateCompanyReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/uims.orgms.api.Company/CreateCompany")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateCompany(ctx, req.(*CreateCompanyReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateCompanyReply)
		return ctx.Result(200, reply)
	}
}

func _Company_BatchCreateCompany0_HTTP_Handler(srv CompanyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateCompanyReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/uims.orgms.api.Company/BatchCreateCompany")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.BatchCreateCompany(ctx, req.(*CreateCompanyReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*BatchCreateCompanyReply)
		return ctx.Result(200, reply)
	}
}

func _Company_GetCompanyList0_HTTP_Handler(srv CompanyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetCompanyListReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/uims.orgms.api.Company/GetCompanyList")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetCompanyList(ctx, req.(*GetCompanyListReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetCompanyListReply)
		return ctx.Result(200, reply)
	}
}

func _Company_GetCompany0_HTTP_Handler(srv CompanyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetCompanyReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/uims.orgms.api.Company/GetCompany")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetCompany(ctx, req.(*GetCompanyReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetCompanyReply)
		return ctx.Result(200, reply)
	}
}

func _Company_UpdateCompany0_HTTP_Handler(srv CompanyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateCompanyReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/uims.orgms.api.Company/UpdateCompany")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateCompany(ctx, req.(*UpdateCompanyReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateCompanyReply)
		return ctx.Result(200, reply)
	}
}

func _Company_DeleteCompany0_HTTP_Handler(srv CompanyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DelCompanyReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/uims.orgms.api.Company/DeleteCompany")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteCompany(ctx, req.(*DelCompanyReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DelCompanyReply)
		return ctx.Result(200, reply)
	}
}

func _Company_OrderCompany0_HTTP_Handler(srv CompanyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in OrderCompanyReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/uims.orgms.api.Company/OrderCompany")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.OrderCompany(ctx, req.(*OrderCompanyReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*OrderCompanyReply)
		return ctx.Result(200, reply)
	}
}

type CompanyHTTPClient interface {
	BatchCreateCompany(ctx context.Context, req *CreateCompanyReq, opts ...http.CallOption) (rsp *BatchCreateCompanyReply, err error)
	CreateCompany(ctx context.Context, req *CreateCompanyReq, opts ...http.CallOption) (rsp *CreateCompanyReply, err error)
	DeleteCompany(ctx context.Context, req *DelCompanyReq, opts ...http.CallOption) (rsp *DelCompanyReply, err error)
	GetCompany(ctx context.Context, req *GetCompanyReq, opts ...http.CallOption) (rsp *GetCompanyReply, err error)
	GetCompanyList(ctx context.Context, req *GetCompanyListReq, opts ...http.CallOption) (rsp *GetCompanyListReply, err error)
	OrderCompany(ctx context.Context, req *OrderCompanyReq, opts ...http.CallOption) (rsp *OrderCompanyReply, err error)
	UpdateCompany(ctx context.Context, req *UpdateCompanyReq, opts ...http.CallOption) (rsp *UpdateCompanyReply, err error)
}

type CompanyHTTPClientImpl struct {
	cc *http.Client
}

func NewCompanyHTTPClient(client *http.Client) CompanyHTTPClient {
	return &CompanyHTTPClientImpl{client}
}

func (c *CompanyHTTPClientImpl) BatchCreateCompany(ctx context.Context, in *CreateCompanyReq, opts ...http.CallOption) (*BatchCreateCompanyReply, error) {
	var out BatchCreateCompanyReply
	pattern := "/uims/org/company/batch"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/uims.orgms.api.Company/BatchCreateCompany"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CompanyHTTPClientImpl) CreateCompany(ctx context.Context, in *CreateCompanyReq, opts ...http.CallOption) (*CreateCompanyReply, error) {
	var out CreateCompanyReply
	pattern := "/uims/org/company"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/uims.orgms.api.Company/CreateCompany"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CompanyHTTPClientImpl) DeleteCompany(ctx context.Context, in *DelCompanyReq, opts ...http.CallOption) (*DelCompanyReply, error) {
	var out DelCompanyReply
	pattern := "/uims/org/company/{company_id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/uims.orgms.api.Company/DeleteCompany"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CompanyHTTPClientImpl) GetCompany(ctx context.Context, in *GetCompanyReq, opts ...http.CallOption) (*GetCompanyReply, error) {
	var out GetCompanyReply
	pattern := "/uims/org/company/{company_id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/uims.orgms.api.Company/GetCompany"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CompanyHTTPClientImpl) GetCompanyList(ctx context.Context, in *GetCompanyListReq, opts ...http.CallOption) (*GetCompanyListReply, error) {
	var out GetCompanyListReply
	pattern := "/uims/org/company"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/uims.orgms.api.Company/GetCompanyList"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CompanyHTTPClientImpl) OrderCompany(ctx context.Context, in *OrderCompanyReq, opts ...http.CallOption) (*OrderCompanyReply, error) {
	var out OrderCompanyReply
	pattern := "/uims/org/company"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/uims.orgms.api.Company/OrderCompany"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PATCH", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CompanyHTTPClientImpl) UpdateCompany(ctx context.Context, in *UpdateCompanyReq, opts ...http.CallOption) (*UpdateCompanyReply, error) {
	var out UpdateCompanyReply
	pattern := "/uims/org/company"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/uims.orgms.api.Company/UpdateCompany"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

type DepartmentHTTPServer interface {
	AssignDepartmentUser(context.Context, *AssignDepartmentUserReq) (*AssignDepartmentUserReply, error)
	BatchCreateDepartment(context.Context, *BatchCreateDepartmentReq) (*BatchCreateDepartmentReply, error)
	CreateDepartment(context.Context, *CreateDepartmentReq) (*CreateDepartmentReply, error)
	DeleteDepartment(context.Context, *DelDepartmentReq) (*DelDepartmentReply, error)
	GetDepartment(context.Context, *GetDepartmentReq) (*GetDepartmentReply, error)
	GetDepartmentList(context.Context, *GetDepartmentListReq) (*GetCompanyListReply, error)
	MergeDepartment(context.Context, *MergeDepartmentReq) (*MergeDepartmentReply, error)
	MoveDepartment(context.Context, *MoveDepartmentReq) (*MoveDepartmentReply, error)
	OrderDepartment(context.Context, *OrderDepartmentReq) (*OrderDepartmentReply, error)
	UpdateDepartment(context.Context, *UpdateDepartmentReq) (*UpdateDepartmentReply, error)
}

func RegisterDepartmentHTTPServer(s *http.Server, srv DepartmentHTTPServer) {
	r := s.Route("/")
	r.GET("/uims/orgms/department/{department_id}", _Department_GetDepartment0_HTTP_Handler(srv))
	r.GET("/uims/orgms/department", _Department_GetDepartmentList0_HTTP_Handler(srv))
	r.POST("/uims/orgms/department", _Department_CreateDepartment0_HTTP_Handler(srv))
	r.POST("/uims/orgms/department/batch", _Department_BatchCreateDepartment0_HTTP_Handler(srv))
	r.PUT("/uims/orgms/department", _Department_UpdateDepartment0_HTTP_Handler(srv))
	r.DELETE("/uims/orgms/department/{department_id}", _Department_DeleteDepartment0_HTTP_Handler(srv))
	r.PATCH("/uims/orgms/department/user", _Department_AssignDepartmentUser0_HTTP_Handler(srv))
	r.PATCH("/uims/orgms/department/order", _Department_OrderDepartment0_HTTP_Handler(srv))
	r.PATCH("/uims/orgms/department/move", _Department_MoveDepartment0_HTTP_Handler(srv))
	r.PATCH("/uims/orgms/department/merge", _Department_MergeDepartment0_HTTP_Handler(srv))
}

func _Department_GetDepartment0_HTTP_Handler(srv DepartmentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetDepartmentReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/uims.orgms.api.Department/GetDepartment")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetDepartment(ctx, req.(*GetDepartmentReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetDepartmentReply)
		return ctx.Result(200, reply)
	}
}

func _Department_GetDepartmentList0_HTTP_Handler(srv DepartmentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetDepartmentListReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/uims.orgms.api.Department/GetDepartmentList")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetDepartmentList(ctx, req.(*GetDepartmentListReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetCompanyListReply)
		return ctx.Result(200, reply)
	}
}

func _Department_CreateDepartment0_HTTP_Handler(srv DepartmentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateDepartmentReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/uims.orgms.api.Department/CreateDepartment")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateDepartment(ctx, req.(*CreateDepartmentReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateDepartmentReply)
		return ctx.Result(200, reply)
	}
}

func _Department_BatchCreateDepartment0_HTTP_Handler(srv DepartmentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in BatchCreateDepartmentReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/uims.orgms.api.Department/BatchCreateDepartment")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.BatchCreateDepartment(ctx, req.(*BatchCreateDepartmentReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*BatchCreateDepartmentReply)
		return ctx.Result(200, reply)
	}
}

func _Department_UpdateDepartment0_HTTP_Handler(srv DepartmentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateDepartmentReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/uims.orgms.api.Department/UpdateDepartment")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateDepartment(ctx, req.(*UpdateDepartmentReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateDepartmentReply)
		return ctx.Result(200, reply)
	}
}

func _Department_DeleteDepartment0_HTTP_Handler(srv DepartmentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DelDepartmentReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/uims.orgms.api.Department/DeleteDepartment")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteDepartment(ctx, req.(*DelDepartmentReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DelDepartmentReply)
		return ctx.Result(200, reply)
	}
}

func _Department_AssignDepartmentUser0_HTTP_Handler(srv DepartmentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AssignDepartmentUserReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/uims.orgms.api.Department/AssignDepartmentUser")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AssignDepartmentUser(ctx, req.(*AssignDepartmentUserReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AssignDepartmentUserReply)
		return ctx.Result(200, reply)
	}
}

func _Department_OrderDepartment0_HTTP_Handler(srv DepartmentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in OrderDepartmentReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/uims.orgms.api.Department/OrderDepartment")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.OrderDepartment(ctx, req.(*OrderDepartmentReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*OrderDepartmentReply)
		return ctx.Result(200, reply)
	}
}

func _Department_MoveDepartment0_HTTP_Handler(srv DepartmentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in MoveDepartmentReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/uims.orgms.api.Department/MoveDepartment")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.MoveDepartment(ctx, req.(*MoveDepartmentReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*MoveDepartmentReply)
		return ctx.Result(200, reply)
	}
}

func _Department_MergeDepartment0_HTTP_Handler(srv DepartmentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in MergeDepartmentReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/uims.orgms.api.Department/MergeDepartment")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.MergeDepartment(ctx, req.(*MergeDepartmentReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*MergeDepartmentReply)
		return ctx.Result(200, reply)
	}
}

type DepartmentHTTPClient interface {
	AssignDepartmentUser(ctx context.Context, req *AssignDepartmentUserReq, opts ...http.CallOption) (rsp *AssignDepartmentUserReply, err error)
	BatchCreateDepartment(ctx context.Context, req *BatchCreateDepartmentReq, opts ...http.CallOption) (rsp *BatchCreateDepartmentReply, err error)
	CreateDepartment(ctx context.Context, req *CreateDepartmentReq, opts ...http.CallOption) (rsp *CreateDepartmentReply, err error)
	DeleteDepartment(ctx context.Context, req *DelDepartmentReq, opts ...http.CallOption) (rsp *DelDepartmentReply, err error)
	GetDepartment(ctx context.Context, req *GetDepartmentReq, opts ...http.CallOption) (rsp *GetDepartmentReply, err error)
	GetDepartmentList(ctx context.Context, req *GetDepartmentListReq, opts ...http.CallOption) (rsp *GetCompanyListReply, err error)
	MergeDepartment(ctx context.Context, req *MergeDepartmentReq, opts ...http.CallOption) (rsp *MergeDepartmentReply, err error)
	MoveDepartment(ctx context.Context, req *MoveDepartmentReq, opts ...http.CallOption) (rsp *MoveDepartmentReply, err error)
	OrderDepartment(ctx context.Context, req *OrderDepartmentReq, opts ...http.CallOption) (rsp *OrderDepartmentReply, err error)
	UpdateDepartment(ctx context.Context, req *UpdateDepartmentReq, opts ...http.CallOption) (rsp *UpdateDepartmentReply, err error)
}

type DepartmentHTTPClientImpl struct {
	cc *http.Client
}

func NewDepartmentHTTPClient(client *http.Client) DepartmentHTTPClient {
	return &DepartmentHTTPClientImpl{client}
}

func (c *DepartmentHTTPClientImpl) AssignDepartmentUser(ctx context.Context, in *AssignDepartmentUserReq, opts ...http.CallOption) (*AssignDepartmentUserReply, error) {
	var out AssignDepartmentUserReply
	pattern := "/uims/orgms/department/user"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/uims.orgms.api.Department/AssignDepartmentUser"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PATCH", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DepartmentHTTPClientImpl) BatchCreateDepartment(ctx context.Context, in *BatchCreateDepartmentReq, opts ...http.CallOption) (*BatchCreateDepartmentReply, error) {
	var out BatchCreateDepartmentReply
	pattern := "/uims/orgms/department/batch"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/uims.orgms.api.Department/BatchCreateDepartment"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DepartmentHTTPClientImpl) CreateDepartment(ctx context.Context, in *CreateDepartmentReq, opts ...http.CallOption) (*CreateDepartmentReply, error) {
	var out CreateDepartmentReply
	pattern := "/uims/orgms/department"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/uims.orgms.api.Department/CreateDepartment"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DepartmentHTTPClientImpl) DeleteDepartment(ctx context.Context, in *DelDepartmentReq, opts ...http.CallOption) (*DelDepartmentReply, error) {
	var out DelDepartmentReply
	pattern := "/uims/orgms/department/{department_id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/uims.orgms.api.Department/DeleteDepartment"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DepartmentHTTPClientImpl) GetDepartment(ctx context.Context, in *GetDepartmentReq, opts ...http.CallOption) (*GetDepartmentReply, error) {
	var out GetDepartmentReply
	pattern := "/uims/orgms/department/{department_id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/uims.orgms.api.Department/GetDepartment"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DepartmentHTTPClientImpl) GetDepartmentList(ctx context.Context, in *GetDepartmentListReq, opts ...http.CallOption) (*GetCompanyListReply, error) {
	var out GetCompanyListReply
	pattern := "/uims/orgms/department"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/uims.orgms.api.Department/GetDepartmentList"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DepartmentHTTPClientImpl) MergeDepartment(ctx context.Context, in *MergeDepartmentReq, opts ...http.CallOption) (*MergeDepartmentReply, error) {
	var out MergeDepartmentReply
	pattern := "/uims/orgms/department/merge"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/uims.orgms.api.Department/MergeDepartment"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PATCH", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DepartmentHTTPClientImpl) MoveDepartment(ctx context.Context, in *MoveDepartmentReq, opts ...http.CallOption) (*MoveDepartmentReply, error) {
	var out MoveDepartmentReply
	pattern := "/uims/orgms/department/move"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/uims.orgms.api.Department/MoveDepartment"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PATCH", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DepartmentHTTPClientImpl) OrderDepartment(ctx context.Context, in *OrderDepartmentReq, opts ...http.CallOption) (*OrderDepartmentReply, error) {
	var out OrderDepartmentReply
	pattern := "/uims/orgms/department/order"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/uims.orgms.api.Department/OrderDepartment"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PATCH", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DepartmentHTTPClientImpl) UpdateDepartment(ctx context.Context, in *UpdateDepartmentReq, opts ...http.CallOption) (*UpdateDepartmentReply, error) {
	var out UpdateDepartmentReply
	pattern := "/uims/orgms/department"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/uims.orgms.api.Department/UpdateDepartment"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

type UserHTTPServer interface {
	BatchCreateUser(context.Context, *BatchCreateUserReq) (*BatchCreateUserReply, error)
	CreateUser(context.Context, *CreateUserReq) (*CreateUserReply, error)
	DelUser(context.Context, *DelUserReq) (*DelUserReply, error)
	GetUserInfoReq(context.Context, *GetUserReq) (*GetUserReply, error)
}

func RegisterUserHTTPServer(s *http.Server, srv UserHTTPServer) {
	r := s.Route("/")
	r.POST("/uims/orgms/user", _User_CreateUser0_HTTP_Handler(srv))
	r.POST("/uims/orgms/user/batch", _User_BatchCreateUser0_HTTP_Handler(srv))
	r.DELETE("/uims/orgms/user/{uid}", _User_DelUser0_HTTP_Handler(srv))
	r.GET("/uims/orgms/user/{uid}", _User_GetUserInfoReq0_HTTP_Handler(srv))
}

func _User_CreateUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateUserReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/uims.orgms.api.User/CreateUser")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateUser(ctx, req.(*CreateUserReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateUserReply)
		return ctx.Result(200, reply)
	}
}

func _User_BatchCreateUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in BatchCreateUserReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/uims.orgms.api.User/BatchCreateUser")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.BatchCreateUser(ctx, req.(*BatchCreateUserReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*BatchCreateUserReply)
		return ctx.Result(200, reply)
	}
}

func _User_DelUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DelUserReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/uims.orgms.api.User/DelUser")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DelUser(ctx, req.(*DelUserReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DelUserReply)
		return ctx.Result(200, reply)
	}
}

func _User_GetUserInfoReq0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/uims.orgms.api.User/GetUserInfoReq")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUserInfoReq(ctx, req.(*GetUserReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUserReply)
		return ctx.Result(200, reply)
	}
}

type UserHTTPClient interface {
	BatchCreateUser(ctx context.Context, req *BatchCreateUserReq, opts ...http.CallOption) (rsp *BatchCreateUserReply, err error)
	CreateUser(ctx context.Context, req *CreateUserReq, opts ...http.CallOption) (rsp *CreateUserReply, err error)
	DelUser(ctx context.Context, req *DelUserReq, opts ...http.CallOption) (rsp *DelUserReply, err error)
	GetUserInfoReq(ctx context.Context, req *GetUserReq, opts ...http.CallOption) (rsp *GetUserReply, err error)
}

type UserHTTPClientImpl struct {
	cc *http.Client
}

func NewUserHTTPClient(client *http.Client) UserHTTPClient {
	return &UserHTTPClientImpl{client}
}

func (c *UserHTTPClientImpl) BatchCreateUser(ctx context.Context, in *BatchCreateUserReq, opts ...http.CallOption) (*BatchCreateUserReply, error) {
	var out BatchCreateUserReply
	pattern := "/uims/orgms/user/batch"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/uims.orgms.api.User/BatchCreateUser"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) CreateUser(ctx context.Context, in *CreateUserReq, opts ...http.CallOption) (*CreateUserReply, error) {
	var out CreateUserReply
	pattern := "/uims/orgms/user"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/uims.orgms.api.User/CreateUser"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) DelUser(ctx context.Context, in *DelUserReq, opts ...http.CallOption) (*DelUserReply, error) {
	var out DelUserReply
	pattern := "/uims/orgms/user/{uid}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/uims.orgms.api.User/DelUser"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) GetUserInfoReq(ctx context.Context, in *GetUserReq, opts ...http.CallOption) (*GetUserReply, error) {
	var out GetUserReply
	pattern := "/uims/orgms/user/{uid}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/uims.orgms.api.User/GetUserInfoReq"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
