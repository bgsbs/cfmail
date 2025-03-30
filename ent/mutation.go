// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
	"tmail/ent/attachment"
	"tmail/ent/envelope"
	"tmail/ent/predicate"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeAttachment = "Attachment"
	TypeEnvelope   = "Envelope"
)

// AttachmentMutation represents an operation that mutates the Attachment nodes in the graph.
type AttachmentMutation struct {
	config
	op            Op
	typ           string
	id            *string
	filename      *string
	filepath      *string
	contentType   *string
	clearedFields map[string]struct{}
	owner         *int
	clearedowner  bool
	done          bool
	oldValue      func(context.Context) (*Attachment, error)
	predicates    []predicate.Attachment
}

var _ ent.Mutation = (*AttachmentMutation)(nil)

// attachmentOption allows management of the mutation configuration using functional options.
type attachmentOption func(*AttachmentMutation)

// newAttachmentMutation creates new mutation for the Attachment entity.
func newAttachmentMutation(c config, op Op, opts ...attachmentOption) *AttachmentMutation {
	m := &AttachmentMutation{
		config:        c,
		op:            op,
		typ:           TypeAttachment,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withAttachmentID sets the ID field of the mutation.
func withAttachmentID(id string) attachmentOption {
	return func(m *AttachmentMutation) {
		var (
			err   error
			once  sync.Once
			value *Attachment
		)
		m.oldValue = func(ctx context.Context) (*Attachment, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Attachment.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withAttachment sets the old Attachment of the mutation.
func withAttachment(node *Attachment) attachmentOption {
	return func(m *AttachmentMutation) {
		m.oldValue = func(context.Context) (*Attachment, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m AttachmentMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m AttachmentMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Attachment entities.
func (m *AttachmentMutation) SetID(id string) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *AttachmentMutation) ID() (id string, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *AttachmentMutation) IDs(ctx context.Context) ([]string, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []string{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Attachment.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetFilename sets the "filename" field.
func (m *AttachmentMutation) SetFilename(s string) {
	m.filename = &s
}

// Filename returns the value of the "filename" field in the mutation.
func (m *AttachmentMutation) Filename() (r string, exists bool) {
	v := m.filename
	if v == nil {
		return
	}
	return *v, true
}

// OldFilename returns the old "filename" field's value of the Attachment entity.
// If the Attachment object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AttachmentMutation) OldFilename(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldFilename is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldFilename requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldFilename: %w", err)
	}
	return oldValue.Filename, nil
}

// ResetFilename resets all changes to the "filename" field.
func (m *AttachmentMutation) ResetFilename() {
	m.filename = nil
}

// SetFilepath sets the "filepath" field.
func (m *AttachmentMutation) SetFilepath(s string) {
	m.filepath = &s
}

// Filepath returns the value of the "filepath" field in the mutation.
func (m *AttachmentMutation) Filepath() (r string, exists bool) {
	v := m.filepath
	if v == nil {
		return
	}
	return *v, true
}

// OldFilepath returns the old "filepath" field's value of the Attachment entity.
// If the Attachment object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AttachmentMutation) OldFilepath(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldFilepath is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldFilepath requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldFilepath: %w", err)
	}
	return oldValue.Filepath, nil
}

// ResetFilepath resets all changes to the "filepath" field.
func (m *AttachmentMutation) ResetFilepath() {
	m.filepath = nil
}

// SetContentType sets the "contentType" field.
func (m *AttachmentMutation) SetContentType(s string) {
	m.contentType = &s
}

// ContentType returns the value of the "contentType" field in the mutation.
func (m *AttachmentMutation) ContentType() (r string, exists bool) {
	v := m.contentType
	if v == nil {
		return
	}
	return *v, true
}

// OldContentType returns the old "contentType" field's value of the Attachment entity.
// If the Attachment object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AttachmentMutation) OldContentType(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldContentType is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldContentType requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldContentType: %w", err)
	}
	return oldValue.ContentType, nil
}

// ResetContentType resets all changes to the "contentType" field.
func (m *AttachmentMutation) ResetContentType() {
	m.contentType = nil
}

// SetOwnerID sets the "owner" edge to the Envelope entity by id.
func (m *AttachmentMutation) SetOwnerID(id int) {
	m.owner = &id
}

// ClearOwner clears the "owner" edge to the Envelope entity.
func (m *AttachmentMutation) ClearOwner() {
	m.clearedowner = true
}

// OwnerCleared reports if the "owner" edge to the Envelope entity was cleared.
func (m *AttachmentMutation) OwnerCleared() bool {
	return m.clearedowner
}

// OwnerID returns the "owner" edge ID in the mutation.
func (m *AttachmentMutation) OwnerID() (id int, exists bool) {
	if m.owner != nil {
		return *m.owner, true
	}
	return
}

// OwnerIDs returns the "owner" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// OwnerID instead. It exists only for internal usage by the builders.
func (m *AttachmentMutation) OwnerIDs() (ids []int) {
	if id := m.owner; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetOwner resets all changes to the "owner" edge.
func (m *AttachmentMutation) ResetOwner() {
	m.owner = nil
	m.clearedowner = false
}

// Where appends a list predicates to the AttachmentMutation builder.
func (m *AttachmentMutation) Where(ps ...predicate.Attachment) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the AttachmentMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *AttachmentMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Attachment, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *AttachmentMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *AttachmentMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Attachment).
func (m *AttachmentMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *AttachmentMutation) Fields() []string {
	fields := make([]string, 0, 3)
	if m.filename != nil {
		fields = append(fields, attachment.FieldFilename)
	}
	if m.filepath != nil {
		fields = append(fields, attachment.FieldFilepath)
	}
	if m.contentType != nil {
		fields = append(fields, attachment.FieldContentType)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *AttachmentMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case attachment.FieldFilename:
		return m.Filename()
	case attachment.FieldFilepath:
		return m.Filepath()
	case attachment.FieldContentType:
		return m.ContentType()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *AttachmentMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case attachment.FieldFilename:
		return m.OldFilename(ctx)
	case attachment.FieldFilepath:
		return m.OldFilepath(ctx)
	case attachment.FieldContentType:
		return m.OldContentType(ctx)
	}
	return nil, fmt.Errorf("unknown Attachment field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *AttachmentMutation) SetField(name string, value ent.Value) error {
	switch name {
	case attachment.FieldFilename:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetFilename(v)
		return nil
	case attachment.FieldFilepath:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetFilepath(v)
		return nil
	case attachment.FieldContentType:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetContentType(v)
		return nil
	}
	return fmt.Errorf("unknown Attachment field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *AttachmentMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *AttachmentMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *AttachmentMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Attachment numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *AttachmentMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *AttachmentMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *AttachmentMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Attachment nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *AttachmentMutation) ResetField(name string) error {
	switch name {
	case attachment.FieldFilename:
		m.ResetFilename()
		return nil
	case attachment.FieldFilepath:
		m.ResetFilepath()
		return nil
	case attachment.FieldContentType:
		m.ResetContentType()
		return nil
	}
	return fmt.Errorf("unknown Attachment field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *AttachmentMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.owner != nil {
		edges = append(edges, attachment.EdgeOwner)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *AttachmentMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case attachment.EdgeOwner:
		if id := m.owner; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *AttachmentMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *AttachmentMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *AttachmentMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedowner {
		edges = append(edges, attachment.EdgeOwner)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *AttachmentMutation) EdgeCleared(name string) bool {
	switch name {
	case attachment.EdgeOwner:
		return m.clearedowner
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *AttachmentMutation) ClearEdge(name string) error {
	switch name {
	case attachment.EdgeOwner:
		m.ClearOwner()
		return nil
	}
	return fmt.Errorf("unknown Attachment unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *AttachmentMutation) ResetEdge(name string) error {
	switch name {
	case attachment.EdgeOwner:
		m.ResetOwner()
		return nil
	}
	return fmt.Errorf("unknown Attachment edge %s", name)
}

// EnvelopeMutation represents an operation that mutates the Envelope nodes in the graph.
type EnvelopeMutation struct {
	config
	op                 Op
	typ                string
	id                 *int
	to                 *string
	from               *string
	subject            *string
	content            *string
	created_at         *time.Time
	clearedFields      map[string]struct{}
	attachments        map[string]struct{}
	removedattachments map[string]struct{}
	clearedattachments bool
	done               bool
	oldValue           func(context.Context) (*Envelope, error)
	predicates         []predicate.Envelope
}

var _ ent.Mutation = (*EnvelopeMutation)(nil)

// envelopeOption allows management of the mutation configuration using functional options.
type envelopeOption func(*EnvelopeMutation)

// newEnvelopeMutation creates new mutation for the Envelope entity.
func newEnvelopeMutation(c config, op Op, opts ...envelopeOption) *EnvelopeMutation {
	m := &EnvelopeMutation{
		config:        c,
		op:            op,
		typ:           TypeEnvelope,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withEnvelopeID sets the ID field of the mutation.
func withEnvelopeID(id int) envelopeOption {
	return func(m *EnvelopeMutation) {
		var (
			err   error
			once  sync.Once
			value *Envelope
		)
		m.oldValue = func(ctx context.Context) (*Envelope, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Envelope.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withEnvelope sets the old Envelope of the mutation.
func withEnvelope(node *Envelope) envelopeOption {
	return func(m *EnvelopeMutation) {
		m.oldValue = func(context.Context) (*Envelope, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m EnvelopeMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m EnvelopeMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *EnvelopeMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *EnvelopeMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Envelope.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetTo sets the "to" field.
func (m *EnvelopeMutation) SetTo(s string) {
	m.to = &s
}

// To returns the value of the "to" field in the mutation.
func (m *EnvelopeMutation) To() (r string, exists bool) {
	v := m.to
	if v == nil {
		return
	}
	return *v, true
}

// OldTo returns the old "to" field's value of the Envelope entity.
// If the Envelope object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EnvelopeMutation) OldTo(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldTo is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldTo requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTo: %w", err)
	}
	return oldValue.To, nil
}

// ResetTo resets all changes to the "to" field.
func (m *EnvelopeMutation) ResetTo() {
	m.to = nil
}

// SetFrom sets the "from" field.
func (m *EnvelopeMutation) SetFrom(s string) {
	m.from = &s
}

// From returns the value of the "from" field in the mutation.
func (m *EnvelopeMutation) From() (r string, exists bool) {
	v := m.from
	if v == nil {
		return
	}
	return *v, true
}

// OldFrom returns the old "from" field's value of the Envelope entity.
// If the Envelope object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EnvelopeMutation) OldFrom(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldFrom is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldFrom requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldFrom: %w", err)
	}
	return oldValue.From, nil
}

// ResetFrom resets all changes to the "from" field.
func (m *EnvelopeMutation) ResetFrom() {
	m.from = nil
}

// SetSubject sets the "subject" field.
func (m *EnvelopeMutation) SetSubject(s string) {
	m.subject = &s
}

// Subject returns the value of the "subject" field in the mutation.
func (m *EnvelopeMutation) Subject() (r string, exists bool) {
	v := m.subject
	if v == nil {
		return
	}
	return *v, true
}

// OldSubject returns the old "subject" field's value of the Envelope entity.
// If the Envelope object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EnvelopeMutation) OldSubject(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldSubject is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldSubject requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldSubject: %w", err)
	}
	return oldValue.Subject, nil
}

// ResetSubject resets all changes to the "subject" field.
func (m *EnvelopeMutation) ResetSubject() {
	m.subject = nil
}

// SetContent sets the "content" field.
func (m *EnvelopeMutation) SetContent(s string) {
	m.content = &s
}

// Content returns the value of the "content" field in the mutation.
func (m *EnvelopeMutation) Content() (r string, exists bool) {
	v := m.content
	if v == nil {
		return
	}
	return *v, true
}

// OldContent returns the old "content" field's value of the Envelope entity.
// If the Envelope object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EnvelopeMutation) OldContent(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldContent is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldContent requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldContent: %w", err)
	}
	return oldValue.Content, nil
}

// ResetContent resets all changes to the "content" field.
func (m *EnvelopeMutation) ResetContent() {
	m.content = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *EnvelopeMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *EnvelopeMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Envelope entity.
// If the Envelope object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EnvelopeMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *EnvelopeMutation) ResetCreatedAt() {
	m.created_at = nil
}

// AddAttachmentIDs adds the "attachments" edge to the Attachment entity by ids.
func (m *EnvelopeMutation) AddAttachmentIDs(ids ...string) {
	if m.attachments == nil {
		m.attachments = make(map[string]struct{})
	}
	for i := range ids {
		m.attachments[ids[i]] = struct{}{}
	}
}

// ClearAttachments clears the "attachments" edge to the Attachment entity.
func (m *EnvelopeMutation) ClearAttachments() {
	m.clearedattachments = true
}

// AttachmentsCleared reports if the "attachments" edge to the Attachment entity was cleared.
func (m *EnvelopeMutation) AttachmentsCleared() bool {
	return m.clearedattachments
}

// RemoveAttachmentIDs removes the "attachments" edge to the Attachment entity by IDs.
func (m *EnvelopeMutation) RemoveAttachmentIDs(ids ...string) {
	if m.removedattachments == nil {
		m.removedattachments = make(map[string]struct{})
	}
	for i := range ids {
		delete(m.attachments, ids[i])
		m.removedattachments[ids[i]] = struct{}{}
	}
}

// RemovedAttachments returns the removed IDs of the "attachments" edge to the Attachment entity.
func (m *EnvelopeMutation) RemovedAttachmentsIDs() (ids []string) {
	for id := range m.removedattachments {
		ids = append(ids, id)
	}
	return
}

// AttachmentsIDs returns the "attachments" edge IDs in the mutation.
func (m *EnvelopeMutation) AttachmentsIDs() (ids []string) {
	for id := range m.attachments {
		ids = append(ids, id)
	}
	return
}

// ResetAttachments resets all changes to the "attachments" edge.
func (m *EnvelopeMutation) ResetAttachments() {
	m.attachments = nil
	m.clearedattachments = false
	m.removedattachments = nil
}

// Where appends a list predicates to the EnvelopeMutation builder.
func (m *EnvelopeMutation) Where(ps ...predicate.Envelope) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the EnvelopeMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *EnvelopeMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Envelope, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *EnvelopeMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *EnvelopeMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Envelope).
func (m *EnvelopeMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *EnvelopeMutation) Fields() []string {
	fields := make([]string, 0, 5)
	if m.to != nil {
		fields = append(fields, envelope.FieldTo)
	}
	if m.from != nil {
		fields = append(fields, envelope.FieldFrom)
	}
	if m.subject != nil {
		fields = append(fields, envelope.FieldSubject)
	}
	if m.content != nil {
		fields = append(fields, envelope.FieldContent)
	}
	if m.created_at != nil {
		fields = append(fields, envelope.FieldCreatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *EnvelopeMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case envelope.FieldTo:
		return m.To()
	case envelope.FieldFrom:
		return m.From()
	case envelope.FieldSubject:
		return m.Subject()
	case envelope.FieldContent:
		return m.Content()
	case envelope.FieldCreatedAt:
		return m.CreatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *EnvelopeMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case envelope.FieldTo:
		return m.OldTo(ctx)
	case envelope.FieldFrom:
		return m.OldFrom(ctx)
	case envelope.FieldSubject:
		return m.OldSubject(ctx)
	case envelope.FieldContent:
		return m.OldContent(ctx)
	case envelope.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Envelope field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *EnvelopeMutation) SetField(name string, value ent.Value) error {
	switch name {
	case envelope.FieldTo:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTo(v)
		return nil
	case envelope.FieldFrom:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetFrom(v)
		return nil
	case envelope.FieldSubject:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetSubject(v)
		return nil
	case envelope.FieldContent:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetContent(v)
		return nil
	case envelope.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Envelope field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *EnvelopeMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *EnvelopeMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *EnvelopeMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Envelope numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *EnvelopeMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *EnvelopeMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *EnvelopeMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Envelope nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *EnvelopeMutation) ResetField(name string) error {
	switch name {
	case envelope.FieldTo:
		m.ResetTo()
		return nil
	case envelope.FieldFrom:
		m.ResetFrom()
		return nil
	case envelope.FieldSubject:
		m.ResetSubject()
		return nil
	case envelope.FieldContent:
		m.ResetContent()
		return nil
	case envelope.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	}
	return fmt.Errorf("unknown Envelope field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *EnvelopeMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.attachments != nil {
		edges = append(edges, envelope.EdgeAttachments)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *EnvelopeMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case envelope.EdgeAttachments:
		ids := make([]ent.Value, 0, len(m.attachments))
		for id := range m.attachments {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *EnvelopeMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedattachments != nil {
		edges = append(edges, envelope.EdgeAttachments)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *EnvelopeMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case envelope.EdgeAttachments:
		ids := make([]ent.Value, 0, len(m.removedattachments))
		for id := range m.removedattachments {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *EnvelopeMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedattachments {
		edges = append(edges, envelope.EdgeAttachments)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *EnvelopeMutation) EdgeCleared(name string) bool {
	switch name {
	case envelope.EdgeAttachments:
		return m.clearedattachments
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *EnvelopeMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Envelope unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *EnvelopeMutation) ResetEdge(name string) error {
	switch name {
	case envelope.EdgeAttachments:
		m.ResetAttachments()
		return nil
	}
	return fmt.Errorf("unknown Envelope edge %s", name)
}
