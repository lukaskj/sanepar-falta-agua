package types

type TFaltaAguaResponse struct {
	Mensagem         string `json:"Mensagem,omitempty"`
	NormalizacaoData string `json:"NormalizacaoData,omitempty"`
	NormalizacaoHora string `json:"NormalizacaoHora,omitempty"`
	PrevisaoData     string `json:"PrevisaoData,omitempty"`
	PrevisaoHora     string `json:"PrevisaoHora,omitempty"`
}

func (t* TFaltaAguaResponse) GetId() string {
	return t.PrevisaoData + "-" + t.PrevisaoHora + "|" + t.NormalizacaoData + "-" + t.NormalizacaoHora
}