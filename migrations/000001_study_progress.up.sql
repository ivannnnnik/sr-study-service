CREATE TABLE IF NOT EXISTS study_progress (
    id               UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id          UUID        NOT NULL,
    question_id      UUID        NOT NULL,
    ease_factor      DOUBLE PRECISION NOT NULL DEFAULT 2.5,
    interval         INT         NOT NULL DEFAULT 0,
    repetitions      INT         NOT NULL DEFAULT 0,
    next_review_at   TIMESTAMP   NOT NULL DEFAULT now(),
    last_reviewed_at TIMESTAMP   NOT NULL DEFAULT now(),
    UNIQUE (user_id, question_id)
);