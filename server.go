package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"html/template"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

func OkResponse(c *gin.Context, data interface{}) {
	c.JSON(200, data)
}

func ErrorResponse(c *gin.Context, code int, error error) {
	c.JSON(code, gin.H{
		"error": error.Error(),
	})
}

func runServer() {
	r := gin.Default()

	templateRoute := r.Group("/template")
	{
		type UploadTemplateForm struct {
			Template *multipart.FileHeader `form:"template" binding:"required"`
		}
		templateRoute.POST("/upload", func(c *gin.Context) {

			var form UploadTemplateForm
			if err := c.ShouldBind(&form); err != nil {
				c.String(http.StatusBadRequest, "bad request")
				ErrorResponse(c, http.StatusBadRequest, err)
				return
			}
			isValid, err := ValidateFilename(form.Template.Filename)
			if err != nil {
				ErrorResponse(c, http.StatusBadRequest, err)
				return
			}
			if !isValid {
				ErrorResponse(c, http.StatusBadRequest, errors.New("invalid filename. It must have letters, numbers and \"_\" and \"-\" characters."))
				return
			}

			file, err := form.Template.Open()
			defer file.Close()
			if err != nil {
				ErrorResponse(c, http.StatusBadRequest, err)
				return
			}
			err = PutTemplate(file, form.Template.Filename, form.Template.Size)
			if err != nil {
				ErrorResponse(c, http.StatusBadRequest, err)
				return
			}

			OkResponse(c, gin.H{
				"message": "template uploaded",
			})
		})

		templateRoute.DELETE("/delete", func(c *gin.Context) {
			filename := c.Query("filename")
			if filename == "" {
				ErrorResponse(c, 400, errors.New("No filename provided"))
				return
			}
			isValid, err := ValidateFilename(filename)
			if err != nil {
				ErrorResponse(c, http.StatusBadRequest, err)
				return
			}
			if !isValid {
				ErrorResponse(c, http.StatusBadRequest, errors.New("invalid filename. It must have letters, numbers and \"_\" and \"-\" characters."))
				return
			}
			err = DeleteTemplate(filename)
			if err != nil {
				ErrorResponse(c, http.StatusBadRequest, err)
				return
			}
			OkResponse(c, gin.H{
				"message": "template deleted",
			})
		})
	}

	emailRoute := r.Group("/email")
	{
		type SendEmailInput struct {
			TemplateName string       `json:"template_name" binding:"required"`
			TemplateData *interface{} `json:"template_data"`
			EmailSepcs   MJInput      `json:"email_specs" binding:"required"`
		}
		emailRoute.POST("/send", func(c *gin.Context) {
			var input SendEmailInput
			err := c.ShouldBindJSON(&input)
			if err != nil {
				ErrorResponse(c, http.StatusBadRequest, err)
				return
			}

			object, err := GetTemplate(input.TemplateName)
			if err != nil {
				ErrorResponse(c, http.StatusBadRequest, err)
				return
			}
			dataInBytes, err := ioutil.ReadAll(object)
			if err != nil {
				ErrorResponse(c, http.StatusBadRequest, err)
				return
			}

			tmpl, err := template.New(input.TemplateName).Parse(string(dataInBytes))
			if err != nil {
				ErrorResponse(c, http.StatusBadRequest, err)
				return
			}
			var buff bytes.Buffer

			// merging: template = => final email <= = data
			err = tmpl.Execute(&buff, input.TemplateData)
			if err != nil {
				ErrorResponse(c, http.StatusBadRequest, err)
				return
			}
			input.EmailSepcs.HTMLPart = buff.String()

			// validating from data
			if input.EmailSepcs.From == nil {
				input.EmailSepcs.From = &RecipientInputPart{
					Email: mjSenderEmail,
					Name:  mjSenderName,
				}
			}

			err = SendEmail(&input.EmailSepcs)
			if err != nil {
				ErrorResponse(c, http.StatusInternalServerError, err)
				return
			}

			OkResponse(c, gin.H{
				"message": "email sent",
			})
		})
	}

	err := r.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err)
	}
}
