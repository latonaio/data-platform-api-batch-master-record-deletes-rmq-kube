package dpfm_api_caller

import (
	dpfm_api_input_reader "data-platform-api-batch-master-record-deletes-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-batch-master-record-deletes-rmq-kube/DPFM_API_Output_Formatter"
	"fmt"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) BatchRead(
	input *dpfm_api_input_reader.SDC,
	log *logger.Logger,
) *dpfm_api_output_formatter.Batch {
	where := fmt.Sprintf("WHERE batch.Product = \"%s\"", input.Batch.Product)
	where := fmt.Sprintf("WHERE batch.BusinessPartner = %d", input.Batch.BusinessPartner)
	where := fmt.Sprintf("WHERE batch.Plant = \"%s\"", input.Batch.Plant)
	where := fmt.Sprintf("WHERE batch.Batch = \"%s\"", input.Batch.Batch)
	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_batch_master_record_batch_data as batch 
		` + where + ` ;`)
	if err != nil {
		log.Error("%+v", err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToBatch(rows)
	if err != nil {
		log.Error("%+v", err)
		return nil
	}

	return data
}
