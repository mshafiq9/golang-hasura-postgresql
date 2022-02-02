SET check_function_bodies = false;
CREATE SCHEMA hasurapg;
CREATE TABLE hasurapg.articles (
    id integer NOT NULL,
    title text NOT NULL,
    description text NOT NULL,
    content text NOT NULL
);
CREATE SEQUENCE hasurapg.articles_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER SEQUENCE hasurapg.articles_id_seq OWNED BY hasurapg.articles.id;
ALTER TABLE ONLY hasurapg.articles ALTER COLUMN id SET DEFAULT nextval('hasurapg.articles_id_seq'::regclass);
ALTER TABLE ONLY hasurapg.articles
    ADD CONSTRAINT articles_pkey PRIMARY KEY (id);
