package tomodels

import (
	"payment_service/models"
	"payment_service/structs/responses"
)

func CustomerResponseToModel(cust *responses.Customers) models.Customers {
	if cust == nil {
		return models.Customers{}
	}

	result := models.Customers{
		CustomerId:           cust.CustomerId,
		FullName:             cust.FullName,
		ImagePath:            cust.ImagePath,
		Email:                cust.Email,
		PhoneNumber:          cust.PhoneNumber,
		Location:             cust.Location,
		IdentificationNumber: cust.IdentificationNumber,
		Nickname:             cust.Nickname,
		Dob:                  cust.Dob,
		DateCreated:          cust.DateCreated,
		DateModified:         cust.DateModified,
		CreatedBy:            cust.CreatedBy,
		ModifiedBy:           cust.ModifiedBy,
		Active:               cust.Active,
		LastTxnDate:          cust.LastTxnDate,
	}

	if cust.IdentificationType != nil {
		result.IdentificationType = cust.IdentificationType.IdentificationTypeId
	}

	if cust.Branch != nil {
		result.Branch = &models.Branches{
			BranchId:     cust.Branch.BranchId,
			Branch:       cust.Branch.Branch,
			Location:     cust.Branch.Location,
			PhoneNumber:  cust.Branch.PhoneNumber,
			Active:       cust.Branch.Active,
			DateCreated:  cust.Branch.DateCreated,
			DateModified: cust.Branch.DateModified,
			CreatedBy:    cust.Branch.CreatedBy,
			ModifiedBy:   cust.Branch.ModifiedBy,
		}
	}

	if cust.Shop != nil {
		result.ShopId = cust.Shop.ShopId
	}

	if cust.CustomerCategory != nil {
		result.CustomerCategory = &models.Customer_categories{
			CustomerCategoryId: cust.CustomerCategory.CustomerCategoryId,
			Category:           cust.CustomerCategory.Category,
			Description:        cust.CustomerCategory.Description,
			DateCreated:        cust.CustomerCategory.DateCreated,
			DateModified:       cust.CustomerCategory.DateModified,
			CreatedBy:          cust.CustomerCategory.CreatedBy,
			ModifiedBy:         cust.CustomerCategory.ModifiedBy,
			Active:             cust.CustomerCategory.Active,
		}
	}

	return result
}

func UserResponseToModel(user *responses.Users) models.Users {
	if user == nil {
		return models.Users{}
	}

	result := models.Users{
		UserId:        user.UserId,
		ImagePath:     user.ImagePath,
		UserType:      user.UserType,
		FullName:      user.FullName,
		Username:      user.Username,
		Password:      user.Password,
		Email:         user.Email,
		PhoneNumber:   user.PhoneNumber,
		Gender:        user.Gender,
		Dob:           user.Dob,
		Address:       user.Address,
		IdType:        user.IdType,
		IdNumber:      user.IdNumber,
		MaritalStatus: user.MaritalStatus,
		Active:        user.Active,
		IsVerified:    user.IsVerified,
		DateCreated:   user.DateCreated,
		DateModified:  user.DateModified,
		CreatedBy:     user.CreatedBy,
		ModifiedBy:    user.ModifiedBy,
	}

	return result
}
