package viewmodels

//Brands Brands Model
type Brands struct {
	Title  string
	Active string
	Banner string
	Brands []Brand
}

//Brand Brand type
type Brand struct {
	ImageURL    string
	Title       string
	Description string
	id          int
}

//GetBrands function returns Brands collection
func GetBrands() Brands {
	result := Brands{
		Title:  "Watches an E-Commerce online Shopping Category Flat Bootstrap Responsive Website Template| Home :: w3layouts",
		Active: "Brands",
		Banner: "Brands",
	}

	sday := Brand{
		id:          1,
		ImageURL:    "2nd-day.jpg",
		Title:       "Lorem Ipsum",
		Description: "Lorem ipsum dolor sit amet, consectetuer adipiscing elit, sed diam nonummy nibh euismod tincidunt ut laoreet dolore magna aliquam erat volutpat. Ut wisi enim",
	}

	mih := Brand{
		id:          2,
		ImageURL:    "mih-jeans.jpg",
		Title:       "Lorem Ipsum",
		Description: "Lorem ipsum dolor sit amet, consectetuer adipiscing elit, sed diam nonummy nibh euismod tincidunt ut laoreet dolore magna aliquam erat volutpat. Ut wisi enim",
	}

	gstar := Brand{
		id:          3,
		ImageURL:    "g-star-raw.jpg",
		Title:       "Lorem Ipsum",
		Description: "Lorem ipsum dolor sit amet, consectetuer adipiscing elit, sed diam nonummy nibh euismod tincidunt ut laoreet dolore magna aliquam erat volutpat. Ut wisi enim",
	}

	weekday := Brand{
		id:          4,
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
