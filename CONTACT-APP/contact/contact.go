package contact

import (
	"contactApp/contactinfo"
	"errors"
	"fmt"
)

type Contact struct {
	ContactID    int
	Fname        string
	Lname        string
	IsActive     bool
	ContactInfos []*contactinfo.ContactInfo
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

// NewContact factory function
func NewContact(firstname, lastname string, contactid int) (*Contact, error) {
	err := validName(firstname)
	if err != nil {
		return nil, err
	}
	err = validName(lastname)
	if err != nil {
		return nil, err
	}
	contact := &Contact{
		ContactID:    contactid, // Assuming contacts is a package-level variable or slice
		Fname:        firstname,
		Lname:        lastname,
		IsActive:     true,
		ContactInfos: nil,
	}

	return contact,nil
}

//UPDATE CONTACT

func (targetContact *Contact) UpdateContact(contactID int, parameter string, newValue interface{}, contacts []*Contact) error {

	switch parameter {
	case "firstname":
		if firstname, ok := newValue.(string); ok && firstname != "" {
			targetContact.Fname = firstname
		} else {
			return errors.New("invalid value type for firstname")
		}

	case "lastname":
		if lastname, ok := newValue.(string); ok && lastname != "" {
			targetContact.Lname = lastname
		} else {
			return errors.New("invalid value type for lastname")
		}

	case "isActive":
		if isActive, ok := newValue.(bool); ok {
			targetContact.IsActive = isActive
		} else {
			return errors.New("invalid value type for isActive")
		}

	default:
		return errors.New("invalid parameter")
	}

	return nil
}

// DELETE CONTACT OPERATION (SOFT DELETE)
func (c *Contact) DeleteContact() error {
	if !c.IsActive {
		return errors.New("contact is already inactive")
	}

	c.IsActive = false
	return nil
}

// CREATE CONTACT INFO
func (c *Contact) CreateContactInfo(infoType, value string) error {
	if !c.IsActive {
		return errors.New("contact is not active, cannot add details")
	}
	contactinfoid:=0

	if len(c.ContactInfos)!=0{
		contactinfoid=c.ContactInfos[len(c.ContactInfos)-1].ContactInfoID
		contactinfoid++
	}
	newInfo := contactinfo.NewContactInfo(infoType, value, contactinfoid) //<--Actual contactinfo object creation logic in contactinfo package

	c.ContactInfos = append(c.ContactInfos, newInfo)
	return nil
}

// READ CONTACT INFO
func (c *Contact) GetContactInfo(infoID int) (*contactinfo.ContactInfo, error) {
	if !c.IsActive {
		return nil, errors.New("contact is inactive, cannot read details")
	}

	for _, info := range c.ContactInfos {
		if info.ContactInfoID == infoID {
			return info, nil
		}
	}
	return nil, errors.New("contact detail not found")
}

// UPDATE CONTACT INFO
func (c *Contact) UpdateContactInfo(infoID int, parameter string, newValue interface{}) error {
	if !c.IsActive {
		return errors.New("contact is inactive, cannot update infos")
	}

	for _, info := range c.ContactInfos {
		if info.ContactInfoID == infoID {
			return info.UpdateContactInfo(parameter, newValue) // actual updatecontactinfo logic in contactinfo package
		}
	}
	return errors.New("contact info not found")
}

func (c *Contact) PrintDetails() {
	fmt.Printf("ContactID: %d\n", c.ContactID)
	fmt.Printf("Name: %s %s\n", c.Fname, c.Lname)
	fmt.Printf("IsActive: %t\n", c.IsActive)
	fmt.Println("Contact Details:")

	for _, detail := range c.ContactInfos {
		detail.PrintDetails()
	}
	fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^")
}
