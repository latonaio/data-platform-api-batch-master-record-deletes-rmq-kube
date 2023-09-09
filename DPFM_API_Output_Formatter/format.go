package dpfm_api_output_formatter

import (
	"database/sql"
	"fmt"
)

func ConvertToBatch(rows *sql.Rows) (*Batch, error) {
	defer rows.Close()
	batch := Batch{}
	i := 0

	for rows.Next() {
		i++
		err := rows.Scan(
			&batch.Product,
			&batch.BusinessPartner,
			&batch.Plant,
			&batch.Batch,
			&batch.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &batch, err
		}

	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &batch, nil
	}

	return &batch, nil
}
