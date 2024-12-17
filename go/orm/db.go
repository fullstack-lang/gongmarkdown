// generated code - do not edit
package orm

import (
	"errors"
	"fmt"
	"strconv"
	"sync"

	"github.com/fullstack-lang/gongmarkdown/go/db"
)

// Ensure DBLite implements DBInterface
var _ db.DBInterface = &DBLite{}

// DBLite is an in-memory database implementation of DBInterface
type DBLite struct {
	// Mutex to protect shared resources
	mu sync.RWMutex

	// insertion point definitions

	anotherdummydataDBs map[uint]*AnotherDummyDataDB

	nextIDAnotherDummyDataDB uint

	cellDBs map[uint]*CellDB

	nextIDCellDB uint

	dummydataDBs map[uint]*DummyDataDB

	nextIDDummyDataDB uint

	elementDBs map[uint]*ElementDB

	nextIDElementDB uint

	markdowncontentDBs map[uint]*MarkdownContentDB

	nextIDMarkdownContentDB uint

	rowDBs map[uint]*RowDB

	nextIDRowDB uint
}

// NewDBLite creates a new instance of DBLite
func NewDBLite() *DBLite {
	return &DBLite{
		// insertion point maps init

		anotherdummydataDBs: make(map[uint]*AnotherDummyDataDB),

		cellDBs: make(map[uint]*CellDB),

		dummydataDBs: make(map[uint]*DummyDataDB),

		elementDBs: make(map[uint]*ElementDB),

		markdowncontentDBs: make(map[uint]*MarkdownContentDB),

		rowDBs: make(map[uint]*RowDB),
	}
}

// Create inserts a new record into the database
func (db *DBLite) Create(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gongmarkdown/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point create
	case *AnotherDummyDataDB:
		db.nextIDAnotherDummyDataDB++
		v.ID = db.nextIDAnotherDummyDataDB
		db.anotherdummydataDBs[v.ID] = v
	case *CellDB:
		db.nextIDCellDB++
		v.ID = db.nextIDCellDB
		db.cellDBs[v.ID] = v
	case *DummyDataDB:
		db.nextIDDummyDataDB++
		v.ID = db.nextIDDummyDataDB
		db.dummydataDBs[v.ID] = v
	case *ElementDB:
		db.nextIDElementDB++
		v.ID = db.nextIDElementDB
		db.elementDBs[v.ID] = v
	case *MarkdownContentDB:
		db.nextIDMarkdownContentDB++
		v.ID = db.nextIDMarkdownContentDB
		db.markdowncontentDBs[v.ID] = v
	case *RowDB:
		db.nextIDRowDB++
		v.ID = db.nextIDRowDB
		db.rowDBs[v.ID] = v
	default:
		return nil, errors.New("github.com/fullstack-lang/gongmarkdown/go, unsupported type in Create")
	}
	return db, nil
}

// Unscoped sets the unscoped flag for soft-deletes (not used in this implementation)
func (db *DBLite) Unscoped() (db.DBInterface, error) {
	return db, nil
}

// Model is a placeholder in this implementation
func (db *DBLite) Model(instanceDB any) (db.DBInterface, error) {
	// Not implemented as types are handled directly
	return db, nil
}

// Delete removes a record from the database
func (db *DBLite) Delete(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gongmarkdown/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *AnotherDummyDataDB:
		delete(db.anotherdummydataDBs, v.ID)
	case *CellDB:
		delete(db.cellDBs, v.ID)
	case *DummyDataDB:
		delete(db.dummydataDBs, v.ID)
	case *ElementDB:
		delete(db.elementDBs, v.ID)
	case *MarkdownContentDB:
		delete(db.markdowncontentDBs, v.ID)
	case *RowDB:
		delete(db.rowDBs, v.ID)
	default:
		return nil, errors.New("github.com/fullstack-lang/gongmarkdown/go, unsupported type in Delete")
	}
	return db, nil
}

// Save updates or inserts a record into the database
func (db *DBLite) Save(instanceDB any) (db.DBInterface, error) {

	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gongmarkdown/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *AnotherDummyDataDB:
		db.anotherdummydataDBs[v.ID] = v
		return db, nil
	case *CellDB:
		db.cellDBs[v.ID] = v
		return db, nil
	case *DummyDataDB:
		db.dummydataDBs[v.ID] = v
		return db, nil
	case *ElementDB:
		db.elementDBs[v.ID] = v
		return db, nil
	case *MarkdownContentDB:
		db.markdowncontentDBs[v.ID] = v
		return db, nil
	case *RowDB:
		db.rowDBs[v.ID] = v
		return db, nil
	default:
		return nil, errors.New("github.com/fullstack-lang/gongmarkdown/go, Save: unsupported type")
	}
}

// Updates modifies an existing record in the database
func (db *DBLite) Updates(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gongmarkdown/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *AnotherDummyDataDB:
		if existing, ok := db.anotherdummydataDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db AnotherDummyData github.com/fullstack-lang/gongmarkdown/go, record not found")
		}
	case *CellDB:
		if existing, ok := db.cellDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Cell github.com/fullstack-lang/gongmarkdown/go, record not found")
		}
	case *DummyDataDB:
		if existing, ok := db.dummydataDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db DummyData github.com/fullstack-lang/gongmarkdown/go, record not found")
		}
	case *ElementDB:
		if existing, ok := db.elementDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Element github.com/fullstack-lang/gongmarkdown/go, record not found")
		}
	case *MarkdownContentDB:
		if existing, ok := db.markdowncontentDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db MarkdownContent github.com/fullstack-lang/gongmarkdown/go, record not found")
		}
	case *RowDB:
		if existing, ok := db.rowDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Row github.com/fullstack-lang/gongmarkdown/go, record not found")
		}
	default:
		return nil, errors.New("github.com/fullstack-lang/gongmarkdown/go, unsupported type in Updates")
	}
	return db, nil
}

// Find retrieves all records of a type from the database
func (db *DBLite) Find(instanceDBs any) (db.DBInterface, error) {

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch ptr := instanceDBs.(type) {
	// insertion point find
	case *[]AnotherDummyDataDB:
		*ptr = make([]AnotherDummyDataDB, 0, len(db.anotherdummydataDBs))
		for _, v := range db.anotherdummydataDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]CellDB:
		*ptr = make([]CellDB, 0, len(db.cellDBs))
		for _, v := range db.cellDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]DummyDataDB:
		*ptr = make([]DummyDataDB, 0, len(db.dummydataDBs))
		for _, v := range db.dummydataDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]ElementDB:
		*ptr = make([]ElementDB, 0, len(db.elementDBs))
		for _, v := range db.elementDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]MarkdownContentDB:
		*ptr = make([]MarkdownContentDB, 0, len(db.markdowncontentDBs))
		for _, v := range db.markdowncontentDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]RowDB:
		*ptr = make([]RowDB, 0, len(db.rowDBs))
		for _, v := range db.rowDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	default:
		return nil, errors.New("github.com/fullstack-lang/gongmarkdown/go, Find: unsupported type")
	}
}

// First retrieves the first record of a type from the database
func (db *DBLite) First(instanceDB any, conds ...any) (db.DBInterface, error) {
	if len(conds) != 1 {
		return nil, errors.New("github.com/fullstack-lang/gongmarkdown/go, Do not process when conds is not a single parameter")
	}

	var i uint64
	var err error

	switch cond := conds[0].(type) {
	case string:
		i, err = strconv.ParseUint(cond, 10, 32) // Base 10, 32-bit unsigned int
		if err != nil {
			return nil, errors.New("github.com/fullstack-lang/gongmarkdown/go, conds[0] is not a string number")
		}
	case uint64:
		i = cond
	case uint:
		i = uint64(cond)
	default:
		return nil, errors.New("github.com/fullstack-lang/gongmarkdown/go, conds[0] is not a string or uint64")
	}

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch instanceDB.(type) {
	// insertion point first
	case *AnotherDummyDataDB:
		tmp, ok := db.anotherdummydataDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First AnotherDummyData Unkown entry %d", i))
		}

		anotherdummydataDB, _ := instanceDB.(*AnotherDummyDataDB)
		*anotherdummydataDB = *tmp
		
	case *CellDB:
		tmp, ok := db.cellDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Cell Unkown entry %d", i))
		}

		cellDB, _ := instanceDB.(*CellDB)
		*cellDB = *tmp
		
	case *DummyDataDB:
		tmp, ok := db.dummydataDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First DummyData Unkown entry %d", i))
		}

		dummydataDB, _ := instanceDB.(*DummyDataDB)
		*dummydataDB = *tmp
		
	case *ElementDB:
		tmp, ok := db.elementDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Element Unkown entry %d", i))
		}

		elementDB, _ := instanceDB.(*ElementDB)
		*elementDB = *tmp
		
	case *MarkdownContentDB:
		tmp, ok := db.markdowncontentDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First MarkdownContent Unkown entry %d", i))
		}

		markdowncontentDB, _ := instanceDB.(*MarkdownContentDB)
		*markdowncontentDB = *tmp
		
	case *RowDB:
		tmp, ok := db.rowDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Row Unkown entry %d", i))
		}

		rowDB, _ := instanceDB.(*RowDB)
		*rowDB = *tmp
		
	default:
		return nil, errors.New("github.com/fullstack-lang/gongmarkdown/go, Unkown type")
	}
	
	return db, nil
}

