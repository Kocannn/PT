package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kocannn/api-programming-tadulako/domain"
)


func NewUserHandler(userServices domain.UserServices) *UserHandler{
  return &UserHandler{
    UserService: userServices,
  }
}

type UserHandler struct {
	UserService domain.UserServices
}


func (h *UserHandler) Register(c *gin.Context){
  var user domain.User

  if err := c.ShouldBindJSON(&user); err != nil{
    c.JSON(http.StatusInternalServerError, gin.H{
      "error": err.Error(),
    })
    return
  }

  
  if err := h.UserService.Register(&user); err != nil{
    c.JSON(http.StatusInternalServerError, gin.H{
      "error": err.Error(),
    })
    return
  }

  

  c.JSON(http.StatusOK, gin.H{
    "message": "User created successfully",
  })
}

func (h *UserHandler) Login(c *gin.Context){
  var loginRequest struct{
    Email string `json:"email"`
    Password string `json:"password"`
  }

  if err := c.ShouldBindJSON(&loginRequest); err != nil{
    c.JSON(http.StatusBadRequest, gin.H{
      "error": err.Error(),
    })
    return
  }

  token, err := h.UserService.Login(loginRequest.Email, loginRequest.Password)
  if err != nil{
    c.JSON(http.StatusInternalServerError, gin.H{
     "error": err.Error(),
    })
    return
  }
  c.SetCookie("Authorization", token, 3600 * 24 * 30, "", "", false, true )

  c.JSON(http.StatusOK, gin.H{
    "token": token,
  })
}