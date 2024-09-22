package main

func main() {
	app := App{}
	app.Initailize(DbUser, DbPassword, DbName)
	app.Run("localhost:8080")
}
