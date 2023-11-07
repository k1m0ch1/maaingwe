package utils

type CheckInTemplate struct {
	Message						string					`yaml:"message"`
	LatLong						string					`yaml:"latlong"`
	LocationType				int						`yaml:"locationType"`
}

type CheckIn struct{
	Location					string					`json:"location"`
	Message						string					`json:"message",yaml:"message"`
	LatLong						string					`json:"latlng",yaml:"latlong"`
	Token						string					`json:"token"`
	LocationType				int						`json:"location_type",yaml:"locationType"` // 1 for Office, 2 for Home, 3 for Field Duty
	InOut						int						`json:"in_out"` // 1 Checkin, 2 Checkout
	UDID 						string					`json:"udid"`
	Purpose						string					`json:"purpose"`
	ID							string					`json:"checkin_id"` // for checkout purpose
}

type CheckOut struct{
	LocationType				int						`yaml:"locationType"`
	Message						string					`yaml:"message"`
	LatLong						string					`yaml:"latlong"`
}

type Scheduler struct{
	CheckIn						string					`yaml:"checkin"`
	CheckOut					string					`yaml:"checkout"`
	// CheckInRandom			[]string	`yaml:"checkinRandom"`
	// CheckOutRandom			[]string	`yaml:"checkoutRandom"`
}

type AppConfig struct{
	Token						string					`yaml:"token"`
	Hostname					string					`yaml:"hostname"`
	CheckIn						CheckInTemplate			`yaml:"checkin"`
	CheckOut 					CheckOut				`yaml:"checkout"`
	Scheduler					Scheduler				`yaml:"scheduler"`
}

type Token struct {
	Token						string					`json:"token"`
}

type CheckInResponse struct{
	Status						int						`json:"status"`
	Message						string					`json:"message"`
}

type CheckInID struct{
	ID							string					`json:"id"`
	Date						string					`json:"date"`
	LastAction					int						`json:"last_action"` // 1 for checkin, 2 for checkout
}

type CheckInIDResponse struct{
	Status						int						`json:"status"`
	Message						CheckInID				`json:"message"`
	Error						string					`json:"error"`
}

type Auth struct {
	QRCode						string					`json:"qrcode"`
	UDID						string					`json:"udid"`
}

type OrgStandardField struct {
	DesignationName				string					`json:"designation_name"`
	TopDepartment				string					`json:"top_department"`
}

type UserDetails struct {
	Name						string					`json:"name"`
	Email						string					`json:"email"`
	UserID						string					`json:"user_id"`
	TenantID					string					`json:"tenant_id"`
	MongoID						string					`json:"mongo_id"`
	Designation					string					`json:"designation"`
	Department					string					`json:"department"`
	BusinessUnit				string					`json:"business_unit"`
	Mobile						string					`json:"mobile"`
	Office						string					`json:"office"`
	OfficeAddress				string					`json:"office_address"`
	DateOfBirth					string					`json:"dob"`
	DateOfJoin					string					`json:"doj"`
	EmployeeNo					string					`json:"employee_no"`
	ManagerName					string					`json:"manager_name"`
	PictureOne					string					`json:"pic48"`
	PictureTwo					string					`json:"pic320"`
	PictureThree				string					`json:"pic25"`
	ProfileTag					string					`json:"profile_tag"`
	OrgStandardField			OrgStandardField		`json:"org_standard_fields"`
}

type UserDetailsProfile struct {
	ProfilePicture				string					`json:"Profile Picture"`
	EmailID						string					`json:"Email ID"`
	Company						string					`json:"Company"`
	EmployeeSubType				string					`json:"Employee Sub Type"`
	ExperienceRole				string					`json:"Experience in current role"`
	FunctionalHead				string					`json:"Functional Head"`
	LocationType				string					`json:"Head Office"`
}

type UserDetailsProfileNT struct {
	ProfilePicture				string					`json:"Profile Picture"`
	EmailID						string					`json:"Email ID"`
	Company						string					`json:"Company"`
	EmployeeSubType				string					`json:"Employee Sub Type"`
	ExperienceRole				string					`json:"Experience in current role"`
	FunctionalHead				string					`json:"Functional Head"`
	LocationType				string					`json:"Head Office"`
}

type CurrentUserProfileResponse struct {
	Status						int					`json:"status"`
	Message						string					`json:"message"`
	UserDetails					UserDetails				`json:"user_details"`
	UserDetailsProfile			UserDetailsProfile		`json:"user_details_profile"`
	UserDetailsProfileNT		UserDetailsProfileNT	`json:"user_details_profile_non_translated"`
}

type AuthResponse struct{
	ErrorCode					int						`json:"error_code"`
	Status						int						`json:"status"`
	Token						string					`json:"token"`
	UserID						string					`json:"user_id"`
	TenantID					string					`json:"tenant_id"`
	Expires						int						`json:"expires"`
	IsManager					bool					`json:"is_manager"`
	Message						string					`json:"message"`
	UserDetails					UserDetails				`json:"user_details"`
}

