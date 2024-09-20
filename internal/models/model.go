package models

type Producto struct {
	prd_lvl_child  int         `json:"prd_lvl_child"`
	prd_lvl_number string      `json:"prd_lvl_number"`
	descripcion    string      `json:"descripcion"`
	proveedor      int         `json:"proveedor"`
	precio         int         `json:"precio"`
	dcreacion      string      `json:"dcreacion"`
	atributos      []Atributos `json:"atributos"`
}

type Atributos struct {
	tipo     string `json:"tipo"`
	subtipo  string `json:"subtipo"`
	atributo string `json:"atributo"`
}
