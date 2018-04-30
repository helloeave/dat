package runner

import (
	"database/sql"
	"testing"
	"time"

	"github.com/helloeave/dat/dat"
	"github.com/lib/pq/hstore"
	"github.com/stretchr/testify/require"
	"gopkg.in/stretchr/testify.v1/assert"
)

func TestUpsertReal(t *testing.T) {
	// Insert by specifying values
	s := beginTxWithFixtures()
	defer s.AutoRollback()

	var id int64
	err := s.Upsert("people").
		Columns("name", "email").
		Values("mario", "mgutz@mgutz.com").
		Where("name = $1", "mario").
		Returning("id").
		QueryScalar(&id)
	assert.NoError(t, err)
	assert.True(t, id > 0)

	var id2 int64
	err = s.Upsert("people").
		Columns("name", "email").
		Values("mario", "mario@foo.com").
		Where("name = $1", "mario").
		Returning("id").
		QueryScalar(&id2)
	assert.NoError(t, err)
	assert.Equal(t, id, id2)

	// Insert by specifying a record (ptr to struct)
	person := Person{Name: "Barack"}
	person.Email.Valid = true
	person.Email.String = "obama1@whitehouse.gov"

	err = s.
		Upsert("people").
		Columns("name", "email").
		Record(&person).
		Where("name = $1", "Barack").
		Returning("id", "created_at").
		QueryStruct(&person)
	assert.NoError(t, err)
	assert.True(t, person.ID > 0)
	assert.NotEqual(t, person.CreatedAt, dat.NullTime{})
}

func TestUpsertReal_Nullability_Record(t *testing.T) {
	// Insert by specifying values
	s := beginTxWithFixtures()
	defer s.AutoRollback()

	// with nil values
	person := Person{
		Nullable:    nil,
		NullableMap: hstore.Hstore{},
		NullableAt:  nil,
	}
	err := s.
		Upsert("people").
		Columns("nullable", "nullable_map", "nullable_at").
		Record(person).
		Where("name = 'Mario'").
		Returning("id", "nullable", "nullable_map", "nullable_at").
		QueryStruct(&person)

	require.NoError(t, err)
	assert.True(t, person.ID > 0)
	assert.Nil(t, person.Nullable)
	assert.Nil(t, person.NullableMap.Map)
	assert.Nil(t, person.NullableAt)

	// with non-nil values
	oneOneEighteen := time.Date(2018, time.January, 1, 0, 0, 0, 0, time.UTC)
	yep := "yep"
	person = Person{
		Nullable: &yep,
		NullableMap: hstore.Hstore{
			Map: map[string]sql.NullString{
				"key1": {String: "value1", Valid: true},
			},
		},
		NullableAt: &oneOneEighteen,
	}
	person.Email.String = "sara@helloeave.com"
	person.Email.Valid = true

	err = s.
		Upsert("people").
		Columns("email", "nullable", "nullable_map", "nullable_at").
		Record(person).
		Where("name = 'Mario'").
		Returning("id", "email", "nullable", "nullable_map", "nullable_at").
		QueryStruct(&person)

	require.NoError(t, err)
	assert.True(t, person.ID > 0)
	assert.Equal(t, &yep, person.Nullable)
	assert.True(t, person.Email.Valid)
	assert.Equal(t, "sara@helloeave.com", person.Email.String)
	assert.True(t, oneOneEighteen.Equal(*person.NullableAt))
	assert.Equal(t, sql.NullString{String: "value1", Valid: true}, person.NullableMap.Map["key1"])
}

func TestUpsertReal_Nullability_PtrToRecord(t *testing.T) {
	// Insert by specifying values
	s := beginTxWithFixtures()
	defer s.AutoRollback()

	// with nil values
	person := Person{
		Nullable:    nil,
		NullableMap: hstore.Hstore{},
		NullableAt:  nil,
	}
	err := s.
		Upsert("people").
		Columns("nullable", "nullable_map", "nullable_at").
		Record(&person).
		Where("name = 'Mario'").
		Returning("id", "nullable", "nullable_map", "nullable_at").
		QueryStruct(&person)

	require.NoError(t, err)
	assert.True(t, person.ID > 0)
	assert.Nil(t, person.Nullable)
	assert.Nil(t, person.NullableMap.Map)
	assert.Nil(t, person.NullableAt)

	// with non-nil values
	oneOneEighteen := time.Date(2018, time.January, 1, 0, 0, 0, 0, time.UTC)
	yep := "yep"
	person = Person{
		Nullable: &yep,
		NullableMap: hstore.Hstore{
			Map: map[string]sql.NullString{
				"key1": {String: "value1", Valid: true},
			},
		},
		NullableAt: &oneOneEighteen,
	}
	person.Email.String = "sara@helloeave.com"
	person.Email.Valid = true

	err = s.
		Upsert("people").
		Columns("email", "nullable", "nullable_map", "nullable_at").
		Record(&person).
		Where("name = 'Mario'").
		Returning("id", "email", "nullable", "nullable_map", "nullable_at").
		QueryStruct(&person)

	require.NoError(t, err)
	assert.True(t, person.ID > 0)
	assert.Equal(t, &yep, person.Nullable)
	assert.True(t, person.Email.Valid)
	assert.Equal(t, "sara@helloeave.com", person.Email.String)
	assert.True(t, oneOneEighteen.Equal(*person.NullableAt))
	assert.Equal(t, sql.NullString{String: "value1", Valid: true}, person.NullableMap.Map["key1"])
}
