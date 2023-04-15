SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: prayer_times; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.prayer_times (
    date date NOT NULL,
    zone character varying(255) NOT NULL,
    imsak timestamp without time zone NOT NULL,
    fajr timestamp without time zone NOT NULL,
    syuruk timestamp without time zone NOT NULL,
    dhuhr timestamp without time zone NOT NULL,
    asr timestamp without time zone NOT NULL,
    maghrib timestamp without time zone NOT NULL,
    isha timestamp without time zone NOT NULL
);


--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.schema_migrations (
    version character varying(128) NOT NULL
);


--
-- Name: prayer_times prayer_times_pk; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.prayer_times
    ADD CONSTRAINT prayer_times_pk PRIMARY KEY (date, zone);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- PostgreSQL database dump complete
--


--
-- Dbmate schema migrations
--

INSERT INTO public.schema_migrations (version) VALUES
    ('20230415122935');
