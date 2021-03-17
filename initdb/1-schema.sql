CREATE TABLE leason (
    id BIGSERIAL NOT NULL,
    name VARCHAR(32) NOT NULL,
    id_course BIGINT NOT NULL,
    schedule VARCHAR(32)
);


ALTER TABLE leason ADD CONSTRAINT leason_pkey PRIMARY KEY (id);

CREATE TABLE course (
    id BIGSERIAL NOT NULL,
    name VARCHAR(32) NOT NULL,
    id_user BIGINT NOT NULL,
    students jsonb NOT NULL,
    schedule VARCHAR(32)
);


ALTER TABLE course ADD CONSTRAINT course_pkey PRIMARY KEY (id);

CREATE TABLE user (
    id BIGSERIAL NOT NULL,
    email VARCHAR(32) NOT NULL,
    password VARCHAR(32),
    phone VARCHAR(32),
    data jsonb NOT NULL/* json with courses and status it */
);


ALTER TABLE user ADD CONSTRAINT user_pkey PRIMARY KEY (id);

CREATE TABLE additional_material (
    id BIGSERIAL NOT NULL,
    data jsonb NOT NULL,
    id_leason BIGINT NOT NULL
);


ALTER TABLE additional_material ADD CONSTRAINT additional_material_pkey PRIMARY KEY (id);

CREATE TABLE video (
    id BIGSERIAL NOT NULL,
    name VARCHAR(32) NOT NULL,
    id_leason BIGINT NOT NULL
);


ALTER TABLE video ADD CONSTRAINT video_pkey PRIMARY KEY (id);

CREATE TABLE email_template (
    id BIGSERIAL,
    name VARCHAR(32) NOT NULL,
    id_course BIGINT NOT NULL,
    text TEXT NOT NULL
);


ALTER TABLE email_template ADD CONSTRAINT email_template_pkey PRIMARY KEY (id);

ALTER TABLE leason ADD CONSTRAINT leason_id_course_fkey FOREIGN KEY (id_course) REFERENCES course(id);
ALTER TABLE course ADD CONSTRAINT course_id_user_fkey FOREIGN KEY (id_user) REFERENCES user(id);
ALTER TABLE additional_material ADD CONSTRAINT additional_material_id_leason_fkey FOREIGN KEY (id_leason) REFERENCES leason(id);
ALTER TABLE video ADD CONSTRAINT video_id_leason_fkey FOREIGN KEY (id_leason) REFERENCES leason(id);
ALTER TABLE email_template ADD CONSTRAINT email_template_id_course_fkey FOREIGN KEY (id_course) REFERENCES course(id);
