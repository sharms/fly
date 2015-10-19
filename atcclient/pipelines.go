package atcclient

import (
	"fmt"
	"net/http"

	"github.com/concourse/atc"
)

func (handler AtcHandler) ListPipelines() ([]atc.Pipeline, error) {
	var pipelines []atc.Pipeline
	err := handler.client.Send(Request{
		RequestName: atc.ListPipelines,
	}, Response{
		Result: &pipelines,
	})

	return pipelines, err
}

func (handler AtcHandler) DeletePipeline(pipelineName string) error {
	params := map[string]string{"pipeline_name": pipelineName}
	err := handler.client.Send(Request{
		RequestName: atc.DeletePipeline,
		Params:      params,
	}, Response{})

	if ure, ok := err.(UnexpectedResponseError); ok {
		if ure.StatusCode == http.StatusNotFound {
			return fmt.Errorf("`%s` does not exist", pipelineName)
		}
	}

	return err
}
