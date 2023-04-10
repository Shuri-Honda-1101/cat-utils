// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/Shuri-Honda-1101/cat-utils/ent/migrate"
	"github.com/google/uuid"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/Shuri-Honda-1101/cat-utils/ent/cat"
	"github.com/Shuri-Honda-1101/cat-utils/ent/toilet"
	"github.com/Shuri-Honda-1101/cat-utils/ent/user"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Cat is the client for interacting with the Cat builders.
	Cat *CatClient
	// Toilet is the client for interacting with the Toilet builders.
	Toilet *ToiletClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Cat = NewCatClient(c.config)
	c.Toilet = NewToiletClient(c.config)
	c.User = NewUserClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:    ctx,
		config: cfg,
		Cat:    NewCatClient(cfg),
		Toilet: NewToiletClient(cfg),
		User:   NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:    ctx,
		config: cfg,
		Cat:    NewCatClient(cfg),
		Toilet: NewToiletClient(cfg),
		User:   NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Cat.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Cat.Use(hooks...)
	c.Toilet.Use(hooks...)
	c.User.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Cat.Intercept(interceptors...)
	c.Toilet.Intercept(interceptors...)
	c.User.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *CatMutation:
		return c.Cat.mutate(ctx, m)
	case *ToiletMutation:
		return c.Toilet.mutate(ctx, m)
	case *UserMutation:
		return c.User.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// CatClient is a client for the Cat schema.
type CatClient struct {
	config
}

// NewCatClient returns a client for the Cat from the given config.
func NewCatClient(c config) *CatClient {
	return &CatClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `cat.Hooks(f(g(h())))`.
func (c *CatClient) Use(hooks ...Hook) {
	c.hooks.Cat = append(c.hooks.Cat, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `cat.Intercept(f(g(h())))`.
func (c *CatClient) Intercept(interceptors ...Interceptor) {
	c.inters.Cat = append(c.inters.Cat, interceptors...)
}

// Create returns a builder for creating a Cat entity.
func (c *CatClient) Create() *CatCreate {
	mutation := newCatMutation(c.config, OpCreate)
	return &CatCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Cat entities.
func (c *CatClient) CreateBulk(builders ...*CatCreate) *CatCreateBulk {
	return &CatCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Cat.
func (c *CatClient) Update() *CatUpdate {
	mutation := newCatMutation(c.config, OpUpdate)
	return &CatUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CatClient) UpdateOne(ca *Cat) *CatUpdateOne {
	mutation := newCatMutation(c.config, OpUpdateOne, withCat(ca))
	return &CatUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CatClient) UpdateOneID(id uuid.UUID) *CatUpdateOne {
	mutation := newCatMutation(c.config, OpUpdateOne, withCatID(id))
	return &CatUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Cat.
func (c *CatClient) Delete() *CatDelete {
	mutation := newCatMutation(c.config, OpDelete)
	return &CatDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CatClient) DeleteOne(ca *Cat) *CatDeleteOne {
	return c.DeleteOneID(ca.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *CatClient) DeleteOneID(id uuid.UUID) *CatDeleteOne {
	builder := c.Delete().Where(cat.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CatDeleteOne{builder}
}

// Query returns a query builder for Cat.
func (c *CatClient) Query() *CatQuery {
	return &CatQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeCat},
		inters: c.Interceptors(),
	}
}

// Get returns a Cat entity by its id.
func (c *CatClient) Get(ctx context.Context, id uuid.UUID) (*Cat, error) {
	return c.Query().Where(cat.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CatClient) GetX(ctx context.Context, id uuid.UUID) *Cat {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOwner queries the owner edge of a Cat.
func (c *CatClient) QueryOwner(ca *Cat) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ca.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(cat.Table, cat.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, cat.OwnerTable, cat.OwnerColumn),
		)
		fromV = sqlgraph.Neighbors(ca.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryToilets queries the toilets edge of a Cat.
func (c *CatClient) QueryToilets(ca *Cat) *ToiletQuery {
	query := (&ToiletClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ca.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(cat.Table, cat.FieldID, id),
			sqlgraph.To(toilet.Table, toilet.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, cat.ToiletsTable, cat.ToiletsColumn),
		)
		fromV = sqlgraph.Neighbors(ca.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CatClient) Hooks() []Hook {
	return c.hooks.Cat
}

// Interceptors returns the client interceptors.
func (c *CatClient) Interceptors() []Interceptor {
	return c.inters.Cat
}

func (c *CatClient) mutate(ctx context.Context, m *CatMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&CatCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&CatUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&CatUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&CatDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Cat mutation op: %q", m.Op())
	}
}

// ToiletClient is a client for the Toilet schema.
type ToiletClient struct {
	config
}

// NewToiletClient returns a client for the Toilet from the given config.
func NewToiletClient(c config) *ToiletClient {
	return &ToiletClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `toilet.Hooks(f(g(h())))`.
func (c *ToiletClient) Use(hooks ...Hook) {
	c.hooks.Toilet = append(c.hooks.Toilet, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `toilet.Intercept(f(g(h())))`.
func (c *ToiletClient) Intercept(interceptors ...Interceptor) {
	c.inters.Toilet = append(c.inters.Toilet, interceptors...)
}

// Create returns a builder for creating a Toilet entity.
func (c *ToiletClient) Create() *ToiletCreate {
	mutation := newToiletMutation(c.config, OpCreate)
	return &ToiletCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Toilet entities.
func (c *ToiletClient) CreateBulk(builders ...*ToiletCreate) *ToiletCreateBulk {
	return &ToiletCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Toilet.
func (c *ToiletClient) Update() *ToiletUpdate {
	mutation := newToiletMutation(c.config, OpUpdate)
	return &ToiletUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ToiletClient) UpdateOne(t *Toilet) *ToiletUpdateOne {
	mutation := newToiletMutation(c.config, OpUpdateOne, withToilet(t))
	return &ToiletUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ToiletClient) UpdateOneID(id int) *ToiletUpdateOne {
	mutation := newToiletMutation(c.config, OpUpdateOne, withToiletID(id))
	return &ToiletUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Toilet.
func (c *ToiletClient) Delete() *ToiletDelete {
	mutation := newToiletMutation(c.config, OpDelete)
	return &ToiletDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ToiletClient) DeleteOne(t *Toilet) *ToiletDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ToiletClient) DeleteOneID(id int) *ToiletDeleteOne {
	builder := c.Delete().Where(toilet.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ToiletDeleteOne{builder}
}

// Query returns a query builder for Toilet.
func (c *ToiletClient) Query() *ToiletQuery {
	return &ToiletQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeToilet},
		inters: c.Interceptors(),
	}
}

// Get returns a Toilet entity by its id.
func (c *ToiletClient) Get(ctx context.Context, id int) (*Toilet, error) {
	return c.Query().Where(toilet.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ToiletClient) GetX(ctx context.Context, id int) *Toilet {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryCat queries the cat edge of a Toilet.
func (c *ToiletClient) QueryCat(t *Toilet) *CatQuery {
	query := (&CatClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(toilet.Table, toilet.FieldID, id),
			sqlgraph.To(cat.Table, cat.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, toilet.CatTable, toilet.CatColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ToiletClient) Hooks() []Hook {
	return c.hooks.Toilet
}

// Interceptors returns the client interceptors.
func (c *ToiletClient) Interceptors() []Interceptor {
	return c.inters.Toilet
}

func (c *ToiletClient) mutate(ctx context.Context, m *ToiletMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ToiletCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ToiletUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ToiletUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ToiletDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Toilet mutation op: %q", m.Op())
	}
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `user.Intercept(f(g(h())))`.
func (c *UserClient) Intercept(interceptors ...Interceptor) {
	c.inters.User = append(c.inters.User, interceptors...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id uuid.UUID) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id uuid.UUID) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeUser},
		inters: c.Interceptors(),
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id uuid.UUID) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id uuid.UUID) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryCats queries the cats edge of a User.
func (c *UserClient) QueryCats(u *User) *CatQuery {
	query := (&CatClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(cat.Table, cat.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.CatsTable, user.CatsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// Interceptors returns the client interceptors.
func (c *UserClient) Interceptors() []Interceptor {
	return c.inters.User
}

func (c *UserClient) mutate(ctx context.Context, m *UserMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UserCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UserUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UserDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown User mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Cat, Toilet, User []ent.Hook
	}
	inters struct {
		Cat, Toilet, User []ent.Interceptor
	}
)
