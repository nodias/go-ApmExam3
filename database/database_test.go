package database

func Example_PostgresAccess_Get() {
	p := NewPostgreAccess()
	p.Get("nodias")
	//Output:
}
