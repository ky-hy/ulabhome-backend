package repository

// MySQLのエラーコード一覧
// https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html
const (
	// ErrCodeMySQLDuplicateEntry はMySQL系のDUPLICATEエラーコード(Error number: 1062; Symbol: ER_DUP_ENTRY; SQLSTATE: 23000)
	ErrCodeMySQLDuplicateEntry  = 1062
	ErrCodeMySQLNoReferencedRow = 1452
)
