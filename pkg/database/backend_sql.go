package database;

import (
	"fmt"
	"database/sql"
	"github.com/upsilonproject/upsilon-golib-database/pkg/models"

)

func PrepareMetricInsert(dbMetrics *sql.DB) (*sql.Stmt) {
	sql := "INSERT INTO metrics (extracted, updated, serviceCheckResultId, serviceIdentifier, name, value) VALUES (now(), ?, ?, ?, ?, ?)";
	stmt, err := dbMetrics.Prepare(sql);

	if (err != nil) {
		panic(err);
	}

	return stmt;
}

func InsertMetric(stmt *sql.Stmt, result models.ServiceResult, metric models.Metric) {
	fmt.Println("inserting metric", result.Id, "called", metric.Name, "updated", result.Updated);

	_, err := stmt.Exec(result.Updated, result.Id, result.Identifier, metric.Name, metric.Value)

	if (err != nil) {
		panic(err);
	}
}

