package dao

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewOrgmsCompanyModel,
	NewOrgmsDepartmentModel,
	NewOrgmsUserExtraModel,
	NewOrgmsUserDepartmentModel,
	NewOrgmsExtAttrModel,
)