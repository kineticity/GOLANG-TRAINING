package user

import (
	"contactApp/contact"
	"contactApp/contactinfo"
	"errors"
	"fmt"
)

var userID int = 0
var allAdmin []*User
var allStaff []*User

// User structure
type User struct {
	UserID    int
	Firstname string
	Lastname  string
	IsAdmin   bool
	IsActive  bool
	Contacts  []*contact.Contact
}

// Factory ADMIN
func NewAdmin(firstname, lastname string) (*User,error) {
	err:=validName(firstname)
	if err!=nil{
		return nil,err
	}
	err=validName(lastname)
	if err!=nil{
		return nil,err
	}

	adminUser := &User{
		UserID:    userID,
		Firstname: firstname,
		Lastname:  lastname,
		IsAdmin:   true,
		IsActive:  true,
		Contacts:  nil,
	}

	userID++
	allAdmin = append(allAdmin, adminUser)
	return adminUser,nil
}

// Factory STAFF OR CREATE STAFF OPERATION FOR ADMIN
func (u *User) NewStaff(firstname, lastname string) (*User, error) {
	err:=validName(firstname)
	if err!=nil{
		return nil,err
	}
	err=validName(lastname)
	if err!=nil{
		return nil,err
	}
	if !u.IsAdmin || !u.IsActive {
		return nil, errors.New("only active Admins can create users")
	} else {
		staffUser := &User{
			UserID:    userID,
			Firstname: firstname,
			Lastname:  lastname,
			IsAdmin:   false,
			IsActive:  true,
			Contacts:  nil,
		}
		userID++
		allStaff = append(allStaff, staffUser)
		return staffUser, nil
	}
}

//READ USERS FOR ADMIN
func (u *User) GetUsers() ([]*User, error) {
	if !u.IsAdmin || !u.IsActive {
		return nil, errors.New("only active Admins can read users")
	}
	return allStaff, nil
}

// Validation function for names
func validName(name interface{}) error {
	nameValue, ok := name.(string)
	if !ok {
		return errors.New("invalid value type for name")
	}
	if nameValue == "" {
		return errors.New("name cannot be empty")
	}
	return nil
}

// VALIDATION FN. FOR UPDATEUSER ARGS
func validateUserParams(userID int, parameter string, newValue interface{}) error {
	if userID <= 0 {
		return errors.New("invalid userID provided")
	}

	validParameters := map[string]string{
		"firstname": "string",
		"lastname":  "string",
		"isAdmin":   "bool",
		"isActive":  "bool",
	}

	expectedType, valid := validParameters[parameter]
	if !valid {
		return errors.New("invalid parameter provided, must be one of 'firstname', 'lastname', 'isAdmin', or 'isActive'")
	}

	switch expectedType {
	case "string":
		_, ok := newValue.(string)
		if !ok || newValue.(string) == "" {
			return errors.New("invalid value for " + parameter + ": must be a non-empty string")
		}
	case "bool":
		_, ok := newValue.(bool)
		if !ok {
			return errors.New("invalid value for " + parameter + ": must be a boolean")
		}
	default:
		return errors.New("unexpected parameter type")
	}

	return nil
}

// UpdateUserByParameter method for admins
func (u *User) UpdateUserByParameter(userID int, parameter string, newValue interface{}) error {
	// Ensure the user is an admin before allowing updates
	if !u.IsAdmin {
		return errors.New("only admins can update staff user values")
	}

	// Validate the parameters
	if err := validateUserParams(userID, parameter, newValue); err != nil {
		return err
	}

	// Find the target user
	var targetUser *User
	for _, user := range allStaff {
		if user.UserID == userID {
			targetUser = user
			break
		}
	}

	if targetUser == nil {
		return errors.New("staff user not found")
	}

	// Update the parameters of the target user
	switch parameter {
	case "firstname":
		firstname, _ := newValue.(string)
		targetUser.Firstname = firstname

	case "lastname":
		lastname, _ := newValue.(string)
		targetUser.Lastname = lastname

	case "isAdmin":
		isAdmin, _ := newValue.(bool)
		targetUser.IsAdmin = isAdmin

	case "isActive":
		isActive, _ := newValue.(bool)
		targetUser.IsActive = isActive

	default:
		return errors.New("invalid parameter") // This case should be unreachable due to validation
	}

	return nil
}

//DELETE USER OPERATION FOR ADMIN
func (u *User) DeleteUser(userID int) error {
	if !u.IsAdmin || !u.IsActive {
		return errors.New("only active Admins can delete users")
	}
	for i, user := range allStaff {
		if user.UserID == userID {
			allStaff[i].IsActive = false
			return nil
		}
	}
	return errors.New("user not found")
}

//CREATE CONTACT OPERATION FOR STAFF
func (u *User) CreateContact(firstname, lastname string) error {
	if u.IsAdmin || !u.IsActive {
		return errors.New("only active Staff can create contacts")
	}

	contactid:=0

	if len(u.Contacts)!=0{
		contactid=u.Contacts[len(u.Contacts)-1].ContactID
		contactid++
	}


	newcontact,err := contact.NewContact(firstname, lastname,contactid)
	if err!=nil{
		return err
	}
	u.Contacts = append(u.Contacts, newcontact)

	return nil
}


//READ CONTACTS OPERATION FOR STAFF
func (u *User) GetContacts() ([]*contact.Contact, error) {
	if u.IsAdmin || !u.IsActive {
		return nil, errors.New("only active Staff can read contacts")
	}
	for _, contact := range u.Contacts {
		fmt.Println(contact)

	}
	return u.Contacts, nil
}


// VALIDATION FN FOR UPDATEcontact ARGUMENTS
func validateContactParams(contactID int, parameter string) error {
	if contactID <= 0 {
		return errors.New("invalid contactID provided")
	}

	validParameters := map[string]bool{
		"firstname":   true,
		"lastname":   true,
		"isActive": true,
	}

	if _, valid := validParameters[parameter]; !valid {
		return errors.New("invalid parameter provided, must be one of 'email', 'phone', 'address', or 'name'")
	}

	return nil
}

// UpdateContact method for User
func (u *User) UpdateContact(contactID int, parameter string, newValue interface{}) error {
	if u.IsAdmin || !u.IsActive {
		return errors.New("only active staff can update contacts")
	}

	if err := validateContactParams(contactID, parameter); err != nil {
		return err
	}

	var targetContact *contact.Contact
	for _, c := range u.Contacts {
		if c.ContactID == contactID {
			targetContact = c
			break
		}
	}

	if targetContact == nil {
		return errors.New("contact not found")
	}

	return targetContact.UpdateContact(contactID, parameter, newValue, u.Contacts)
}

//DELETE CONTACT OPERATION FOR STAFF
func (u *User) DeleteContact(contactID int) error {
	if u.IsAdmin || !u.IsActive {
		return errors.New("only active Staff can delete contacts")
	}

	var targetContact *contact.Contact
	for _, c := range u.Contacts {
		if c.ContactID == contactID {
			targetContact = c
			break
		}
	}

	if targetContact == nil {
		return errors.New("contact not found")
	}

	if err := targetContact.DeleteContact(); err != nil { //Actual delete contact logic in contact package
		return err
	}

	return nil
}

//CREATE CONTACT INFO OPERATION FOR STAFF
func (u *User) CreateContactInfo(contactID int, infoType, value string) error {
	if u.IsAdmin || !u.IsActive {
		return errors.New("only active Staff can create contact details")
	}

	var targetContact *contact.Contact
	for _, c := range u.Contacts {
		if c.ContactID == contactID {
			targetContact = c
			break
		}
	}
	if targetContact == nil {
		return errors.New("contact not found")
	}

	return targetContact.CreateContactInfo(infoType, value) 
	//actual create contactinfo logic in contactinfo package, contact is validated in contact package
}

// READ CONTACT INFO OPERATION FOR STAFF
func (u *User) GetContactInfo(contactID int, infoID int) (*contactinfo.ContactInfo, error) {
	if u.IsAdmin || !u.IsActive {
		return nil, errors.New("only active Staff can read contact details")
	}
	var targetContact *contact.Contact
	for _, c := range u.Contacts {
		if c.ContactID == contactID {
			targetContact = c
			break
		}
	}

	if targetContact == nil {
		return nil, errors.New("contact not found")
	}

	return targetContact.GetContactInfo(infoID)
}

// contactinfo validation
func validateContactInfoParams(contactID int, infoID int, parameter string) error {
	if contactID < 0 {
		return errors.New("invalid contactID provided")
	}

	if infoID < 0 {
		return errors.New("invalid infoID provided")
	}

	validParameters := map[string]bool{
		"type":   true,
		"value":   true,
	}

	if _, valid := validParameters[parameter]; !valid {
		return errors.New("invalid parameter provided, must be one of 'email', 'phone', or 'address'")
	}

	return nil
}

// UPDATE CONTACT INFO FOR STAFF
func (u *User) UpdateContactInfo(contactID int, infoID int, parameter string, newValue interface{}) error {
	if u.IsAdmin || !u.IsActive {
		return errors.New("only active Staff can update contact infos")
	}

	if err := validateContactInfoParams(contactID, infoID, parameter); err != nil {
		return err
	}

	var targetContact *contact.Contact
	for _, c := range u.Contacts {
		if c.ContactID == contactID {
			targetContact = c
			break
		}
	}

	if targetContact == nil {
		return errors.New("contact not found")
	}

	return targetContact.UpdateContactInfo(infoID, parameter, newValue)
}

// DELETE CONTACTINFO OPERATION FOR STAFF
func (u *User) DeleteContactInfo(contactID, infoID int) error {
	if u.IsAdmin || !u.IsActive { //make fn
		return errors.New("only active staff can delete contact information")
	}

	var targetContact *contact.Contact //make fn
	for _, c := range u.Contacts {
		if c.ContactID == contactID {
			targetContact = c
			break
		}
	}

	if targetContact == nil {
		return errors.New("contact not found")
	}

	err := targetContact.DeleteContactInfo(infoID)
	if err != nil {
		return err
	}

	return nil
}


func (u *User) PrintDetails() {
	fmt.Printf("UserID: %d\n", u.UserID)
	fmt.Printf("Name: %s %s\n", u.Firstname, u.Lastname)
	fmt.Printf("IsAdmin: %t\n", u.IsAdmin)
	fmt.Printf("IsActive: %t\n", u.IsActive)
	fmt.Println("Contacts:")

	for _, contact := range u.Contacts {
		contact.PrintDetails() 
	}
	fmt.Println("-----------------------")
	fmt.Println("-----------------------")
}
