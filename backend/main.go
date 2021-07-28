package main
 
import (
   "context"
   "fmt"
   "log"
 
   "github.com/wipha26/app/controllers"
   _ "github.com/wipha26/app/docs"
   "github.com/wipha26/app/ent"
   "github.com/wipha26/app/ent/gender"
   "github.com/gin-contrib/cors"
   "github.com/gin-gonic/gin"
   _ "github.com/mattn/go-sqlite3"
   swaggerFiles "github.com/swaggo/files"
   ginSwagger "github.com/swaggo/gin-swagger"
)

type Users struct {
	User []User
}

type User struct {
	Name  string
	Idemployee string
   Gender int

}

type Genders struct {
	Gender []Gender
}

type Gender struct {
	Gender string
}

// @title SUT SA Example API
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/
 
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
 
// @host localhost:8080
// @BasePath /api/v1
 
// @securityDefinitions.basic BasicAuth
 
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
 
// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information
 
// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information
 
// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information
 
// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {
   router := gin.Default()
   router.Use(cors.Default())
 
   client, err := ent.Open("sqlite3", "file:dinformation.db?cache=shared&_fk=1")
   if err != nil {
       log.Fatalf("fail to open sqlite3: %v", err)
   }
   defer client.Close()
 
   if err := client.Schema.Create(context.Background()); err != nil {
       log.Fatalf("failed creating schema resources: %v", err)
   }
 
   v1 := router.Group("/api/v1")
   controllers.NewUserController(v1, client)
   controllers.NewGenderController(v1, client)

   // Set Genders Data
	genders := Genders{
		Gender: []Gender{
			Gender{"ชาย"},
			Gender{"หญิง"},
		},
	}

	for _, g := range genders.Gender {
		client.Gender.
			Create().
			SetGender(g.Gender).
			Save(context.Background())
	}

   //Set User Data
   users := Users{
      User: []User{
         User{"123123","สมหญิง ใจสู้",2},
         User{"123124","สมชาย ใจสู้",1},

      },
   }

   for _, u := range users.User{
        
      g,err := client.Gender.
               Query().
               Where(gender.IDEQ(int(u.Gender))).
               Only(context.Background())
            
      if err != nil {
            fmt.Println(err.Error())
            return
         }


      client.User.
         Create().
         SetIdemployee(u.Idemployee).
         SetName(u.Name).
         SetGender(g).
         Save(context.Background())
   }
 
   router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
   router.Run()
}