// Code generated by SQLBoiler 4.11.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Roles", testRoles)
	t.Run("UserRoles", testUserRoles)
	t.Run("Users", testUsers)
}

func TestDelete(t *testing.T) {
	t.Run("Roles", testRolesDelete)
	t.Run("UserRoles", testUserRolesDelete)
	t.Run("Users", testUsersDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Roles", testRolesQueryDeleteAll)
	t.Run("UserRoles", testUserRolesQueryDeleteAll)
	t.Run("Users", testUsersQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Roles", testRolesSliceDeleteAll)
	t.Run("UserRoles", testUserRolesSliceDeleteAll)
	t.Run("Users", testUsersSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Roles", testRolesExists)
	t.Run("UserRoles", testUserRolesExists)
	t.Run("Users", testUsersExists)
}

func TestFind(t *testing.T) {
	t.Run("Roles", testRolesFind)
	t.Run("UserRoles", testUserRolesFind)
	t.Run("Users", testUsersFind)
}

func TestBind(t *testing.T) {
	t.Run("Roles", testRolesBind)
	t.Run("UserRoles", testUserRolesBind)
	t.Run("Users", testUsersBind)
}

func TestOne(t *testing.T) {
	t.Run("Roles", testRolesOne)
	t.Run("UserRoles", testUserRolesOne)
	t.Run("Users", testUsersOne)
}

func TestAll(t *testing.T) {
	t.Run("Roles", testRolesAll)
	t.Run("UserRoles", testUserRolesAll)
	t.Run("Users", testUsersAll)
}

func TestCount(t *testing.T) {
	t.Run("Roles", testRolesCount)
	t.Run("UserRoles", testUserRolesCount)
	t.Run("Users", testUsersCount)
}

func TestHooks(t *testing.T) {
	t.Run("Roles", testRolesHooks)
	t.Run("UserRoles", testUserRolesHooks)
	t.Run("Users", testUsersHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Roles", testRolesInsert)
	t.Run("Roles", testRolesInsertWhitelist)
	t.Run("UserRoles", testUserRolesInsert)
	t.Run("UserRoles", testUserRolesInsertWhitelist)
	t.Run("Users", testUsersInsert)
	t.Run("Users", testUsersInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("UserRoleToRoleUsingRole", testUserRoleToOneRoleUsingRole)
	t.Run("UserRoleToRoleUsingUser", testUserRoleToOneRoleUsingUser)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("RoleToUserRoles", testRoleToManyUserRoles)
	t.Run("RoleToUserUserRoles", testRoleToManyUserUserRoles)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("UserRoleToRoleUsingUserRoles", testUserRoleToOneSetOpRoleUsingRole)
	t.Run("UserRoleToRoleUsingUserUserRoles", testUserRoleToOneSetOpRoleUsingUser)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("RoleToUserRoles", testRoleToManyAddOpUserRoles)
	t.Run("RoleToUserUserRoles", testRoleToManyAddOpUserUserRoles)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("Roles", testRolesReload)
	t.Run("UserRoles", testUserRolesReload)
	t.Run("Users", testUsersReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Roles", testRolesReloadAll)
	t.Run("UserRoles", testUserRolesReloadAll)
	t.Run("Users", testUsersReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Roles", testRolesSelect)
	t.Run("UserRoles", testUserRolesSelect)
	t.Run("Users", testUsersSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Roles", testRolesUpdate)
	t.Run("UserRoles", testUserRolesUpdate)
	t.Run("Users", testUsersUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Roles", testRolesSliceUpdateAll)
	t.Run("UserRoles", testUserRolesSliceUpdateAll)
	t.Run("Users", testUsersSliceUpdateAll)
}
