package database

var connection string

// tiap kali package ini diakses maka fungsi init ini akan dijalankan
func init() {
	connection = "MySQL"
}

func GetConnection() string {
	return connection
}
