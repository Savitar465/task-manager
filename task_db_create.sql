-- Created by Vertabelo (http://vertabelo.com)
-- Last modification date: 2024-07-23 06:37:31.625

-- tables
-- Table: board_stages
CREATE TABLE board_stages (
    board_stage_id varchar(200)  NOT NULL,
    stages_stage_id varchar(200)  NOT NULL,
    boards_board_id varchar(200)  NOT NULL,
    CONSTRAINT board_stages_pk PRIMARY KEY (board_stage_id)
);

CREATE INDEX boards_idx_board_stages ON board_stages (board_stage_id);

CREATE INDEX stages_idx_board_stages ON board_stages (stages_stage_id);

-- Table: boards
CREATE TABLE boards (
    board_id varchar(200)  NOT NULL,
    name varchar(200)  NOT NULL,
    _status int  NOT NULL,
    created_by varchar(200)  NOT NULL,
    created_at datetime  NOT NULL,
    updated_by varchar(200)  NOT NULL,
    updated_at datetime  NOT NULL,
    CONSTRAINT boards_pk PRIMARY KEY (board_id)
);

CREATE INDEX boards_idx_1 ON boards (board_id);

-- Table: issue_links
CREATE TABLE issue_links (
    issue_link_id varchar(200)  NOT NULL,
    issues_issue_id bigint  NOT NULL,
    issues_issue_dependent_id bigint  NOT NULL,
    created_by varchar(200)  NOT NULL,
    created_at datetime  NOT NULL,
    updated_by varchar(200)  NOT NULL,
    updated_at datetime  NOT NULL,
    CONSTRAINT issue_links_pk PRIMARY KEY (issue_link_id)
);

CREATE INDEX issues_idx_links ON issue_links (issues_issue_id);

CREATE INDEX issue_links_idx_dependent ON issue_links (issues_issue_dependent_id);

-- Table: issues
CREATE TABLE issues (
    issue_id bigint  NOT NULL,
    types_type_id varchar(200)  NOT NULL,
    title varchar(200)  NOT NULL,
    description varchar(500)  NOT NULL,
    start_date datetime  NOT NULL,
    due_date datetime  NOT NULL,
    stages_stage_id varchar(200)  NOT NULL,
    boards_board_id varchar(200)  NOT NULL,
    assigned_user_id varchar(200)  NOT NULL,
    _status int  NOT NULL,
    created_by varchar(200)  NOT NULL,
    created_at datetime  NOT NULL,
    updated_by varchar(200)  NOT NULL,
    updated_at datetime  NOT NULL,
    CONSTRAINT issues_pk PRIMARY KEY (issue_id)
);

CREATE INDEX types_idx_issues ON issues (types_type_id);

CREATE INDEX stages_idx_issues ON issues (stages_stage_id);

CREATE INDEX boards_idx_issues ON issues (boards_board_id);

-- Table: issues_variables
CREATE TABLE issues_variables (
    issue_variable_id varchar(200)  NOT NULL,
    issues_issue_id bigint  NOT NULL,
    variables_variable_id varchar(200)  NOT NULL,
    created_by varchar(200)  NOT NULL,
    created_at datetime  NOT NULL,
    updated_by varchar(200)  NOT NULL,
    updated_at datetime  NOT NULL,
    CONSTRAINT issues_variables_pk PRIMARY KEY (issue_variable_id)
);

CREATE INDEX issues_idx_issues_variables ON issues_variables (issues_issue_id);

CREATE INDEX variables_idx_issues_variables ON issues_variables (variables_variable_id);

-- Table: options
CREATE TABLE options (
    option_id varchar(200)  NOT NULL,
    variables_variable_id varchar(200)  NOT NULL,
    name varchar(500)  NOT NULL,
    _status int  NOT NULL,
    created_by varchar(200)  NOT NULL,
    created_at datetime  NOT NULL,
    updated_by varchar(200)  NOT NULL,
    updated_at datetime  NOT NULL,
    CONSTRAINT options_pk PRIMARY KEY (option_id)
);

CREATE INDEX variables_idx_opciones ON options (variables_variable_id);

-- Table: stages
CREATE TABLE stages (
    stage_id varchar(200)  NOT NULL,
    name varchar(200)  NOT NULL,
    _status int  NOT NULL,
    created_by varchar(200)  NOT NULL,
    created_at datetime  NOT NULL,
    updated_by varchar(200)  NOT NULL,
    updated_at datetime  NOT NULL,
    stage_type int  NOT NULL COMMENT 'finished=1, pending=0',
    CONSTRAINT stages_pk PRIMARY KEY (stage_id)
);

-- Table: types
CREATE TABLE types (
    type_id varchar(200)  NOT NULL,
    type char(2)  NOT NULL,
    description varchar(500)  NOT NULL,
    _status int  NOT NULL,
    created_by varchar(200)  NOT NULL,
    created_at datetime  NOT NULL,
    updated_by varchar(200)  NOT NULL,
    updated_at datetime  NOT NULL,
    CONSTRAINT types_pk PRIMARY KEY (type_id)
);

-- Table: values
CREATE TABLE `values` (
    value_id varchar(200)  NOT NULL,
    issues_variables_values_id varchar(200)  NOT NULL,
    value varchar(800)  NOT NULL,
    CONSTRAINT values_pk PRIMARY KEY (value_id)
);

-- Table: variables
CREATE TABLE variables (
    variable_id varchar(200)  NOT NULL,
    type_variable varchar(20)  NOT NULL,
    name varchar(200)  NOT NULL,
    types_type_id varchar(200)  NOT NULL,
    required boolean  NOT NULL,
    multiple boolean  NOT NULL,
    _status int  NOT NULL,
    created_by varchar(200)  NOT NULL,
    created_at datetime  NOT NULL,
    updated_by varchar(200)  NOT NULL,
    updated_at datetime  NOT NULL,
    CONSTRAINT variables_pk PRIMARY KEY (variable_id)
);

CREATE INDEX types_idx_variables ON variables (types_type_id);

-- foreign keys
-- Reference: board_stages_boards (table: board_stages)
ALTER TABLE board_stages ADD CONSTRAINT board_stages_boards FOREIGN KEY board_stages_boards (boards_board_id)
    REFERENCES boards (board_id);

-- Reference: board_stages_stages (table: board_stages)
ALTER TABLE board_stages ADD CONSTRAINT board_stages_stages FOREIGN KEY board_stages_stages (stages_stage_id)
    REFERENCES stages (stage_id);

-- Reference: issue_variables_types (table: variables)
ALTER TABLE variables ADD CONSTRAINT issue_variables_types FOREIGN KEY issue_variables_types (types_type_id)
    REFERENCES types (type_id);

-- Reference: issues_boards (table: issues)
ALTER TABLE issues ADD CONSTRAINT issues_boards FOREIGN KEY issues_boards (boards_board_id)
    REFERENCES boards (board_id);

-- Reference: issues_stages (table: issues)
ALTER TABLE issues ADD CONSTRAINT issues_stages FOREIGN KEY issues_stages (stages_stage_id)
    REFERENCES stages (stage_id);

-- Reference: issues_types (table: issues)
ALTER TABLE issues ADD CONSTRAINT issues_types FOREIGN KEY issues_types (types_type_id)
    REFERENCES types (type_id);

-- Reference: issues_variables_issues (table: issues_variables)
ALTER TABLE issues_variables ADD CONSTRAINT issues_variables_issues FOREIGN KEY issues_variables_issues (issues_issue_id)
    REFERENCES issues (issue_id);

-- Reference: issues_variables_variables (table: issues_variables)
ALTER TABLE issues_variables ADD CONSTRAINT issues_variables_variables FOREIGN KEY issues_variables_variables (variables_variable_id)
    REFERENCES variables (variable_id);

-- Reference: links_issues (table: issue_links)
ALTER TABLE issue_links ADD CONSTRAINT links_issues FOREIGN KEY links_issues (issues_issue_id)
    REFERENCES issues (issue_id);

-- Reference: links_issues_dependent (table: issue_links)
ALTER TABLE issue_links ADD CONSTRAINT links_issues_dependent FOREIGN KEY links_issues_dependent (issues_issue_dependent_id)
    REFERENCES issues (issue_id);

-- Reference: options_variables (table: options)
ALTER TABLE options ADD CONSTRAINT options_variables FOREIGN KEY options_variables (variables_variable_id)
    REFERENCES variables (variable_id);

-- Reference: valores_issues_variables (table: values)
ALTER TABLE `values` ADD CONSTRAINT valores_issues_variables FOREIGN KEY valores_issues_variables (issues_variables_values_id)
    REFERENCES issues_variables (issue_variable_id);

-- End of file.

