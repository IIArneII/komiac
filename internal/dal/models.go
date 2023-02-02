package dal

type Medicament struct {
	SMNN         string `db:"smnn"`
	MNN          string `db:"mnn"`
	Form         string `db:"mnn"`
	Dosage       string `db:"form"`
	ConsumerUnit string `db:"consumer_unit"`
}

type Customer struct {
	OID string `db:"oid"`
}

type Department struct {
	OID         string `db:"oid"`
	CustomerOID string `db:"customer_oid"`
}

type Application struct {
	ID            int    `db:"id"`
	DepartmentOID string `db:"department_oid"`
	Mnemocode     string `db:"mnemocode"`
}
