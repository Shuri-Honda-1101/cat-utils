// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Shuri-Honda-1101/cat-utils/ent/cat"
	"github.com/Shuri-Honda-1101/cat-utils/ent/predicate"
	"github.com/Shuri-Honda-1101/cat-utils/ent/toilet"
	"github.com/Shuri-Honda-1101/cat-utils/ent/user"
	"github.com/google/uuid"
)

// CatQuery is the builder for querying Cat entities.
type CatQuery struct {
	config
	ctx         *QueryContext
	order       []OrderFunc
	inters      []Interceptor
	predicates  []predicate.Cat
	withOwner   *UserQuery
	withToilets *ToiletQuery
	withFKs     bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CatQuery builder.
func (cq *CatQuery) Where(ps ...predicate.Cat) *CatQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit the number of records to be returned by this query.
func (cq *CatQuery) Limit(limit int) *CatQuery {
	cq.ctx.Limit = &limit
	return cq
}

// Offset to start from.
func (cq *CatQuery) Offset(offset int) *CatQuery {
	cq.ctx.Offset = &offset
	return cq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cq *CatQuery) Unique(unique bool) *CatQuery {
	cq.ctx.Unique = &unique
	return cq
}

// Order specifies how the records should be ordered.
func (cq *CatQuery) Order(o ...OrderFunc) *CatQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// QueryOwner chains the current query on the "owner" edge.
func (cq *CatQuery) QueryOwner() *UserQuery {
	query := (&UserClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(cat.Table, cat.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, cat.OwnerTable, cat.OwnerColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryToilets chains the current query on the "toilets" edge.
func (cq *CatQuery) QueryToilets() *ToiletQuery {
	query := (&ToiletClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(cat.Table, cat.FieldID, selector),
			sqlgraph.To(toilet.Table, toilet.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, cat.ToiletsTable, cat.ToiletsColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Cat entity from the query.
// Returns a *NotFoundError when no Cat was found.
func (cq *CatQuery) First(ctx context.Context) (*Cat, error) {
	nodes, err := cq.Limit(1).All(setContextOp(ctx, cq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{cat.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *CatQuery) FirstX(ctx context.Context) *Cat {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Cat ID from the query.
// Returns a *NotFoundError when no Cat ID was found.
func (cq *CatQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = cq.Limit(1).IDs(setContextOp(ctx, cq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{cat.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cq *CatQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Cat entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Cat entity is found.
// Returns a *NotFoundError when no Cat entities are found.
func (cq *CatQuery) Only(ctx context.Context) (*Cat, error) {
	nodes, err := cq.Limit(2).All(setContextOp(ctx, cq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{cat.Label}
	default:
		return nil, &NotSingularError{cat.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *CatQuery) OnlyX(ctx context.Context) *Cat {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Cat ID in the query.
// Returns a *NotSingularError when more than one Cat ID is found.
// Returns a *NotFoundError when no entities are found.
func (cq *CatQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = cq.Limit(2).IDs(setContextOp(ctx, cq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{cat.Label}
	default:
		err = &NotSingularError{cat.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *CatQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Cats.
func (cq *CatQuery) All(ctx context.Context) ([]*Cat, error) {
	ctx = setContextOp(ctx, cq.ctx, "All")
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Cat, *CatQuery]()
	return withInterceptors[[]*Cat](ctx, cq, qr, cq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cq *CatQuery) AllX(ctx context.Context) []*Cat {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Cat IDs.
func (cq *CatQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if cq.ctx.Unique == nil && cq.path != nil {
		cq.Unique(true)
	}
	ctx = setContextOp(ctx, cq.ctx, "IDs")
	if err = cq.Select(cat.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *CatQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *CatQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, cq.ctx, "Count")
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cq, querierCount[*CatQuery](), cq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cq *CatQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *CatQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, cq.ctx, "Exist")
	switch _, err := cq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *CatQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CatQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *CatQuery) Clone() *CatQuery {
	if cq == nil {
		return nil
	}
	return &CatQuery{
		config:      cq.config,
		ctx:         cq.ctx.Clone(),
		order:       append([]OrderFunc{}, cq.order...),
		inters:      append([]Interceptor{}, cq.inters...),
		predicates:  append([]predicate.Cat{}, cq.predicates...),
		withOwner:   cq.withOwner.Clone(),
		withToilets: cq.withToilets.Clone(),
		// clone intermediate query.
		sql:  cq.sql.Clone(),
		path: cq.path,
	}
}

// WithOwner tells the query-builder to eager-load the nodes that are connected to
// the "owner" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CatQuery) WithOwner(opts ...func(*UserQuery)) *CatQuery {
	query := (&UserClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withOwner = query
	return cq
}

// WithToilets tells the query-builder to eager-load the nodes that are connected to
// the "toilets" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CatQuery) WithToilets(opts ...func(*ToiletQuery)) *CatQuery {
	query := (&ToiletClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withToilets = query
	return cq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Cat.Query().
//		GroupBy(cat.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cq *CatQuery) GroupBy(field string, fields ...string) *CatGroupBy {
	cq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CatGroupBy{build: cq}
	grbuild.flds = &cq.ctx.Fields
	grbuild.label = cat.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Cat.Query().
//		Select(cat.FieldName).
//		Scan(ctx, &v)
func (cq *CatQuery) Select(fields ...string) *CatSelect {
	cq.ctx.Fields = append(cq.ctx.Fields, fields...)
	sbuild := &CatSelect{CatQuery: cq}
	sbuild.label = cat.Label
	sbuild.flds, sbuild.scan = &cq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CatSelect configured with the given aggregations.
func (cq *CatQuery) Aggregate(fns ...AggregateFunc) *CatSelect {
	return cq.Select().Aggregate(fns...)
}

func (cq *CatQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cq); err != nil {
				return err
			}
		}
	}
	for _, f := range cq.ctx.Fields {
		if !cat.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cq.path != nil {
		prev, err := cq.path(ctx)
		if err != nil {
			return err
		}
		cq.sql = prev
	}
	return nil
}

func (cq *CatQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Cat, error) {
	var (
		nodes       = []*Cat{}
		withFKs     = cq.withFKs
		_spec       = cq.querySpec()
		loadedTypes = [2]bool{
			cq.withOwner != nil,
			cq.withToilets != nil,
		}
	)
	if cq.withOwner != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, cat.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Cat).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Cat{config: cq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := cq.withOwner; query != nil {
		if err := cq.loadOwner(ctx, query, nodes, nil,
			func(n *Cat, e *User) { n.Edges.Owner = e }); err != nil {
			return nil, err
		}
	}
	if query := cq.withToilets; query != nil {
		if err := cq.loadToilets(ctx, query, nodes,
			func(n *Cat) { n.Edges.Toilets = []*Toilet{} },
			func(n *Cat, e *Toilet) { n.Edges.Toilets = append(n.Edges.Toilets, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cq *CatQuery) loadOwner(ctx context.Context, query *UserQuery, nodes []*Cat, init func(*Cat), assign func(*Cat, *User)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Cat)
	for i := range nodes {
		if nodes[i].user_cats == nil {
			continue
		}
		fk := *nodes[i].user_cats
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_cats" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (cq *CatQuery) loadToilets(ctx context.Context, query *ToiletQuery, nodes []*Cat, init func(*Cat), assign func(*Cat, *Toilet)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Cat)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Toilet(func(s *sql.Selector) {
		s.Where(sql.InValues(cat.ToiletsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.cat_toilets
		if fk == nil {
			return fmt.Errorf(`foreign-key "cat_toilets" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "cat_toilets" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (cq *CatQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
	_spec.Node.Columns = cq.ctx.Fields
	if len(cq.ctx.Fields) > 0 {
		_spec.Unique = cq.ctx.Unique != nil && *cq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *CatQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(cat.Table, cat.Columns, sqlgraph.NewFieldSpec(cat.FieldID, field.TypeUUID))
	_spec.From = cq.sql
	if unique := cq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if cq.path != nil {
		_spec.Unique = true
	}
	if fields := cq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, cat.FieldID)
		for i := range fields {
			if fields[i] != cat.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cq *CatQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(cat.Table)
	columns := cq.ctx.Fields
	if len(columns) == 0 {
		columns = cat.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cq.ctx.Unique != nil && *cq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	for _, p := range cq.order {
		p(selector)
	}
	if offset := cq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CatGroupBy is the group-by builder for Cat entities.
type CatGroupBy struct {
	selector
	build *CatQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *CatGroupBy) Aggregate(fns ...AggregateFunc) *CatGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the selector query and scans the result into the given value.
func (cgb *CatGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cgb.build.ctx, "GroupBy")
	if err := cgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CatQuery, *CatGroupBy](ctx, cgb.build, cgb, cgb.build.inters, v)
}

func (cgb *CatGroupBy) sqlScan(ctx context.Context, root *CatQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cgb.fns))
	for _, fn := range cgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cgb.flds)+len(cgb.fns))
		for _, f := range *cgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CatSelect is the builder for selecting fields of Cat entities.
type CatSelect struct {
	*CatQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cs *CatSelect) Aggregate(fns ...AggregateFunc) *CatSelect {
	cs.fns = append(cs.fns, fns...)
	return cs
}

// Scan applies the selector query and scans the result into the given value.
func (cs *CatSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cs.ctx, "Select")
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CatQuery, *CatSelect](ctx, cs.CatQuery, cs, cs.inters, v)
}

func (cs *CatSelect) sqlScan(ctx context.Context, root *CatQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cs.fns))
	for _, fn := range cs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
