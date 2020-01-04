package models

import "../config"

var (
	db  = config.DBInit()
	idb = config.InDB{DB: db}
)