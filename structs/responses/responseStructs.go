package responses

import (
	"time"
)

type Currencies struct {
	CurrencyId   int64
	Symbol       string
	Currency     string
	Active       int
	DateCreated  time.Time
	DateModified time.Time
	CreatedBy    int
	ModifiedBy   int
}

type Countries struct {
	CountryId       int64
	Country         string
	Description     string
	CountryCode     string
	DefaultCurrency *Currencies
	DateCreated     time.Time
	DateModified    time.Time
	CreatedBy       int
	ModifiedBy      int
}

type Branches struct {
	BranchId     int64
	Branch       string
	Country      *Countries
	Location     string
	PhoneNumber  string
	Active       int
	DateCreated  time.Time
	DateModified time.Time
	CreatedBy    int
	ModifiedBy   int
}

type Identification_types struct {
	IdentificationTypeId int64
	Name                 string
	Code                 string
	DateCreated          time.Time
	DateModified         time.Time
	CreatedBy            int
	ModifiedBy           int
	Active               int
}

type Shops struct {
	ShopId              int64
	ShopName            string
	ShopDescription     string
	ShopAssistantName   string
	ShopAssistantNumber string
	PhoneNumber         string
	Email               string
	Image               string
	DateCreated         time.Time
	DateModified        time.Time
	CreatedBy           int
	ModifiedBy          int
	Active              int
}

type Customer_categories struct {
	CustomerCategoryId int64
	Category           string
	Description        string
	DateCreated        time.Time
	DateModified       time.Time
	CreatedBy          int
	ModifiedBy         int
	Active             int
}

type Customers struct {
	CustomerId           int64
	CustomerNumber       string
	FullName             string
	ImagePath            string
	Email                string
	PhoneNumber          string
	Gender               string
	Location             string
	IdentificationType   *Identification_types
	IdentificationNumber string
	Branch               *Branches
	Shop                 *Shops
	CustomerCategory     *Customer_categories
	Nickname             string
	Dob                  time.Time
	DateCreated          time.Time
	DateModified         time.Time
	CreatedBy            int
	ModifiedBy           int
	Active               int
	LastTxnDate          time.Time
	EmergencyContacts    []*Customer_emergency_contacts
	Guarantors           []*Customer_guarantors
}

type Customer_guarantors struct {
	CustomerGuarantorId int64
	Name                string
	Contact             string
	Customer            *Customers
	DateCreated         time.Time
	DateModified        time.Time
	CreatedBy           int
	ModifiedBy          int
}

type Customer_emergency_contacts struct {
	CustomerEmergencyContactId int64
	Name                       string
	Contact                    string
	Customer                   *Customers
	DateCreated                time.Time
	DateModified               time.Time
	CreatedBy                  int
	ModifiedBy                 int
}

type Users struct {
	UserId        int64
	UserDetails   *UserExtraDetails
	ImagePath     string
	UserType      int
	FullName      string
	Username      string
	Password      string
	Email         string
	PhoneNumber   string
	Gender        string
	Dob           time.Time
	Address       string
	IdType        string
	IdNumber      string
	MaritalStatus string
	Active        int
	Role          *Roles
	IsVerified    bool
	DateCreated   time.Time
	DateModified  time.Time
	CreatedBy     int
	ModifiedBy    int
}

type UserExtraDetails struct {
	UserDetailsId int64
	Branch        *Branches
	Shop          *Shops
	Nickname      string
	DateCreated   time.Time
	DateModified  time.Time
	CreatedBy     int
	ModifiedBy    int
	Active        int
}

type Roles struct {
	RoleId       int64
	Role         string
	Description  string
	DateCreated  time.Time
	DateModified time.Time
	CreatedBy    int
	ModifiedBy   int
	Active       int
}

type CustomerResponseDTO struct {
	StatusCode int
	Customer   *Customers
	StatusDesc string
}

type UserResponseDTO struct {
	StatusCode int
	User       *Users
	StatusDesc string
}
