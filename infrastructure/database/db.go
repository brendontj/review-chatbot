package database

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/brendontj/review-chatbot/core/gateway"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
)

type Database struct {
	conn *pgx.Conn
}

func New() *Database {
	return &Database{}
}

func (d *Database) Connect() {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:pg123@localhost:5432/chatbot")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	d.conn = conn
}

func (d *Database) Disconnect() {
	d.conn.Close(context.Background())
}

func (d *Database) GetLastWorkflowNonFinalizedWithSteps() (gateway.WorkflowWithStepsModel, error) {
	queryLastWorkflow := `
     SELECT 
        workflows.id, 
        workflows.type, 
     FROM 
        workflows 
     ORDER BY 
        workflows.createdAt DESC 
     LIMIT 1
  `

	row, err := d.conn.Query(context.Background(), queryLastWorkflow)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	var workflow WorkflowWithStepsModel
	row.Scan(&workflow.ID, &workflow.TypeF)

	queryStepsFromWorkflow := `
	 SELECT
		steps.id,
		steps.step_order,
		answers.message
	 FROM
		steps
	 LEFT JOIN answers 
	 	ON steps.id = answers.step_id
	 WHERE steps.workflow_id = $1
	 ORDER BY steps.step_order ASC
	`

	rows, err := d.conn.Query(context.Background(), queryStepsFromWorkflow, workflow.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var step StepModel
		rows.Scan(&step.ID, &step.OrderF, &step.AnswerF)
		workflow.StepsF = append(workflow.StepsF, step)
	}

	return workflow, nil
}

func (d *Database) SaveStepAnswer(stepID uuid.UUID, answer string) error {
	query := `
    INSERT INTO answers (step_id, message) 
    VALUES ($1, $2)
 `

	_, err := d.conn.Exec(context.Background(), query, stepID, answer)
	if err != nil {
		return err
	}

	return nil
}

func (d *Database) SaveWorkflow(workflowType string) (gateway.WorkflowWithStepsModel, error) {
	query := `
   INSERT INTO workflows (id, type, createdAt) 
   VALUES ($1, $2, $3)
   RETURNING id, type, createdAt
 `
	id := uuid.New()

	row := d.conn.QueryRow(context.Background(), query, id, workflowType, time.Now())

	var workflow WorkflowWithStepsModel
	err := row.Scan(&workflow.ID, &workflow.TypeF, []StepModel{})
	if err != nil {
		return WorkflowWithStepsModel{}, err
	}

	return workflow, nil
}

func (d *Database) SaveStep(workflowID uuid.UUID, stepOrder int) error {
	query := `
  INSERT INTO steps (id, workflow_id, step_order, createdAt) 
  VALUES ($1, $2, $3, $4)
 `

	id := uuid.New()

	_, err := d.conn.Exec(context.Background(), query, id, workflowID, stepOrder, time.Now())
	if err != nil {
		return err
	}

	return nil
}

func (d *Database) SaveReview(workflow gateway.WorkflowWithStepsModel, reviewText string) error {
	query := `
 INSERT INTO reviews (id, workflow_id, product_name, review_text, rating) 
 VALUES ($1, $2, $3, $4, $5)
 `

	id := uuid.New()

	var (
		productName string
		rating      int
	)

	for _, s := range workflow.Steps() {
		switch s.Order() {
		case 0:
			productName = s.Answer()
		case 1:
			intAnswer, err := strconv.Atoi(s.Answer())
			if err != nil {
				return err
			}
			rating = intAnswer
		default:
			return fmt.Errorf("invalid step order")
		}

	}

	_, err := d.conn.Exec(context.Background(), query, id, workflow.Id(), productName, reviewText, rating)
	if err != nil {
		return err
	}

	return nil
}
