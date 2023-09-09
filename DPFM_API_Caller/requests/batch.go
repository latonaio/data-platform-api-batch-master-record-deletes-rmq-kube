package requests

type Batch struct {
	Product             string `json:Product`
	BusinessPartner     int    `json:BusinessPartner`
	Plant               string `json:Plant`
	Batch               string `json:Batch`
	IsMarkedForDeletion *bool  `json:IsMarkedForDeletion`
}
