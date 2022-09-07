--
-- PostgreSQL database dump
--

-- Dumped from database version 10.1
-- Dumped by pg_dump version 10.1

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

SET search_path = public, pg_catalog;

--
-- Data for Name: auctions; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO auctions VALUES (1, 1, '2019-08-02', '2019-12-02', 1, 1, '1', 1, 1, 1);
INSERT INTO auctions VALUES (3, 2, '2019-06-05', '2019-08-05', 2, 100000, 'note', 0, 1, 2);
INSERT INTO auctions VALUES (4, 2, '2019-06-05', '2019-08-05', 2, 100000, 'note', 0, 2, 0);
INSERT INTO auctions VALUES (2, 2, '2019-06-05', '2019-08-05', 2, 4000000, 'note', 2, 1, 2);
INSERT INTO auctions VALUES (5, 2, '2019-06-05', '2019-08-05', 2, 100000, 'note', 0, 1, 0);


--
-- Data for Name: biddings; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO biddings VALUES (2, 2, 2, 2);
INSERT INTO biddings VALUES (1, 1, 1, 3);


--
-- Data for Name: collateral; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO collateral VALUES (2, 2, 162, 'ahmad', 'senayan', 1400000, 88888, 1, 'dx.docx');
INSERT INTO collateral VALUES (3, 4, 163, 'anji', 'jakarta', 18000000, 30000000, 2, 'result.docx');
INSERT INTO collateral VALUES (1, 4, 4, 'anji', 'jakarta', 18000000, 30000000, 2, 'result.docx');


--
-- Data for Name: logging; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: m_option; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO m_option VALUES (1, 1, 161, 'Gedung');
INSERT INTO m_option VALUES (2, 1, 162, 'Gudang');
INSERT INTO m_option VALUES (3, 1, 163, 'Ruko/Rukan/Kios ');
INSERT INTO m_option VALUES (4, 1, 175, 'Properti Komersil Lainnya');
INSERT INTO m_option VALUES (5, 1, 176, 'Rumah Tinggal');
INSERT INTO m_option VALUES (6, 1, 177, 'Apartemen / Rusun');
INSERT INTO m_option VALUES (7, 1, 187, 'Tanah ');
INSERT INTO m_option VALUES (8, 1, 189, 'Kendaraan Bermotor');
INSERT INTO m_option VALUES (9, 2, 1, 'Self Selling');
INSERT INTO m_option VALUES (10, 2, 2, 'Balai Lelang');
INSERT INTO m_option VALUES (11, 3, 1, 'Opened');
INSERT INTO m_option VALUES (12, 3, 2, 'On Process');
INSERT INTO m_option VALUES (13, 3, 3, 'Sold Out');
INSERT INTO m_option VALUES (14, 3, 4, 'Expired');
INSERT INTO m_option VALUES (15, 4, 1, 'Interested');
INSERT INTO m_option VALUES (16, 4, 2, 'Joining');
INSERT INTO m_option VALUES (17, 4, 3, 'Accepted');
INSERT INTO m_option VALUES (18, 4, 4, 'Declined');
INSERT INTO m_option VALUES (19, 5, 1, 'Inactived');
INSERT INTO m_option VALUES (20, 5, 2, 'Actived');
INSERT INTO m_option VALUES (21, 5, 3, 'Deactived');
INSERT INTO m_option VALUES (22, 3, 5, 'AcceptedBidder(updated)');


--
-- Data for Name: partneraccount; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO partneraccount VALUES (1, 'dennis', 'dennis', 'dennis@yahoo.com', 1, 3, '13013201230');
INSERT INTO partneraccount VALUES (2, 'riza', 'riza', 'riza@yahoo.com', 1, 2, '1132132321312');
INSERT INTO partneraccount VALUES (3, 'riza1', 'riza1', 'riza1@yahoo.com', 1, 2, '1132132321312');


--
-- Data for Name: useraccount; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO useraccount VALUES (2, 'maman', 'maman', 'maman@yahoo.com', 2, '31132312132');
INSERT INTO useraccount VALUES (3, 'riza2', 'riza2', 'riza2@yahoo.com', 2, '1132132321312');
INSERT INTO useraccount VALUES (1, 'dennis', 'dennis', 'manullang_d@yahoo.com', 2, '132132132132');
INSERT INTO useraccount VALUES (5, 'riza3', 'riza3', 'riza3@yahoo.com', 2, '1132132321312');


--
-- Name: auctions_auction_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('auctions_auction_id_seq', 5, true);


--
-- Name: biddings_bid_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('biddings_bid_id_seq', 2, true);


--
-- Name: collateral_coll_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('collateral_coll_id_seq', 3, true);


--
-- Name: m_option_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('m_option_id_seq', 1, true);


--
-- Name: partneraccount_partner_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('partneraccount_partner_id_seq', 7, true);


--
-- Name: useraccount_user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('useraccount_user_id_seq', 5, true);


--
-- PostgreSQL database dump complete
--

