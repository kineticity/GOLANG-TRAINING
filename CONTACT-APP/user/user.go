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
func NewAdmin(firstname, lastname string) *User {
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
	return adminUser
}

// Factory STAFF OR CREATE STAFF OPERATION FOR ADMIN
func (u *User) NewStaff(firstname, lastname string) (*User, error) {
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
func (u *User) ReadUsers() ([]*User, error) {
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

//UPDATE USER BY PARAMETER OPERATION FOR ADMINS
func (u *User) UpdateUserByParameter(userID int, parameter string, newValue interface{}) error {
	// Ensure the user is an admin before allowing updates
	if !u.IsAdmin {
		return errors.New("only admins can update staff user values")
	}

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
		err := validName(newValue)
		if err != nil {
			return err
		}
		firstname, _ := newValue.(string)
		targetUser.Firstname = firstname

	case "lastname":
		err := validName(newValue)
		if err != nil {
			return err
		}
		lastname, _ := newValue.(string)
		targetUser.Lastname = lastname

	case "isAdmin":
		// Only admins should be allowed to set isAdmin
		isAdmin, ok := newValue.(bool)
		if !ok {
			return errors.New("invalid value type for isAdmin")
		}
		targetUser.IsAdmin = isAdmin

	case "isActive":
		// Allow admins to change isActive status
		isActive, ok := newValue.(bool)
		if !ok {
			return errors.New("invalid value type for isActive")
		}
		targetUser.IsActive = isActive

	default:
		return errors.New("invalid parameter")
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

	newcontact := contact.NewContact(firstname, lastname, len(u.Contacts))
	u.Contacts = append(u.Contacts, newcontact)

	return nil
}


//READ CONTACTS OPERATION FOR STAFF
func (u *User) ReadContacts() ([]*contact.Contact, error) {
	if u.IsAdmin || !u.IsActive {
		return nil, errors.New("only active Staff can read contacts")
	}
	for _, contact := range u.Contacts {
		fmt.Println(contact)

	}
	return u.Contacts, nil
}


//UPDATE CONTACT OPERATION FOR STAFF
func (u *User) UpdateContact(contactID int, parameter string, newValue interface{}) error {
	if u.IsAdmin || !u.IsActive { 
		return errors.New("only active staff can update contacts")
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

	return targetContact.UpdateContact(contactID, parameter, newValue, u.Contacts) //actual update logic in contact package
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
func (u *User) ReadContactInfo(contactID int, infoID int) (*contactinfo.ContactInfo, error) {
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

	return targetContact.ReadContactInfo(infoID)
}

//UPDATE CONTACT INFO OPERATION FOR STAFF
func (u *User) UpdateContactInfo(contactID int, infoID int, parameter string, newValue interface{}) error {
	if u.IsAdmin || !u.IsActive {
		return errors.New("only active Staff can update contact infos")
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
