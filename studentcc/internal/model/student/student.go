package student

type Student struct {
	ID           string
	Name         string
	Email        string
	Phone        string
	Camp         string
	ProjectGrade float64
}

func New(ID, name, email, phone, camp string) *Student {
	return &Student{
		ID:    ID,
		Name:  name,
		Email: email,
		Phone: phone,
		Camp:  camp,
	}
}
