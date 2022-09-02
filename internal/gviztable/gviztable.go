package gviztable

type Cell struct {
	Font    string
	BgColor string
	Color   string
}

type Table struct {
	// specifies horizontal placement.
	// When an object is allocated more space than required,
	// this value determines where the extra space is placed left and right of the object.
	//
	// CENTER aligns the object in the center. (Default)
	// LEFT aligns the object on the left
	// RIGHT aligns the object on the right
	Align          string
	BgColor        string
	Border         string
	CellBorder     string
	CellPadding    string
	CellSpacing    string
	Color          string
	Columns        string
	FixedSize      bool
	GradientTangle string
	Height         string
	Href           string
	Id             string
	Port           string
	Rows           string
	Sides          string
	Style          string
	Target         string
	Title          string
	Tooltip        string
	VAlign         string
	Width          string
}

func New() Table {
	return Table{}
}
