package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

var (
	dburl      = os.Getenv("DB_URL")
	testdburl  = os.Getenv("TEST_DB_URL")
	dbInstance *service
)

func New() Service {
	if dbInstance != nil {
		return dbInstance
	}

	db, err := sql.Open("sqlite3", dburl)
	if err != nil {
		log.Fatal(err)
	}

	if err := initTables(db); err != nil {
		log.Fatal(err)
	}

	if err := fakeData(db); err != nil {
		fmt.Print(err)
	}

	dbInstance = &service{
		db: db,
	}
	return dbInstance
}

func Test() Service {
	if dbInstance != nil {
		return dbInstance
	}

	db, err := sql.Open("sqlite3", testdburl)
	if err != nil {
		log.Fatal(err)
	}

	if err := initTables(db); err != nil {
		log.Fatal(err)
	}

	if err := fakeData(db); err != nil {
		fmt.Print(err)
	}

	dbInstance = &service{
		db: db,
	}
	return dbInstance
}

func initTables(db *sql.DB) error {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS articles(
		id text,
		title text,
		description text,
		content text,
		created_at text,
		updated_at text,
		public integer,
		CONSTRAINT rid_pkey PRIMARY KEY (id)
	)`)
	return err
}

func fakeData(db *sql.DB) error {
	_, err := db.Exec(`
insert into articles (id, title, description, content, created_at, updated_at, public) values ('b0607ed2-d7c0-4934-b1c9-fed9d1d30777', 'Markdown Lists', 'Qucik example of markdown lists for display test purpose.', '## Lists

Unordered
+ Sub-lists are made by indenting 2 spaces:
  - Marker character change forces new list start:
    * Ac tristique libero volutpat at
    + Facilisis in pretium nisl aliquet
    - Nulla volutpat aliquam velit
+ Very easy!

Ordered

1. Lorem ipsum dolor sit amet
2. Consectetur adipiscing elit
3. Integer molestie lorem at massa



Start numbering with offset:

57. foo
1. bar', '2/1/2024', '4/23/2024', true);
insert into articles (id, title, description, content, created_at, updated_at, public) values ('17fcf145-adc8-4357-9a7b-d5497c2a0536', 'Something Wild', 'Nullam porttitor lacus at turpis. Donec posuere metus vitae ipsum.', 'Duis bibendum. Morbi non quam nec dui luctus rutrum. Nulla tellus. In sagittis dui vel nisl. Duis ac nibh. Fusce lacus purus, aliquet at, feugiat non, pretium quis, lectus. Suspendisse potenti.', '9/28/2023', '7/16/2024', false);
insert into articles (id, title, description, content, created_at, updated_at, public) values ('71d7d61e-2ead-4b66-909b-3ed3764e4448', 'Snake of June, A (Rokugatsu no hebi)', 'Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Vivamus vestibulum sagittis sapien.', 'Morbi porttitor lorem id ligula. Suspendisse ornare consequat lectus. In est risus, auctor sed, tristique in, tempus sit amet, sem. Fusce consequat. Nulla nisl. Nunc nisl. Duis bibendum, felis sed interdum venenatis, turpis enim blandit mi, in porttitor pede justo eu massa.', '7/13/2024', '10/6/2023', true);
insert into articles (id, title, description, content, created_at, updated_at, public) values ('d82bc108-ae3a-441f-8803-e193538f0c6e', 'Open House ', 'Sed sagittis. Nam congue, risus semper porta volutpat, quam pede lobortis ligula, sit amet eleifend pede libero quis orci.', 'Quisque porta volutpat erat. Quisque erat eros, viverra eget, congue eget, semper rutrum, nulla. Nunc purus. Phasellus in felis. Donec semper sapien a libero. Nam dui.', '7/3/2024', '7/28/2024', true);
insert into articles (id, title, description, content, created_at, updated_at, public) values ('c1fe4a3f-7111-4254-aefa-b7946725299c', 'Valentino: The Last Emperor', 'Morbi vestibulum, velit id pretium iaculis, diam erat fermentum justo, nec condimentum neque sapien placerat ante.', 'Curabitur in libero ut massa volutpat convallis. Morbi odio odio, elementum eu, interdum eu, tincidunt in, leo. Maecenas pulvinar lobortis est. Phasellus sit amet erat. Nulla tempus.', '3/31/2024', '2/7/2024', false);
insert into articles (id, title, description, content, created_at, updated_at, public) values ('478ec462-9c94-47b6-9451-d75913648bcf', 'Goemon', 'Duis bibendum, felis sed interdum venenatis, turpis enim blandit mi, in porttitor pede justo eu massa.', 'Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Etiam vel augue. Vestibulum rutrum rutrum neque. Aenean auctor gravida sem. Praesent id massa id nisl venenatis lacinia. Aenean sit amet justo.', '12/13/2023', '6/13/2024', false);
insert into articles (id, title, description, content, created_at, updated_at, public) values ('e50f9f00-2413-48b1-93f2-1e7f3b3bc727', 'Tale of Cinema (Geuk jang jeon)', 'Suspendisse ornare consequat lectus.', 'Maecenas tristique, est et tempus semper, est quam pharetra magna, ac consequat metus sapien ut nunc. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Mauris viverra diam vitae quam. Suspendisse potenti. Nullam porttitor lacus at turpis. Donec posuere metus vitae ipsum. Aliquam non mauris. Morbi non lectus.', '10/29/2023', '5/31/2024', true);
`)
	return err
}
