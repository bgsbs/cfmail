// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"tmail/ent/attachment"
	"tmail/ent/envelope"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Attachment is the model entity for the Attachment schema.
type Attachment struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Filename holds the value of the "filename" field.
	Filename string `json:"filename,omitempty"`
	// Filepath holds the value of the "filepath" field.
	Filepath string `json:"filepath,omitempty"`
	// ContentType holds the value of the "contentType" field.
	ContentType string `json:"contentType,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AttachmentQuery when eager-loading is set.
	Edges                AttachmentEdges `json:"edges"`
	envelope_attachments *int
	selectValues         sql.SelectValues
}

// AttachmentEdges holds the relations/edges for other nodes in the graph.
type AttachmentEdges struct {
	// Owner holds the value of the owner edge.
	Owner *Envelope `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AttachmentEdges) OwnerOrErr() (*Envelope, error) {
	if e.Owner != nil {
		return e.Owner, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: envelope.Label}
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Attachment) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case attachment.FieldID, attachment.FieldFilename, attachment.FieldFilepath, attachment.FieldContentType:
			values[i] = new(sql.NullString)
		case attachment.ForeignKeys[0]: // envelope_attachments
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Attachment fields.
func (a *Attachment) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case attachment.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				a.ID = value.String
			}
		case attachment.FieldFilename:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field filename", values[i])
			} else if value.Valid {
				a.Filename = value.String
			}
		case attachment.FieldFilepath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field filepath", values[i])
			} else if value.Valid {
				a.Filepath = value.String
			}
		case attachment.FieldContentType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field contentType", values[i])
			} else if value.Valid {
				a.ContentType = value.String
			}
		case attachment.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field envelope_attachments", value)
			} else if value.Valid {
				a.envelope_attachments = new(int)
				*a.envelope_attachments = int(value.Int64)
			}
		default:
			a.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Attachment.
// This includes values selected through modifiers, order, etc.
func (a *Attachment) Value(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// QueryOwner queries the "owner" edge of the Attachment entity.
func (a *Attachment) QueryOwner() *EnvelopeQuery {
	return NewAttachmentClient(a.config).QueryOwner(a)
}

// Update returns a builder for updating this Attachment.
// Note that you need to call Attachment.Unwrap() before calling this method if this Attachment
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Attachment) Update() *AttachmentUpdateOne {
	return NewAttachmentClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Attachment entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Attachment) Unwrap() *Attachment {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Attachment is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Attachment) String() string {
	var builder strings.Builder
	builder.WriteString("Attachment(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("filename=")
	builder.WriteString(a.Filename)
	builder.WriteString(", ")
	builder.WriteString("filepath=")
	builder.WriteString(a.Filepath)
	builder.WriteString(", ")
	builder.WriteString("contentType=")
	builder.WriteString(a.ContentType)
	builder.WriteByte(')')
	return builder.String()
}

// Attachments is a parsable slice of Attachment.
type Attachments []*Attachment
