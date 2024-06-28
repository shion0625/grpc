// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

// Product is an object representing the database table.
type Product struct {
	ID         int    `boil:"id" json:"id" toml:"id" yaml:"id"`
	ObjID      string `boil:"obj_id" json:"obj_id" toml:"obj_id" yaml:"obj_id"`
	Name       string `boil:"name" json:"name" toml:"name" yaml:"name"`
	Price      int    `boil:"price" json:"price" toml:"price" yaml:"price"`
	CategoryID string `boil:"category_id" json:"category_id" toml:"category_id" yaml:"category_id"`

	R *productR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L productL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ProductColumns = struct {
	ID         string
	ObjID      string
	Name       string
	Price      string
	CategoryID string
}{
	ID:         "id",
	ObjID:      "obj_id",
	Name:       "name",
	Price:      "price",
	CategoryID: "category_id",
}

var ProductTableColumns = struct {
	ID         string
	ObjID      string
	Name       string
	Price      string
	CategoryID string
}{
	ID:         "product.id",
	ObjID:      "product.obj_id",
	Name:       "product.name",
	Price:      "product.price",
	CategoryID: "product.category_id",
}

// Generated where

var ProductWhere = struct {
	ID         whereHelperint
	ObjID      whereHelperstring
	Name       whereHelperstring
	Price      whereHelperint
	CategoryID whereHelperstring
}{
	ID:         whereHelperint{field: "`product`.`id`"},
	ObjID:      whereHelperstring{field: "`product`.`obj_id`"},
	Name:       whereHelperstring{field: "`product`.`name`"},
	Price:      whereHelperint{field: "`product`.`price`"},
	CategoryID: whereHelperstring{field: "`product`.`category_id`"},
}

// ProductRels is where relationship names are stored.
var ProductRels = struct {
	Category string
}{
	Category: "Category",
}

// productR is where relationships are stored.
type productR struct {
	Category *Category `boil:"Category" json:"Category" toml:"Category" yaml:"Category"`
}

// NewStruct creates a new relationship struct
func (*productR) NewStruct() *productR {
	return &productR{}
}

func (r *productR) GetCategory() *Category {
	if r == nil {
		return nil
	}
	return r.Category
}

// productL is where Load methods for each relationship are stored.
type productL struct{}

var (
	productAllColumns            = []string{"id", "obj_id", "name", "price", "category_id"}
	productColumnsWithoutDefault = []string{"obj_id", "name", "price", "category_id"}
	productColumnsWithDefault    = []string{"id"}
	productPrimaryKeyColumns     = []string{"id"}
	productGeneratedColumns      = []string{}
)

type (
	// ProductSlice is an alias for a slice of pointers to Product.
	// This should almost always be used instead of []Product.
	ProductSlice []*Product
	// ProductHook is the signature for custom Product hook methods
	ProductHook func(context.Context, boil.ContextExecutor, *Product) error

	productQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	productType                 = reflect.TypeOf(&Product{})
	productMapping              = queries.MakeStructMapping(productType)
	productPrimaryKeyMapping, _ = queries.BindMapping(productType, productMapping, productPrimaryKeyColumns)
	productInsertCacheMut       sync.RWMutex
	productInsertCache          = make(map[string]insertCache)
	productUpdateCacheMut       sync.RWMutex
	productUpdateCache          = make(map[string]updateCache)
	productUpsertCacheMut       sync.RWMutex
	productUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var productAfterSelectMu sync.Mutex
var productAfterSelectHooks []ProductHook

var productBeforeInsertMu sync.Mutex
var productBeforeInsertHooks []ProductHook
var productAfterInsertMu sync.Mutex
var productAfterInsertHooks []ProductHook

var productBeforeUpdateMu sync.Mutex
var productBeforeUpdateHooks []ProductHook
var productAfterUpdateMu sync.Mutex
var productAfterUpdateHooks []ProductHook

var productBeforeDeleteMu sync.Mutex
var productBeforeDeleteHooks []ProductHook
var productAfterDeleteMu sync.Mutex
var productAfterDeleteHooks []ProductHook

var productBeforeUpsertMu sync.Mutex
var productBeforeUpsertHooks []ProductHook
var productAfterUpsertMu sync.Mutex
var productAfterUpsertHooks []ProductHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Product) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Product) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Product) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Product) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Product) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Product) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Product) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Product) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Product) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddProductHook registers your hook function for all future operations.
func AddProductHook(hookPoint boil.HookPoint, productHook ProductHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		productAfterSelectMu.Lock()
		productAfterSelectHooks = append(productAfterSelectHooks, productHook)
		productAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		productBeforeInsertMu.Lock()
		productBeforeInsertHooks = append(productBeforeInsertHooks, productHook)
		productBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		productAfterInsertMu.Lock()
		productAfterInsertHooks = append(productAfterInsertHooks, productHook)
		productAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		productBeforeUpdateMu.Lock()
		productBeforeUpdateHooks = append(productBeforeUpdateHooks, productHook)
		productBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		productAfterUpdateMu.Lock()
		productAfterUpdateHooks = append(productAfterUpdateHooks, productHook)
		productAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		productBeforeDeleteMu.Lock()
		productBeforeDeleteHooks = append(productBeforeDeleteHooks, productHook)
		productBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		productAfterDeleteMu.Lock()
		productAfterDeleteHooks = append(productAfterDeleteHooks, productHook)
		productAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		productBeforeUpsertMu.Lock()
		productBeforeUpsertHooks = append(productBeforeUpsertHooks, productHook)
		productBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		productAfterUpsertMu.Lock()
		productAfterUpsertHooks = append(productAfterUpsertHooks, productHook)
		productAfterUpsertMu.Unlock()
	}
}

// One returns a single product record from the query.
func (q productQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Product, error) {
	o := &Product{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for product")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Product records from the query.
func (q productQuery) All(ctx context.Context, exec boil.ContextExecutor) (ProductSlice, error) {
	var o []*Product

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Product slice")
	}

	if len(productAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Product records in the query.
func (q productQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count product rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q productQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if product exists")
	}

	return count > 0, nil
}

// Category pointed to by the foreign key.
func (o *Product) Category(mods ...qm.QueryMod) categoryQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`obj_id` = ?", o.CategoryID),
	}

	queryMods = append(queryMods, mods...)

	return Categories(queryMods...)
}

// LoadCategory allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (productL) LoadCategory(ctx context.Context, e boil.ContextExecutor, singular bool, maybeProduct interface{}, mods queries.Applicator) error {
	var slice []*Product
	var object *Product

	if singular {
		var ok bool
		object, ok = maybeProduct.(*Product)
		if !ok {
			object = new(Product)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeProduct)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeProduct))
			}
		}
	} else {
		s, ok := maybeProduct.(*[]*Product)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeProduct)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeProduct))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &productR{}
		}
		args[object.CategoryID] = struct{}{}

	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &productR{}
			}

			args[obj.CategoryID] = struct{}{}

		}
	}

	if len(args) == 0 {
		return nil
	}

	argsSlice := make([]interface{}, len(args))
	i := 0
	for arg := range args {
		argsSlice[i] = arg
		i++
	}

	query := NewQuery(
		qm.From(`category`),
		qm.WhereIn(`category.obj_id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Category")
	}

	var resultSlice []*Category
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Category")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for category")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for category")
	}

	if len(categoryAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Category = foreign
		if foreign.R == nil {
			foreign.R = &categoryR{}
		}
		foreign.R.Products = append(foreign.R.Products, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.CategoryID == foreign.ObjID {
				local.R.Category = foreign
				if foreign.R == nil {
					foreign.R = &categoryR{}
				}
				foreign.R.Products = append(foreign.R.Products, local)
				break
			}
		}
	}

	return nil
}

// SetCategory of the product to the related item.
// Sets o.R.Category to related.
// Adds o to related.R.Products.
func (o *Product) SetCategory(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Category) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `product` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"category_id"}),
		strmangle.WhereClause("`", "`", 0, productPrimaryKeyColumns),
	)
	values := []interface{}{related.ObjID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.CategoryID = related.ObjID
	if o.R == nil {
		o.R = &productR{
			Category: related,
		}
	} else {
		o.R.Category = related
	}

	if related.R == nil {
		related.R = &categoryR{
			Products: ProductSlice{o},
		}
	} else {
		related.R.Products = append(related.R.Products, o)
	}

	return nil
}

// Products retrieves all the records using an executor.
func Products(mods ...qm.QueryMod) productQuery {
	mods = append(mods, qm.From("`product`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`product`.*"})
	}

	return productQuery{q}
}

// FindProduct retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindProduct(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Product, error) {
	productObj := &Product{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `product` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, productObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from product")
	}

	if err = productObj.doAfterSelectHooks(ctx, exec); err != nil {
		return productObj, err
	}

	return productObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Product) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no product provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(productColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	productInsertCacheMut.RLock()
	cache, cached := productInsertCache[key]
	productInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			productAllColumns,
			productColumnsWithDefault,
			productColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(productType, productMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(productType, productMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `product` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `product` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `product` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, productPrimaryKeyColumns))
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
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into product")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == productMapping["id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for product")
	}

CacheNoHooks:
	if !cached {
		productInsertCacheMut.Lock()
		productInsertCache[key] = cache
		productInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Product.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Product) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	productUpdateCacheMut.RLock()
	cache, cached := productUpdateCache[key]
	productUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			productAllColumns,
			productPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update product, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `product` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, productPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(productType, productMapping, append(wl, productPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update product row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for product")
	}

	if !cached {
		productUpdateCacheMut.Lock()
		productUpdateCache[key] = cache
		productUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q productQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for product")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for product")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ProductSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), productPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `product` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, productPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in product slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all product")
	}
	return rowsAff, nil
}

var mySQLProductUniqueColumns = []string{
	"id",
	"obj_id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Product) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no product provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(productColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLProductUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
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
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	productUpsertCacheMut.RLock()
	cache, cached := productUpsertCache[key]
	productUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			productAllColumns,
			productColumnsWithDefault,
			productColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			productAllColumns,
			productPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert product, could not build update column list")
		}

		ret := strmangle.SetComplement(productAllColumns, strmangle.SetIntersect(insert, update))

		cache.query = buildUpsertQueryMySQL(dialect, "`product`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `product` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(productType, productMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(productType, productMapping, ret)
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
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for product")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == productMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(productType, productMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for product")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for product")
	}

CacheNoHooks:
	if !cached {
		productUpsertCacheMut.Lock()
		productUpsertCache[key] = cache
		productUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Product record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Product) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Product provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), productPrimaryKeyMapping)
	sql := "DELETE FROM `product` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from product")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for product")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q productQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no productQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from product")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for product")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ProductSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(productBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), productPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `product` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, productPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from product slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for product")
	}

	if len(productAfterDeleteHooks) != 0 {
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
func (o *Product) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindProduct(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ProductSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ProductSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), productPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `product`.* FROM `product` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, productPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ProductSlice")
	}

	*o = slice

	return nil
}

// ProductExists checks if the Product row exists.
func ProductExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `product` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if product exists")
	}

	return exists, nil
}

// Exists checks if the Product row exists.
func (o *Product) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return ProductExists(ctx, exec, o.ID)
}