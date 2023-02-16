CREATE TABLE logics  (
    id uuid PRIMARY KEY NOT NULL,
    expression text NOT NULL,
    expression_code text NOT NULL,
    created_at timestamptz NOT NULL DEFAULT NOW(),
    updated_at timestamptz NOT NULL DEFAULT NOW(),
    UNIQUE(expression, expression_code)
);