// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/wiisportsresort/chat/server/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// AvatarURL holds the value of the "avatar_url" field.
	AvatarURL *string `json:"avatar_url,omitempty"`
	// SecretHash holds the value of the "secret_hash" field.
	SecretHash *string `json:"-"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges UserEdges `json:"edges"`
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Messages holds the value of the messages edge.
	Messages []*Message `json:"messages,omitempty"`
	// PrivateMessages holds the value of the private_messages edge.
	PrivateMessages []*PrivateMessage `json:"private_messages,omitempty"`
	// PrivateMessagesReceived holds the value of the private_messages_received edge.
	PrivateMessagesReceived []*PrivateMessage `json:"private_messages_received,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// MessagesOrErr returns the Messages value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) MessagesOrErr() ([]*Message, error) {
	if e.loadedTypes[0] {
		return e.Messages, nil
	}
	return nil, &NotLoadedError{edge: "messages"}
}

// PrivateMessagesOrErr returns the PrivateMessages value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) PrivateMessagesOrErr() ([]*PrivateMessage, error) {
	if e.loadedTypes[1] {
		return e.PrivateMessages, nil
	}
	return nil, &NotLoadedError{edge: "private_messages"}
}

// PrivateMessagesReceivedOrErr returns the PrivateMessagesReceived value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) PrivateMessagesReceivedOrErr() ([]*PrivateMessage, error) {
	if e.loadedTypes[2] {
		return e.PrivateMessagesReceived, nil
	}
	return nil, &NotLoadedError{edge: "private_messages_received"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldID, user.FieldUsername, user.FieldAvatarURL, user.FieldSecretHash:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type User", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				u.ID = value.String
			}
		case user.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				u.Username = value.String
			}
		case user.FieldAvatarURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field avatar_url", values[i])
			} else if value.Valid {
				u.AvatarURL = new(string)
				*u.AvatarURL = value.String
			}
		case user.FieldSecretHash:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field secret_hash", values[i])
			} else if value.Valid {
				u.SecretHash = new(string)
				*u.SecretHash = value.String
			}
		}
	}
	return nil
}

// QueryMessages queries the "messages" edge of the User entity.
func (u *User) QueryMessages() *MessageQuery {
	return (&UserClient{config: u.config}).QueryMessages(u)
}

// QueryPrivateMessages queries the "private_messages" edge of the User entity.
func (u *User) QueryPrivateMessages() *PrivateMessageQuery {
	return (&UserClient{config: u.config}).QueryPrivateMessages(u)
}

// QueryPrivateMessagesReceived queries the "private_messages_received" edge of the User entity.
func (u *User) QueryPrivateMessagesReceived() *PrivateMessageQuery {
	return (&UserClient{config: u.config}).QueryPrivateMessagesReceived(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", username=")
	builder.WriteString(u.Username)
	if v := u.AvatarURL; v != nil {
		builder.WriteString(", avatar_url=")
		builder.WriteString(*v)
	}
	builder.WriteString(", secret_hash=<sensitive>")
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}