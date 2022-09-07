package models

type Collateral struct {
	Coll_id int `orm:"pk;auto"`
	Acc_id  int
	Type_id int	
	Owner_name string
	Coll_location string
	Initial_col_price int 
	Final_col_price int
	Ljk_id int
	Document_path string
}

func (a *Collateral) TableName() string {
	return "collateral"
}