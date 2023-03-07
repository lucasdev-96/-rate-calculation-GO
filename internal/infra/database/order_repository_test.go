package database

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/suite"
)

type OrderRepositoryTestSuit struct {
	suite.Suite
	Db *sql.DB
}

func (suite *OrderRepositoryTestSuit) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)
	db.Exec("CREATE TABLE orders (id varchar(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL, PRIMARY KEY (id))")
	suite.Db = db
}

func (suite *OrderRepositoryTestSuit) TearDownSuite() {
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(OrderRepositoryTestSuit))
}

// func (suite *OrderRepositoryTestSuit) TestSave() {
// 	order, err := entity.NewOrder("123", 10.0, 1.0)
// 	suite.NoError(err)
// 	suite.NoError(order.CalculateFinalPrice())
// 	repo := NewOrderRepository(suite.Db)
// 	err = repo.Save(order)
// 	suite.NoError(err)
// }
