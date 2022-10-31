--
-- PostgreSQL database dump
--

-- Dumped from database version 14.5 (Ubuntu 14.5-1.pgdg20.04+1)
-- Dumped by pg_dump version 14.5 (Ubuntu 14.5-1.pgdg20.04+1)

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
-- Name: cities; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.cities (
    id bigint NOT NULL,
    name text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.cities OWNER TO postgres;

--
-- Name: cities_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.cities_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.cities_id_seq OWNER TO postgres;

--
-- Name: cities_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.cities_id_seq OWNED BY public.cities.id;


--
-- Name: game_chances; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.game_chances (
    id bigint NOT NULL,
    user_id bigint,
    chance bigint,
    counter bigint,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.game_chances OWNER TO postgres;

--
-- Name: game_chances_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.game_chances_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.game_chances_id_seq OWNER TO postgres;

--
-- Name: game_chances_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.game_chances_id_seq OWNED BY public.game_chances.id;


--
-- Name: house_photos; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.house_photos (
    id bigint NOT NULL,
    house_id bigint,
    photo text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.house_photos OWNER TO postgres;

--
-- Name: house_photos_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.house_photos_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.house_photos_id_seq OWNER TO postgres;

--
-- Name: house_photos_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.house_photos_id_seq OWNED BY public.house_photos.id;


--
-- Name: house_statuses; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.house_statuses (
    id bigint NOT NULL,
    status text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.house_statuses OWNER TO postgres;

--
-- Name: house_statuses_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.house_statuses_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.house_statuses_id_seq OWNER TO postgres;

--
-- Name: house_statuses_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.house_statuses_id_seq OWNED BY public.house_statuses.id;


--
-- Name: houses; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.houses (
    id bigint NOT NULL,
    name text,
    user_id bigint,
    price bigint,
    description text,
    city_id bigint,
    max_guest bigint,
    house_status_id bigint,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.houses OWNER TO postgres;

--
-- Name: houses_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.houses_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.houses_id_seq OWNER TO postgres;

--
-- Name: houses_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.houses_id_seq OWNED BY public.houses.id;


--
-- Name: pickup_statuses; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.pickup_statuses (
    id bigint NOT NULL,
    status text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.pickup_statuses OWNER TO postgres;

--
-- Name: pickup_statuses_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.pickup_statuses_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.pickup_statuses_id_seq OWNER TO postgres;

--
-- Name: pickup_statuses_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.pickup_statuses_id_seq OWNED BY public.pickup_statuses.id;


--
-- Name: pickups; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.pickups (
    id bigint NOT NULL,
    reservation_id bigint,
    pickup_status_id bigint,
    pickup_address text,
    city_id bigint,
    pickup_price bigint,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.pickups OWNER TO postgres;

--
-- Name: pickups_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.pickups_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.pickups_id_seq OWNER TO postgres;

--
-- Name: pickups_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.pickups_id_seq OWNED BY public.pickups.id;


--
-- Name: reservations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.reservations (
    id bigint NOT NULL,
    house_id bigint,
    user_id bigint,
    reservation_code text,
    check_in_date timestamp with time zone,
    check_out_date timestamp with time zone,
    total_price bigint,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.reservations OWNER TO postgres;

--
-- Name: reservations_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.reservations_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.reservations_id_seq OWNER TO postgres;

--
-- Name: reservations_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.reservations_id_seq OWNED BY public.reservations.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    name text,
    email text,
    password text,
    address text,
    city_id bigint,
    role text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: wallets; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.wallets (
    id bigint NOT NULL,
    user_id bigint,
    balance bigint,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.wallets OWNER TO postgres;

--
-- Name: wallets_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.wallets_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.wallets_id_seq OWNER TO postgres;

--
-- Name: wallets_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.wallets_id_seq OWNED BY public.wallets.id;


--
-- Name: cities id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cities ALTER COLUMN id SET DEFAULT nextval('public.cities_id_seq'::regclass);


--
-- Name: game_chances id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.game_chances ALTER COLUMN id SET DEFAULT nextval('public.game_chances_id_seq'::regclass);


--
-- Name: house_photos id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.house_photos ALTER COLUMN id SET DEFAULT nextval('public.house_photos_id_seq'::regclass);


--
-- Name: house_statuses id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.house_statuses ALTER COLUMN id SET DEFAULT nextval('public.house_statuses_id_seq'::regclass);


--
-- Name: houses id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.houses ALTER COLUMN id SET DEFAULT nextval('public.houses_id_seq'::regclass);


--
-- Name: pickup_statuses id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.pickup_statuses ALTER COLUMN id SET DEFAULT nextval('public.pickup_statuses_id_seq'::regclass);


--
-- Name: pickups id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.pickups ALTER COLUMN id SET DEFAULT nextval('public.pickups_id_seq'::regclass);


--
-- Name: reservations id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reservations ALTER COLUMN id SET DEFAULT nextval('public.reservations_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Name: wallets id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wallets ALTER COLUMN id SET DEFAULT nextval('public.wallets_id_seq'::regclass);


--
-- Data for Name: cities; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.cities (id, name, created_at, updated_at, deleted_at) FROM stdin;
1	Bandung	2022-10-30 20:00:01.637251+07	2022-10-30 20:00:01.637251+07	\N
2	Batam	2022-10-30 20:00:01.637251+07	2022-10-30 20:00:01.637251+07	\N
3	Bekasi	2022-10-30 20:00:01.637251+07	2022-10-30 20:00:01.637251+07	\N
4	Bogor	2022-10-30 20:00:01.637251+07	2022-10-30 20:00:01.637251+07	\N
5	Depok	2022-10-30 20:00:01.637251+07	2022-10-30 20:00:01.637251+07	\N
6	Jakarta Barat	2022-10-30 20:00:01.637251+07	2022-10-30 20:00:01.637251+07	\N
7	Jakarta Pusat	2022-10-30 20:00:01.637251+07	2022-10-30 20:00:01.637251+07	\N
8	Jakarta Selatan	2022-10-30 20:00:01.637251+07	2022-10-30 20:00:01.637251+07	\N
9	Jakarta Timur	2022-10-30 20:00:01.637251+07	2022-10-30 20:00:01.637251+07	\N
10	Jakarta Utara	2022-10-30 20:00:01.637251+07	2022-10-30 20:00:01.637251+07	\N
11	Makassar	2022-10-30 20:00:01.637251+07	2022-10-30 20:00:01.637251+07	\N
12	Malang	2022-10-30 20:00:01.637251+07	2022-10-30 20:00:01.637251+07	\N
13	Medan	2022-10-30 20:00:01.637251+07	2022-10-30 20:00:01.637251+07	\N
14	Palembang	2022-10-30 20:00:01.637251+07	2022-10-30 20:00:01.637251+07	\N
15	Pekanbaru	2022-10-30 20:00:01.637251+07	2022-10-30 20:00:01.637251+07	\N
16	Semarang	2022-10-30 20:00:01.637251+07	2022-10-30 20:00:01.637251+07	\N
17	Serang	2022-10-30 20:00:01.637251+07	2022-10-30 20:00:01.637251+07	\N
18	Surabaya	2022-10-30 20:00:01.637251+07	2022-10-30 20:00:01.637251+07	\N
19	Tangerang	2022-10-30 20:00:01.637251+07	2022-10-30 20:00:01.637251+07	\N
20	Tangerang Selatan	2022-10-30 20:00:01.637251+07	2022-10-30 20:00:01.637251+07	\N
\.


--
-- Data for Name: game_chances; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.game_chances (id, user_id, chance, counter, created_at, updated_at, deleted_at) FROM stdin;
1	1	0	0	2022-10-30 20:00:51.655007+07	2022-10-30 20:00:51.655007+07	\N
2	2	0	0	2022-10-30 20:02:37.932673+07	2022-10-30 20:02:37.932673+07	\N
\.


--
-- Data for Name: house_photos; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.house_photos (id, house_id, photo, created_at, updated_at, deleted_at) FROM stdin;
1	1	https://res.cloudinary.com/dbhvedme6/image/upload/v1667135285/oxc7zzsniqpzueuhxehe.jpg	2022-10-30 20:08:08.012442+07	2022-10-30 20:08:08.012442+07	\N
2	1	https://res.cloudinary.com/dbhvedme6/image/upload/v1667135286/itcvfg2fgh8lcprytdes.jpg	2022-10-30 20:08:08.014771+07	2022-10-30 20:08:08.014771+07	\N
3	1	https://res.cloudinary.com/dbhvedme6/image/upload/v1667135286/idsppcr8ovb9g2waiv7r.jpg	2022-10-30 20:08:08.016229+07	2022-10-30 20:08:08.016229+07	\N
4	1	https://res.cloudinary.com/dbhvedme6/image/upload/v1667135286/wgqa2pj6gjdxhfnigurw.jpg	2022-10-30 20:08:08.017673+07	2022-10-30 20:08:08.017673+07	\N
5	1	https://res.cloudinary.com/dbhvedme6/image/upload/v1667135287/all9lcobfn3u12k5123w.jpg	2022-10-30 20:08:08.019036+07	2022-10-30 20:08:08.019036+07	\N
6	1	https://res.cloudinary.com/dbhvedme6/image/upload/v1667135287/bivjmoyfmqa96ucw41ch.jpg	2022-10-30 20:08:08.020346+07	2022-10-30 20:08:08.020346+07	\N
7	2	https://res.cloudinary.com/dbhvedme6/image/upload/v1667135781/dwx5q2xlwbmcjeifjjd7.webp	2022-10-30 20:16:21.99669+07	2022-10-30 20:16:21.99669+07	\N
8	2	https://res.cloudinary.com/dbhvedme6/image/upload/v1667135781/vljrijfmeccueot6hzqa.webp	2022-10-30 20:16:21.998389+07	2022-10-30 20:16:21.998389+07	\N
9	2	https://res.cloudinary.com/dbhvedme6/image/upload/v1667135781/ijobqhfgiuhd1dbh9gtb.webp	2022-10-30 20:16:21.999718+07	2022-10-30 20:16:21.999718+07	\N
10	3	https://res.cloudinary.com/dbhvedme6/image/upload/v1667136037/gbofnflzy2xds5f8zepa.webp	2022-10-30 20:20:38.641573+07	2022-10-30 20:20:38.641573+07	\N
11	3	https://res.cloudinary.com/dbhvedme6/image/upload/v1667136038/mdesyho7gcakwsep4ml0.webp	2022-10-30 20:20:38.642812+07	2022-10-30 20:20:38.642812+07	\N
12	3	https://res.cloudinary.com/dbhvedme6/image/upload/v1667136038/ygg4jzj3cmrhvzb0jzz1.webp	2022-10-30 20:20:38.643954+07	2022-10-30 20:20:38.643954+07	\N
13	4	https://res.cloudinary.com/dbhvedme6/image/upload/v1667136317/wdnkhderqfa9rraswqkv.webp	2022-10-30 20:25:18.563622+07	2022-10-30 20:25:18.563622+07	\N
14	4	https://res.cloudinary.com/dbhvedme6/image/upload/v1667136317/ojoffdngv8aveirujrzp.webp	2022-10-30 20:25:18.565011+07	2022-10-30 20:25:18.565011+07	\N
15	4	https://res.cloudinary.com/dbhvedme6/image/upload/v1667136317/r1ekmdr66gmdi6livekz.webp	2022-10-30 20:25:18.566266+07	2022-10-30 20:25:18.566266+07	\N
16	4	https://res.cloudinary.com/dbhvedme6/image/upload/v1667136318/fytq2cmamgzzgez1r4l3.webp	2022-10-30 20:25:18.56742+07	2022-10-30 20:25:18.56742+07	\N
17	4	https://res.cloudinary.com/dbhvedme6/image/upload/v1667136318/ytrerul21xhafrgetyba.webp	2022-10-30 20:25:18.568466+07	2022-10-30 20:25:18.568466+07	\N
18	5	https://res.cloudinary.com/dbhvedme6/image/upload/v1667137215/u9a2mueg0x0itkaxbogn.webp	2022-10-30 20:40:17.787407+07	2022-10-30 20:40:17.787407+07	\N
19	5	https://res.cloudinary.com/dbhvedme6/image/upload/v1667137216/fnyglo4i1p2dcax3ywut.webp	2022-10-30 20:40:17.788393+07	2022-10-30 20:40:17.788393+07	\N
20	5	https://res.cloudinary.com/dbhvedme6/image/upload/v1667137216/uqecwtn5uv9qhywlcwba.webp	2022-10-30 20:40:17.789371+07	2022-10-30 20:40:17.789371+07	\N
21	5	https://res.cloudinary.com/dbhvedme6/image/upload/v1667137216/kz9dlzhm1tyb2zpsyequ.webp	2022-10-30 20:40:17.790456+07	2022-10-30 20:40:17.790456+07	\N
22	6	https://res.cloudinary.com/dbhvedme6/image/upload/v1667137479/xbzfhaby8k1fhj9v8psn.webp	2022-10-30 20:44:40.389387+07	2022-10-30 20:44:40.389387+07	\N
23	6	https://res.cloudinary.com/dbhvedme6/image/upload/v1667137479/q6cuilooc2ydszar0r3t.webp	2022-10-30 20:44:40.390554+07	2022-10-30 20:44:40.390554+07	\N
24	6	https://res.cloudinary.com/dbhvedme6/image/upload/v1667137479/v64b64gwviyzjranb8zt.webp	2022-10-30 20:44:40.391639+07	2022-10-30 20:44:40.391639+07	\N
25	6	https://res.cloudinary.com/dbhvedme6/image/upload/v1667137479/u1ofryufkfgfzzb4edep.webp	2022-10-30 20:44:40.392659+07	2022-10-30 20:44:40.392659+07	\N
26	7	https://res.cloudinary.com/dbhvedme6/image/upload/v1667137754/ezffguqxqzgw7416oooc.webp	2022-10-30 20:49:16.075882+07	2022-10-30 20:49:16.075882+07	\N
27	7	https://res.cloudinary.com/dbhvedme6/image/upload/v1667137755/qaa17cjgf75hjrztoonn.webp	2022-10-30 20:49:16.077113+07	2022-10-30 20:49:16.077113+07	\N
28	8	https://res.cloudinary.com/dbhvedme6/image/upload/v1667137934/xizwdt0up9bkeizpe8jv.webp	2022-10-30 20:52:17.245398+07	2022-10-30 20:52:17.245398+07	\N
29	8	https://res.cloudinary.com/dbhvedme6/image/upload/v1667137935/ro8zeuda2jw7ha9khrey.webp	2022-10-30 20:52:17.246721+07	2022-10-30 20:52:17.246721+07	\N
30	8	https://res.cloudinary.com/dbhvedme6/image/upload/v1667137935/df1bqcowzkh8owpi2vzn.webp	2022-10-30 20:52:17.247934+07	2022-10-30 20:52:17.247934+07	\N
31	8	https://res.cloudinary.com/dbhvedme6/image/upload/v1667137936/zwrrje7rt0kqwrmrlfhw.webp	2022-10-30 20:52:17.249101+07	2022-10-30 20:52:17.249101+07	\N
32	8	https://res.cloudinary.com/dbhvedme6/image/upload/v1667137935/h0qc7nbrarsnj9xpkcbr.webp	2022-10-30 20:52:17.250488+07	2022-10-30 20:52:17.250488+07	\N
33	9	https://res.cloudinary.com/dbhvedme6/image/upload/v1667138134/hrmjsqcmtjzeohdfascu.webp	2022-10-30 20:55:36.225013+07	2022-10-30 20:55:36.225013+07	\N
34	9	https://res.cloudinary.com/dbhvedme6/image/upload/v1667138134/xdom3abn5idoy3kmet0d.webp	2022-10-30 20:55:36.226266+07	2022-10-30 20:55:36.226266+07	\N
35	9	https://res.cloudinary.com/dbhvedme6/image/upload/v1667138135/gscuh9rrdatpvmq0xsvv.webp	2022-10-30 20:55:36.227384+07	2022-10-30 20:55:36.227384+07	\N
36	10	https://res.cloudinary.com/dbhvedme6/image/upload/v1667138233/p9qdqekxusqb5gkwpexa.webp	2022-10-30 20:57:14.640145+07	2022-10-30 20:57:14.640145+07	\N
37	10	https://res.cloudinary.com/dbhvedme6/image/upload/v1667138233/yp9aiezfbk9nwun2mngp.webp	2022-10-30 20:57:14.641361+07	2022-10-30 20:57:14.641361+07	\N
38	10	https://res.cloudinary.com/dbhvedme6/image/upload/v1667138233/chbcfy5aupydleeuge6j.webp	2022-10-30 20:57:14.642606+07	2022-10-30 20:57:14.642606+07	\N
39	10	https://res.cloudinary.com/dbhvedme6/image/upload/v1667138234/bjqrtydnyojtzxphk9md.webp	2022-10-30 20:57:14.643927+07	2022-10-30 20:57:14.643927+07	\N
40	11	https://res.cloudinary.com/dbhvedme6/image/upload/v1667138352/c1rpuqvohqqpzn13gp5e.webp	2022-10-30 20:59:14.085063+07	2022-10-30 20:59:14.085063+07	\N
41	11	https://res.cloudinary.com/dbhvedme6/image/upload/v1667138352/bxfc2d0eoqmaiib6whx9.webp	2022-10-30 20:59:14.086908+07	2022-10-30 20:59:14.086908+07	\N
42	11	https://res.cloudinary.com/dbhvedme6/image/upload/v1667138353/od0rozcvey3hs8ofce0l.webp	2022-10-30 20:59:14.088282+07	2022-10-30 20:59:14.088282+07	\N
43	12	https://res.cloudinary.com/dbhvedme6/image/upload/v1667138615/qkzybambqwfpcmaxxh3g.webp	2022-10-30 21:03:37.641753+07	2022-10-30 21:03:37.641753+07	\N
44	12	https://res.cloudinary.com/dbhvedme6/image/upload/v1667138616/ikkv82nnqqjxjbyp5k34.webp	2022-10-30 21:03:37.643454+07	2022-10-30 21:03:37.643454+07	\N
45	12	https://res.cloudinary.com/dbhvedme6/image/upload/v1667138616/nzubnkkiochiysf88t2q.webp	2022-10-30 21:03:37.644784+07	2022-10-30 21:03:37.644784+07	\N
46	12	https://res.cloudinary.com/dbhvedme6/image/upload/v1667138617/q7yg3kkjctt3sbwoandd.webp	2022-10-30 21:03:37.646108+07	2022-10-30 21:03:37.646108+07	\N
47	13	https://res.cloudinary.com/dbhvedme6/image/upload/v1667138921/ravmc7jahtgivbuo526l.webp	2022-10-30 21:08:42.90615+07	2022-10-30 21:08:42.90615+07	\N
48	13	https://res.cloudinary.com/dbhvedme6/image/upload/v1667138921/s4om48m6ym56q4vpom3t.webp	2022-10-30 21:08:42.907447+07	2022-10-30 21:08:42.907447+07	\N
49	13	https://res.cloudinary.com/dbhvedme6/image/upload/v1667138922/us8dyghxvkpmnz9mchtf.webp	2022-10-30 21:08:42.908558+07	2022-10-30 21:08:42.908558+07	\N
50	14	https://res.cloudinary.com/dbhvedme6/image/upload/v1667139221/xbu0rkfgdiwqzimiwpxb.webp	2022-10-30 21:13:43.400481+07	2022-10-30 21:13:43.400481+07	\N
51	14	https://res.cloudinary.com/dbhvedme6/image/upload/v1667139221/w2vk1rnbgdajnonwpvor.webp	2022-10-30 21:13:43.402131+07	2022-10-30 21:13:43.402131+07	\N
52	14	https://res.cloudinary.com/dbhvedme6/image/upload/v1667139222/xajbkewomkthpbghe8vm.webp	2022-10-30 21:13:43.403545+07	2022-10-30 21:13:43.403545+07	\N
53	14	https://res.cloudinary.com/dbhvedme6/image/upload/v1667139222/jwwch8ssafc2zurlrzkq.webp	2022-10-30 21:13:43.404776+07	2022-10-30 21:13:43.404776+07	\N
54	15	https://res.cloudinary.com/dbhvedme6/image/upload/v1667139429/filjjbvsnaeo18sri4qk.webp	2022-10-30 21:17:11.350379+07	2022-10-30 21:17:11.350379+07	\N
55	15	https://res.cloudinary.com/dbhvedme6/image/upload/v1667139429/owkxe3kx0k5dlt8zu2tb.webp	2022-10-30 21:17:11.351892+07	2022-10-30 21:17:11.351892+07	\N
56	15	https://res.cloudinary.com/dbhvedme6/image/upload/v1667139429/smqqws9b3wu4qqe4mlnn.webp	2022-10-30 21:17:11.353092+07	2022-10-30 21:17:11.353092+07	\N
57	15	https://res.cloudinary.com/dbhvedme6/image/upload/v1667139430/jlhfkapzkl7nxjqqyk0i.webp	2022-10-30 21:17:11.355277+07	2022-10-30 21:17:11.355277+07	\N
58	16	https://res.cloudinary.com/dbhvedme6/image/upload/v1667139563/a4zfkn6kct0vxjj2q3iz.webp	2022-10-30 21:19:24.270543+07	2022-10-30 21:19:24.270543+07	\N
59	16	https://res.cloudinary.com/dbhvedme6/image/upload/v1667139563/i9xhibeuegdazjwic3ep.webp	2022-10-30 21:19:24.271834+07	2022-10-30 21:19:24.271834+07	\N
60	16	https://res.cloudinary.com/dbhvedme6/image/upload/v1667139563/cnmf1b9bejlhtoej9t1b.webp	2022-10-30 21:19:24.272835+07	2022-10-30 21:19:24.272835+07	\N
62	17	https://res.cloudinary.com/dbhvedme6/image/upload/v1667139564/crya8s3lzom3ejsdd4de_lsvzjs.webp	2022-10-30 21:19:24.965691+07	2022-10-30 21:19:24.965691+07	\N
63	17	https://res.cloudinary.com/dbhvedme6/image/upload/v1667139564/tbe5vdrgbplxtkivonj5_o8k1jk.webp	2022-10-30 21:19:24.966654+07	2022-10-30 21:19:24.966654+07	\N
61	17	https://res.cloudinary.com/dbhvedme6/image/upload/v1667139564/ft97vms1ozxzlqtme9jk_cjw0kf.webp	2022-10-30 21:19:24.964726+07	2022-10-30 21:19:24.964726+07	\N
64	18	https://res.cloudinary.com/dbhvedme6/image/upload/v1667140597/u9zm03fymjylwkps6u5v.webp	2022-10-30 21:36:38.363808+07	2022-10-30 21:36:38.363808+07	\N
65	18	https://res.cloudinary.com/dbhvedme6/image/upload/v1667140597/dx5dv0ms9fiothzlxguj.webp	2022-10-30 21:36:38.365336+07	2022-10-30 21:36:38.365336+07	\N
66	18	https://res.cloudinary.com/dbhvedme6/image/upload/v1667140597/kqq9ygbwgdfnpziqja2b.webp	2022-10-30 21:36:38.36673+07	2022-10-30 21:36:38.36673+07	\N
67	18	https://res.cloudinary.com/dbhvedme6/image/upload/v1667140597/puqrngy302hggztbdkyc.webp	2022-10-30 21:36:38.368123+07	2022-10-30 21:36:38.368123+07	\N
68	19	https://res.cloudinary.com/dbhvedme6/image/upload/v1667140758/s2aphncijnl6lid3q683.webp	2022-10-30 21:39:20.490368+07	2022-10-30 21:39:20.490368+07	\N
69	19	https://res.cloudinary.com/dbhvedme6/image/upload/v1667140759/oey0za0uqcknc67nnfif.webp	2022-10-30 21:39:20.491851+07	2022-10-30 21:39:20.491851+07	\N
70	19	https://res.cloudinary.com/dbhvedme6/image/upload/v1667140758/nnhqgwiwxmgzs7en2o17.webp	2022-10-30 21:39:20.493141+07	2022-10-30 21:39:20.493141+07	\N
71	19	https://res.cloudinary.com/dbhvedme6/image/upload/v1667140759/agwhwojs7rlvybcaq7tb.webp	2022-10-30 21:39:20.494439+07	2022-10-30 21:39:20.494439+07	\N
72	19	https://res.cloudinary.com/dbhvedme6/image/upload/v1667140759/x4zkps1k4ijcmfkqcrcv.webp	2022-10-30 21:39:20.49582+07	2022-10-30 21:39:20.49582+07	\N
73	20	https://res.cloudinary.com/dbhvedme6/image/upload/v1667141051/iijgixwizb1kgbychkxo.webp	2022-10-30 21:44:11.615686+07	2022-10-30 21:44:11.615686+07	\N
74	20	https://res.cloudinary.com/dbhvedme6/image/upload/v1667141051/zyr9l9jh76dopm1gmcvq.webp	2022-10-30 21:44:11.617044+07	2022-10-30 21:44:11.617044+07	\N
75	20	https://res.cloudinary.com/dbhvedme6/image/upload/v1667141051/p1wipt3lpnh4r9xiqbcp.webp	2022-10-30 21:44:11.618352+07	2022-10-30 21:44:11.618352+07	\N
76	20	https://res.cloudinary.com/dbhvedme6/image/upload/v1667141051/wx9zeeyop5h5u2qcqnfn.webp	2022-10-30 21:44:11.619668+07	2022-10-30 21:44:11.619668+07	\N
77	21	https://res.cloudinary.com/dbhvedme6/image/upload/v1667141149/iilrkvitcylgrbals032.webp	2022-10-30 21:45:51.206473+07	2022-10-30 21:45:51.206473+07	\N
78	21	https://res.cloudinary.com/dbhvedme6/image/upload/v1667141149/pyujsvf9tiyisr3p6fcw.webp	2022-10-30 21:45:51.207938+07	2022-10-30 21:45:51.207938+07	\N
79	21	https://res.cloudinary.com/dbhvedme6/image/upload/v1667141150/jds1qqlyt8hbqx4uwlwz.webp	2022-10-30 21:45:51.209315+07	2022-10-30 21:45:51.209315+07	\N
80	21	https://res.cloudinary.com/dbhvedme6/image/upload/v1667141150/icxzdvzzcd0sbfszkctx.webp	2022-10-30 21:45:51.210583+07	2022-10-30 21:45:51.210583+07	\N
81	22	https://res.cloudinary.com/dbhvedme6/image/upload/v1667141266/tlcfyhkwpzfyp5olxkdj.webp	2022-10-30 21:47:48.981621+07	2022-10-30 21:47:48.981621+07	\N
82	22	https://res.cloudinary.com/dbhvedme6/image/upload/v1667141266/p3ars5klyvgghjja7neg.webp	2022-10-30 21:47:48.983004+07	2022-10-30 21:47:48.983004+07	\N
83	22	https://res.cloudinary.com/dbhvedme6/image/upload/v1667141266/s9m49fjc4i7kad5kejis.webp	2022-10-30 21:47:48.984428+07	2022-10-30 21:47:48.984428+07	\N
84	22	https://res.cloudinary.com/dbhvedme6/image/upload/v1667141267/jmndynejs8emchxvgufc.webp	2022-10-30 21:47:48.985779+07	2022-10-30 21:47:48.985779+07	\N
85	22	https://res.cloudinary.com/dbhvedme6/image/upload/v1667141267/qrds9df6yttxsvygphi3.webp	2022-10-30 21:47:48.987174+07	2022-10-30 21:47:48.987174+07	\N
86	23	https://res.cloudinary.com/dbhvedme6/image/upload/v1667141498/m6npvjucugtcykhz5p9c.webp	2022-10-30 21:51:39.521303+07	2022-10-30 21:51:39.521303+07	\N
87	23	https://res.cloudinary.com/dbhvedme6/image/upload/v1667141498/vnzmoqfykwjaduepe44x.webp	2022-10-30 21:51:39.522689+07	2022-10-30 21:51:39.522689+07	\N
88	23	https://res.cloudinary.com/dbhvedme6/image/upload/v1667141498/zswaetw59o2bbwniu7pl.webp	2022-10-30 21:51:39.523895+07	2022-10-30 21:51:39.523895+07	\N
89	23	https://res.cloudinary.com/dbhvedme6/image/upload/v1667141498/lkb2dvdsntgaivxwvcqm.webp	2022-10-30 21:51:39.525185+07	2022-10-30 21:51:39.525185+07	\N
90	23	https://res.cloudinary.com/dbhvedme6/image/upload/v1667141499/qimcwzyf2lbzargskatz.webp	2022-10-30 21:51:39.526592+07	2022-10-30 21:51:39.526592+07	\N
91	24	https://res.cloudinary.com/dbhvedme6/image/upload/v1667141667/nuehaaueh16fxsf3bzeg.webp	2022-10-30 21:54:28.612182+07	2022-10-30 21:54:28.612182+07	\N
92	24	https://res.cloudinary.com/dbhvedme6/image/upload/v1667141667/v05k1be9bcssg3ghnb4g.webp	2022-10-30 21:54:28.613666+07	2022-10-30 21:54:28.613666+07	\N
93	24	https://res.cloudinary.com/dbhvedme6/image/upload/v1667141667/x7sgmhrhwyy3wss6xszi.webp	2022-10-30 21:54:28.615007+07	2022-10-30 21:54:28.615007+07	\N
94	24	https://res.cloudinary.com/dbhvedme6/image/upload/v1667141667/x1pypl8qplfi7zjympmq.webp	2022-10-30 21:54:28.616401+07	2022-10-30 21:54:28.616401+07	\N
95	24	https://res.cloudinary.com/dbhvedme6/image/upload/v1667141668/lmbphuogwbj4pj1mh7fp.webp	2022-10-30 21:54:28.617757+07	2022-10-30 21:54:28.617757+07	\N
\.


--
-- Data for Name: house_statuses; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.house_statuses (id, status, created_at, updated_at, deleted_at) FROM stdin;
1	Active	2022-10-30 20:00:13.570586+07	2022-10-30 20:00:13.570586+07	\N
2	Inactive	2022-10-30 20:00:13.570586+07	2022-10-30 20:00:13.570586+07	\N
3	Deactivated by Admin	2022-10-30 20:00:13.570586+07	2022-10-30 20:00:13.570586+07	\N
\.


--
-- Data for Name: houses; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.houses (id, name, user_id, price, description, city_id, max_guest, house_status_id, created_at, updated_at, deleted_at) FROM stdin;
1	Orchid House	2	150000	Clean, comfortable and safe, close to restaurants which are about 300 meters from the residence.	6	5	1	2022-10-30 20:08:08.0083+07	2022-10-30 20:08:08.0083+07	\N
2	Centro House	2	200000	A furnished co-living space in the city centre. The listing has super fast WiFi, work desk, laundry facilities (dryer & washer) and an all-inclusive kitchen.	6	2	1	2022-10-30 20:16:21.993398+07	2022-10-30 20:16:21.993398+07	\N
3	One Apartment	2	120000	Forget your worries in this spacious and serene space	6	2	1	2022-10-30 20:20:38.638756+07	2022-10-30 20:20:38.638756+07	\N
4	One Condo	2	150000	A cozy place to stay where guests can relax	6	2	1	2022-10-30 20:25:18.560566+07	2022-10-30 20:25:18.560566+07	\N
5	Blu Apartment	2	250000	Amazing city view and close to shopping district	6	2	1	2022-10-30 20:40:17.785042+07	2022-10-30 20:40:17.785042+07	\N
6	New Apartment	2	250000	A fully furnished apartment with two bedrooms.	1	4	1	2022-10-30 20:44:40.388122+07	2022-10-30 20:44:40.388122+07	\N
7	Life Apartment	2	110000	A fully furnished apartment with one bedroom.	1	2	1	2022-10-30 20:49:16.073108+07	2022-10-30 20:49:16.073108+07	\N
8	Cozy Apartment	2	200000	This fully equipped furnished apartment is located above a shopping mall and it has direct access to the mall.	1	4	1	2022-10-30 20:52:17.242678+07	2022-10-30 20:52:34.825153+07	\N
9	Wave Apartment	2	150000	Fully furnished apartment	1	2	1	2022-10-30 20:55:36.222413+07	2022-10-30 20:55:36.222413+07	\N
10	Prime Apartment	2	250000	Provides a 24/7 security system, parking lot, swimming pool, fitness centers, round the clock help desk, 24-hour mini mart, laundromat, restaurants, pharmacy, and much more!	1	4	1	2022-10-30 20:57:14.637458+07	2022-10-30 20:57:14.637458+07	\N
11	Tree Apartment	2	170000	The apartment provides spacious & comfortable space in the city. It has wifi, 2 cable TVs (40" and 32"), DVD player as well as ready-to-use appliances & utensils to ensure comfort. It's perfect for business traveler or those traveling with family.	1	2	1	2022-10-30 20:59:14.082184+07	2022-10-30 20:59:14.082184+07	\N
12	Lumi House	2	300000	This house consists of one bedroom, one bathroom, kitchen space and living room.	1	2	1	2022-10-30 21:03:37.638719+07	2022-10-30 21:03:37.638719+07	\N
13	Day Apartment	2	250000	Clean, comfortable, suitable for family and friends	1	3	1	2022-10-30 21:08:42.903498+07	2022-10-30 21:08:42.903498+07	\N
14	Night Apartment	2	300000	Comfortable for family and friends	1	4	1	2022-10-30 21:13:43.397512+07	2022-10-30 21:13:43.397512+07	\N
15	Work Apartment	2	150000	A fully furnished apartment consist of 1 bedroom, a living room, kitchen and bathroom	1	2	1	2022-10-30 21:17:11.347288+07	2022-10-30 21:17:11.347288+07	\N
16	De Apartment	2	250000	Near everywhere	1	4	1	2022-10-30 21:19:24.268153+07	2022-10-30 21:19:24.268153+07	\N
17	Soft Apartment	2	250000	No pets allowed. Wifi & Netflix are available for free.	1	4	1	2022-10-30 21:19:24.962485+07	2022-10-30 21:21:28.364608+07	\N
18	Modern House	2	300000	Clean and comfortable house	6	4	1	2022-10-30 21:36:38.360594+07	2022-10-30 21:36:38.360594+07	\N
19	Luxury House	2	500000	Have fun with the family and friends at this stylish place.	6	6	1	2022-10-30 21:39:20.487128+07	2022-10-30 21:39:20.487128+07	\N
20	Dream House	2	250000	Sleep well in a comfortable bed with hotel grade soft linens, fluffy towels. Free fast WiFi, cable TV with 170 channels.	6	4	1	2022-10-30 21:44:11.614073+07	2022-10-30 21:44:11.614073+07	\N
21	Comfy House	2	200000	Comfortable and clean house	6	2	1	2022-10-30 21:45:51.203566+07	2022-10-30 21:45:51.203566+07	\N
22	De House	2	500000	Comfortable house with modern luxury interior design	6	4	1	2022-10-30 21:47:48.978551+07	2022-10-30 21:47:48.978551+07	\N
23	Skylight House	2	350000	From our skylight, rooftop and attic bed you and your family can enjoy star gazing at nights, sunrise, scenic mountain view during the day and cool weather in the mornings.	1	6	1	2022-10-30 21:51:39.519349+07	2022-10-30 21:51:39.519349+07	\N
24	Homy Apartment	2	350000	A cozy and spacious apartment	1	3	1	2022-10-30 21:54:28.609141+07	2022-10-30 21:54:28.609141+07	\N
\.


--
-- Data for Name: pickup_statuses; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.pickup_statuses (id, status, created_at, updated_at, deleted_at) FROM stdin;
1	Pending Admin	2022-10-30 20:00:08.635743+07	2022-10-30 20:00:08.635743+07	\N
2	Awaiting Check-in Date	2022-10-30 20:00:08.635743+07	2022-10-30 20:00:08.635743+07	\N
3	On the Way to Pickup	2022-10-30 20:00:08.635743+07	2022-10-30 20:00:08.635743+07	\N
4	On the Way to Reservation	2022-10-30 20:00:08.635743+07	2022-10-30 20:00:08.635743+07	\N
5	Completed	2022-10-30 20:00:08.635743+07	2022-10-30 20:00:08.635743+07	\N
\.


--
-- Data for Name: pickups; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.pickups (id, reservation_id, pickup_status_id, pickup_address, city_id, pickup_price, created_at, updated_at, deleted_at) FROM stdin;
\.


--
-- Data for Name: reservations; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.reservations (id, house_id, user_id, reservation_code, check_in_date, check_out_date, total_price, created_at, updated_at, deleted_at) FROM stdin;
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, name, email, password, address, city_id, role, created_at, updated_at, deleted_at) FROM stdin;
1	Admin	admin@email.com	$2a$10$iRVwzHeCUNA2QitVMZq7meP4mwAegSDzyYFM8fu37bOZeOqlAhiBS	Admin Street	1	admin	2022-10-30 20:00:51.648981+07	2022-10-30 20:00:51.648981+07	\N
2	Acong Suherman	acong.suherman@email.com	$2a$10$I/x3H3wCbHhbBKyqWbhqDer0c3LgCLGB/Mt9Bhv1zwmYtVUsIuqAq	Merpati Street	6	host	2022-10-30 20:02:37.930032+07	2022-10-30 20:02:46.529543+07	\N
\.


--
-- Data for Name: wallets; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.wallets (id, user_id, balance, created_at, updated_at, deleted_at) FROM stdin;
1	1	0	2022-10-30 20:00:51.653084+07	2022-10-30 20:00:51.653084+07	\N
2	2	0	2022-10-30 20:02:37.931532+07	2022-10-30 20:02:37.931532+07	\N
\.


--
-- Name: cities_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.cities_id_seq', 20, true);


--
-- Name: game_chances_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.game_chances_id_seq', 2, true);


--
-- Name: house_photos_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.house_photos_id_seq', 95, true);


--
-- Name: house_statuses_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.house_statuses_id_seq', 3, true);


--
-- Name: houses_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.houses_id_seq', 24, true);


--
-- Name: pickup_statuses_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.pickup_statuses_id_seq', 5, true);


--
-- Name: pickups_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.pickups_id_seq', 1, false);


--
-- Name: reservations_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.reservations_id_seq', 1, false);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 2, true);


--
-- Name: wallets_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.wallets_id_seq', 2, true);


--
-- Name: cities cities_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cities
    ADD CONSTRAINT cities_pkey PRIMARY KEY (id);


--
-- Name: game_chances game_chances_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.game_chances
    ADD CONSTRAINT game_chances_pkey PRIMARY KEY (id);


--
-- Name: house_photos house_photos_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.house_photos
    ADD CONSTRAINT house_photos_pkey PRIMARY KEY (id);


--
-- Name: house_statuses house_statuses_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.house_statuses
    ADD CONSTRAINT house_statuses_pkey PRIMARY KEY (id);


--
-- Name: houses houses_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.houses
    ADD CONSTRAINT houses_pkey PRIMARY KEY (id);


--
-- Name: pickup_statuses pickup_statuses_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.pickup_statuses
    ADD CONSTRAINT pickup_statuses_pkey PRIMARY KEY (id);


--
-- Name: pickups pickups_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.pickups
    ADD CONSTRAINT pickups_pkey PRIMARY KEY (id);


--
-- Name: reservations reservations_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reservations
    ADD CONSTRAINT reservations_pkey PRIMARY KEY (id);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: wallets wallets_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wallets
    ADD CONSTRAINT wallets_pkey PRIMARY KEY (id);


--
-- Name: idx_cities_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_cities_deleted_at ON public.cities USING btree (deleted_at);


--
-- Name: idx_game_chances_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_game_chances_deleted_at ON public.game_chances USING btree (deleted_at);


--
-- Name: idx_house_photos_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_house_photos_deleted_at ON public.house_photos USING btree (deleted_at);


--
-- Name: idx_house_statuses_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_house_statuses_deleted_at ON public.house_statuses USING btree (deleted_at);


--
-- Name: idx_houses_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_houses_deleted_at ON public.houses USING btree (deleted_at);


--
-- Name: idx_pickup_statuses_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_pickup_statuses_deleted_at ON public.pickup_statuses USING btree (deleted_at);


--
-- Name: idx_pickups_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_pickups_deleted_at ON public.pickups USING btree (deleted_at);


--
-- Name: idx_reservations_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_reservations_deleted_at ON public.reservations USING btree (deleted_at);


--
-- Name: idx_users_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_users_deleted_at ON public.users USING btree (deleted_at);


--
-- Name: idx_wallets_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wallets_deleted_at ON public.wallets USING btree (deleted_at);


--
-- Name: game_chances fk_game_chances_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.game_chances
    ADD CONSTRAINT fk_game_chances_user FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: house_photos fk_house_photos_house; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.house_photos
    ADD CONSTRAINT fk_house_photos_house FOREIGN KEY (house_id) REFERENCES public.houses(id);


--
-- Name: houses fk_houses_city; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.houses
    ADD CONSTRAINT fk_houses_city FOREIGN KEY (city_id) REFERENCES public.cities(id);


--
-- Name: houses fk_houses_house_status; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.houses
    ADD CONSTRAINT fk_houses_house_status FOREIGN KEY (house_status_id) REFERENCES public.house_statuses(id);


--
-- Name: houses fk_houses_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.houses
    ADD CONSTRAINT fk_houses_user FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: pickups fk_pickups_city; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.pickups
    ADD CONSTRAINT fk_pickups_city FOREIGN KEY (city_id) REFERENCES public.cities(id);


--
-- Name: pickups fk_pickups_pickup_status; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.pickups
    ADD CONSTRAINT fk_pickups_pickup_status FOREIGN KEY (pickup_status_id) REFERENCES public.pickup_statuses(id);


--
-- Name: pickups fk_pickups_reservation; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.pickups
    ADD CONSTRAINT fk_pickups_reservation FOREIGN KEY (reservation_id) REFERENCES public.reservations(id);


--
-- Name: reservations fk_reservations_house; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reservations
    ADD CONSTRAINT fk_reservations_house FOREIGN KEY (house_id) REFERENCES public.houses(id);


--
-- Name: reservations fk_reservations_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reservations
    ADD CONSTRAINT fk_reservations_user FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: users fk_users_city; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT fk_users_city FOREIGN KEY (city_id) REFERENCES public.cities(id);


--
-- Name: wallets fk_wallets_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wallets
    ADD CONSTRAINT fk_wallets_user FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- PostgreSQL database dump complete
--

