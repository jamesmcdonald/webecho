package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

func Logger(c <-chan RequestInfo) error {
	hostname, _ := os.Hostname()

	// Get configuration from the environment
	// TODO check that these are actually set

	var protocol string
	var address string
	if instance := os.Getenv("GCLOUD_SQL"); instance == "" {
		protocol = "tcp"
		port := os.Getenv("MYSQL_PORT")
		if port == "" {
			port = "3306"
		}
		address = fmt.Sprintf("%s:%s", os.Getenv("MYSQL_HOST"), port)
	} else {
		protocol = "unix"
		address = fmt.Sprintf("/opt/cloud_sql_proxy/run/%s:%s:%s",
			os.Getenv("GCLOUD_PROJECT"), os.Getenv("GCLOUD_LOCATION"), instance)
	}
	username := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	dbname := os.Getenv("MYSQL_DB")

	dsn := fmt.Sprintf("%s:%s@%s(%s)/%s", username, password, protocol, address, dbname)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	for {
		ri := <-c
		// Convert to unix time and from unix time. This will result in a UTC timestamp in the database.
		_, err := db.Exec("insert into echostats values (from_unixtime(?), ?, ?, ?, ?)", ri.timestamp.Unix(), ri.url, hostname, ri.status, ri.duration)
		if err != nil {
			log.Printf("Logger failed to insert: %s", err)
		}
	}
}
