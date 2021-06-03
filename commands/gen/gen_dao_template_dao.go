package gen

const templateDaoDaoIndexContent = `
// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"{TplImportPrefix}/dao/internal"
)

// {TplTableNameCamelLowerCase}Dao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type {TplTableNameCamelLowerCase}Dao struct {
	internal.{TplTableNameCamelCase}Dao
}

var (
	// {TplTableNameCamelCase} is globally public accessible object for table {TplTableName} operations.
	{TplTableNameCamelCase} {TplTableNameCamelLowerCase}Dao
)

func init() {
	{TplTableNameCamelCase} = {TplTableNameCamelLowerCase}Dao{
		internal.{TplTableNameCamelCase},
	}
}

// Fill with you ideas below.

`

const templateDaoDaoInternalContent = `
// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
)

// {TplTableNameCamelCase}Dao is the manager for logic model data accessing
// and custom defined data operations functions management.
type {TplTableNameCamelCase}Dao struct {
	gmvc.M                                      // M is the core and embedded struct that inherits all chaining operations from gdb.Model.
	DB      gdb.DB                              // DB is the raw underlying database management object.
	Table   string                              // Table is the table name of the DAO.
	Columns {TplTableNameCamelLowerCase}Columns // Columns contains all the columns of Table that for convenient usage.
}

// {TplTableNameCamelCase}Columns defines and stores column names for table {TplTableName}.
type {TplTableNameCamelLowerCase}Columns struct {
	{TplColumnDefine}
}

var (
	// {TplTableNameCamelCase} is globally public accessible object for table {TplTableName} operations.
	{TplTableNameCamelCase} {TplTableNameCamelCase}Dao
)

func init() {
	{TplTableNameCamelCase} = {TplTableNameCamelCase}Dao{
		M:     g.DB("{TplGroupName}").Model("{TplTableName}").Safe(),
		DB:    g.DB("{TplGroupName}"),
		Table: "{TplTableName}",
		Columns: {TplTableNameCamelLowerCase}Columns{
			{TplColumnNames}
		},
	}
}
`
