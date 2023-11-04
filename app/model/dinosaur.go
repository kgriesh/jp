package model

type Dinosaur struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type Dinosaurs struct {
	Dinosaur []Dinosaur `json:"dinosaur"`
}

// func (p *Dinosaur) Bind(r *http.Request) error {
// 	if p.Name == "" {
// 		return fmt.Errorf("name is a required field")
// 	}
// 	return nil
// }
// func (*Dinosaurs) Render(w http.ResponseWriter, r *http.Request) error {
// 	return nil
// }
// func (*Dinosaur) Render(w http.ResponseWriter, r *http.Request) error {
// 	return nil
// }
