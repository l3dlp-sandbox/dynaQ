package dq

import (
	"context"
	"database/sql"

	conServ "github.com/syke99/dq/pkg/conn"
	dbServ "github.com/syke99/dq/pkg/db"
	stmntServ "github.com/syke99/dq/pkg/stmnt"
	txserv "github.com/syke99/dq/pkg/tx"
)

type Dq struct {
	db           *sql.DB
	dbService    dbServ.DataBase
	stmntService stmntServ.Statement
	txService    txserv.Transaction
	conService   conServ.Connection
}

func NewDq(db *sql.DB) Dq {
	dbService := dbServ.NewDbService(db)

	return Dq{
		db:           db,
		dbService:    dbService.(dbServ.DataBase),
		stmntService: stmntServ.Statement{},
		txService:    txserv.Transaction{},
		conService:   conServ.Connection{},
	}
}

func (dq Dq) NewDqPreparedStatement(query string) (Dq, error) {
	stm, err := dq.db.Prepare(query)
	if err != nil {
		return dq, err
	}

	stmntService := stmntServ.NewPreparedStatementService(stm)

	dq.stmntService = stmntService.(stmntServ.Statement)

	return dq, nil
}

func (dq Dq) NewDqPreparedStatementWithContext(ctx context.Context, query string) (Dq, error) {
	stm, err := dq.db.PrepareContext(ctx, query)
	if err != nil {
		return dq, err
	}

	stmntService := stmntServ.NewPreparedStatementService(stm)

	dq.stmntService = stmntService.(stmntServ.Statement)

	return dq, nil
}

func (dq Dq) NewDqTransaction(tx *sql.Tx) Dq {
	txService := txserv.NewTransactionService(tx)

	dq.txService = txService.(txserv.Transaction)

	return dq
}

func (dq Dq) NewDqConn(con *sql.Conn) Dq {
	conService := conServ.NewConnectionService(con)

	dq.conService = conService.(conServ.Connection)

	return dq
}

func (dq Dq) DatabaseQuery(query string, args ...interface{}) ([]map[string]interface{}, error) {
	return dq.dbService.Query(query, args)
}

func (dq Dq) DatabaseQueryRow(query string, args ...interface{}) (map[string]interface{}, error) {
	return dq.dbService.QueryRow(query, args)
}

func (dq Dq) DatabaseQueryContext(ctx context.Context, query string, args ...interface{}) ([]map[string]interface{}, error) {
	return dq.dbService.QueryWithContext(ctx, query, args)
}

func (dq Dq) DatabaseQueryRowContext(ctx context.Context, query string, args ...interface{}) (map[string]interface{}, error) {
	return dq.dbService.QueryRowWithContext(ctx, query, args)
}

func (dq Dq) PreparedStatementQuery(query string, args ...interface{}) ([]map[string]interface{}, error) {
	return dq.stmntService.Query(query, args)
}

func (dq Dq) PreparedStatementQueryRow(query string, args ...interface{}) (map[string]interface{}, error) {
	return dq.stmntService.QueryRow(query, args)
}

func (dq Dq) PreparedStatementQueryContext(ctx context.Context, query string, args ...interface{}) ([]map[string]interface{}, error) {
	return dq.stmntService.QueryWithContext(ctx, query, args)
}

func (dq Dq) PreparedStatementQueryRowContext(ctx context.Context, query string) (map[string]interface{}, error) {
	return dq.stmntService.QueryRowWithContext(ctx, query)
}

func (dq Dq) TransactionQuery(query string, args ...interface{}) ([]map[string]interface{}, error) {
	return dq.txService.Query(query, args)
}

func (dq Dq) TransactionQueryRow(query string, args ...interface{}) (map[string]interface{}, error) {
	return dq.txService.QueryRow(query, args)
}

func (dq Dq) TransactionQueryContext(ctx context.Context, query string, args ...interface{}) ([]map[string]interface{}, error) {
	return dq.txService.QueryWithContext(ctx, query, args)
}

func (dq Dq) TransactionQueryRowContext(ctx context.Context, query string, args ...interface{}) (map[string]interface{}, error) {
	return dq.txService.QueryRowWithContext(ctx, query, args)
}

func (dq Dq) ConnectionQueryContext(ctx context.Context, query string, args ...interface{}) ([]map[string]interface{}, error) {
	return dq.conService.QueryWithContext(ctx, query, args)
}

func (dq Dq) ConnectionQueryRowContext(ctx context.Context, query string, args ...interface{}) (map[string]interface{}, error) {
	return dq.conService.QueryRowWithContext(ctx, query, args)
}
