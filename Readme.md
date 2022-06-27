## Getting Started

To start the application locally, run the following commands:

```bash
cd ./infra
docker-compose -f docker-compose.dev.yaml up -d  
```

Open [http://localhost:3000](http://localhost:3000) with your browser to see the result.

## SQL Script
```bash
-- public.users definition

-- Drop table

-- DROP TABLE public.users;

CREATE TABLE public.users (
	id text NOT NULL,
	username varchar(64) NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	CONSTRAINT users_pkey PRIMARY KEY (id)
);
CREATE UNIQUE INDEX idx_users_username ON public.users USING btree (username);


-- public.chats definition

-- Drop table

-- DROP TABLE public.chats;

CREATE TABLE public.chats (
	id text NOT NULL,
	message varchar(64) NOT NULL,
	user_id varchar(64) NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	CONSTRAINT chats_pkey PRIMARY KEY (id),
	CONSTRAINT fk_users_chats FOREIGN KEY (user_id) REFERENCES public.users(id)
);  
```

## Reason for using SQL database

I use SQL because of the extensive support with GORM library as well as the atomicity attribute
