CREATE DATABASE chatbot;

\c chatbot;

CREATE TABLE workflows (
   id UUID PRIMARY KEY,
   type VARCHAR,
   createdAt TIMESTAMP
);

CREATE TABLE finalized_workflows (
   workflow_id UUID REFERENCES workflows(id),
   createdAt TIMESTAMP
);

CREATE TABLE steps (
   id UUID PRIMARY KEY,
   workflow_id UUID REFERENCES workflows(id),
   step_order INT,
   createdAt TIMESTAMP
);

CREATE TABLE answers (
   step_id UUID REFERENCES steps(id),
   message VARCHAR
);

CREATE TABLE reviews (
 id UUID PRIMARY KEY,
 workflow_id UUID REFERENCES workflows(id),
 product_name VARCHAR,
 review_text VARCHAR,
 rating INT
); 
