// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/wiisportsresort/chat/server/ent/migrate"

	"github.com/wiisportsresort/chat/server/ent/message"
	"github.com/wiisportsresort/chat/server/ent/privatemessage"
	"github.com/wiisportsresort/chat/server/ent/user"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Message is the client for interacting with the Message builders.
	Message *MessageClient
	// PrivateMessage is the client for interacting with the PrivateMessage builders.
	PrivateMessage *PrivateMessageClient
	// User is the client for interacting with the User builders.
	User *UserClient
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
	c.Message = NewMessageClient(c.config)
	c.PrivateMessage = NewPrivateMessageClient(c.config)
	c.User = NewUserClient(c.config)
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
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:            ctx,
		config:         cfg,
		Message:        NewMessageClient(cfg),
		PrivateMessage: NewPrivateMessageClient(cfg),
		User:           NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
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
		ctx:            ctx,
		config:         cfg,
		Message:        NewMessageClient(cfg),
		PrivateMessage: NewPrivateMessageClient(cfg),
		User:           NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Message.
//		Query().
//		Count(ctx)
//
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
	c.Message.Use(hooks...)
	c.PrivateMessage.Use(hooks...)
	c.User.Use(hooks...)
}

// MessageClient is a client for the Message schema.
type MessageClient struct {
	config
}

// NewMessageClient returns a client for the Message from the given config.
func NewMessageClient(c config) *MessageClient {
	return &MessageClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `message.Hooks(f(g(h())))`.
func (c *MessageClient) Use(hooks ...Hook) {
	c.hooks.Message = append(c.hooks.Message, hooks...)
}

// Create returns a create builder for Message.
func (c *MessageClient) Create() *MessageCreate {
	mutation := newMessageMutation(c.config, OpCreate)
	return &MessageCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Message entities.
func (c *MessageClient) CreateBulk(builders ...*MessageCreate) *MessageCreateBulk {
	return &MessageCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Message.
func (c *MessageClient) Update() *MessageUpdate {
	mutation := newMessageMutation(c.config, OpUpdate)
	return &MessageUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MessageClient) UpdateOne(m *Message) *MessageUpdateOne {
	mutation := newMessageMutation(c.config, OpUpdateOne, withMessage(m))
	return &MessageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MessageClient) UpdateOneID(id string) *MessageUpdateOne {
	mutation := newMessageMutation(c.config, OpUpdateOne, withMessageID(id))
	return &MessageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Message.
func (c *MessageClient) Delete() *MessageDelete {
	mutation := newMessageMutation(c.config, OpDelete)
	return &MessageDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *MessageClient) DeleteOne(m *Message) *MessageDeleteOne {
	return c.DeleteOneID(m.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *MessageClient) DeleteOneID(id string) *MessageDeleteOne {
	builder := c.Delete().Where(message.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MessageDeleteOne{builder}
}

// Query returns a query builder for Message.
func (c *MessageClient) Query() *MessageQuery {
	return &MessageQuery{
		config: c.config,
	}
}

// Get returns a Message entity by its id.
func (c *MessageClient) Get(ctx context.Context, id string) (*Message, error) {
	return c.Query().Where(message.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MessageClient) GetX(ctx context.Context, id string) *Message {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryAuthor queries the author edge of a Message.
func (c *MessageClient) QueryAuthor(m *Message) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(message.Table, message.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, message.AuthorTable, message.AuthorColumn),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryParent queries the parent edge of a Message.
func (c *MessageClient) QueryParent(m *Message) *MessageQuery {
	query := &MessageQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(message.Table, message.FieldID, id),
			sqlgraph.To(message.Table, message.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, message.ParentTable, message.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryChildren queries the children edge of a Message.
func (c *MessageClient) QueryChildren(m *Message) *MessageQuery {
	query := &MessageQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(message.Table, message.FieldID, id),
			sqlgraph.To(message.Table, message.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, message.ChildrenTable, message.ChildrenColumn),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *MessageClient) Hooks() []Hook {
	return c.hooks.Message
}

// PrivateMessageClient is a client for the PrivateMessage schema.
type PrivateMessageClient struct {
	config
}

// NewPrivateMessageClient returns a client for the PrivateMessage from the given config.
func NewPrivateMessageClient(c config) *PrivateMessageClient {
	return &PrivateMessageClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `privatemessage.Hooks(f(g(h())))`.
func (c *PrivateMessageClient) Use(hooks ...Hook) {
	c.hooks.PrivateMessage = append(c.hooks.PrivateMessage, hooks...)
}

// Create returns a create builder for PrivateMessage.
func (c *PrivateMessageClient) Create() *PrivateMessageCreate {
	mutation := newPrivateMessageMutation(c.config, OpCreate)
	return &PrivateMessageCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of PrivateMessage entities.
func (c *PrivateMessageClient) CreateBulk(builders ...*PrivateMessageCreate) *PrivateMessageCreateBulk {
	return &PrivateMessageCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for PrivateMessage.
func (c *PrivateMessageClient) Update() *PrivateMessageUpdate {
	mutation := newPrivateMessageMutation(c.config, OpUpdate)
	return &PrivateMessageUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PrivateMessageClient) UpdateOne(pm *PrivateMessage) *PrivateMessageUpdateOne {
	mutation := newPrivateMessageMutation(c.config, OpUpdateOne, withPrivateMessage(pm))
	return &PrivateMessageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PrivateMessageClient) UpdateOneID(id string) *PrivateMessageUpdateOne {
	mutation := newPrivateMessageMutation(c.config, OpUpdateOne, withPrivateMessageID(id))
	return &PrivateMessageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for PrivateMessage.
func (c *PrivateMessageClient) Delete() *PrivateMessageDelete {
	mutation := newPrivateMessageMutation(c.config, OpDelete)
	return &PrivateMessageDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *PrivateMessageClient) DeleteOne(pm *PrivateMessage) *PrivateMessageDeleteOne {
	return c.DeleteOneID(pm.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *PrivateMessageClient) DeleteOneID(id string) *PrivateMessageDeleteOne {
	builder := c.Delete().Where(privatemessage.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PrivateMessageDeleteOne{builder}
}

// Query returns a query builder for PrivateMessage.
func (c *PrivateMessageClient) Query() *PrivateMessageQuery {
	return &PrivateMessageQuery{
		config: c.config,
	}
}

// Get returns a PrivateMessage entity by its id.
func (c *PrivateMessageClient) Get(ctx context.Context, id string) (*PrivateMessage, error) {
	return c.Query().Where(privatemessage.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PrivateMessageClient) GetX(ctx context.Context, id string) *PrivateMessage {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryAuthor queries the author edge of a PrivateMessage.
func (c *PrivateMessageClient) QueryAuthor(pm *PrivateMessage) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pm.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(privatemessage.Table, privatemessage.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, privatemessage.AuthorTable, privatemessage.AuthorColumn),
		)
		fromV = sqlgraph.Neighbors(pm.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryRecipient queries the recipient edge of a PrivateMessage.
func (c *PrivateMessageClient) QueryRecipient(pm *PrivateMessage) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pm.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(privatemessage.Table, privatemessage.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, privatemessage.RecipientTable, privatemessage.RecipientColumn),
		)
		fromV = sqlgraph.Neighbors(pm.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryParent queries the parent edge of a PrivateMessage.
func (c *PrivateMessageClient) QueryParent(pm *PrivateMessage) *PrivateMessageQuery {
	query := &PrivateMessageQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pm.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(privatemessage.Table, privatemessage.FieldID, id),
			sqlgraph.To(privatemessage.Table, privatemessage.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, privatemessage.ParentTable, privatemessage.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(pm.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryChildren queries the children edge of a PrivateMessage.
func (c *PrivateMessageClient) QueryChildren(pm *PrivateMessage) *PrivateMessageQuery {
	query := &PrivateMessageQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pm.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(privatemessage.Table, privatemessage.FieldID, id),
			sqlgraph.To(privatemessage.Table, privatemessage.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, privatemessage.ChildrenTable, privatemessage.ChildrenColumn),
		)
		fromV = sqlgraph.Neighbors(pm.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PrivateMessageClient) Hooks() []Hook {
	return c.hooks.PrivateMessage
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

// Create returns a create builder for User.
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
func (c *UserClient) UpdateOneID(id string) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserClient) DeleteOneID(id string) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id string) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id string) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryMessages queries the messages edge of a User.
func (c *UserClient) QueryMessages(u *User) *MessageQuery {
	query := &MessageQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(message.Table, message.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.MessagesTable, user.MessagesColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryPrivateMessages queries the private_messages edge of a User.
func (c *UserClient) QueryPrivateMessages(u *User) *PrivateMessageQuery {
	query := &PrivateMessageQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(privatemessage.Table, privatemessage.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.PrivateMessagesTable, user.PrivateMessagesColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryPrivateMessagesReceived queries the private_messages_received edge of a User.
func (c *UserClient) QueryPrivateMessagesReceived(u *User) *PrivateMessageQuery {
	query := &PrivateMessageQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(privatemessage.Table, privatemessage.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.PrivateMessagesReceivedTable, user.PrivateMessagesReceivedColumn),
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