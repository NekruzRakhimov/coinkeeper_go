package db

const (
	CreatingOperation = `INSERT INTO operations (operation_type_id, category_id, amount, description, user_id)
						VALUES ($1, $2, $3, $4, $5)`
	GetAllOperations = `SELECT o.id,
							   ot.title as operation_type,
							   oc.title as category,
							   o.amount,
							   o.created_at,
							   o.description,
							   u.full_name as user_full_name
						FROM operations o
								 JOIN operation_categories oc ON o.category_id = oc.id
								 JOIN operation_types ot ON o.operation_type_id = ot.id
								 JOIN users u ON o.user_id = u.id
						WHERE o.is_deleted = FALSE;`
	GetTotalAmountByOperationType = `SELECT sum(o.amount) FROM operations o WHERE operation_type_id = ?;`
	GetOperationByID              = `SELECT o.id,
										   o.operation_type_id,
										   o.category_id,
										   o.amount,
										   o.created_at,
										   o.description,
										   u.full_name as user_full_name
									FROM operations o
											 JOIN users u ON o.user_id = u.id
									WHERE o.is_deleted = FALSE
									  AND o.id = ?;`
	UpdateOperationByID = `UPDATE operations
							SET operation_type_id = $1,
								category_id       = $2,
								amount            = $3
							WHERE id = $4;`
	SoftDeleteOperationByID = `UPDATE operations SET is_deleted = true WHERE id = $1;`
)
