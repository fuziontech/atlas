// Copyright 2021-present The Atlas Authors. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"ariga.io/atlas/cmd/atlas/internal/migrate/ent/migrate"

	"ariga.io/atlas/cmd/atlas/internal/migrate/ent/attempt"
	"ariga.io/atlas/cmd/atlas/internal/migrate/ent/revision"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Attempt is the client for interacting with the Attempt builders.
	Attempt *AttemptClient
	// Revision is the client for interacting with the Revision builders.
	Revision *RevisionClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Attempt = NewAttemptClient(c.config)
	c.Revision = NewRevisionClient(c.config)
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
		ctx:      ctx,
		config:   cfg,
		Attempt:  NewAttemptClient(cfg),
		Revision: NewRevisionClient(cfg),
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
		ctx:      ctx,
		config:   cfg,
		Attempt:  NewAttemptClient(cfg),
		Revision: NewRevisionClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Attempt.
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
	c.Attempt.Use(hooks...)
	c.Revision.Use(hooks...)
}

// AttemptClient is a client for the Attempt schema.
type AttemptClient struct {
	config
}

// NewAttemptClient returns a client for the Attempt from the given config.
func NewAttemptClient(c config) *AttemptClient {
	return &AttemptClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `attempt.Hooks(f(g(h())))`.
func (c *AttemptClient) Use(hooks ...Hook) {
	c.hooks.Attempt = append(c.hooks.Attempt, hooks...)
}

// Create returns a builder for creating a Attempt entity.
func (c *AttemptClient) Create() *AttemptCreate {
	mutation := newAttemptMutation(c.config, OpCreate)
	return &AttemptCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Attempt entities.
func (c *AttemptClient) CreateBulk(builders ...*AttemptCreate) *AttemptCreateBulk {
	return &AttemptCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Attempt.
func (c *AttemptClient) Update() *AttemptUpdate {
	mutation := newAttemptMutation(c.config, OpUpdate)
	return &AttemptUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AttemptClient) UpdateOne(a *Attempt) *AttemptUpdateOne {
	mutation := newAttemptMutation(c.config, OpUpdateOne, withAttempt(a))
	return &AttemptUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AttemptClient) UpdateOneID(id int) *AttemptUpdateOne {
	mutation := newAttemptMutation(c.config, OpUpdateOne, withAttemptID(id))
	return &AttemptUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Attempt.
func (c *AttemptClient) Delete() *AttemptDelete {
	mutation := newAttemptMutation(c.config, OpDelete)
	return &AttemptDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *AttemptClient) DeleteOne(a *Attempt) *AttemptDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *AttemptClient) DeleteOneID(id int) *AttemptDeleteOne {
	builder := c.Delete().Where(attempt.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AttemptDeleteOne{builder}
}

// Query returns a query builder for Attempt.
func (c *AttemptClient) Query() *AttemptQuery {
	return &AttemptQuery{
		config: c.config,
	}
}

// Get returns a Attempt entity by its id.
func (c *AttemptClient) Get(ctx context.Context, id int) (*Attempt, error) {
	return c.Query().Where(attempt.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AttemptClient) GetX(ctx context.Context, id int) *Attempt {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *AttemptClient) Hooks() []Hook {
	return c.hooks.Attempt
}

// RevisionClient is a client for the Revision schema.
type RevisionClient struct {
	config
}

// NewRevisionClient returns a client for the Revision from the given config.
func NewRevisionClient(c config) *RevisionClient {
	return &RevisionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `revision.Hooks(f(g(h())))`.
func (c *RevisionClient) Use(hooks ...Hook) {
	c.hooks.Revision = append(c.hooks.Revision, hooks...)
}

// Create returns a builder for creating a Revision entity.
func (c *RevisionClient) Create() *RevisionCreate {
	mutation := newRevisionMutation(c.config, OpCreate)
	return &RevisionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Revision entities.
func (c *RevisionClient) CreateBulk(builders ...*RevisionCreate) *RevisionCreateBulk {
	return &RevisionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Revision.
func (c *RevisionClient) Update() *RevisionUpdate {
	mutation := newRevisionMutation(c.config, OpUpdate)
	return &RevisionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *RevisionClient) UpdateOne(r *Revision) *RevisionUpdateOne {
	mutation := newRevisionMutation(c.config, OpUpdateOne, withRevision(r))
	return &RevisionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *RevisionClient) UpdateOneID(id string) *RevisionUpdateOne {
	mutation := newRevisionMutation(c.config, OpUpdateOne, withRevisionID(id))
	return &RevisionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Revision.
func (c *RevisionClient) Delete() *RevisionDelete {
	mutation := newRevisionMutation(c.config, OpDelete)
	return &RevisionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *RevisionClient) DeleteOne(r *Revision) *RevisionDeleteOne {
	return c.DeleteOneID(r.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *RevisionClient) DeleteOneID(id string) *RevisionDeleteOne {
	builder := c.Delete().Where(revision.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &RevisionDeleteOne{builder}
}

// Query returns a query builder for Revision.
func (c *RevisionClient) Query() *RevisionQuery {
	return &RevisionQuery{
		config: c.config,
	}
}

// Get returns a Revision entity by its id.
func (c *RevisionClient) Get(ctx context.Context, id string) (*Revision, error) {
	return c.Query().Where(revision.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *RevisionClient) GetX(ctx context.Context, id string) *Revision {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *RevisionClient) Hooks() []Hook {
	return c.hooks.Revision
}
