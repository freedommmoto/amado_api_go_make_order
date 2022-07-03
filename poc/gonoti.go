package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/lib/pq"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "root"
	password = "secret"
	dbname   = "amado"
)

// notifier encapsulates the state of the listener connection.
type notifier struct {
	listener *pq.Listener
	failed   chan error
}

// newNotifier creates a new notifier for given PostgreSQL credentials.
func newNotifier(dsn, channelName string) (*notifier, error) {
	n := &notifier{failed: make(chan error, 2)}

	listener := pq.NewListener(
		dsn,
		3*time.Second, time.Minute,
		n.logListener)

	if err := listener.Listen(channelName); err != nil {
		listener.Close()
		log.Println("ERROR!:", err)
		return nil, err
	}
	fmt.Println("Successfully connected ja !")

	n.listener = listener
	return n, nil
}

// logListener is the state change callback for the listener.
func (n *notifier) logListener(event pq.ListenerEventType, err error) {

	fmt.Println("in logListener")

	if err != nil {
		log.Printf("listener error: %s\n", err)
	}
	if event == pq.ListenerEventConnectionAttemptFailed {
		n.failed <- err
	}
}

// fetch is the main loop of the notifier to receive data from
// the database in JSON-FORMAT and send it down the send channel.
func (n *notifier) fetch(data chan []byte) error {
	//var fetchCounter uint64
	addLogIntoFile("start fetch funcion")
	for {
		select {
		case e := <-n.listener.Notify:
			if e == nil {
				continue
			}
			addLogIntoFile(e.Extra)
			// fetchCounter++
			// data <- []byte(e.Extra)
			//log.Println("FETCHED DAta", []byte(e.Extra))
		case err := <-n.failed:
			return err
		case <-time.After(time.Minute):
			go n.listener.Ping()
		}
	}
}

// close closes the notifier.
func (n *notifier) close() error {
	close(n.failed)
	return n.listener.Close()
}

func addLogIntoFile(text string) {
	f, err := os.Create("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err2 := f.WriteString(text + "\n")
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println("done add text :", text)

}

func main() {
	psqlInfo := basicTestConnection()
	notifier, err := newNotifier(psqlInfo, "order_progress_event")
	if err != nil {
		fmt.Println("error after call new noti", err)
	}
	ch := make(chan []byte)
	notifier.fetch(ch)
}

func basicTestConnection() string {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=amado sslmode=disable",
		host, port, user, password)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	defer db.Close()
	fmt.Println("Successfully connected test")
	return psqlInfo
}

/*
create table if not exists production_item_wip (
                                     id serial primary key,
                                     test varchar,
                                     insert_time timestamp default NOW()
);

drop trigger if exists production_stage on production_item_wip;
drop function if exists  fn_production_stage_modified;

create
    or replace function fn_production_stage_modified() returns trigger as $psql$
begin
    perform pg_notify(
            'order_progress_event',
            'time_for_update'
        );
    return new;
end;$psql$ language plpgsql;

create trigger production_stage before
    insert
    on production_item_wip for each row
    execute procedure fn_production_stage_modified();

insert into production_item_wip (test) values ('asdf');
select count(1) from production_item_wip;


--
--
--
--
--

LISTEN order_progress_event;
notify order_progress_event, 'notify_test 111';

SELECT pg_notify('order_progress_event','notify_test 1213123');
SELECT pg_notify('order_progress_event', 'notify_test');

*/