package gateway

import "github.com/brendontj/review-chatbot/infrastructure/template"

type Template interface {
	FindByType(t string) template.WorkflowTemplate
}
