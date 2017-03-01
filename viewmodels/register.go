package viewmodels

// Register struct to work with registration form
type Register struct {
	Title  string
	Banner string
	Active string
	User   User
}

// User struct is an sub type to handle user information
type User struct {
	Fname      string
	Lname      string
	Email      string
	NewsLetter bool
	Password   string
}

// GetRegistration function to get and process the registration event of user
func GetRegistration() Register {
	result := Register{
		Title:  "Watches an E-Commerce online Shopping Category Flat Bootstrap Responsive Website Template| Home :: w3layouts",
		Active: "Register",
		Banner: "Register",
	}

	return result
}
