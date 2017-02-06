package viewmodels

/*Index... required data to be passed to view layer*/
type Index struct {
	Title  string
	Active string
	Banner string
}

/*GetIndex... Get the needed data for the home page*/
func GetIndex() Index {
	result := Index{
		Title:  "Watches an E-Commerce online Shopping Category Flat Bootstrap Responsive Website Template| Home :: w3layouts",
		Active: "Index",
		Banner: "Index",
	}
	return result
}
