package viewmodels

//ProductDetail structure to specify details
type ProductDetail struct {
	Title       string
	Active      string
	Banner      string
	ImageURL    string
	Price       float32
	Description string
}

func GetProducts() ProductDetail {

}

//GetBrands...Brands collection
func GetProducts() Products {
	result := Brands{
		Title:  "Watches an E-Commerce online Shopping Category Flat Bootstrap Responsive Website Template| Home :: w3layouts",
		Active: "Brands",
		Banner: "Brands",
	}

	sday := Brand{
		ImageURL:    "2nd-day.jpg",
		Title:       "Lorem Ipsum",
		Description: "Lorem ipsum dolor sit amet, consectetuer adipiscing elit, sed diam nonummy nibh euismod tincidunt ut laoreet dolore magna aliquam erat volutpat. Ut wisi enim",
	}

	mih := Brand{
		ImageURL:    "mih-jeans.jpg",
		Title:       "Lorem Ipsum",
		Description: "Lorem ipsum dolor sit amet, consectetuer adipiscing elit, sed diam nonummy nibh euismod tincidunt ut laoreet dolore magna aliquam erat volutpat. Ut wisi enim",
	}

	gstar := Brand{
		ImageURL:    "g-star-raw.jpg",
		Title:       "Lorem Ipsum",
		Description: "Lorem ipsum dolor sit amet, consectetuer adipiscing elit, sed diam nonummy nibh euismod tincidunt ut laoreet dolore magna aliquam erat volutpat. Ut wisi enim",
	}

	weekday := Brand{
		ImageURL:    "weekday1.jpg",
		Title:       "Lorem Ipsum",
		Description: "Lorem ipsum dolor sit amet, consectetuer adipiscing elit, sed diam nonummy nibh euismod tincidunt ut laoreet dolore magna aliquam erat volutpat. Ut wisi enim",
	}

	result.Brands = []Brand{
		sday,
		mih,
		gstar,
		weekday,
	}

	return result
}
