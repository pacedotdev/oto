package services

// Page describes a page of data.
type Page struct {
	// Cursor is the cursor to start at.
	// example: "cursor-123456"
	Cursor string
	// OrderField is the field to use to order the results.
	// example: "Name"
	OrderField string
	// OrderAsc is whether to order the field in an ascending order or not.
	// example: "ASC"
	// options: ["ASC", "DESC"]
	OrderAsc bool
}
