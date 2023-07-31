package main

import "jwt-auth/database"

func main() {
	// Initialize Database
	database.Connect("root:root@tcp(localhost:3306)/jwt_demo?parseTime=true")
	database.Migrate()
}
