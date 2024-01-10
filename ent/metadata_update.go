// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dskydiver/ipfs-metadata-fetcher/ent/metadata"
	"github.com/dskydiver/ipfs-metadata-fetcher/ent/predicate"
)

// MetadataUpdate is the builder for updating Metadata entities.
type MetadataUpdate struct {
	config
	hooks    []Hook
	mutation *MetadataMutation
}

// Where appends a list predicates to the MetadataUpdate builder.
func (mu *MetadataUpdate) Where(ps ...predicate.Metadata) *MetadataUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetCid sets the "cid" field.
func (mu *MetadataUpdate) SetCid(s string) *MetadataUpdate {
	mu.mutation.SetCid(s)
	return mu
}

// SetNillableCid sets the "cid" field if the given value is not nil.
func (mu *MetadataUpdate) SetNillableCid(s *string) *MetadataUpdate {
	if s != nil {
		mu.SetCid(*s)
	}
	return mu
}

// SetImage sets the "image" field.
func (mu *MetadataUpdate) SetImage(s string) *MetadataUpdate {
	mu.mutation.SetImage(s)
	return mu
}

// SetNillableImage sets the "image" field if the given value is not nil.
func (mu *MetadataUpdate) SetNillableImage(s *string) *MetadataUpdate {
	if s != nil {
		mu.SetImage(*s)
	}
	return mu
}

// SetDescription sets the "description" field.
func (mu *MetadataUpdate) SetDescription(s string) *MetadataUpdate {
	mu.mutation.SetDescription(s)
	return mu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (mu *MetadataUpdate) SetNillableDescription(s *string) *MetadataUpdate {
	if s != nil {
		mu.SetDescription(*s)
	}
	return mu
}

// SetName sets the "name" field.
func (mu *MetadataUpdate) SetName(s string) *MetadataUpdate {
	mu.mutation.SetName(s)
	return mu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (mu *MetadataUpdate) SetNillableName(s *string) *MetadataUpdate {
	if s != nil {
		mu.SetName(*s)
	}
	return mu
}

// SetCreatedAt sets the "created_at" field.
func (mu *MetadataUpdate) SetCreatedAt(t time.Time) *MetadataUpdate {
	mu.mutation.SetCreatedAt(t)
	return mu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mu *MetadataUpdate) SetNillableCreatedAt(t *time.Time) *MetadataUpdate {
	if t != nil {
		mu.SetCreatedAt(*t)
	}
	return mu
}

// SetDeletedAt sets the "deleted_at" field.
func (mu *MetadataUpdate) SetDeletedAt(t time.Time) *MetadataUpdate {
	mu.mutation.SetDeletedAt(t)
	return mu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (mu *MetadataUpdate) SetNillableDeletedAt(t *time.Time) *MetadataUpdate {
	if t != nil {
		mu.SetDeletedAt(*t)
	}
	return mu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (mu *MetadataUpdate) ClearDeletedAt() *MetadataUpdate {
	mu.mutation.ClearDeletedAt()
	return mu
}

// Mutation returns the MetadataMutation object of the builder.
func (mu *MetadataUpdate) Mutation() *MetadataMutation {
	return mu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MetadataUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, mu.sqlSave, mu.mutation, mu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MetadataUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MetadataUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MetadataUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mu *MetadataUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(metadata.Table, metadata.Columns, sqlgraph.NewFieldSpec(metadata.FieldID, field.TypeString))
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.Cid(); ok {
		_spec.SetField(metadata.FieldCid, field.TypeString, value)
	}
	if value, ok := mu.mutation.Image(); ok {
		_spec.SetField(metadata.FieldImage, field.TypeString, value)
	}
	if value, ok := mu.mutation.Description(); ok {
		_spec.SetField(metadata.FieldDescription, field.TypeString, value)
	}
	if value, ok := mu.mutation.Name(); ok {
		_spec.SetField(metadata.FieldName, field.TypeString, value)
	}
	if value, ok := mu.mutation.CreatedAt(); ok {
		_spec.SetField(metadata.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := mu.mutation.DeletedAt(); ok {
		_spec.SetField(metadata.FieldDeletedAt, field.TypeTime, value)
	}
	if mu.mutation.DeletedAtCleared() {
		_spec.ClearField(metadata.FieldDeletedAt, field.TypeTime)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{metadata.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mu.mutation.done = true
	return n, nil
}

// MetadataUpdateOne is the builder for updating a single Metadata entity.
type MetadataUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MetadataMutation
}

// SetCid sets the "cid" field.
func (muo *MetadataUpdateOne) SetCid(s string) *MetadataUpdateOne {
	muo.mutation.SetCid(s)
	return muo
}

// SetNillableCid sets the "cid" field if the given value is not nil.
func (muo *MetadataUpdateOne) SetNillableCid(s *string) *MetadataUpdateOne {
	if s != nil {
		muo.SetCid(*s)
	}
	return muo
}

// SetImage sets the "image" field.
func (muo *MetadataUpdateOne) SetImage(s string) *MetadataUpdateOne {
	muo.mutation.SetImage(s)
	return muo
}

// SetNillableImage sets the "image" field if the given value is not nil.
func (muo *MetadataUpdateOne) SetNillableImage(s *string) *MetadataUpdateOne {
	if s != nil {
		muo.SetImage(*s)
	}
	return muo
}

// SetDescription sets the "description" field.
func (muo *MetadataUpdateOne) SetDescription(s string) *MetadataUpdateOne {
	muo.mutation.SetDescription(s)
	return muo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (muo *MetadataUpdateOne) SetNillableDescription(s *string) *MetadataUpdateOne {
	if s != nil {
		muo.SetDescription(*s)
	}
	return muo
}

// SetName sets the "name" field.
func (muo *MetadataUpdateOne) SetName(s string) *MetadataUpdateOne {
	muo.mutation.SetName(s)
	return muo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (muo *MetadataUpdateOne) SetNillableName(s *string) *MetadataUpdateOne {
	if s != nil {
		muo.SetName(*s)
	}
	return muo
}

// SetCreatedAt sets the "created_at" field.
func (muo *MetadataUpdateOne) SetCreatedAt(t time.Time) *MetadataUpdateOne {
	muo.mutation.SetCreatedAt(t)
	return muo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (muo *MetadataUpdateOne) SetNillableCreatedAt(t *time.Time) *MetadataUpdateOne {
	if t != nil {
		muo.SetCreatedAt(*t)
	}
	return muo
}

// SetDeletedAt sets the "deleted_at" field.
func (muo *MetadataUpdateOne) SetDeletedAt(t time.Time) *MetadataUpdateOne {
	muo.mutation.SetDeletedAt(t)
	return muo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (muo *MetadataUpdateOne) SetNillableDeletedAt(t *time.Time) *MetadataUpdateOne {
	if t != nil {
		muo.SetDeletedAt(*t)
	}
	return muo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (muo *MetadataUpdateOne) ClearDeletedAt() *MetadataUpdateOne {
	muo.mutation.ClearDeletedAt()
	return muo
}

// Mutation returns the MetadataMutation object of the builder.
func (muo *MetadataUpdateOne) Mutation() *MetadataMutation {
	return muo.mutation
}

// Where appends a list predicates to the MetadataUpdate builder.
func (muo *MetadataUpdateOne) Where(ps ...predicate.Metadata) *MetadataUpdateOne {
	muo.mutation.Where(ps...)
	return muo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MetadataUpdateOne) Select(field string, fields ...string) *MetadataUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Metadata entity.
func (muo *MetadataUpdateOne) Save(ctx context.Context) (*Metadata, error) {
	return withHooks(ctx, muo.sqlSave, muo.mutation, muo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MetadataUpdateOne) SaveX(ctx context.Context) *Metadata {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MetadataUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MetadataUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (muo *MetadataUpdateOne) sqlSave(ctx context.Context) (_node *Metadata, err error) {
	_spec := sqlgraph.NewUpdateSpec(metadata.Table, metadata.Columns, sqlgraph.NewFieldSpec(metadata.FieldID, field.TypeString))
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Metadata.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, metadata.FieldID)
		for _, f := range fields {
			if !metadata.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != metadata.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.Cid(); ok {
		_spec.SetField(metadata.FieldCid, field.TypeString, value)
	}
	if value, ok := muo.mutation.Image(); ok {
		_spec.SetField(metadata.FieldImage, field.TypeString, value)
	}
	if value, ok := muo.mutation.Description(); ok {
		_spec.SetField(metadata.FieldDescription, field.TypeString, value)
	}
	if value, ok := muo.mutation.Name(); ok {
		_spec.SetField(metadata.FieldName, field.TypeString, value)
	}
	if value, ok := muo.mutation.CreatedAt(); ok {
		_spec.SetField(metadata.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := muo.mutation.DeletedAt(); ok {
		_spec.SetField(metadata.FieldDeletedAt, field.TypeTime, value)
	}
	if muo.mutation.DeletedAtCleared() {
		_spec.ClearField(metadata.FieldDeletedAt, field.TypeTime)
	}
	_node = &Metadata{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{metadata.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	muo.mutation.done = true
	return _node, nil
}
