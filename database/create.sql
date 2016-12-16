create table users (
	id serial not null primary key,
	name text not null default '',
	date timestamptz not null default localtimestamp,
	last_ip text not null default '',
	has_stripe bool not null default false,
	stripe_token text not null default '',
	has_paypal bool not null default false,
	paypal_name text not null default '',
	country text not null default 'AU',
	email text not null default '',
	mobile text not null default '',
	gender text not null default 'F',
	newsletter bool default true,
	sms_alerts bool default false
);

create table address (
	id serial not null primary key,
	date timestamptz not null default localtimestamp,
	user_id int not null default 0,
	country text not null default '',
	state_code text not null default '',
	postcode text not null default '',
	name text not null default '',
	address_1 text not null default '',
	address_2 text not null default ''
);

create table sale_order (
	id serial not null primary key,
	date timestamptz not null default localtimestamp,
	user_id int not null default 0,
	address_id int not null default 0
);

create table sale_order_item (
	order_id int not null,
	item_id int not null,
	qty int not null
);

create table item (
	id serial not null primary key,
	sku text not null default '',
	price numeric(8,2),
	name text not null default '',
	descr text not null default '',
	image text not null default '',
	volume_ml int not null,
	weight_g int not null,
	shipping_volume_ml int not null,
	shipping_weight_g int not null
);
create unique index item_sku on item (sku);

create table blog (
	id serial not null primary key,
	image text not null default '', 
	date timestamptz not null default localtimestamp,
	name text not null default '',
	title text not null default '',
	content text not null default '',
	share_twitter int not null default 0,
	share_facebook int not null default 0,
	share_instagram int not null default 0,
	share_google_plus int not null default 0
);

create table tags (
	blog_id int not null,
	tag text
);

create index tags_tag on tags (tag);

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
\.

