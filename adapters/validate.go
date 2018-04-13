package adapters

import (
	"encoding/json"
	"fmt"
	"github.com/tyndyll/pii/domain"
	"io/ioutil"
	"net/http"
)

type ValidationHTTPAdapter struct {
	ProcessorFunctions map[string]domain.Processor
}

func (adapter *ValidationHTTPAdapter) Validate(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	validationRequest := &ValidationRequest{}
	if err := json.Unmarshal(body, validationRequest); err != nil {
		http.Error(w, fmt.Sprintf("Error: could not unmarshal JSON. %s", err.Error()), http.StatusBadRequest)
		return
	}

	for _, processor := range validationRequest.Processors {
		if _, found := adapter.ProcessorFunctions[processor]; !found {
			http.Error(w, fmt.Sprintf("Error: processor %s does not exist", processor), http.StatusBadRequest)
			return
		}
	}

	response := &ValidationResponse{

		Warnings: []*domain.ProcessorResult{},
	}

	for _, processor := range validationRequest.Processors {
		for _, data := range validationRequest.Data {
			s := data.(string)
			result, err := adapter.ProcessorFunctions[processor].Process(&s, nil)
			if result != nil {
				response.Warnings = append(response.Warnings, result)
				response.ItemWarnings++
			} else if err != nil {
				response.ItemErrors++
			}
		}
		response.ItemCount++
	}

	data, err := json.Marshal(response)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: could not marshal JSON. %s", err.Error()), http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

func (adapter *ValidationHTTPAdapter) ProcessorInformation(w http.ResponseWriter, req *http.Request) {
	//parts := strings.Split(req.URL.Path, `/`)

}

type ValidationRequest struct {
	Processors []string
	Data       []interface{}
}

type ValidationResponse struct {
	ItemCount    int64
	ItemWarnings int64
	ItemErrors   int64
	Warnings     []*domain.ProcessorResult
}
