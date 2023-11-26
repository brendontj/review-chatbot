package template

type WorkflowStep struct {
	Order int
	Text  string
}

type WorkflowTemplate struct {
	Type  string
	Steps []WorkflowStep
}

func (wt WorkflowTemplate) ExistStepWithOrderEqual(orderIdentifier int) bool {
	for _, v := range wt.Steps {
		if v.Order == orderIdentifier {
			return true
		}
	}
	return false
}

type WorkflowsTemplate []WorkflowTemplate

func (w WorkflowsTemplate) FindByType(t string) WorkflowTemplate {
	for _, workflow := range w {
		if workflow.Type == t {
			return workflow
		}
	}

	return WorkflowTemplate{}
}

func GenerateWorkflowsTemplate() WorkflowsTemplate {
	return []WorkflowTemplate{
		{
			Type: "review-signal",
			Steps: []WorkflowStep{
				{
					Order: 0,

					Text: "Hello, I'm a review chatbot. I will help you with the product review from your customer. What is the product name?",
				},
				{
					Order: 1,
					Text:  "In a scale of 1 to 5, how do you rate our product?",
				},
				{
					Order: 2,
					Text:  "Please write your review here",
				},
			},
		},
	}
}
