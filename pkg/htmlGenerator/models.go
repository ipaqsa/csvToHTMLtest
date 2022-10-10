package htmlGenerator

type Workbook struct {
	Name        string `csv:"0" json:"name"`
	Address     string `csv:"1" json:"address"`
	Postcode    string `csv:"2" json:"postcode"`
	Phone       string `csv:"3" json:"phone"`
	CreditLimit string `csv:"4" json:"creditlimit"`
	Birthday    string `csv:"5" json:"birthday"`
}
