package main

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"

	"github.com/fullstack-lang/gongmarkdown/go/controllers"
	"github.com/fullstack-lang/gongmarkdown/go/models"
	"github.com/fullstack-lang/gongmarkdown/go/orm"

	// gong stack for model analysis
	gong_controllers "github.com/fullstack-lang/gong/go/controllers"
	gong_models "github.com/fullstack-lang/gong/go/models"
	gong_orm "github.com/fullstack-lang/gong/go/orm"
	_ "github.com/fullstack-lang/gong/ng"

	// for diagrams
	gongdoc_controllers "github.com/fullstack-lang/gongdoc/go/controllers"
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
	gongdoc_orm "github.com/fullstack-lang/gongdoc/go/orm"
	_ "github.com/fullstack-lang/gongdoc/ng"

	gongmarkdown "github.com/fullstack-lang/gongmarkdown"
)

var (
	logDBFlag  = flag.Bool("logDB", false, "log mode for db")
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")

	marshallOnStartup = flag.String("marshallOnStartup", "", "at startup, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")
	unmarshall        = flag.String("unmarshall", "", "unmarshall data from marshall name and '.go' (must be lowercased without spaces), If unmarshall arg is '', no unmarshalling")
	marshallOnCommit  = flag.String("marshallOnCommit", "", "on all commits, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")

	diagrams = flag.Bool("diagrams", true, "parse/analysis go/models and go/diagrams (takes a few seconds)")

	generateDocument = flag.Bool("generateDocument", false, "generates a markdown document")
)

// InjectionGateway is the singloton that stores all functions
// that can set the objects the stage
// InjectionGateway stores function as a map of names
var InjectionGateway = make(map[string](func()))

// hook marhalling to stage
type BeforeCommitImplementation struct {
}

func (impl *BeforeCommitImplementation) BeforeCommit(stage *models.StageStruct) {
	file, err := os.Create(fmt.Sprintf("./%s.go", *marshallOnCommit))
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	models.Stage.Checkout()
	models.Stage.Marshall(file, "gongmarkdown/go/models", "main")
}

func main() {

	log.SetPrefix("gongmarkdown: ")
	log.SetFlags(0)

	// parse program arguments
	flag.Parse()

	// setup controlers
	if !*logGINFlag {
		myfile, _ := os.Create("/tmp/server.log")
		gin.DefaultWriter = myfile
	}
	r := gin.Default()
	r.Use(cors.Default())

	// setup GORM
	db := orm.SetupModels(*logDBFlag, ":memory:")
	dbDB, _ := db.DB()
	dbDB.SetMaxOpenConns(1)

	// add gongdocatabase
	gongdoc_orm.AutoMigrate(db)
	gong_orm.AutoMigrate(db)

	// generate injection code from the stage
	if *marshallOnStartup != "" {

		if strings.Contains(*marshallOnStartup, " ") {
			log.Fatalln(*marshallOnStartup + " must not contains blank spaces")
		}
		if strings.ToLower(*marshallOnStartup) != *marshallOnStartup {
			log.Fatalln(*marshallOnStartup + " must be lowercases")
		}

		file, err := os.Create(fmt.Sprintf("./%s.go", *marshallOnStartup))
		if err != nil {
			log.Fatal(err.Error())
		}
		defer file.Close()

		models.Stage.Checkout()
		models.Stage.Marshall(file, "gongmarkdown/go/models", "main")
		os.Exit(0)
	}

	// setup the stage by injecting the code from code database
	if *unmarshall != "" {
		models.Stage.Checkout()
		models.Stage.Reset()
		models.Stage.Commit()
		if InjectionGateway[*unmarshall] != nil {
			InjectionGateway[*unmarshall]()
		}
		models.Stage.Commit()
	}

	// hook automatic marshall to go code at every commit
	if *marshallOnCommit != "" {
		hook := new(BeforeCommitImplementation)
		models.Stage.OnInitCommitFromFrontCallback = hook
	}

	if *diagrams {
		// load package to analyse
		modelPkg := &gong_models.ModelPkg{}

		gong_models.Walk("../../models", modelPkg)
		modelPkg.SerializeToStage()
		gong_models.Stage.Commit()

		// create the diagrams
		// prepare the model views
		pkgelt := new(gongdoc_models.Pkgelt)

		// first, get all gong struct in the model
		for gongStruct := range gong_models.Stage.GongStructs {

			// let create the gong struct in the gongdoc models
			// and put the numbre of instances
			gongStruct_ := (&gongdoc_models.GongStruct{Name: gongStruct.Name}).Stage()
			nbInstances, ok := models.Stage.Map_GongStructName_InstancesNb[gongStruct.Name]
			if ok {
				gongStruct_.NbInstances = nbInstances
			}
		}

		// classdiagram can only be fully in memory when they are Unmarshalled
		// for instance, the Name of diagrams or the Name of the Link
		pkgelt.Unmarshall(modelPkg.PkgPath, "../../diagrams")
		pkgelt.SerializeToStage()
	}

	controllers.RegisterControllers(r)
	gongdoc_controllers.RegisterControllers(r)
	gong_controllers.RegisterControllers(r)
	gongdoc_models.Stage.Commit()

	if *generateDocument {
		markdownContent := (&models.MarkdownContent{Name: "Dummy"}).Stage()
		markdownContent.Name = "Dummy"

		root := (&models.Element{Name: "Root"}).Stage()
		root.Type = models.PARAGRAPH
		markdownContent.Root = root

		// table := (&models.Element{Name: "Table"}).Stage()
		// table.Content = "Synthetic table"
		// table.Type = models.TABLE

		dummyData1 := (&models.DummyData{}).Stage()
		dummyData1.Name = "dummyData1"
		dummyData1.DummyString = "dummyData1 string"
		dummyData1.DummyBool = true
		dummyData1.DummyDuration = time.Hour + time.Minute
		dummyData1.DummyInt = 42
		dummyData1.DummyFloat = 1.62

		dummyData2 := (&models.DummyData{}).Stage()
		dummyData2.Name = "dummyData2"
		dummyData2.DummyString = "dummyData2 string"
		dummyData2.DummyBool = true
		dummyData2.DummyDuration = 3*time.Hour + 34*time.Minute
		dummyData2.DummyInt = 43
		dummyData2.DummyFloat = 5.77

		dummies := []*models.DummyData{dummyData1, dummyData2}

		another := (&models.AnotherDummyData{Name: "another"}).Stage()
		dummyData1.DummyPointerToGongStruct = another

		table := models.GenerateTableOfDummnies()
		root.SubElements = append(root.SubElements, table)

		// generic call

		tableGeneric := GenerateTableFromDummyDatas(dummies)
		root.SubElements = append(root.SubElements, tableGeneric)

	}

	// fetch the document singloton
	var singloton *models.MarkdownContent
	for s := range models.Stage.MarkdownContents {
		singloton = s
	}
	singloton.UpdateContent()

	models.Stage.Commit()

	// provide the static route for the angular pages
	r.Use(static.Serve("/", EmbedFolder(gongmarkdown.NgDistNg, "ng/dist/ng")))
	r.NoRoute(func(c *gin.Context) {
		fmt.Println(c.Request.URL.Path, "doesn't exists, redirect on /")
		c.Redirect(http.StatusMovedPermanently, "/")
		c.Abort()
	})

	log.Printf("Server ready serve on localhost:8080")
	r.Run()
}

type embedFileSystem struct {
	http.FileSystem
}

func (e embedFileSystem) Exists(prefix string, path string) bool {
	_, err := e.Open(path)
	return err == nil
}

func EmbedFolder(fsEmbed embed.FS, targetPath string) static.ServeFileSystem {
	fsys, err := fs.Sub(fsEmbed, targetPath)
	if err != nil {
		panic(err)
	}
	return embedFileSystem{
		FileSystem: http.FS(fsys),
	}
}
