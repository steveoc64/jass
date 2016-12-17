--
-- PostgreSQL database dump
--

SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'SQL_ASCII';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = public, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: address; Type: TABLE; Schema: public; Owner: steve; Tablespace: 
--

CREATE TABLE address (
    id integer NOT NULL,
    date timestamp with time zone DEFAULT ('now'::text)::timestamp without time zone NOT NULL,
    user_id integer DEFAULT 0 NOT NULL,
    country text DEFAULT ''::text NOT NULL,
    state_code text DEFAULT ''::text NOT NULL,
    postcode text DEFAULT ''::text NOT NULL,
    name text DEFAULT ''::text NOT NULL,
    address_1 text DEFAULT ''::text NOT NULL,
    address_2 text DEFAULT ''::text NOT NULL
);


ALTER TABLE public.address OWNER TO steve;

--
-- Name: address_id_seq; Type: SEQUENCE; Schema: public; Owner: steve
--

CREATE SEQUENCE address_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.address_id_seq OWNER TO steve;

--
-- Name: address_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: steve
--

ALTER SEQUENCE address_id_seq OWNED BY address.id;


--
-- Name: blog; Type: TABLE; Schema: public; Owner: steve; Tablespace: 
--

CREATE TABLE blog (
    id integer NOT NULL,
    image text DEFAULT ''::text NOT NULL,
    date timestamp with time zone DEFAULT ('now'::text)::timestamp without time zone NOT NULL,
    name text DEFAULT ''::text NOT NULL,
    title text DEFAULT ''::text NOT NULL,
    content text DEFAULT ''::text NOT NULL,
    share_twitter integer DEFAULT 0 NOT NULL,
    share_facebook integer DEFAULT 0 NOT NULL,
    share_instagram integer DEFAULT 0 NOT NULL,
    share_google_plus integer DEFAULT 0 NOT NULL
);


ALTER TABLE public.blog OWNER TO steve;

--
-- Name: blog_id_seq; Type: SEQUENCE; Schema: public; Owner: steve
--

CREATE SEQUENCE blog_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.blog_id_seq OWNER TO steve;

--
-- Name: blog_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: steve
--

ALTER SEQUENCE blog_id_seq OWNED BY blog.id;


--
-- Name: item; Type: TABLE; Schema: public; Owner: steve; Tablespace: 
--

CREATE TABLE item (
    id integer NOT NULL,
    sku text DEFAULT ''::text NOT NULL,
    price numeric(8,2),
    name text DEFAULT ''::text NOT NULL,
    descr text DEFAULT ''::text NOT NULL,
    image text DEFAULT ''::text NOT NULL,
    volume_ml integer NOT NULL,
    weight_g integer NOT NULL,
    shipping_volume_ml integer NOT NULL,
    shipping_weight_g integer NOT NULL
);


ALTER TABLE public.item OWNER TO steve;

--
-- Name: item_id_seq; Type: SEQUENCE; Schema: public; Owner: steve
--

CREATE SEQUENCE item_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.item_id_seq OWNER TO steve;

--
-- Name: item_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: steve
--

ALTER SEQUENCE item_id_seq OWNED BY item.id;


--
-- Name: sale_order; Type: TABLE; Schema: public; Owner: steve; Tablespace: 
--

CREATE TABLE sale_order (
    id integer NOT NULL,
    date timestamp with time zone DEFAULT ('now'::text)::timestamp without time zone NOT NULL,
    user_id integer DEFAULT 0 NOT NULL,
    address_id integer DEFAULT 0 NOT NULL
);


ALTER TABLE public.sale_order OWNER TO steve;

--
-- Name: sale_order_id_seq; Type: SEQUENCE; Schema: public; Owner: steve
--

CREATE SEQUENCE sale_order_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.sale_order_id_seq OWNER TO steve;

--
-- Name: sale_order_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: steve
--

ALTER SEQUENCE sale_order_id_seq OWNED BY sale_order.id;


--
-- Name: sale_order_item; Type: TABLE; Schema: public; Owner: steve; Tablespace: 
--

CREATE TABLE sale_order_item (
    order_id integer NOT NULL,
    item_id integer NOT NULL,
    qty integer NOT NULL
);


ALTER TABLE public.sale_order_item OWNER TO steve;

--
-- Name: tags; Type: TABLE; Schema: public; Owner: steve; Tablespace: 
--

CREATE TABLE tags (
    blog_id integer NOT NULL,
    tag text
);


ALTER TABLE public.tags OWNER TO steve;

--
-- Name: users; Type: TABLE; Schema: public; Owner: steve; Tablespace: 
--

CREATE TABLE users (
    id integer NOT NULL,
    name text DEFAULT ''::text NOT NULL,
    date timestamp with time zone DEFAULT ('now'::text)::timestamp without time zone NOT NULL,
    last_ip text DEFAULT ''::text NOT NULL,
    has_stripe boolean DEFAULT false NOT NULL,
    stripe_token text DEFAULT ''::text NOT NULL,
    has_paypal boolean DEFAULT false NOT NULL,
    paypal_name text DEFAULT ''::text NOT NULL,
    country text DEFAULT 'AU'::text NOT NULL,
    email text DEFAULT ''::text NOT NULL,
    mobile text DEFAULT ''::text NOT NULL,
    gender text DEFAULT 'F'::text NOT NULL,
    newsletter boolean DEFAULT true,
    sms_alerts boolean DEFAULT false
);


ALTER TABLE public.users OWNER TO steve;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: steve
--

CREATE SEQUENCE users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO steve;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: steve
--

ALTER SEQUENCE users_id_seq OWNED BY users.id;


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: steve
--

ALTER TABLE ONLY address ALTER COLUMN id SET DEFAULT nextval('address_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: steve
--

ALTER TABLE ONLY blog ALTER COLUMN id SET DEFAULT nextval('blog_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: steve
--

ALTER TABLE ONLY item ALTER COLUMN id SET DEFAULT nextval('item_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: steve
--

ALTER TABLE ONLY sale_order ALTER COLUMN id SET DEFAULT nextval('sale_order_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: steve
--

ALTER TABLE ONLY users ALTER COLUMN id SET DEFAULT nextval('users_id_seq'::regclass);


--
-- Data for Name: address; Type: TABLE DATA; Schema: public; Owner: steve
--

COPY address (id, date, user_id, country, state_code, postcode, name, address_1, address_2) FROM stdin;
\.


--
-- Name: address_id_seq; Type: SEQUENCE SET; Schema: public; Owner: steve
--

SELECT pg_catalog.setval('address_id_seq', 1, false);


--
-- Data for Name: blog; Type: TABLE DATA; Schema: public; Owner: steve
--

COPY blog (id, image, date, name, title, content, share_twitter, share_facebook, share_instagram, share_google_plus) FROM stdin;
4	model-004.jpg	2016-12-16 17:00:59.789943+10:30	The perfect imperfections.	True essense of beauty.	There is nothing more rare, nor more beautiful, than a woman being unapologetically herself; comfortable in her perfect imperfection. To me, that is the true essence of beauty.	0	0	0	0
2	model-005.jpg	2016-12-16 16:43:52.52166+10:30	I stared at her	When I see you the world stops.	The first time I saw you, my heart fell. The second time I saw you, my heart fell. The third time fourth time fifth time and every time since, my heart has fallen.\nI stared at her.\nYou are the most beautiful woman I have ever seen. Your hair, your eyes, your lips, your body that you haven’t grown into, the way you walk, smile, laugh, the way your cheeks drop when you’re mad or upset, the way you drag your feet when you’re tired. Every single thing about you is beautiful.\nI stared at her.\nWhen I see you the world stops. It stops and all that exists for me is you and my eyes staring at you. There’s nothing else. No noise, no other people, no thoughts or worries, no yesterday, no tomorrow. The world just stops and it is a beautiful place and there is only you. Just you, and my eyes staring at you.\nI stared.\nWhen you’re gone, the world starts again, and I don’t like it as much. I can live in it, but I don’t like it. I just walk around in it and wait to see you again and wait for it to stop again. I love it when it stops. It’s the best ******* thing I’ve ever known or ever felt, the best thing, and that, beautiful Girl, is why I stare at you.	0	0	0	0
3	model-008.jpg	2016-12-16 16:53:46.398967+10:30	Destroying the Spell.	He did not dare approach her.	To him she seemed so beautiful, so seductive, so different from ordinary people, that he could not understand why no one was as disturbed as he by the clicking of her heels on the paving stones, why no one else’s heart was wild with the breeze stirred by the sighs of her veils, why everyone did not go mad with the movements of her braid, the flight of her hands, the gold of her laughter. He had not missed a single one of her gestures, not one of the indications of her character, but he did not dare approach her for fear of destroying the spell.	0	0	0	0
1	model-000.jpg	2016-12-16 16:38:30.312353+10:30	She was art.	She was my personal gallery.	She was art. The type of art that makes you feel a thousand things at the same time. The type of art that everyone wants to see. Her smile more worth than a Picasso. Eyes gorgeous like the starry night. My Mona Lisa. She was my personal gallery.	0	0	0	0
5	model-003.jpg	2016-12-17 11:36:28.373041+10:30	Deep down to her soul.	No, she wasn’t beautiful.	She was beautiful, but not like those girls in the magazines. She was beautiful, for the way she thought. She was beautiful, for the sparkle in her eyes when she talked about something she loved. She was beautiful, for her ability to make other people smile even if she was sad. No, she wasn’t beautiful for something as temporary as her looks. She was beautiful, deep down to her soul.	0	0	0	0
\.


--
-- Name: blog_id_seq; Type: SEQUENCE SET; Schema: public; Owner: steve
--

SELECT pg_catalog.setval('blog_id_seq', 5, true);


--
-- Data for Name: item; Type: TABLE DATA; Schema: public; Owner: steve
--

COPY item (id, sku, price, name, descr, image, volume_ml, weight_g, shipping_volume_ml, shipping_weight_g) FROM stdin;
\.


--
-- Name: item_id_seq; Type: SEQUENCE SET; Schema: public; Owner: steve
--

SELECT pg_catalog.setval('item_id_seq', 1, false);


--
-- Data for Name: sale_order; Type: TABLE DATA; Schema: public; Owner: steve
--

COPY sale_order (id, date, user_id, address_id) FROM stdin;
\.


--
-- Name: sale_order_id_seq; Type: SEQUENCE SET; Schema: public; Owner: steve
--

SELECT pg_catalog.setval('sale_order_id_seq', 1, false);


--
-- Data for Name: sale_order_item; Type: TABLE DATA; Schema: public; Owner: steve
--

COPY sale_order_item (order_id, item_id, qty) FROM stdin;
\.


--
-- Data for Name: tags; Type: TABLE DATA; Schema: public; Owner: steve
--

COPY tags (blog_id, tag) FROM stdin;
2	#jass
2	#chamelee
2	#perfumes
2	#worldofjass
2	#myjass
2	#hidenomore
1	#jass
1	#chamelee
1	#perfumes
1	#worldofjass
1	#myjass
1	#hidenomore
3	#jass
3	#chamelee
3	#perfumes
3	#worldofjass
3	#myjass
3	#hidenomore
4	#jass
4	#chamelee
4	#perfumes
4	#worldofjass
4	#myjass
4	#hidenomore
5	#jass
5	#chamelee
5	#perfumes
5	#worldofjass
5	#myjass
5	#hidenomore
5	#jass
5	#chamelee
5	#perfumes
5	#worldofjass
5	#myjass
5	#hidenomore
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: steve
--

COPY users (id, name, date, last_ip, has_stripe, stripe_token, has_paypal, paypal_name, country, email, mobile, gender, newsletter, sms_alerts) FROM stdin;
\.


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: steve
--

SELECT pg_catalog.setval('users_id_seq', 1, false);


--
-- Name: address_pkey; Type: CONSTRAINT; Schema: public; Owner: steve; Tablespace: 
--

ALTER TABLE ONLY address
    ADD CONSTRAINT address_pkey PRIMARY KEY (id);


--
-- Name: blog_pkey; Type: CONSTRAINT; Schema: public; Owner: steve; Tablespace: 
--

ALTER TABLE ONLY blog
    ADD CONSTRAINT blog_pkey PRIMARY KEY (id);


--
-- Name: item_pkey; Type: CONSTRAINT; Schema: public; Owner: steve; Tablespace: 
--

ALTER TABLE ONLY item
    ADD CONSTRAINT item_pkey PRIMARY KEY (id);


--
-- Name: sale_order_pkey; Type: CONSTRAINT; Schema: public; Owner: steve; Tablespace: 
--

ALTER TABLE ONLY sale_order
    ADD CONSTRAINT sale_order_pkey PRIMARY KEY (id);


--
-- Name: users_pkey; Type: CONSTRAINT; Schema: public; Owner: steve; Tablespace: 
--

ALTER TABLE ONLY users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: item_sku; Type: INDEX; Schema: public; Owner: steve; Tablespace: 
--

CREATE UNIQUE INDEX item_sku ON item USING btree (sku);


--
-- Name: tags_tag; Type: INDEX; Schema: public; Owner: steve; Tablespace: 
--

CREATE INDEX tags_tag ON tags USING btree (tag);


--
-- Name: public; Type: ACL; Schema: -; Owner: postgres
--

REVOKE ALL ON SCHEMA public FROM PUBLIC;
REVOKE ALL ON SCHEMA public FROM postgres;
GRANT ALL ON SCHEMA public TO postgres;
GRANT ALL ON SCHEMA public TO PUBLIC;


--
-- PostgreSQL database dump complete
--

