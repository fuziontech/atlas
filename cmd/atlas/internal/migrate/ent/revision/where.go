// Copyright 2021-present The Atlas Authors. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package revision

import (
	"time"

	"ariga.io/atlas/cmd/atlas/internal/migrate/ent/predicate"
	"ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v migrate.RevisionType) predicate.Revision {
	vc := uint(v)
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), vc))
	})
}

// Applied applies equality check predicate on the "applied" field. It's identical to AppliedEQ.
func Applied(v int) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldApplied), v))
	})
}

// Total applies equality check predicate on the "total" field. It's identical to TotalEQ.
func Total(v int) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTotal), v))
	})
}

// ExecutedAt applies equality check predicate on the "executed_at" field. It's identical to ExecutedAtEQ.
func ExecutedAt(v time.Time) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExecutedAt), v))
	})
}

// ExecutionTime applies equality check predicate on the "execution_time" field. It's identical to ExecutionTimeEQ.
func ExecutionTime(v time.Duration) predicate.Revision {
	vc := int64(v)
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExecutionTime), vc))
	})
}

// Error applies equality check predicate on the "error" field. It's identical to ErrorEQ.
func Error(v string) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldError), v))
	})
}

// Hash applies equality check predicate on the "hash" field. It's identical to HashEQ.
func Hash(v string) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHash), v))
	})
}

// OperatorVersion applies equality check predicate on the "operator_version" field. It's identical to OperatorVersionEQ.
func OperatorVersion(v string) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOperatorVersion), v))
	})
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDescription), v))
	})
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Revision {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDescription), v...))
	})
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Revision {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDescription), v...))
	})
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDescription), v))
	})
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDescription), v))
	})
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDescription), v))
	})
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDescription), v))
	})
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDescription), v))
	})
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDescription), v))
	})
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDescription), v))
	})
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDescription), v))
	})
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDescription), v))
	})
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v migrate.RevisionType) predicate.Revision {
	vc := uint(v)
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), vc))
	})
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v migrate.RevisionType) predicate.Revision {
	vc := uint(v)
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldType), vc))
	})
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...migrate.RevisionType) predicate.Revision {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = uint(vs[i])
	}
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldType), v...))
	})
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...migrate.RevisionType) predicate.Revision {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = uint(vs[i])
	}
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldType), v...))
	})
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v migrate.RevisionType) predicate.Revision {
	vc := uint(v)
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldType), vc))
	})
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v migrate.RevisionType) predicate.Revision {
	vc := uint(v)
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldType), vc))
	})
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v migrate.RevisionType) predicate.Revision {
	vc := uint(v)
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldType), vc))
	})
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v migrate.RevisionType) predicate.Revision {
	vc := uint(v)
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldType), vc))
	})
}

// AppliedEQ applies the EQ predicate on the "applied" field.
func AppliedEQ(v int) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldApplied), v))
	})
}

// AppliedNEQ applies the NEQ predicate on the "applied" field.
func AppliedNEQ(v int) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldApplied), v))
	})
}

// AppliedIn applies the In predicate on the "applied" field.
func AppliedIn(vs ...int) predicate.Revision {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldApplied), v...))
	})
}

// AppliedNotIn applies the NotIn predicate on the "applied" field.
func AppliedNotIn(vs ...int) predicate.Revision {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldApplied), v...))
	})
}

// AppliedGT applies the GT predicate on the "applied" field.
func AppliedGT(v int) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldApplied), v))
	})
}

// AppliedGTE applies the GTE predicate on the "applied" field.
func AppliedGTE(v int) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldApplied), v))
	})
}

// AppliedLT applies the LT predicate on the "applied" field.
func AppliedLT(v int) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldApplied), v))
	})
}

// AppliedLTE applies the LTE predicate on the "applied" field.
func AppliedLTE(v int) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldApplied), v))
	})
}

// TotalEQ applies the EQ predicate on the "total" field.
func TotalEQ(v int) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTotal), v))
	})
}

// TotalNEQ applies the NEQ predicate on the "total" field.
func TotalNEQ(v int) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTotal), v))
	})
}

// TotalIn applies the In predicate on the "total" field.
func TotalIn(vs ...int) predicate.Revision {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTotal), v...))
	})
}

// TotalNotIn applies the NotIn predicate on the "total" field.
func TotalNotIn(vs ...int) predicate.Revision {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTotal), v...))
	})
}

// TotalGT applies the GT predicate on the "total" field.
func TotalGT(v int) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTotal), v))
	})
}

// TotalGTE applies the GTE predicate on the "total" field.
func TotalGTE(v int) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTotal), v))
	})
}

// TotalLT applies the LT predicate on the "total" field.
func TotalLT(v int) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTotal), v))
	})
}

// TotalLTE applies the LTE predicate on the "total" field.
func TotalLTE(v int) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTotal), v))
	})
}

// ExecutedAtEQ applies the EQ predicate on the "executed_at" field.
func ExecutedAtEQ(v time.Time) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExecutedAt), v))
	})
}

// ExecutedAtNEQ applies the NEQ predicate on the "executed_at" field.
func ExecutedAtNEQ(v time.Time) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldExecutedAt), v))
	})
}

// ExecutedAtIn applies the In predicate on the "executed_at" field.
func ExecutedAtIn(vs ...time.Time) predicate.Revision {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldExecutedAt), v...))
	})
}

// ExecutedAtNotIn applies the NotIn predicate on the "executed_at" field.
func ExecutedAtNotIn(vs ...time.Time) predicate.Revision {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldExecutedAt), v...))
	})
}

// ExecutedAtGT applies the GT predicate on the "executed_at" field.
func ExecutedAtGT(v time.Time) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldExecutedAt), v))
	})
}

// ExecutedAtGTE applies the GTE predicate on the "executed_at" field.
func ExecutedAtGTE(v time.Time) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldExecutedAt), v))
	})
}

// ExecutedAtLT applies the LT predicate on the "executed_at" field.
func ExecutedAtLT(v time.Time) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldExecutedAt), v))
	})
}

// ExecutedAtLTE applies the LTE predicate on the "executed_at" field.
func ExecutedAtLTE(v time.Time) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldExecutedAt), v))
	})
}

// ExecutionTimeEQ applies the EQ predicate on the "execution_time" field.
func ExecutionTimeEQ(v time.Duration) predicate.Revision {
	vc := int64(v)
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExecutionTime), vc))
	})
}

// ExecutionTimeNEQ applies the NEQ predicate on the "execution_time" field.
func ExecutionTimeNEQ(v time.Duration) predicate.Revision {
	vc := int64(v)
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldExecutionTime), vc))
	})
}

// ExecutionTimeIn applies the In predicate on the "execution_time" field.
func ExecutionTimeIn(vs ...time.Duration) predicate.Revision {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int64(vs[i])
	}
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldExecutionTime), v...))
	})
}

// ExecutionTimeNotIn applies the NotIn predicate on the "execution_time" field.
func ExecutionTimeNotIn(vs ...time.Duration) predicate.Revision {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int64(vs[i])
	}
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldExecutionTime), v...))
	})
}

// ExecutionTimeGT applies the GT predicate on the "execution_time" field.
func ExecutionTimeGT(v time.Duration) predicate.Revision {
	vc := int64(v)
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldExecutionTime), vc))
	})
}

// ExecutionTimeGTE applies the GTE predicate on the "execution_time" field.
func ExecutionTimeGTE(v time.Duration) predicate.Revision {
	vc := int64(v)
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldExecutionTime), vc))
	})
}

// ExecutionTimeLT applies the LT predicate on the "execution_time" field.
func ExecutionTimeLT(v time.Duration) predicate.Revision {
	vc := int64(v)
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldExecutionTime), vc))
	})
}

// ExecutionTimeLTE applies the LTE predicate on the "execution_time" field.
func ExecutionTimeLTE(v time.Duration) predicate.Revision {
	vc := int64(v)
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldExecutionTime), vc))
	})
}

// ErrorEQ applies the EQ predicate on the "error" field.
func ErrorEQ(v string) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldError), v))
	})
}

// ErrorNEQ applies the NEQ predicate on the "error" field.
func ErrorNEQ(v string) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldError), v))
	})
}

// ErrorIn applies the In predicate on the "error" field.
func ErrorIn(vs ...string) predicate.Revision {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldError), v...))
	})
}

// ErrorNotIn applies the NotIn predicate on the "error" field.
func ErrorNotIn(vs ...string) predicate.Revision {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldError), v...))
	})
}

// ErrorGT applies the GT predicate on the "error" field.
func ErrorGT(v string) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldError), v))
	})
}

// ErrorGTE applies the GTE predicate on the "error" field.
func ErrorGTE(v string) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldError), v))
	})
}

// ErrorLT applies the LT predicate on the "error" field.
func ErrorLT(v string) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldError), v))
	})
}

// ErrorLTE applies the LTE predicate on the "error" field.
func ErrorLTE(v string) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldError), v))
	})
}

// ErrorContains applies the Contains predicate on the "error" field.
func ErrorContains(v string) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldError), v))
	})
}

// ErrorHasPrefix applies the HasPrefix predicate on the "error" field.
func ErrorHasPrefix(v string) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldError), v))
	})
}

// ErrorHasSuffix applies the HasSuffix predicate on the "error" field.
func ErrorHasSuffix(v string) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldError), v))
	})
}

// ErrorIsNil applies the IsNil predicate on the "error" field.
func ErrorIsNil() predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldError)))
	})
}

// ErrorNotNil applies the NotNil predicate on the "error" field.
func ErrorNotNil() predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldError)))
	})
}

// ErrorEqualFold applies the EqualFold predicate on the "error" field.
func ErrorEqualFold(v string) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldError), v))
	})
}

// ErrorContainsFold applies the ContainsFold predicate on the "error" field.
func ErrorContainsFold(v string) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldError), v))
	})
}

// HashEQ applies the EQ predicate on the "hash" field.
func HashEQ(v string) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHash), v))
	})
}

// HashNEQ applies the NEQ predicate on the "hash" field.
func HashNEQ(v string) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHash), v))
	})
}

// HashIn applies the In predicate on the "hash" field.
func HashIn(vs ...string) predicate.Revision {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldHash), v...))
	})
}

// HashNotIn applies the NotIn predicate on the "hash" field.
func HashNotIn(vs ...string) predicate.Revision {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldHash), v...))
	})
}

// HashGT applies the GT predicate on the "hash" field.
func HashGT(v string) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldHash), v))
	})
}

// HashGTE applies the GTE predicate on the "hash" field.
func HashGTE(v string) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldHash), v))
	})
}

// HashLT applies the LT predicate on the "hash" field.
func HashLT(v string) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldHash), v))
	})
}

// HashLTE applies the LTE predicate on the "hash" field.
func HashLTE(v string) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldHash), v))
	})
}

// HashContains applies the Contains predicate on the "hash" field.
func HashContains(v string) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldHash), v))
	})
}

// HashHasPrefix applies the HasPrefix predicate on the "hash" field.
func HashHasPrefix(v string) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldHash), v))
	})
}

// HashHasSuffix applies the HasSuffix predicate on the "hash" field.
func HashHasSuffix(v string) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldHash), v))
	})
}

// HashEqualFold applies the EqualFold predicate on the "hash" field.
func HashEqualFold(v string) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldHash), v))
	})
}

// HashContainsFold applies the ContainsFold predicate on the "hash" field.
func HashContainsFold(v string) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldHash), v))
	})
}

// PartialHashesIsNil applies the IsNil predicate on the "partial_hashes" field.
func PartialHashesIsNil() predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldPartialHashes)))
	})
}

// PartialHashesNotNil applies the NotNil predicate on the "partial_hashes" field.
func PartialHashesNotNil() predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldPartialHashes)))
	})
}

// OperatorVersionEQ applies the EQ predicate on the "operator_version" field.
func OperatorVersionEQ(v string) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOperatorVersion), v))
	})
}

// OperatorVersionNEQ applies the NEQ predicate on the "operator_version" field.
func OperatorVersionNEQ(v string) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOperatorVersion), v))
	})
}

// OperatorVersionIn applies the In predicate on the "operator_version" field.
func OperatorVersionIn(vs ...string) predicate.Revision {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldOperatorVersion), v...))
	})
}

// OperatorVersionNotIn applies the NotIn predicate on the "operator_version" field.
func OperatorVersionNotIn(vs ...string) predicate.Revision {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldOperatorVersion), v...))
	})
}

// OperatorVersionGT applies the GT predicate on the "operator_version" field.
func OperatorVersionGT(v string) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOperatorVersion), v))
	})
}

// OperatorVersionGTE applies the GTE predicate on the "operator_version" field.
func OperatorVersionGTE(v string) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOperatorVersion), v))
	})
}

// OperatorVersionLT applies the LT predicate on the "operator_version" field.
func OperatorVersionLT(v string) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOperatorVersion), v))
	})
}

// OperatorVersionLTE applies the LTE predicate on the "operator_version" field.
func OperatorVersionLTE(v string) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOperatorVersion), v))
	})
}

// OperatorVersionContains applies the Contains predicate on the "operator_version" field.
func OperatorVersionContains(v string) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOperatorVersion), v))
	})
}

// OperatorVersionHasPrefix applies the HasPrefix predicate on the "operator_version" field.
func OperatorVersionHasPrefix(v string) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOperatorVersion), v))
	})
}

// OperatorVersionHasSuffix applies the HasSuffix predicate on the "operator_version" field.
func OperatorVersionHasSuffix(v string) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOperatorVersion), v))
	})
}

// OperatorVersionEqualFold applies the EqualFold predicate on the "operator_version" field.
func OperatorVersionEqualFold(v string) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOperatorVersion), v))
	})
}

// OperatorVersionContainsFold applies the ContainsFold predicate on the "operator_version" field.
func OperatorVersionContainsFold(v string) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOperatorVersion), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Revision) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Revision) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Revision) predicate.Revision {
	return predicate.Revision(func(s *sql.Selector) {
		p(s.Not())
	})
}
