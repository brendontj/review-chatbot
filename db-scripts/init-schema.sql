CREATE DATABASE chatbot;

\c chatbot;

CREATE TABLE workflows (
 id UUID PRIMARY KEY,
 session_id UUID,
 name VARCHAR,
 version UUID,
 createdAt TIMESTAMP,
 updatedAt TIMESTAMP
);

CREATE TABLE steps (
 id UUID PRIMARY KEY,
 step_order INT,
 createdAt TIMESTAMP,
 updatedAt TIMESTAMP
);

CREATE TABLE sentMessages (
 workflow_id UUID REFERENCES workflows(id),
 step_id UUID REFERENCES steps(id),
 PRIMARY KEY (workflow_id, step_id)
);

CREATE TABLE answers (
 workflow_id UUID REFERENCES workflows(id),
 step_id UUID REFERENCES steps(id),
 message VARCHAR,
 PRIMARY KEY (workflow_id, step_id)
);

CREATE TABLE reviews (
 id UUID PRIMARY KEY,
 workflow_id UUID REFERENCES workflows(id),
 product_name VARCHAR,
 review_text VARCHAR,
 rating INT
); 
