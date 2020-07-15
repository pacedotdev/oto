package services

// Page describes a page of data.
type Page struct {
	// Cursor is the cursor to start at.
	Cursor string
	// OrderField is the field to use to order the results.
	OrderField string
	// OrderAsc is whether to order the field in an ascending order or not.
	OrderAsc bool
}
