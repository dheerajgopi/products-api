package main

func main() {
	app := App{}
	// app.Initialize(
	// 	os.Getenv("APP_DB_USERNAME"),
	// 	os.Getenv("APP_DB_PASSWORD"),
	// 	os.Getenv("APP_DB_NAME"),
	// )
	app.Initialize(
		"postgres",
		"docker",
		"golangdb",
	)

	app.Run(":8080")
}
