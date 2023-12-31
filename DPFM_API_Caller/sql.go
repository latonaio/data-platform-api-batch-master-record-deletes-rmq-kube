package dpfm_api_caller

import (
	dpfm_api_input_reader "data-platform-api-batch-master-record-deletes-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-batch-master-record-deletes-rmq-kube/DPFM_API_Output_Formatter"
	"fmt"
	"strings"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) BatchRead(
	input *dpfm_api_input_reader.SDC,
	log *logger.Logger,
) *dpfm_api_output_formatter.Batch {

	where := strings.Join([]string{
		fmt.Sprintf("WHERE Batch.Product = \"%s\" ", input.Batch.Product),
		fmt.Sprintf("AND Batch.BusinessPartner = %d ", input.Batch.BusinessPartner),
		fmt.Sprintf("AND Batch.Plant = \"%s\" ", input.Batch.Plant),
		fmt.Sprintf("AND Batch.Batch = \"%s\" ", input.Batch.Batch),
	}, "")

	rows, err := c.db.Query(
		`SELECT 
    	batch.product,
    	batch.businessPartner,
		batch.plant,
    	batch.batch,
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_batch_master_record_general_data as batch
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
