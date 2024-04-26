package dao

type DAOGroup struct {
	UserDao
}

var DAOGroupApp = new(DAOGroup)
