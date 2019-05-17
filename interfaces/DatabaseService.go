package interfaces

type DatabaseService interface {
	Migrate()
	DropTables()
}
