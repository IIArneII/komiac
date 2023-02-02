package entities

type Application struct {
	ID          int
	Department  Department
	Mnemocode   string
	Medicaments []Medicament
}
