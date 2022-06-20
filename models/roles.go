// Code generated by SQLBoiler 4.11.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Role is an object representing the database table.
type Role struct {
	ID       string    `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name     string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	CreateAt time.Time `boil:"create_at" json:"create_at" toml:"create_at" yaml:"create_at"`
	UpdateAt time.Time `boil:"update_at" json:"update_at" toml:"update_at" yaml:"update_at"`

	R *roleR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L roleL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var RoleColumns = struct {
	ID       string
	Name     string
	CreateAt string
	UpdateAt string
}{
	ID:       "id",
	Name:     "name",
	CreateAt: "create_at",
	UpdateAt: "update_at",
}

var RoleTableColumns = struct {
	ID       string
	Name     string
	CreateAt string
	UpdateAt string
}{
	ID:       "roles.id",
	Name:     "roles.name",
	CreateAt: "roles.create_at",
	UpdateAt: "roles.update_at",
}

// Generated where

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var RoleWhere = struct {
	ID       whereHelperstring
	Name     whereHelperstring
	CreateAt whereHelpertime_Time
	UpdateAt whereHelpertime_Time
}{
	ID:       whereHelperstring{field: "\"roles\".\"id\""},
	Name:     whereHelperstring{field: "\"roles\".\"name\""},
	CreateAt: whereHelpertime_Time{field: "\"roles\".\"create_at\""},
	UpdateAt: whereHelpertime_Time{field: "\"roles\".\"update_at\""},
}

// RoleRels is where relationship names are stored.
var RoleRels = struct {
	UserRoles     string
	UserUserRoles string
}{
	UserRoles:     "UserRoles",
	UserUserRoles: "UserUserRoles",
}

// roleR is where relationships are stored.
type roleR struct {
	UserRoles     UserRoleSlice `boil:"UserRoles" json:"UserRoles" toml:"UserRoles" yaml:"UserRoles"`
	UserUserRoles UserRoleSlice `boil:"UserUserRoles" json:"UserUserRoles" toml:"UserUserRoles" yaml:"UserUserRoles"`
}

// NewStruct creates a new relationship struct
func (*roleR) NewStruct() *roleR {
	return &roleR{}
}

func (r *roleR) GetUserRoles() UserRoleSlice {
	if r == nil {
		return nil
	}
	return r.UserRoles
}

func (r *roleR) GetUserUserRoles() UserRoleSlice {
	if r == nil {
		return nil
	}
	return r.UserUserRoles
}

// roleL is where Load methods for each relationship are stored.
type roleL struct{}

var (
	roleAllColumns            = []string{"id", "name", "create_at", "update_at"}
	roleColumnsWithoutDefault = []string{"name"}
	roleColumnsWithDefault    = []string{"id", "create_at", "update_at"}
	rolePrimaryKeyColumns     = []string{"id"}
	roleGeneratedColumns      = []string{}
)

type (
	// RoleSlice is an alias for a slice of pointers to Role.
	// This should almost always be used instead of []Role.
	RoleSlice []*Role
	// RoleHook is the signature for custom Role hook methods
	RoleHook func(context.Context, boil.ContextExecutor, *Role) error

	roleQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	roleType                 = reflect.TypeOf(&Role{})
	roleMapping              = queries.MakeStructMapping(roleType)
	rolePrimaryKeyMapping, _ = queries.BindMapping(roleType, roleMapping, rolePrimaryKeyColumns)
	roleInsertCacheMut       sync.RWMutex
	roleInsertCache          = make(map[string]insertCache)
	roleUpdateCacheMut       sync.RWMutex
	roleUpdateCache          = make(map[string]updateCache)
	roleUpsertCacheMut       sync.RWMutex
	roleUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var roleAfterSelectHooks []RoleHook

var roleBeforeInsertHooks []RoleHook
var roleAfterInsertHooks []RoleHook

var roleBeforeUpdateHooks []RoleHook
var roleAfterUpdateHooks []RoleHook

var roleBeforeDeleteHooks []RoleHook
var roleAfterDeleteHooks []RoleHook

var roleBeforeUpsertHooks []RoleHook
var roleAfterUpsertHooks []RoleHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Role) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range roleAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Role) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range roleBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Role) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range roleAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Role) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range roleBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Role) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range roleAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Role) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range roleBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Role) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range roleAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Role) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range roleBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Role) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range roleAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddRoleHook registers your hook function for all future operations.
func AddRoleHook(hookPoint boil.HookPoint, roleHook RoleHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		roleAfterSelectHooks = append(roleAfterSelectHooks, roleHook)
	case boil.BeforeInsertHook:
		roleBeforeInsertHooks = append(roleBeforeInsertHooks, roleHook)
	case boil.AfterInsertHook:
		roleAfterInsertHooks = append(roleAfterInsertHooks, roleHook)
	case boil.BeforeUpdateHook:
		roleBeforeUpdateHooks = append(roleBeforeUpdateHooks, roleHook)
	case boil.AfterUpdateHook:
		roleAfterUpdateHooks = append(roleAfterUpdateHooks, roleHook)
	case boil.BeforeDeleteHook:
		roleBeforeDeleteHooks = append(roleBeforeDeleteHooks, roleHook)
	case boil.AfterDeleteHook:
		roleAfterDeleteHooks = append(roleAfterDeleteHooks, roleHook)
	case boil.BeforeUpsertHook:
		roleBeforeUpsertHooks = append(roleBeforeUpsertHooks, roleHook)
	case boil.AfterUpsertHook:
		roleAfterUpsertHooks = append(roleAfterUpsertHooks, roleHook)
	}
}

// One returns a single role record from the query.
func (q roleQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Role, error) {
	o := &Role{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for roles")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Role records from the query.
func (q roleQuery) All(ctx context.Context, exec boil.ContextExecutor) (RoleSlice, error) {
	var o []*Role

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Role slice")
	}

	if len(roleAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Role records in the query.
func (q roleQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count roles rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q roleQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if roles exists")
	}

	return count > 0, nil
}

// UserRoles retrieves all the user_role's UserRoles with an executor.
func (o *Role) UserRoles(mods ...qm.QueryMod) userRoleQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"user_roles\".\"role_id\"=?", o.ID),
	)

	return UserRoles(queryMods...)
}

// UserUserRoles retrieves all the user_role's UserRoles with an executor via user_id column.
func (o *Role) UserUserRoles(mods ...qm.QueryMod) userRoleQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"user_roles\".\"user_id\"=?", o.ID),
	)

	return UserRoles(queryMods...)
}

// LoadUserRoles allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (roleL) LoadUserRoles(ctx context.Context, e boil.ContextExecutor, singular bool, maybeRole interface{}, mods queries.Applicator) error {
	var slice []*Role
	var object *Role

	if singular {
		object = maybeRole.(*Role)
	} else {
		slice = *maybeRole.(*[]*Role)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &roleR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &roleR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`user_roles`),
		qm.WhereIn(`user_roles.role_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load user_roles")
	}

	var resultSlice []*UserRole
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice user_roles")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on user_roles")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for user_roles")
	}

	if len(userRoleAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.UserRoles = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &userRoleR{}
			}
			foreign.R.Role = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.RoleID {
				local.R.UserRoles = append(local.R.UserRoles, foreign)
				if foreign.R == nil {
					foreign.R = &userRoleR{}
				}
				foreign.R.Role = local
				break
			}
		}
	}

	return nil
}

// LoadUserUserRoles allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (roleL) LoadUserUserRoles(ctx context.Context, e boil.ContextExecutor, singular bool, maybeRole interface{}, mods queries.Applicator) error {
	var slice []*Role
	var object *Role

	if singular {
		object = maybeRole.(*Role)
	} else {
		slice = *maybeRole.(*[]*Role)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &roleR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &roleR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`user_roles`),
		qm.WhereIn(`user_roles.user_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load user_roles")
	}

	var resultSlice []*UserRole
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice user_roles")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on user_roles")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for user_roles")
	}

	if len(userRoleAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.UserUserRoles = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &userRoleR{}
			}
			foreign.R.User = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.UserID {
				local.R.UserUserRoles = append(local.R.UserUserRoles, foreign)
				if foreign.R == nil {
					foreign.R = &userRoleR{}
				}
				foreign.R.User = local
				break
			}
		}
	}

	return nil
}

// AddUserRoles adds the given related objects to the existing relationships
// of the role, optionally inserting them as new records.
// Appends related to o.R.UserRoles.
// Sets related.R.Role appropriately.
func (o *Role) AddUserRoles(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*UserRole) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.RoleID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"user_roles\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"role_id"}),
				strmangle.WhereClause("\"", "\"", 2, userRolePrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.RoleID = o.ID
		}
	}

	if o.R == nil {
		o.R = &roleR{
			UserRoles: related,
		}
	} else {
		o.R.UserRoles = append(o.R.UserRoles, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &userRoleR{
				Role: o,
			}
		} else {
			rel.R.Role = o
		}
	}
	return nil
}

// AddUserUserRoles adds the given related objects to the existing relationships
// of the role, optionally inserting them as new records.
// Appends related to o.R.UserUserRoles.
// Sets related.R.User appropriately.
func (o *Role) AddUserUserRoles(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*UserRole) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.UserID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"user_roles\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"user_id"}),
				strmangle.WhereClause("\"", "\"", 2, userRolePrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.UserID = o.ID
		}
	}

	if o.R == nil {
		o.R = &roleR{
			UserUserRoles: related,
		}
	} else {
		o.R.UserUserRoles = append(o.R.UserUserRoles, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &userRoleR{
				User: o,
			}
		} else {
			rel.R.User = o
		}
	}
	return nil
}

// Roles retrieves all the records using an executor.
func Roles(mods ...qm.QueryMod) roleQuery {
	mods = append(mods, qm.From("\"roles\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"roles\".*"})
	}

	return roleQuery{q}
}

// FindRole retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindRole(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Role, error) {
	roleObj := &Role{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"roles\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, roleObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from roles")
	}

	if err = roleObj.doAfterSelectHooks(ctx, exec); err != nil {
		return roleObj, err
	}

	return roleObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Role) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no roles provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(roleColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	roleInsertCacheMut.RLock()
	cache, cached := roleInsertCache[key]
	roleInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			roleAllColumns,
			roleColumnsWithDefault,
			roleColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(roleType, roleMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(roleType, roleMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"roles\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"roles\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into roles")
	}

	if !cached {
		roleInsertCacheMut.Lock()
		roleInsertCache[key] = cache
		roleInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Role.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Role) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	roleUpdateCacheMut.RLock()
	cache, cached := roleUpdateCache[key]
	roleUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			roleAllColumns,
			rolePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update roles, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"roles\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, rolePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(roleType, roleMapping, append(wl, rolePrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update roles row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for roles")
	}

	if !cached {
		roleUpdateCacheMut.Lock()
		roleUpdateCache[key] = cache
		roleUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q roleQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for roles")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for roles")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o RoleSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), rolePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"roles\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, rolePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in role slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all role")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Role) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no roles provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(roleColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	roleUpsertCacheMut.RLock()
	cache, cached := roleUpsertCache[key]
	roleUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			roleAllColumns,
			roleColumnsWithDefault,
			roleColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			roleAllColumns,
			rolePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert roles, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(rolePrimaryKeyColumns))
			copy(conflict, rolePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"roles\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(roleType, roleMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(roleType, roleMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert roles")
	}

	if !cached {
		roleUpsertCacheMut.Lock()
		roleUpsertCache[key] = cache
		roleUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Role record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Role) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Role provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), rolePrimaryKeyMapping)
	sql := "DELETE FROM \"roles\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from roles")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for roles")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q roleQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no roleQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from roles")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for roles")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o RoleSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(roleBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), rolePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"roles\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, rolePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from role slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for roles")
	}

	if len(roleAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Role) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindRole(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *RoleSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := RoleSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), rolePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"roles\".* FROM \"roles\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, rolePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in RoleSlice")
	}

	*o = slice

	return nil
}

// RoleExists checks if the Role row exists.
func RoleExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"roles\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if roles exists")
	}

	return exists, nil
}
