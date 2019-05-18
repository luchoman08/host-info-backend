package interfaces
// DatabaseService provide methods for manage the database
// DatabaseService does not provide methods for manage the database info
type DatabaseService interface {
	Migrate()
	DropTables()
}
