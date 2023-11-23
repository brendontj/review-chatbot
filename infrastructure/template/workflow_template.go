package template

type WorkflowStep struct {
	Order int
	Text string
}

type WorkflowTemplate struct {
	Type  string         
	Steps []WorkflowStep 
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
					Text: "Hello, I'm a review chatbot. I will help you to get a review from your customer. What is your name?",
				},
				{
					Order: 1,
					Text: "What is the product name?",
				},
				{
					Order: 2,
					Text: "In a scale of 1 to 5, how do you rate our product?",
				},
				{
					Order: 3,
					Text: "Please write your review here",
				},
				{
					Order: 4,
					Text: "Thank you for your review. We really appreciate it.",
				},
			},
		},
	}
}
