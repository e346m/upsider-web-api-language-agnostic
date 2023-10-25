// Code generated by SQLBoiler 4.13.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

// Client is an object representing the database table.
type Client struct {
	ID             []byte    `boil:"id" json:"id" toml:"id" yaml:"id"`
	OrganizationID []byte    `boil:"organization_id" json:"organization_id" toml:"organization_id" yaml:"organization_id"`
	Name           string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	Representative string    `boil:"representative" json:"representative" toml:"representative" yaml:"representative"`
	PhoneNumber    string    `boil:"phone_number" json:"phone_number" toml:"phone_number" yaml:"phone_number"`
	Address        string    `boil:"address" json:"address" toml:"address" yaml:"address"`
	CreatedAt      time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt      time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *clientR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L clientL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ClientColumns = struct {
	ID             string
	OrganizationID string
	Name           string
	Representative string
	PhoneNumber    string
	Address        string
	CreatedAt      string
	UpdatedAt      string
}{
	ID:             "id",
	OrganizationID: "organization_id",
	Name:           "name",
	Representative: "representative",
	PhoneNumber:    "phone_number",
	Address:        "address",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
}

var ClientTableColumns = struct {
	ID             string
	OrganizationID string
	Name           string
	Representative string
	PhoneNumber    string
	Address        string
	CreatedAt      string
	UpdatedAt      string
}{
	ID:             "clients.id",
	OrganizationID: "clients.organization_id",
	Name:           "clients.name",
	Representative: "clients.representative",
	PhoneNumber:    "clients.phone_number",
	Address:        "clients.address",
	CreatedAt:      "clients.created_at",
	UpdatedAt:      "clients.updated_at",
}

// Generated where

var ClientWhere = struct {
	ID             whereHelper__byte
	OrganizationID whereHelper__byte
	Name           whereHelperstring
	Representative whereHelperstring
	PhoneNumber    whereHelperstring
	Address        whereHelperstring
	CreatedAt      whereHelpertime_Time
	UpdatedAt      whereHelpertime_Time
}{
	ID:             whereHelper__byte{field: "\"clients\".\"id\""},
	OrganizationID: whereHelper__byte{field: "\"clients\".\"organization_id\""},
	Name:           whereHelperstring{field: "\"clients\".\"name\""},
	Representative: whereHelperstring{field: "\"clients\".\"representative\""},
	PhoneNumber:    whereHelperstring{field: "\"clients\".\"phone_number\""},
	Address:        whereHelperstring{field: "\"clients\".\"address\""},
	CreatedAt:      whereHelpertime_Time{field: "\"clients\".\"created_at\""},
	UpdatedAt:      whereHelpertime_Time{field: "\"clients\".\"updated_at\""},
}

// ClientRels is where relationship names are stored.
var ClientRels = struct {
	Organization string
	Invoices     string
}{
	Organization: "Organization",
	Invoices:     "Invoices",
}

// clientR is where relationships are stored.
type clientR struct {
	Organization *Organization `boil:"Organization" json:"Organization" toml:"Organization" yaml:"Organization"`
	Invoices     InvoiceSlice  `boil:"Invoices" json:"Invoices" toml:"Invoices" yaml:"Invoices"`
}

// NewStruct creates a new relationship struct
func (*clientR) NewStruct() *clientR {
	return &clientR{}
}

func (r *clientR) GetOrganization() *Organization {
	if r == nil {
		return nil
	}
	return r.Organization
}

func (r *clientR) GetInvoices() InvoiceSlice {
	if r == nil {
		return nil
	}
	return r.Invoices
}

// clientL is where Load methods for each relationship are stored.
type clientL struct{}

var (
	clientAllColumns            = []string{"id", "organization_id", "name", "representative", "phone_number", "address", "created_at", "updated_at"}
	clientColumnsWithoutDefault = []string{"id", "organization_id", "name", "representative", "phone_number", "address", "created_at", "updated_at"}
	clientColumnsWithDefault    = []string{}
	clientPrimaryKeyColumns     = []string{"id"}
	clientGeneratedColumns      = []string{}
)

type (
	// ClientSlice is an alias for a slice of pointers to Client.
	// This should almost always be used instead of []Client.
	ClientSlice []*Client
	// ClientHook is the signature for custom Client hook methods
	ClientHook func(context.Context, boil.ContextExecutor, *Client) error

	clientQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	clientType                 = reflect.TypeOf(&Client{})
	clientMapping              = queries.MakeStructMapping(clientType)
	clientPrimaryKeyMapping, _ = queries.BindMapping(clientType, clientMapping, clientPrimaryKeyColumns)
	clientInsertCacheMut       sync.RWMutex
	clientInsertCache          = make(map[string]insertCache)
	clientUpdateCacheMut       sync.RWMutex
	clientUpdateCache          = make(map[string]updateCache)
	clientUpsertCacheMut       sync.RWMutex
	clientUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var clientAfterSelectHooks []ClientHook

var clientBeforeInsertHooks []ClientHook
var clientAfterInsertHooks []ClientHook

var clientBeforeUpdateHooks []ClientHook
var clientAfterUpdateHooks []ClientHook

var clientBeforeDeleteHooks []ClientHook
var clientAfterDeleteHooks []ClientHook

var clientBeforeUpsertHooks []ClientHook
var clientAfterUpsertHooks []ClientHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Client) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clientAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Client) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clientBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Client) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clientAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Client) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clientBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Client) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clientAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Client) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clientBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Client) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clientAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Client) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clientBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Client) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clientAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddClientHook registers your hook function for all future operations.
func AddClientHook(hookPoint boil.HookPoint, clientHook ClientHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		clientAfterSelectHooks = append(clientAfterSelectHooks, clientHook)
	case boil.BeforeInsertHook:
		clientBeforeInsertHooks = append(clientBeforeInsertHooks, clientHook)
	case boil.AfterInsertHook:
		clientAfterInsertHooks = append(clientAfterInsertHooks, clientHook)
	case boil.BeforeUpdateHook:
		clientBeforeUpdateHooks = append(clientBeforeUpdateHooks, clientHook)
	case boil.AfterUpdateHook:
		clientAfterUpdateHooks = append(clientAfterUpdateHooks, clientHook)
	case boil.BeforeDeleteHook:
		clientBeforeDeleteHooks = append(clientBeforeDeleteHooks, clientHook)
	case boil.AfterDeleteHook:
		clientAfterDeleteHooks = append(clientAfterDeleteHooks, clientHook)
	case boil.BeforeUpsertHook:
		clientBeforeUpsertHooks = append(clientBeforeUpsertHooks, clientHook)
	case boil.AfterUpsertHook:
		clientAfterUpsertHooks = append(clientAfterUpsertHooks, clientHook)
	}
}

// One returns a single client record from the query.
func (q clientQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Client, error) {
	o := &Client{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for clients")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Client records from the query.
func (q clientQuery) All(ctx context.Context, exec boil.ContextExecutor) (ClientSlice, error) {
	var o []*Client

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Client slice")
	}

	if len(clientAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Client records in the query.
func (q clientQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count clients rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q clientQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if clients exists")
	}

	return count > 0, nil
}

// Organization pointed to by the foreign key.
func (o *Client) Organization(mods ...qm.QueryMod) organizationQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.OrganizationID),
	}

	queryMods = append(queryMods, mods...)

	return Organizations(queryMods...)
}

// Invoices retrieves all the invoice's Invoices with an executor.
func (o *Client) Invoices(mods ...qm.QueryMod) invoiceQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"invoices\".\"client_id\"=?", o.ID),
	)

	return Invoices(queryMods...)
}

// LoadOrganization allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (clientL) LoadOrganization(ctx context.Context, e boil.ContextExecutor, singular bool, maybeClient interface{}, mods queries.Applicator) error {
	var slice []*Client
	var object *Client

	if singular {
		var ok bool
		object, ok = maybeClient.(*Client)
		if !ok {
			object = new(Client)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeClient)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeClient))
			}
		}
	} else {
		s, ok := maybeClient.(*[]*Client)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeClient)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeClient))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &clientR{}
		}
		if !queries.IsNil(object.OrganizationID) {
			args = append(args, object.OrganizationID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &clientR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.OrganizationID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.OrganizationID) {
				args = append(args, obj.OrganizationID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`organizations`),
		qm.WhereIn(`organizations.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Organization")
	}

	var resultSlice []*Organization
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Organization")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for organizations")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for organizations")
	}

	if len(clientAfterSelectHooks) != 0 {
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
		object.R.Organization = foreign
		if foreign.R == nil {
			foreign.R = &organizationR{}
		}
		foreign.R.Clients = append(foreign.R.Clients, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.OrganizationID, foreign.ID) {
				local.R.Organization = foreign
				if foreign.R == nil {
					foreign.R = &organizationR{}
				}
				foreign.R.Clients = append(foreign.R.Clients, local)
				break
			}
		}
	}

	return nil
}

// LoadInvoices allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (clientL) LoadInvoices(ctx context.Context, e boil.ContextExecutor, singular bool, maybeClient interface{}, mods queries.Applicator) error {
	var slice []*Client
	var object *Client

	if singular {
		var ok bool
		object, ok = maybeClient.(*Client)
		if !ok {
			object = new(Client)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeClient)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeClient))
			}
		}
	} else {
		s, ok := maybeClient.(*[]*Client)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeClient)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeClient))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &clientR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &clientR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ID) {
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
		qm.From(`invoices`),
		qm.WhereIn(`invoices.client_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load invoices")
	}

	var resultSlice []*Invoice
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice invoices")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on invoices")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for invoices")
	}

	if len(invoiceAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Invoices = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &invoiceR{}
			}
			foreign.R.Client = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.ID, foreign.ClientID) {
				local.R.Invoices = append(local.R.Invoices, foreign)
				if foreign.R == nil {
					foreign.R = &invoiceR{}
				}
				foreign.R.Client = local
				break
			}
		}
	}

	return nil
}

// SetOrganization of the client to the related item.
// Sets o.R.Organization to related.
// Adds o to related.R.Clients.
func (o *Client) SetOrganization(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Organization) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"clients\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"organization_id"}),
		strmangle.WhereClause("\"", "\"", 2, clientPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.OrganizationID, related.ID)
	if o.R == nil {
		o.R = &clientR{
			Organization: related,
		}
	} else {
		o.R.Organization = related
	}

	if related.R == nil {
		related.R = &organizationR{
			Clients: ClientSlice{o},
		}
	} else {
		related.R.Clients = append(related.R.Clients, o)
	}

	return nil
}

// AddInvoices adds the given related objects to the existing relationships
// of the client, optionally inserting them as new records.
// Appends related to o.R.Invoices.
// Sets related.R.Client appropriately.
func (o *Client) AddInvoices(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Invoice) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.ClientID, o.ID)
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"invoices\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"client_id"}),
				strmangle.WhereClause("\"", "\"", 2, invoicePrimaryKeyColumns),
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

			queries.Assign(&rel.ClientID, o.ID)
		}
	}

	if o.R == nil {
		o.R = &clientR{
			Invoices: related,
		}
	} else {
		o.R.Invoices = append(o.R.Invoices, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &invoiceR{
				Client: o,
			}
		} else {
			rel.R.Client = o
		}
	}
	return nil
}

// Clients retrieves all the records using an executor.
func Clients(mods ...qm.QueryMod) clientQuery {
	mods = append(mods, qm.From("\"clients\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"clients\".*"})
	}

	return clientQuery{q}
}

// FindClient retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindClient(ctx context.Context, exec boil.ContextExecutor, iD []byte, selectCols ...string) (*Client, error) {
	clientObj := &Client{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"clients\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, clientObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from clients")
	}

	if err = clientObj.doAfterSelectHooks(ctx, exec); err != nil {
		return clientObj, err
	}

	return clientObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Client) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no clients provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(clientColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	clientInsertCacheMut.RLock()
	cache, cached := clientInsertCache[key]
	clientInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			clientAllColumns,
			clientColumnsWithDefault,
			clientColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(clientType, clientMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(clientType, clientMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"clients\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"clients\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into clients")
	}

	if !cached {
		clientInsertCacheMut.Lock()
		clientInsertCache[key] = cache
		clientInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Client.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Client) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	clientUpdateCacheMut.RLock()
	cache, cached := clientUpdateCache[key]
	clientUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			clientAllColumns,
			clientPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update clients, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"clients\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, clientPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(clientType, clientMapping, append(wl, clientPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update clients row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for clients")
	}

	if !cached {
		clientUpdateCacheMut.Lock()
		clientUpdateCache[key] = cache
		clientUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q clientQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for clients")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for clients")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ClientSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), clientPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"clients\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, clientPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in client slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all client")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Client) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no clients provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(clientColumnsWithDefault, o)

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

	clientUpsertCacheMut.RLock()
	cache, cached := clientUpsertCache[key]
	clientUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			clientAllColumns,
			clientColumnsWithDefault,
			clientColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			clientAllColumns,
			clientPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert clients, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(clientPrimaryKeyColumns))
			copy(conflict, clientPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"clients\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(clientType, clientMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(clientType, clientMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert clients")
	}

	if !cached {
		clientUpsertCacheMut.Lock()
		clientUpsertCache[key] = cache
		clientUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Client record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Client) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Client provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), clientPrimaryKeyMapping)
	sql := "DELETE FROM \"clients\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from clients")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for clients")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q clientQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no clientQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from clients")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for clients")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ClientSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(clientBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), clientPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"clients\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, clientPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from client slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for clients")
	}

	if len(clientAfterDeleteHooks) != 0 {
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
func (o *Client) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindClient(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ClientSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ClientSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), clientPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"clients\".* FROM \"clients\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, clientPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ClientSlice")
	}

	*o = slice

	return nil
}

// ClientExists checks if the Client row exists.
func ClientExists(ctx context.Context, exec boil.ContextExecutor, iD []byte) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"clients\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if clients exists")
	}

	return exists, nil
}
