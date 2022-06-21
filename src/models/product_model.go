package models

import (
	"database/sql"
	"fmt"
	"strings"

	"main.go/src/entities"
)

type ProductModel struct {
	Db *sql.DB
}

//Query to describe the table
func (pm ProductModel) DescribeTable() ([]entities.TableStructure, error) {
	query := "select  COLUMN_NAME, DATA_TYPE from INFORMATION_SCHEMA.COLUMNS WHERE  TABLE_NAME = 'cars' "
	rows, err1 := pm.Db.Query(query)
	if err1 != nil {
		return nil, err1
	}
	var str1 string
	var str2 string
	data := make([]entities.TableStructure, 0)
	for rows.Next() {
		rows.Scan(&str1, &str2)
		var x entities.TableStructure
		x.ColName = str1
		x.TypeData = str2
		data = append(data, x)
	}
	return data, nil
}

//Select * from the given table and return the result
func (pm ProductModel) QuerySelectFromGivenColorAndType(filetr1 string, filter2 string) ([]entities.Table1, error) {
	query := " select * from cars where type = ? and color = ?"
	//replace ? with filter1 and filter 2
	rows, err := pm.Db.Query(query)
	var cars []entities.Table1
	for rows.Next() {
		var id string
		var cou string
		car := entities.Table1{}
		rows.Scan(&id, &cou)
		car.Id = id
		car.Country = cou
		cars = append(cars, car)
	}

	return cars, err
}

func (pm ProductModel) QuerySelectNameFromGivenColor(filetr1 string) ([]entities.Table1, error) {
	query := "select name from cars where color = ?"
	//replace ? with filter1
	rows, err := pm.Db.Query(query)
	var cars []entities.Table1
	for rows.Next() {
		var name string
		car := entities.Table1{}
		rows.Scan(&name)
		car.Id = name
		cars = append(cars, car)
	}
	return cars, err
}

func (pm ProductModel) QuerySelectEvertythingFromGivenColor(filter string) ([]entities.Table1, error) {
	query := "select * from cars where color = ?"
	//replace ? with filter1
	rows, err := pm.Db.Query(query)
	var cars []entities.Table1
	for rows.Next() {
		var name string
		car := entities.Table1{}
		rows.Scan(&name)
		car.Id = name
		cars = append(cars, car)
	}
	return cars, err
}

//Select id col from the given table and return the result
func (pm ProductModel) SelectIdFromTable(query string) (*sql.Rows, error) {
	rows, err := pm.Db.Query(query)
	return rows, err
}

//Select * from the given table return and the result
func (pm ProductModel) SelectEveryThingFromTable(query string) (*sql.Rows, error) {
	rows, err := pm.Db.Query(query)
	return rows, err
}

//Select * from the given table where id ("x,"y","z"...) and return the result
func (pm ProductModel) selectfromTableWhere(ids []string, query string) (*sql.Rows, error) {

	queryval := strings.Join(ids, ",")
	str := "(" + queryval + ")"
	query = query + str
	fmt.Println(query)
	rows, err := pm.Db.Query(query)
	return rows, err
}

//Scan table 1 result and put it in the data container
func ScanTable1(rows *sql.Rows, combined *[]entities.Combine) error {
	var col1 string
	var col2 string
	var i int = 0
	for rows.Next() {
		err2 := rows.Scan(&col1, &col2)
		if err2 != nil {
			return err2
		} else {
			var combine entities.Combine
			combine.X.Id = col1
			combine.X.Country = col2
			*combined = append(*combined, combine)
			i++
		}
	}
	return nil
}

//Scan table 2 result and put it in the same data container
func ScanTable2(rows *sql.Rows, combined *[]entities.Combine) error {
	var col1 string
	var col2 string
	var i int = 0
	com := *combined
	for rows.Next() {
		err2 := rows.Scan(&col1, &col2)
		if err2 != nil {
			return err2
		} else {
			if i < len(com) {
				com[i].Y.Id = col1
				com[i].Y.State = col2
				i++
			} else {
				var combine entities.Combine
				combine.Y.Id = col1
				combine.Y.State = col2
				*combined = append(*combined, combine)
			}
		}
	}
	return nil
}

// 1. Select unique list of 'Id's from Table1
// 2. Select Id from table2 where id in  (step 1 result list)
// 3. Select * from Table1 where Id is in (step 2 result list)
// 4. Select * from Table2 where Id is in (step 2 result list)
// 5. Join data from steps 2 and 4 locally and return the result

func (pm ProductModel) FindInnerJoin() (produt []entities.Combine, err error) {

	query := "select id from table1"
	rows, err := pm.SelectIdFromTable(query)
	if err != nil {
		return nil, err
	} else {
		ids := make([]string, 0)
		var i int = 0
		for rows.Next() {
			ids = append(ids, "")
			err2 := rows.Scan(&ids[i])
			if err2 != nil {
				return nil, err
			}
			i++
		}
		query3 := "select id from table2 where id in"
		rows1, _ := pm.selectfromTableWhere(ids, query3)

		id1s := make([]string, 0)
		var i1 int = 0
		for rows1.Next() {
			id1s = append(id1s, "")
			err21 := rows1.Scan(&id1s[i1])
			if err21 != nil {
				return nil, err
			}
			i1++
		}

		query1 := "select * from table1 where id in"
		query2 := "select * from table2 where id in"
		val := make([]entities.Combine, 0)
		rows5, err5 := pm.selectfromTableWhere(id1s, query1)
		if err5 != nil {
			return nil, err5
		}
		errTable1 := ScanTable1(rows5, &val)
		if errTable1 != nil {
			return nil, err
		}
		rows2, err2 := pm.selectfromTableWhere(id1s, query2)
		if err2 != nil {
			return nil, err
		}
		errTable2 := ScanTable2(rows2, &val)
		if errTable2 != nil {
			return nil, err
		}
		return val, nil
	}
}

// 1. Select unique list of 'Id's from Table1
// 2. Select * from Table2 where Id is in (first select result list)
// 3. Select * from Table1
// 4. Join data from steps 2 and 3 locally and return the result

func (pm ProductModel) FindLeftJoin() (produt []entities.Combine, err error) {
	query := "select id from table1"
	rows, err := pm.SelectIdFromTable(query)
	if err != nil {
		return nil, err
	} else {
		var ids []string
		var idx int = 0
		for rows.Next() {
			err2 := rows.Scan(ids[idx])
			if err2 != nil {
				return nil, err
			}
			idx++
		}
		query1 := "select * from table1"
		query2 := "select * from table2 where id in"
		var val *[]entities.Combine
		rows1, err1 := pm.SelectEveryThingFromTable(query1)
		if err1 != nil {
			return nil, err
		}
		errTable1 := ScanTable1(rows1, val)
		if errTable1 != nil {
			return nil, err
		}
		rows2, err2 := pm.selectfromTableWhere(ids, query2)
		if err2 != nil {
			return nil, err
		}
		errTable2 := ScanTable2(rows2, val)
		if errTable2 != nil {
			return nil, err
		}
		return *val, nil
	}
}

// 1. Select unique list of 'Id's from Table2
// 2. Select * from Table1 where Id is in (first select result list)
// 3. Select * from Table2
// 4. Join data from steps 2 and 3 locally and return the result
func (pm ProductModel) FindRightJoin() (produt []entities.Combine, err error) {
	query := "select id from table2"
	rows, err := pm.SelectIdFromTable(query)
	if err != nil {
		return nil, err
	} else {
		ids := make([]string, 0)
		var i int = 0
		for rows.Next() {
			ids = append(ids, "")
			err2 := rows.Scan(&ids[i])
			if err2 != nil {
				return nil, err
			}
			i++
		}
		query1 := "select * from table1 where id in"
		query2 := "select * from table2"
		val := make([]entities.Combine, 0)
		rows1, err2 := pm.selectfromTableWhere(ids, query1)
		if err2 != nil {
			return nil, err
		}
		errTable1 := ScanTable1(rows1, &val)
		if errTable1 != nil {
			return nil, err
		}
		rows2, err1 := pm.SelectEveryThingFromTable(query2)
		if err1 != nil {
			return nil, err
		}
		errTable2 := ScanTable2(rows2, &val)
		if errTable2 != nil {
			return nil, err
		}
		return val, nil
	}
}

// 1. Select * from Table1
// 2. Select * from Table2
// 3. Join data from steps 1 and 2 locally and return the result
func (pm ProductModel) FindFullJoin() (produt []entities.Combine, err error) {
	query1 := "select * from table1"
	query2 := "select * from table2"
	val := make([]entities.Combine, 0)
	rows1, err2 := pm.SelectEveryThingFromTable(query1)
	if err2 != nil {
		return nil, err
	}
	errtable1 := ScanTable1(rows1, &val)
	if errtable1 != nil {
		return nil, err
	}
	rows2, err1 := pm.SelectEveryThingFromTable(query2)
	if err1 != nil {
		return nil, err
	}
	errtable2 := ScanTable2(rows2, &val)
	if errtable2 != nil {
		return nil, err
	}
	return val, nil
}
