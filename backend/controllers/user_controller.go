package controllers
 
import (
   "context"
   "fmt"
   "strconv"
   "github.com/wipha26/app/ent"
   "github.com/wipha26/app/ent/user"
   "github.com/wipha26/app/ent/gender"
   "github.com/gin-gonic/gin"
)
 
// UserController defines the struct for the user controller
type UserController struct {
   client *ent.Client
   router gin.IRouter
}

type User struct {
	Name        string
	Idemployee  string
	Gender      int
}

// CreateUser handles POST requests for adding user entities
// @Summary Create user
// @Description Create user
// @ID create-user
// @Accept   json
// @Produce  json
// @Param user body ent.User true "User entity"
// @Success 200 {object} ent.User
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /users [post]
func (ctl *UserController) CreateUser(c *gin.Context) {
   obj := User{}
   if err := c.ShouldBind(&obj); err != nil {
       c.JSON(400, gin.H{
           "error": "user binding failed",
       })
       return
   }

   g, err := ctl.client.Gender.
			Query().
			Where(gender.IDEQ(int(obj.Gender))).
			Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
				"error": "gender not found",
		})
		return
	}
 
   u, err := ctl.client.User.
       Create().
       SetName(obj.Name).
       SetIdemployee(obj.Idemployee).
	   SetGender(g).
       Save(context.Background())

   if err != nil {
       c.JSON(400, gin.H{
           "status": false,
           "error" : err,
       })
       return
   }
 
   c.JSON(200, gin.H{
       "status" : true,
        "date" : u,
     })
}
// GetUser handles GET requests to retrieve a user entity
// @Summary Get a user entity by ID
// @Description get user by ID
// @ID get-user
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} ent.User
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /users/{id} [get]
func (ctl *UserController) GetUser(c *gin.Context) {
   id, err := strconv.ParseInt(c.Param("id"), 10, 64)
   if err != nil {
       c.JSON(400, gin.H{
           "error": err.Error(),
       })
       return
   }
 
   u, err := ctl.client.User.
       Query().
       Where(user.IDEQ(int(id))).
       Only(context.Background())
   if err != nil {
       c.JSON(404, gin.H{
           "error": err.Error(),
       })
       return
   }
 
   c.JSON(200, u)
}
// ListUser handles request to get a list of user entities
// @Summary List user entities
// @Description list user entities
// @ID list-user
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.User
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /users [get]
func (ctl *UserController) ListUser(c *gin.Context) {
   limitQuery := c.Query("limit")
   limit := 10
   if limitQuery != "" {
       limit64, err := strconv.ParseInt(limitQuery, 10, 64)
       if err == nil {limit = int(limit64)}
   }
 
   offsetQuery := c.Query("offset")
   offset := 0
   if offsetQuery != "" {
       offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
       if err == nil {offset = int(offset64)}
   }
 
   users, err := ctl.client.User.
       Query().
	   WithGender().
       Limit(limit).
       Offset(offset).
       All(context.Background())
   	if err != nil {
       c.JSON(400, gin.H{"error": err.Error(),})
       return
   }
 
   c.JSON(200, users)
}
// DeleteUser handles DELETE requests to delete a user entity
// @Summary Delete a user entity by ID
// @Description get user by ID
// @ID delete-user
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /users/{id} [delete]
func (ctl *UserController) DeleteUser(c *gin.Context) {
   id, err := strconv.ParseInt(c.Param("id"), 10, 64)
   if err != nil {
       c.JSON(400, gin.H{
           "error": err.Error(),
       })
       return
   }
 
   err = ctl.client.User.
       DeleteOneID(int(id)).
       Exec(context.Background())
   if err != nil {
       c.JSON(404, gin.H{
           "error": err.Error(),
       })
       return
   }
 
   c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
}

// NewUserController creates and registers handles for the user controller
func NewUserController(router gin.IRouter, client *ent.Client) *UserController {
   uc := &UserController{
       client: client,
       router: router,
   }
   uc.register()
   return uc
}
 
// InitUserController registers routes to the main engine
func (ctl *UserController) register() {
   users := ctl.router.Group("/users")
 
   users.GET("", ctl.ListUser)
 
   // CRUD
   users.POST("", ctl.CreateUser)
   users.GET(":id", ctl.GetUser)
   users.DELETE(":id", ctl.DeleteUser)
}