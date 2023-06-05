package database;

import (
	"database/sql"
	"github.com/upsilonproject/upsilon-golib-database/pkg/models"

)

func PrepareMetricInsert(dbMetrics *sql.DB) (*sql.Stmt) {
	sql := "INSERT INTO service_metrics (extracted, updated, serviceCheckResultId, serviceIdentifier, name, value) VALUES (now(), ?, ?, ?, ?, ?) ON DUPLICATE KEY UPDATE extracted = now(), updated = ?, value = ?";
	stmt, err := dbMetrics.Prepare(sql);

	if (err != nil) {
		panic(err);
	}

	return stmt;
}

func InsertMetric(stmt *sql.Stmt, result models.ServiceResult, metric models.Metric) {
	_, err := stmt.Exec(result.Updated, result.Id, result.Identifier, metric.Name, metric.Value, result.Updated, metric.Value)

	if (err != nil) {
		panic(err);
	}
}

