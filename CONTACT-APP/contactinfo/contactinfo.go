package contactinfo

import (
	"errors"
	"fmt"
)

type ContactInfo struct {
	ContactInfoID int
	Type          string
	Value         string
}

// NewContactInfo factory function 
func NewContactInfo(infotype, value string, id int) *ContactInfo {
	return &ContactInfo{
		ContactInfoID: id,
		Type:          infotype,
		Value:         value,
	}
}

// UpdateContactInfo method
func (info *ContactInfo) UpdateContactInfo(parameter string, newValue interface{}) error {
	switch parameter {
	case "type":
		newType, ok := newValue.(string)
		if !ok {
			return errors.New("invalid value type for type")
		}
		info.Type = newType

	case "value":
		newValueStr, ok := newValue.(string)
		if !ok {
			return errors.New("invalid value type for value")
		}
		info.Value = newValueStr

	default:
		return errors.New("invalid parameter")
	}
	return nil
}

func (info *ContactInfo) PrintDetails() {
	fmt.Printf("ContactInfoID: %d\n", info.ContactInfoID)
	fmt.Printf("Type: %s\n", info.Type)
	fmt.Printf("Value: %s\n", info.Value)
	fmt.Println("**************************")
}
