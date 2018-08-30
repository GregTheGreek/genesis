package db;

import(
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
	"fmt"
	"errors"
)

type Switch struct {
	Addr 	string
	Iface	string
	Brand	int
	Id		int
}

func GetAllSwitches() Switch[] {
	db := getDB()
	defer db.Close()

	rows, err :=  db.Query(fmt.Sprintf("SELECT id,addr,iface,brand FROM %s",SERVER_TABLE ))
	checkFatal(err)
	defer rows.Close()
	switches := Switch[]{}

	for rows.Next() {
		var swtch Switch
		checkFatal(rows.Scan(&swtch.Id,&swtch.Addr,&swtch.Iface,&swtch.Brand))
		switches = append(switches,swtch)
	}
	return allServers
}

func GetSwitchById(id int) Switch {
	db := getDB()
	defer db.Close()

	rows, err :=  db.Query(fmt.Sprintf("SELECT id,addr,iface,brand FROM %s WHERE id = %d",SWITCH_TABLE,id))
	checkFatal(err)
	defer rows.Close()
	
	var swtch Switch

	for rows.Next() { //This needs to be improved
		checkFatal(rows.Scan(&swtch.Id,&swtch.Addr,&swtch.Iface,&swtch.Brand));
	}

	return swtch
}

func GetSwitchByIP(ip string) (Switch, error) {
	db := getDB()
	defer db.Close()

	rows, err :=  db.Query(fmt.Sprintf("SELECT id,addr,iface,brand FROM %s WHERE addr = %s",SWITCH_TABLE,ip))
	checkFatal(err)
	defer rows.Close()
	
	var swtch Switch

	if len(rows.Columns()) == 0 {
		return nil, errors.New()
	} 

	for rows.Next() { //This needs to be improved
		checkFatal(rows.Scan(&swtch.Id,&swtch.Addr,&swtch.Iface,&swtch.Brand));
	}

	return swtch,nil
}

func InsertSwitch(swtch Switch) int {
	db := getDB()
	defer db.Close()

	tx,err := db.Begin()
	checkFatal(err)

	stmt,err := tx.Prepare(fmt.Sprintf("INSERT INTO %s (id,addr,iface,brand) VALUES (?,?,?,?)",SERVER_TABLE))
	checkFatal(err)

	defer stmt.Close()

	_,err := stmt.Exec(server.Id, server.Addr, server.Iaddr.Ip, server.Iaddr.Gateway, server.Iaddr.Subnet,
					   server.Nodes, server.Max, switch_id,name)
	checkFatal(err)
	tx.Commit()

	tmp := GetSwitchByIP(swtch.Ip)

	return tmp.Id
}