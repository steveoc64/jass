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


--
-- Name: fuzzystrmatch; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS fuzzystrmatch WITH SCHEMA public;


--
-- Name: EXTENSION fuzzystrmatch; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION fuzzystrmatch IS 'determine similarities and distance between strings';


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
-- Name: admin_login; Type: TABLE; Schema: public; Owner: steve; Tablespace: 
--

CREATE TABLE admin_login (
    user_id integer,
    date timestamp with time zone DEFAULT ('now'::text)::timestamp without time zone NOT NULL,
    ip_addr text,
    user_agent text,
    referrer text
);


ALTER TABLE public.admin_login OWNER TO steve;

--
-- Name: admin_users; Type: TABLE; Schema: public; Owner: steve; Tablespace: 
--

CREATE TABLE admin_users (
    id integer NOT NULL,
    username text,
    pw text
);


ALTER TABLE public.admin_users OWNER TO steve;

--
-- Name: admin_users_id_seq; Type: SEQUENCE; Schema: public; Owner: steve
--

CREATE SEQUENCE admin_users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.admin_users_id_seq OWNER TO steve;

--
-- Name: admin_users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: steve
--

ALTER SEQUENCE admin_users_id_seq OWNED BY admin_users.id;


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
    share_google_plus integer DEFAULT 0 NOT NULL,
    post_order integer DEFAULT 0 NOT NULL,
    archived boolean DEFAULT false
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
-- Name: carrier; Type: TABLE; Schema: public; Owner: steve; Tablespace: 
--

CREATE TABLE carrier (
    id integer NOT NULL,
    name text,
    descr text,
    contact_name text,
    phone text,
    email text,
    address text,
    www text
);


ALTER TABLE public.carrier OWNER TO steve;

--
-- Name: carrier_id_seq; Type: SEQUENCE; Schema: public; Owner: steve
--

CREATE SEQUENCE carrier_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.carrier_id_seq OWNER TO steve;

--
-- Name: carrier_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: steve
--

ALTER SEQUENCE carrier_id_seq OWNED BY carrier.id;


--
-- Name: category; Type: TABLE; Schema: public; Owner: steve; Tablespace: 
--

CREATE TABLE category (
    id integer NOT NULL,
    name text,
    descr text
);


ALTER TABLE public.category OWNER TO steve;

--
-- Name: category_id_seq; Type: SEQUENCE; Schema: public; Owner: steve
--

CREATE SEQUENCE category_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.category_id_seq OWNER TO steve;

--
-- Name: category_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: steve
--

ALTER SEQUENCE category_id_seq OWNED BY category.id;


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
-- Name: newsletter; Type: TABLE; Schema: public; Owner: steve; Tablespace: 
--

CREATE TABLE newsletter (
    id integer NOT NULL,
    date timestamp with time zone DEFAULT ('now'::text)::timestamp without time zone NOT NULL,
    name text,
    content text
);


ALTER TABLE public.newsletter OWNER TO steve;

--
-- Name: newsletter_id_seq; Type: SEQUENCE; Schema: public; Owner: steve
--

CREATE SEQUENCE newsletter_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.newsletter_id_seq OWNER TO steve;

--
-- Name: newsletter_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: steve
--

ALTER SEQUENCE newsletter_id_seq OWNED BY newsletter.id;


--
-- Name: postage; Type: TABLE; Schema: public; Owner: steve; Tablespace: 
--

CREATE TABLE postage (
    shipping_code integer NOT NULL,
    region_id integer NOT NULL,
    date timestamp with time zone DEFAULT ('now'::text)::timestamp without time zone NOT NULL,
    cost numeric(12,2),
    price numeric(12,2),
    eta_days integer DEFAULT 5,
    max_days integer DEFAULT 30,
    has_tracking boolean DEFAULT false
);


ALTER TABLE public.postage OWNER TO steve;

--
-- Name: product; Type: TABLE; Schema: public; Owner: steve; Tablespace: 
--

CREATE TABLE product (
    id integer NOT NULL,
    cat_id integer NOT NULL,
    name text,
    descr text,
    image text,
    volume_ml integer DEFAULT 0,
    weight_g integer DEFAULT 0,
    shipping_volume_ml integer DEFAULT 0,
    shipping_weight_g integer DEFAULT 0,
    shipping_code integer,
    price numeric(12,2) DEFAULT 1.0
);


ALTER TABLE public.product OWNER TO steve;

--
-- Name: product_id_seq; Type: SEQUENCE; Schema: public; Owner: steve
--

CREATE SEQUENCE product_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.product_id_seq OWNER TO steve;

--
-- Name: product_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: steve
--

ALTER SEQUENCE product_id_seq OWNED BY product.id;


--
-- Name: ref_hit; Type: TABLE; Schema: public; Owner: steve; Tablespace: 
--

CREATE TABLE ref_hit (
    ref_id integer,
    date timestamp with time zone DEFAULT ('now'::text)::timestamp without time zone NOT NULL,
    ip_addr text,
    user_agent text
);


ALTER TABLE public.ref_hit OWNER TO steve;

--
-- Name: referrer; Type: TABLE; Schema: public; Owner: steve; Tablespace: 
--

CREATE TABLE referrer (
    id integer NOT NULL,
    name text,
    url text
);


ALTER TABLE public.referrer OWNER TO steve;

--
-- Name: referrer_id_seq; Type: SEQUENCE; Schema: public; Owner: steve
--

CREATE SEQUENCE referrer_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.referrer_id_seq OWNER TO steve;

--
-- Name: referrer_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: steve
--

ALTER SEQUENCE referrer_id_seq OWNED BY referrer.id;


--
-- Name: region; Type: TABLE; Schema: public; Owner: steve; Tablespace: 
--

CREATE TABLE region (
    id integer NOT NULL,
    name text,
    descr text
);


ALTER TABLE public.region OWNER TO steve;

--
-- Name: region_id_seq; Type: SEQUENCE; Schema: public; Owner: steve
--

CREATE SEQUENCE region_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.region_id_seq OWNER TO steve;

--
-- Name: region_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: steve
--

ALTER SEQUENCE region_id_seq OWNED BY region.id;


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
-- Name: shipping_code; Type: TABLE; Schema: public; Owner: steve; Tablespace: 
--

CREATE TABLE shipping_code (
    id integer NOT NULL,
    name text,
    descr text
);


ALTER TABLE public.shipping_code OWNER TO steve;

--
-- Name: shipping_code_id_seq; Type: SEQUENCE; Schema: public; Owner: steve
--

CREATE SEQUENCE shipping_code_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.shipping_code_id_seq OWNER TO steve;

--
-- Name: shipping_code_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: steve
--

ALTER SEQUENCE shipping_code_id_seq OWNED BY shipping_code.id;


--
-- Name: tags; Type: TABLE; Schema: public; Owner: steve; Tablespace: 
--

CREATE TABLE tags (
    blog_id integer NOT NULL,
    tag text
);


ALTER TABLE public.tags OWNER TO steve;

--
-- Name: up_hit; Type: TABLE; Schema: public; Owner: steve; Tablespace: 
--

CREATE TABLE up_hit (
    id integer NOT NULL,
    date timestamp with time zone DEFAULT ('now'::text)::timestamp without time zone NOT NULL,
    ip_addr text,
    user_agent text
);


ALTER TABLE public.up_hit OWNER TO steve;

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

ALTER TABLE ONLY admin_users ALTER COLUMN id SET DEFAULT nextval('admin_users_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: steve
--

ALTER TABLE ONLY blog ALTER COLUMN id SET DEFAULT nextval('blog_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: steve
--

ALTER TABLE ONLY carrier ALTER COLUMN id SET DEFAULT nextval('carrier_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: steve
--

ALTER TABLE ONLY category ALTER COLUMN id SET DEFAULT nextval('category_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: steve
--

ALTER TABLE ONLY item ALTER COLUMN id SET DEFAULT nextval('item_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: steve
--

ALTER TABLE ONLY newsletter ALTER COLUMN id SET DEFAULT nextval('newsletter_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: steve
--

ALTER TABLE ONLY product ALTER COLUMN id SET DEFAULT nextval('product_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: steve
--

ALTER TABLE ONLY referrer ALTER COLUMN id SET DEFAULT nextval('referrer_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: steve
--

ALTER TABLE ONLY region ALTER COLUMN id SET DEFAULT nextval('region_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: steve
--

ALTER TABLE ONLY sale_order ALTER COLUMN id SET DEFAULT nextval('sale_order_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: steve
--

ALTER TABLE ONLY shipping_code ALTER COLUMN id SET DEFAULT nextval('shipping_code_id_seq'::regclass);


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
-- Data for Name: admin_login; Type: TABLE DATA; Schema: public; Owner: steve
--

COPY admin_login (user_id, date, ip_addr, user_agent, referrer) FROM stdin;
\.


--
-- Data for Name: admin_users; Type: TABLE DATA; Schema: public; Owner: steve
--

COPY admin_users (id, username, pw) FROM stdin;
\.


--
-- Name: admin_users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: steve
--

SELECT pg_catalog.setval('admin_users_id_seq', 1, false);


--
-- Data for Name: blog; Type: TABLE DATA; Schema: public; Owner: steve
--

COPY blog (id, image, date, name, title, content, share_twitter, share_facebook, share_instagram, share_google_plus, post_order, archived) FROM stdin;
1	model-000.jpg	2016-12-16 16:38:30.312353+10:30	She was art.	She was my personal gallery.	She was art. The type of art that makes you feel a thousand things at the same time. The type of art that everyone wants to see. Her smile more worth than a Picasso. Eyes gorgeous like the starry night. My Mona Lisa. She was my personal gallery.	0	0	0	0	90	f
7	model-009.jpg	2016-12-17 12:03:05.087001+10:30	The Love Affair.	IT IS ALL TOO BEAUTIFUL A THING.	“I can’t over-emphasize how important an exquisite perfume is, to be wrapped and cradled in an enchanting scent upon your skin is a magic all on its own! The notes in that precious liquid will remind you that you love yourself and will tell other people that they ought to love you because you know that you’re worth it. The love affair created by a good perfume between you and other people, you and nature, you and yourself, you and your memories and anticipations and hopes and dreams; it is all too beautiful a thing!”	0	0	0	0	20	f
9	model-002.jpg	2016-12-17 12:06:27.462174+10:30	Flowers & Smokey Sandalwood.	“Her joyful spirit would bring laughter.”	“Her joyful spirit would bring laughter and happiness to anyone in her life with the same natural ease that a rose blooms and sheds its perfume.”	0	0	0	0	40	f
6	model-007.jpg	2016-12-17 12:01:48.85934+10:30	Voice of Assurance.	T’S ALWAYS WONDERFUL TO BE ELEGANT.	“It’s not very easy to grow up into a woman. We are always taught, almost bombarded, with ideals of what we should be at every age in our lives: “This is what you should wear at age twenty”, “That is what you must act like at age twenty-five”, “This is what you should be doing when you are seventeen.” But amidst all the many voices that bark all these orders and set all of these ideals for girls today, there lacks the voice of assurance. There is no comfort and assurance. I want to be able to say, that there are four things admirable for a woman to be, at any age! Whether you are four or forty-four or nineteen! It’s always wonderful to be elegant, it’s always fashionable to have grace, it’s always glamorous to be brave, and it’s always important to own a delectable perfume! Yes, wearing a beautiful fragrance is in style at any age!” 	0	0	0	0	10	f
4	model-004.jpg	2016-12-16 17:00:59.789943+10:30	The perfect imperfections.	True essense of beauty.	There is nothing more rare, nor more beautiful, than a woman being unapologetically herself; comfortable in her perfect imperfection. To me, that is the true essence of beauty.	0	0	0	0	50	f
5	model-003.jpg	2016-12-17 11:36:28.373041+10:30	Deep down to her soul.	No, she wasn’t beautiful.	She was beautiful, but not like those girls in the magazines. She was beautiful, for the way she thought. She was beautiful, for the sparkle in her eyes when she talked about something she loved. She was beautiful, for her ability to make other people smile even if she was sad. No, she wasn’t beautiful for something as temporary as her looks. She was beautiful, deep down to her soul.	0	0	0	0	60	f
3	model-008.jpg	2016-12-16 16:53:46.398967+10:30	Destroying the Spell.	He did not dare approach her.	To him she seemed so beautiful, so seductive, so different from ordinary people, that he could not understand why no one was as disturbed as he by the clicking of her heels on the paving stones, why no one else’s heart was wild with the breeze stirred by the sighs of her veils, why everyone did not go mad with the movements of her braid, the flight of her hands, the gold of her laughter. He had not missed a single one of her gestures, not one of the indications of her character, but he did not dare approach her for fear of destroying the spell.	0	0	0	0	70	f
8	model-006.jpg	2016-12-17 12:04:14.487447+10:30	Love, of Course, is Personal.	Many compositions smell great in the first few minutes.	If you’ve tried several perfumes, you know things can go wrong. Many compositions smell great in the first few minutes, then fade rapidly to a murmur or an unpleasant twang you can never quite wash off. Some seem to attack with what feels like an icepick in the eye. Others smell nice for an hour in the middle but boring at start and finish. Some veer uncomfortably sweet, and some fall to pieces, with various parts hanging there in the air but not really cooperating in any useful way. Some never get around to being much of anything at all. The way you can love a person for one quality despite myriad faults, you can sometimes love a perfume for one particular moment or effect, even if the rest is trash. Yet in the thousands of perfumes that exist, some express their ideas seamlessly and eloquently from top to bottom and give a beautiful view from any angle. A rare subset of them always seem to have something new and interesting to say, even if you encounter them daily. Those are the greats. By these criteria, one can certainly admire a perfume without necessarily loving it. Love, of course, is personal.	0	0	0	0	30	f
2	model-005.jpg	2016-12-16 16:43:52.52166+10:30	I stared at her	When I see you the world stops.	The first time I saw you, my heart fell. The second time I saw you, my heart fell. The third time fourth time fifth time and every time since, my heart has fallen.\nI stared at her.\nYou are the most beautiful woman I have ever seen. Your hair, your eyes, your lips, your body that you haven’t grown into, the way you walk, smile, laugh, the way your cheeks drop when you’re mad or upset, the way you drag your feet when you’re tired. Every single thing about you is beautiful.\nI stared at her.\nWhen I see you the world stops. It stops and all that exists for me is you and my eyes staring at you. There’s nothing else. No noise, no other people, no thoughts or worries, no yesterday, no tomorrow. The world just stops and it is a beautiful place and there is only you. Just you, and my eyes staring at you.\nI stared.\nWhen you’re gone, the world starts again, and I don’t like it as much. I can live in it, but I don’t like it. I just walk around in it and wait to see you again and wait for it to stop again. I love it when it stops. It’s the best ******* thing I’ve ever known or ever felt, the best thing, and that, beautiful Girl, is why I stare at you.	0	0	0	0	80	f
11	model-011.jpg	2017-01-06 15:23:59.455188+10:30	I miss you.	She is wishing for you to be her's forever	When she is quiet, millions of things are running in Her mind. When she is not arguing, She is in deep thought. When she looks at you with eyes full of questions, She is wondering how long you will be around. When she answers “I’m fine” after a few seconds, She is not fine. When you tell her why you are late, She is wondering why you are lying. When she lays on your chest, She is wishing for you to be hers forever. When she calls you everyday, She is seeking your attention. When she texts you everyday, She wants you to reply and tell her that she is missed. When she says “I love you”, She really means it. When she says that she can’t live without you, She has made up her mind that you are her future. When she says “I miss you”, no one in this world can miss you more than she does.\n\n	0	0	0	0	110	f
10	model-010.jpg	2017-01-06 15:22:07.080608+10:30	Lost in the sight of her.	Lost in my thoughts, lost in the sight of her.	It slowly began to dawn on me that I had been staring at her for an impossible amount of time. Lost in my thoughts, lost in the sight of her. But her face didn’t look offended or amused. It almost looked as if she were studying the lines of my face, almost as if she were waiting. I wanted to take her hand. I wanted to brush her cheek with my fingertips. I wanted to tell her that she was the most beautiful thing that I had ever seen. The sight of her yawning to the back of her hand was enough to drive the breath from me. How I sometimes lost the sense of her words in the sweet fluting of her voice. I wanted to say that if she were with me then somehow nothing could ever be wrong for me again. In that breathless second I almost asked her. I felt the question boiling up from my chest. I remember drawing a breath then hesitating–what could I say? Come away with me? Stay with me? Coffee? No. Sudden certainty tightened in my chest like a cold fist. What could I ask her? What could I offer? Nothing. Anything I said would sound foolish, a child’s fantasy. I closed my mouth and looked across the water. Inches away, Chamelee did the same. I could feel the heat of her. She smelled like road dust, and honey, and the smell the air holds seconds before a heavy summer rain.\n\nThe closeness of her was the sweetest, sharpest thing I had ever known.\n\nNeither of us spoke.\n\nI closed my eyes.\n\n	0	0	0	0	100	f
\.


--
-- Name: blog_id_seq; Type: SEQUENCE SET; Schema: public; Owner: steve
--

SELECT pg_catalog.setval('blog_id_seq', 15, true);


--
-- Data for Name: carrier; Type: TABLE DATA; Schema: public; Owner: steve
--

COPY carrier (id, name, descr, contact_name, phone, email, address, www) FROM stdin;
\.


--
-- Name: carrier_id_seq; Type: SEQUENCE SET; Schema: public; Owner: steve
--

SELECT pg_catalog.setval('carrier_id_seq', 1, false);


--
-- Data for Name: category; Type: TABLE DATA; Schema: public; Owner: steve
--

COPY category (id, name, descr) FROM stdin;
2	Skincare	Skincare Products
3	Merchandise	Other Chamelee branded products and merchandise
1	Fragrance	Eau de Parfum - for her, and for him
\.


--
-- Name: category_id_seq; Type: SEQUENCE SET; Schema: public; Owner: steve
--

SELECT pg_catalog.setval('category_id_seq', 6, true);


--
-- Data for Name: item; Type: TABLE DATA; Schema: public; Owner: steve
--

COPY item (id, sku, price, name, descr, image, volume_ml, weight_g, shipping_volume_ml, shipping_weight_g) FROM stdin;
2	JASS002	57.25	Chamelee	Chamelee - TBD	jass_hers-color.png	100	100	150	250
1	JASS001	58.25	Metaphor	Metaphor - ATBD	jass_his-color.png	100	100	150	250
\.


--
-- Name: item_id_seq; Type: SEQUENCE SET; Schema: public; Owner: steve
--

SELECT pg_catalog.setval('item_id_seq', 2, true);


--
-- Data for Name: newsletter; Type: TABLE DATA; Schema: public; Owner: steve
--

COPY newsletter (id, date, name, content) FROM stdin;
\.


--
-- Name: newsletter_id_seq; Type: SEQUENCE SET; Schema: public; Owner: steve
--

SELECT pg_catalog.setval('newsletter_id_seq', 1, false);


--
-- Data for Name: postage; Type: TABLE DATA; Schema: public; Owner: steve
--

COPY postage (shipping_code, region_id, date, cost, price, eta_days, max_days, has_tracking) FROM stdin;
\.


--
-- Data for Name: product; Type: TABLE DATA; Schema: public; Owner: steve
--

COPY product (id, cat_id, name, descr, image, volume_ml, weight_g, shipping_volume_ml, shipping_weight_g, shipping_code, price) FROM stdin;
1	1	Chamelee 50ml	Chamelee 50ml bottle	chamelee-bottle.png	0	0	0	0	0	56.00
2	1	Chamelee 100ml	Chamelee 100ml Bottle	chamelee-bottle.png	0	0	0	0	0	84.00
4	1	Chamelee 22ml		chamelee-bottle.png	0	0	0	0	1	22.00
5	1	Chamelee 6ml	Sample Bottle	chamelee-bottle.png	6	0	0	0	1	18.00
6	3	Umbrella	Chamelee Umbrella	umbrella.jpg	0	0	0	0	1	66.00
7	2	Chamelee Creme		white-soap.jpg	0	0	0	0	1	44.00
8	2	Chamelee Soap		white-choc.jpg	0	0	0	0	1	34.00
9	3	TeeShirt	Chamelee casual wear T-Shirt	white-top.jpg	0	0	0	0	1	34.00
\.


--
-- Name: product_id_seq; Type: SEQUENCE SET; Schema: public; Owner: steve
--

SELECT pg_catalog.setval('product_id_seq', 9, true);


--
-- Data for Name: ref_hit; Type: TABLE DATA; Schema: public; Owner: steve
--

COPY ref_hit (ref_id, date, ip_addr, user_agent) FROM stdin;
\.


--
-- Data for Name: referrer; Type: TABLE DATA; Schema: public; Owner: steve
--

COPY referrer (id, name, url) FROM stdin;
\.


--
-- Name: referrer_id_seq; Type: SEQUENCE SET; Schema: public; Owner: steve
--

SELECT pg_catalog.setval('referrer_id_seq', 1, false);


--
-- Data for Name: region; Type: TABLE DATA; Schema: public; Owner: steve
--

COPY region (id, name, descr) FROM stdin;
1	Australia Metro	Metro areas - capital cities, Australia
2	Australia - Rural	Rural areas in Australia, outside of major cities
3	Asia	Asia - China, Japan, SE Asia
4	Indonesia	Indonesia
6	Middle East	Middle East Countries
7	USA	USA - anywhere
8	Canada	Anywhere in Canada
9	Sth America	All South American Countries
10	Russia	Anywhere in Russia, includes Central Asian states.
11	EU 2	Other EU states
5	EU 1	Europe - preferred states
12	UK	UK
\.


--
-- Name: region_id_seq; Type: SEQUENCE SET; Schema: public; Owner: steve
--

SELECT pg_catalog.setval('region_id_seq', 12, true);


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
-- Data for Name: shipping_code; Type: TABLE DATA; Schema: public; Owner: steve
--

COPY shipping_code (id, name, descr) FROM stdin;
2	Medium	Medium Package - clothing items, etc
3	Large	Large Package - umbrella, etc
4	Heavy	Heavy Package - wholesale carton of perfume, etc
1	Small	Small Package - perfume, etc
\.


--
-- Name: shipping_code_id_seq; Type: SEQUENCE SET; Schema: public; Owner: steve
--

SELECT pg_catalog.setval('shipping_code_id_seq', 6, true);


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
-- Data for Name: up_hit; Type: TABLE DATA; Schema: public; Owner: steve
--

COPY up_hit (id, date, ip_addr, user_agent) FROM stdin;
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
-- Name: admin_users_pkey; Type: CONSTRAINT; Schema: public; Owner: steve; Tablespace: 
--

ALTER TABLE ONLY admin_users
    ADD CONSTRAINT admin_users_pkey PRIMARY KEY (id);


--
-- Name: blog_pkey; Type: CONSTRAINT; Schema: public; Owner: steve; Tablespace: 
--

ALTER TABLE ONLY blog
    ADD CONSTRAINT blog_pkey PRIMARY KEY (id);


--
-- Name: carrier_pkey; Type: CONSTRAINT; Schema: public; Owner: steve; Tablespace: 
--

ALTER TABLE ONLY carrier
    ADD CONSTRAINT carrier_pkey PRIMARY KEY (id);


--
-- Name: category_pkey; Type: CONSTRAINT; Schema: public; Owner: steve; Tablespace: 
--

ALTER TABLE ONLY category
    ADD CONSTRAINT category_pkey PRIMARY KEY (id);


--
-- Name: item_pkey; Type: CONSTRAINT; Schema: public; Owner: steve; Tablespace: 
--

ALTER TABLE ONLY item
    ADD CONSTRAINT item_pkey PRIMARY KEY (id);


--
-- Name: newsletter_pkey; Type: CONSTRAINT; Schema: public; Owner: steve; Tablespace: 
--

ALTER TABLE ONLY newsletter
    ADD CONSTRAINT newsletter_pkey PRIMARY KEY (id);


--
-- Name: product_pkey; Type: CONSTRAINT; Schema: public; Owner: steve; Tablespace: 
--

ALTER TABLE ONLY product
    ADD CONSTRAINT product_pkey PRIMARY KEY (id);


--
-- Name: referrer_pkey; Type: CONSTRAINT; Schema: public; Owner: steve; Tablespace: 
--

ALTER TABLE ONLY referrer
    ADD CONSTRAINT referrer_pkey PRIMARY KEY (id);


--
-- Name: region_pkey; Type: CONSTRAINT; Schema: public; Owner: steve; Tablespace: 
--

ALTER TABLE ONLY region
    ADD CONSTRAINT region_pkey PRIMARY KEY (id);


--
-- Name: sale_order_pkey; Type: CONSTRAINT; Schema: public; Owner: steve; Tablespace: 
--

ALTER TABLE ONLY sale_order
    ADD CONSTRAINT sale_order_pkey PRIMARY KEY (id);


--
-- Name: shipping_code_pkey; Type: CONSTRAINT; Schema: public; Owner: steve; Tablespace: 
--

ALTER TABLE ONLY shipping_code
    ADD CONSTRAINT shipping_code_pkey PRIMARY KEY (id);


--
-- Name: up_hit_pkey; Type: CONSTRAINT; Schema: public; Owner: steve; Tablespace: 
--

ALTER TABLE ONLY up_hit
    ADD CONSTRAINT up_hit_pkey PRIMARY KEY (id);


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

