package db

import (
	"context"
	"database/sql"
	"fmt"
)

// store provides all functions to execute database queries and transactions.
// It also provides a way to create a new store with a transaction.
type Store struct {
	*Queries
	db *sql.DB
}

// NewStore creates a new store.
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// execTx executes a function within a database transaction.
// don't export this function, it's only used internally
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx) // create a new store with the transaction
	err = fn(q)  // execute the function with the new store
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	// commit the transaction
	return tx.Commit()
}

// TransferTxParams contains the input parameters for the TransferTx function.
type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

// TransferTxResult contains the result of the TransferTx function.
type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

// type of empty struct to use as a key for the context's values map
// the second bracket is a new empty obje of that type.
var txKey = struct{}{}

// TransferTx performs a money transfer from one account to another.
// It creates a transfer record and updates account balances within a database transaction.
// Why use closure? Because we want to use the same transaction for all queries.
func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		txName := ctx.Value(txKey)

		fmt.Println(txName, "create transfer")
		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		})

		if err != nil {
			return err
		}

		fmt.Println(txName, "create entry 1")
		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})

		if err != nil {
			return err
		}

		fmt.Println(txName, "create entry 2")
		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})

		if err != nil {
			return err
		}

		// Get account -> update its balance
		// // This is incorrectly without a proper locking mechanism
		// // make test:
		// // >> before transfer 236 632
		// // >> transfer 226 642
		// // >> transfer 226 652
		// //     store_test.go:93:
		// //                 Error Trace:    /Users/avery_yang/Workplace/go/src/simpleBank/db/sqlc/store_test.go:93
		// //                 Error:          Not equal:
		// //                                 expected: 10
		// //                                 actual  : 20
		// //                 Test:           TestTransferTX

		// // Because just a GetAccount Query ,so it doesn't block other transaction from reading the same account record.
		// // Therefore, the second transaction will read the same account record and update the balance based on the old value.

		// account1, err := q.GetAccount(ctx, arg.FromAccountID)
		// if err != nil {
		// 	return err
		// }
		// result.FromAccount, err = q.UpdateAccount(ctx, UpdateAccountParams{
		// 	ID:      arg.FromAccountID,
		// 	Balance: account1.Balance - arg.Amount,
		// })

		// account2, err := q.GetAccount(ctx, arg.ToAccountID)
		// if err != nil {
		// 	return err
		// }
		// result.ToAccount, err = q.UpdateAccount(ctx, UpdateAccountParams{
		// 	ID:      arg.ToAccountID,
		// 	Balance: account2.Balance + arg.Amount,
		// })
		fmt.Println(txName, "get account 1")
		account1, err := q.GetAccountForUpdate(ctx, arg.FromAccountID)
		if err != nil {
			return err
		}
		fmt.Println(txName, "update account 1")
		result.FromAccount, err = q.UpdateAccount(ctx, UpdateAccountParams{
			ID:      arg.FromAccountID,
			Balance: account1.Balance - arg.Amount,
		})

		fmt.Println(txName, "get account 2")
		account2, err := q.GetAccountForUpdate(ctx, arg.ToAccountID)
		if err != nil {
			return err
		}

		fmt.Println(txName, "update account 2")
		result.ToAccount, err = q.UpdateAccount(ctx, UpdateAccountParams{
			ID:      arg.ToAccountID,
			Balance: account2.Balance + arg.Amount,
		})

		return nil
	})
	return result, err
}
