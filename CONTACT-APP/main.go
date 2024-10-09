package main

import (
	"fmt"
	"contactApp/user"
)

func main() {
	// Create Admin user trial
	admin,err := user.NewAdmin("Jack", "Dawson")
	if err==nil{
		fmt.Println("Admin created:")

	}else{
		fmt.Println(err)

	}

	// Create Staff user using Admin trial
	staff1, errr := admin.NewStaff("Rose", "Smith")
	if errr == nil {
		fmt.Println("Staff created:")
		// staff1.PrintDetails()

	} else {
		fmt.Println(errr)
	}

	staff2, errr := admin.NewStaff("Lebron", "James")
	if errr == nil {
		fmt.Println("Staff created:")

	} else {
		fmt.Println(errr)
	}

	// Admin reads all staff trial
	fmt.Println("\nAdmin reading all STAFF:")
	users, err := admin.ReadUsers()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		for _, u := range users {
			u.PrintDetails()
		}
	}

	//admin updating user firstname trial
	fmt.Println("\nAdmin updating Staff user (changing Rose's firstname to 'Mary'):")
	err = admin.UpdateUserByParameter(1, "firstname", "Mary")
	if err != nil {
		fmt.Println("Error updating user:", err)
	} else {
		staff1.PrintDetails()
	}

	// Admin reads all staff trial
	fmt.Println("\nAdmin reading all STAFF:")
	users, err = admin.ReadUsers()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		for _, u := range users {
			u.PrintDetails()
		}
	}

	// Admin soft deletes a staff user
	fmt.Println("\nAdmin deleting a Staff user (Lebron):")
	err = admin.DeleteUser(staff2.UserID) //2
	if err != nil {
		fmt.Println("Error deleting user:", err)
	} else {
		fmt.Println("Staff user deleted.")
	}

	// Admin reads all users
	fmt.Println("\nAdmin reading all STAFF:")
	users, err = admin.ReadUsers()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		for _, u := range users {
			u.PrintDetails()
		}
	}

	//CONTACTS CRUD
	// Staff creates a new contact
	err = staff1.CreateContact("John", "Jacobs")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Contact created by Staff:")
		staff1.PrintDetails()
	}

	// Staff creates a new contact
	err = staff1.CreateContact("Jane", "Doe")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Contact created by Staff:")
		staff1.PrintDetails()
	}

	// Staff Reads contact
	contacts, err := staff1.ReadContacts()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(contacts)

	}

	//staff updates contact firstname
	err = staff1.UpdateContact(1, "firstname", "Max")
	if err != nil {
		fmt.Println("Error updating contact:", err)
	}

	// Staff updates contact lastname
	err = staff1.UpdateContact(1, "lastname", "Fosh")
	if err != nil {
		fmt.Println("Error updating contact:", err)
	}

	// Staff Read contact
	contacts, err = staff1.ReadContacts()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(contacts)

	}

	// Staff deletes a contact
	fmt.Println("Deleting contact with ID 0...")

	err = staff1.DeleteContact(0)
	if err != nil {
		fmt.Println("Error deleting contact:", err)
	}

	// Staff Reads contact
	contacts, err = staff1.ReadContacts()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(contacts)

	}

	// Staff creates contactinfo for contactid 1 Max Fosh
	err = staff1.CreateContactInfo(1, "Email", "max@example.com")
	if err != nil {
		fmt.Println("Error creating contact detail:", err)
	} else {
		fmt.Println("Contact detail created successfully!")
	}

	// Staff Read contact
	contacts, err = staff1.ReadContacts()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(contacts)

	}

	// staff read contact info
	info, err := staff1.ReadContactInfo(1, 0) // Read the info id 0 with contactID 1
	if err != nil {
		fmt.Println("Error reading contact detail:", err)
	} else {
		fmt.Printf("Contact Detail - ID: %d, Type: %s, Value: %s\n", info.ContactInfoID, info.Type, info.Value)
	}

	// Update the contact info's type
	err = staff1.UpdateContactInfo(1, 0, "type", "Mobile") // Update the infoid 0 with ID contact id 1
	if err != nil {
		fmt.Println("Error updating contact info:", err)
	} else {
		fmt.Println("Contact info updated successfully!")
	}
	// Update the contact info's value
	err = staff1.UpdateContactInfo(1, 0, "value", "1234567890") // Update the info id 0 with contactID 1
	if err != nil {
		fmt.Println("Error updating contact info:", err)
	} else {
		fmt.Println("Contact info updated successfully!")
	}

	// staff read the contact info
	info, err = staff1.ReadContactInfo(1, 0) // Read the detail with ID 1
	if err != nil {
		fmt.Println("Error reading contact detail:", err)
	} else {
		fmt.Printf("Contact Detail - ID: %d, Type: %s, Value: %s\n", info.ContactInfoID, info.Type, info.Value)
	}

}
