-- +goose Up
-- +goose StatementBegin   
CREATE TABLE IF NOT EXISTS Category (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);
 
CREATE TABLE IF NOT EXISTS Section (
    id SERIAL PRIMARY KEY,
    category_id INTEGER REFERENCES Category(id),
    name VARCHAR(255) NOT NULL
);
 
CREATE TABLE IF NOT EXISTS Topic (
    id SERIAL PRIMARY KEY,
    section_id INTEGER REFERENCES Section(id),
    name VARCHAR(255) NOT NULL
);
 
CREATE TABLE IF NOT EXISTS Question (
    id SERIAL PRIMARY KEY,
    topic_id INTEGER REFERENCES Topic(id),
    question_title VARCHAR(255) NOT NULL,
    question_text TEXT NOT NULL,
    answer_text TEXT NOT NULL
); 
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin 
DROP TABLE IF EXISTS Question;
 
DROP TABLE IF EXISTS Topic;
 
DROP TABLE IF EXISTS Section;
 
DROP TABLE IF EXISTS Category;
  
-- +goose StatementEnd
